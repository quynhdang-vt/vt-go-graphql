package vtgraphql

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	goSet "github.com/amosmzhang/go-set"
	"mime/multipart"
)

type VeritoneAPIClient interface {
	SetAPIToken(token string) error
	SetUserSessionToken(token string)

	// future
	ToGraphQLQuerySelector(interface{}) (string, error)

	UpdateTaskStatus(ctx context.Context, jobID, taskID string, taskStatus TaskStatus, taskOutput interface{}) error
	GetRecording(ctx context.Context, recordingID string) (*Recording, error)
	CreateAsset(ctx context.Context, recordingID string, data io.Reader, asset Asset) (*Asset, error)
	CreateAssetWithMD5(ctx context.Context, recordingID, md5Hash string, data io.Reader, asset Asset) (*Asset, error)
}

type APIClient struct {
	Config     APIConfig
	Timeout    time.Duration
	RetryDelay time.Duration
	tr         *http.Transport
}

type APIConfig struct {
	Token            string `json:"token" yaml:"token"`
	UserSessionToken string `json:"userSessionToken" yaml:"userSessionToken"`
	BaseURI          string `json:"baseURI" yaml:"baseURI"`
	Version          string `json:"version" yaml:"version"`
	MaxAttempts      int    `json:"maxAttempts" yaml:"maxAttempts"`
	Timeout          string `json:"timeout" yaml:"timeout"`
	RetryDelay       string `json:"retryDelay" yaml:"retryDelay"`
}

type ContextKeyString string

const (
	DefaultURI        = "https://core.aws-dev.veritone.com"
	DefaultVersion    = ""
	DefaultTimeout    = "5s"
	DefaultRetryDelay = "5s"
)

var (
	ErrorToken              = errors.New("API Token invalid")
	ErrorMaxAttempts        = errors.New("Max API call attempts reached")
	ErrorTimeout            = errors.New("API call timed out")
	ErrorHTTP               = errors.New("API call HTTP error")
	ErrorInvalidJobID       = errors.New("API call invalid job ID")
	ErrorInvalidJob         = errors.New("API call invalid job")
	ErrorInvalidTaskID      = errors.New("API call invalid task ID")
	ErrorInvalidTaskStatus  = errors.New("API call invalid task status")
	ErrorInvalidRecordingID = errors.New("API call invalid recording ID")
	ErrorInvalidRecording   = errors.New("API call invalid recording")
	ErrorInvalidMD5Hash     = errors.New("API call invalid md5 hash")
	ErrorInvalidLogin       = errors.New("API call invalid login")
)

var (
	SuccessCodeSet = goSet.New(
		http.StatusOK,
		http.StatusCreated,
		http.StatusNoContent,
	)
)

// for formulating requests easier
type Query struct {
	query     string
	querybody json.RawMessage
}

type Response struct {
	Data json.RawMessage `json:"data"`
}

// get new APIClient from config,
func New(config APIConfig) (*APIClient, error) {
	// fill in missing config with defaults
	if config.BaseURI == "" {
		config.BaseURI = DefaultURI
	}
	if config.Version == "" {
		config.Version = DefaultVersion
	}
	if config.Timeout == "" {
		config.Timeout = DefaultTimeout
	}
	if config.RetryDelay == "" {
		config.RetryDelay = DefaultRetryDelay
	}
	if config.MaxAttempts == 0 {
		config.MaxAttempts = 3
	}

	// validate token and config
	if config.Token == "" {
		return nil, ErrorToken
	}

	// get timeout
	timeout, err := time.ParseDuration(config.Timeout)
	if err != nil {
		return nil, err
	}
	delay, err := time.ParseDuration(config.RetryDelay)
	if err != nil {
		return nil, err
	}

	client := APIClient{
		Config:     config,
		Timeout:    timeout,
		RetryDelay: delay,
		tr:         &http.Transport{},
	}
	return &client, nil
}

// set api token on client
func (this *APIClient) SetAPIToken(token string) {
	this.Config.Token = token
}

// set user session token on client
func (this *APIClient) SetUserSessionToken(token string) {
	this.Config.UserSessionToken = token
}

// Given a type, if it's not a structure just returned
// if it's a structure:
// Spit out  '{'
//   get all the fields
//      spit out the field name + " "
//      if the field is a structure, call this function
// Spit out '}' at the end
func (this *APIClient) ToGraphQLQuerySelector(mytype interface{}) (string, error) {
	return "", nil
}

// make API call
func (this *APIClient) doRequestWithBody(ctx context.Context, reqBody io.Reader) (response *json.RawMessage, err error) {
	// get request url
	u, err := url.Parse(this.Config.BaseURI)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, this.Config.Version)

	// prepare request
	req, err := http.NewRequest(http.MethodPost, u.String(), reqBody)
	if err != nil {
		return nil, err
	}

	bearerToken := ""
	if this.Config.UserSessionToken != "" {
		bearerToken = fmt.Sprintf("Bearer %s", this.Config.UserSessionToken)
	} else {
		bearerToken = fmt.Sprintf("Bearer %s", this.Config.Token)
	}

	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Type", "application/json")

	// call API
	for i := 0; i < this.Config.MaxAttempts; i++ {
		// set timeout on context
		ctx, cancel := context.WithTimeout(ctx, this.Timeout)
		defer cancel()

		err = httpDo(ctx, this.tr, req, func(resp *http.Response, err error) error {
			if err != nil {
				log.Printf("raw error: %v\n", err)
				return ErrorHTTP
			}
			defer resp.Body.Close()

			// read response
			resBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			if len(resBody) == 0 {
				if resp.StatusCode != http.StatusNoContent {
					return fmt.Errorf("Response returned status code %d: %s", resp.StatusCode, string(resBody))
				}

				return nil
			}

			if !SuccessCodeSet.Contains(resp.StatusCode) {
				return fmt.Errorf("Response returned status code %d: %s", resp.StatusCode, string(resBody))
			}

			// deserialize
			resStruct := &Response{}
			err = json.Unmarshal(resBody, resStruct)
			if err != nil {
				return err
			}

			response = &resStruct.Data

			return nil
		})

		// check HTTP request error
		switch err {
		case nil:
			return response, nil
		case ErrorHTTP, ErrorTimeout:
			time.Sleep(this.RetryDelay)
			continue
		default:
			return response, err
		}
	}

	if err != nil {
		return response, err
	}

	return response, ErrorMaxAttempts
}

// make API call
func (this *APIClient) doRequestWithJson(ctx context.Context, request interface{}) (response *json.RawMessage, err error) {
	// get request body, if any
	var reqBody io.Reader
	if request != nil {
		sjson, err := json.Marshal(request)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(sjson)
	}
	return this.doRequestWithBody(ctx, reqBody)
}
func httpDo(ctx context.Context, tr *http.Transport, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	if tr == nil {
		tr = &http.Transport{}
	}

	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c // Wait for f to return.
		return ErrorTimeout
	case err := <-c:
		return err
	}
}

// Creates multi-part requests with data as file contents
func (this *APIClient) doMutiPartRequest(ctx context.Context, params map[string]string, paramName, paramValue string, data io.Reader) (response *json.RawMessage, err error) {

	// get request url
	u, err := url.Parse(this.Config.BaseURI)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, this.Config.Version)

	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)
	part, err := writer.CreateFormFile(paramName, paramValue)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, data)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// prepare request
	req, err := http.NewRequest(http.MethodPost, u.String(), reqBody)
	if err != nil {
		return nil, err
	}
	//req, err := http.NewRequest(http.MethodPost, this.Config., reqBody)
	req.Header.Set("Content-Type", writer.FormDataContentType()) // have to do this!

	bearerToken := ""
	if this.Config.UserSessionToken != "" {
		bearerToken = fmt.Sprintf("Bearer %s", this.Config.UserSessionToken)
	} else {
		bearerToken = fmt.Sprintf("Bearer %s", this.Config.Token)
	}
	req.Header.Add("Authorization", bearerToken)

	// call API -- can refactor this with the doRequest
	for i := 0; i < this.Config.MaxAttempts; i++ {
		// set timeout on context
		ctx, cancel := context.WithTimeout(ctx, this.Timeout)
		defer cancel()

		err = httpDo(ctx, this.tr, req, func(resp *http.Response, err error) error {
			if err != nil {
				log.Printf("raw error: %v\n", err)
				return ErrorHTTP
			}
			defer resp.Body.Close()

			// read response
			resBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			if len(resBody) == 0 {
				if resp.StatusCode != http.StatusNoContent {
					return fmt.Errorf("Response returned status code %d: %s", resp.StatusCode, string(resBody))
				}

				return nil
			}

			if !SuccessCodeSet.Contains(resp.StatusCode) {
				return fmt.Errorf("Response returned status code %d: %s", resp.StatusCode, string(resBody))
			}
			// deserialize
			resStruct := &Response{}
			err = json.Unmarshal(resBody, resStruct)
			if err != nil {
				return err
			}

			response = &resStruct.Data
			return nil
		})

		// check HTTP request error
		switch err {
		case nil:
			return response, nil
		case ErrorHTTP, ErrorTimeout:
			time.Sleep(this.RetryDelay)
			continue
		default:
			return response, err
		}
	}

	if err != nil {
		return response, err
	}

	return response, ErrorMaxAttempts

}

func MarshalMetadataAsJson(m map[string]interface{}) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	for k, v := range m {
		vb, err := json.Marshal(v)
		if err == nil {
			buffer.WriteString(fmt.Sprintf("%s:%s, ", k, string(vb)))
		}
	}
	buffer.WriteString("}")
	return buffer.String(), nil
}

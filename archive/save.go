package archive

import (
	"fmt"
	"encoding/json"
	"strings"
	veritoneSchema "github.com/quynhdang-vt/vt-go-graphql/vtgraphql_schema"
	veritoneAPI "github.com/quynhdang-vt/vt-go-graphql/vtgraphql"
	"log"

	"os"
	"time"
	"context"
)

type InterfaceType interface {
	GetInterfaceName() string
}
type testStruct struct {
	Name string
	Age  int
}

func (t testStruct) GetInterfaceName() string {
	return "MYINTERFACE"
}

type MyEnum string

const FIRST MyEnum = "FIRST"

func test10() {

	v1 := FIRST
	fmt.Printf("v1 = %s", v1)

}

// now that we get the structure set up, let's instantiate something
func test5() {

	_, typeMap, _ := veritoneSchema.GetDefaultSchema()

	buf := ""

	graphQLType := ""
	graphQLType = "EngineList"
	engineResultType := typeMap.RegularTypes[graphQLType]
	buf, _ = veritoneSchema.GenerateGraphQLTypeResponseStringForARequest(&engineResultType, typeMap.RegularTypes, make(map[string]bool), false)
	fmt.Printf("\n\n AHAHAHA %s >>>%s<<<<<\\n", graphQLType, buf)

	graphQLType = "TemporalDataObject"
	temporalDataObjectGraphQL := typeMap.RegularTypes[graphQLType]
	buf, _ = veritoneSchema.GenerateGraphQLTypeResponseStringForARequest(&temporalDataObjectGraphQL, typeMap.RegularTypes, make(map[string]bool), false)

	fmt.Printf("\n\n AHAHAHA %s >>>%s<<<<<\\n", graphQLType, buf)
}

func test6() {
	s := fmt.Sprintf("{ testme(id: \"%s\"){}}", "abc", 1234, 56)
	fmt.Println(s)

}
func test3() {
	//o := &veritoneSchema.SchemaQueryResponse{}
	//	veritoneSchema.IntrospectType(reflect.TypeOf( (*veritoneSchema.SchemaQueryResponse)(nil)).Elem(), " ")

	ss := veritoneSchema.InitializeASchemaQueryResponseStruct()
	fmt.Printf("Everything OK? %v", ss)
}
func test1() {

	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

}

func testRecording(veritoneAPIClient *veritoneAPI.APIClient, recordingID string) {

	recording, err := veritoneAPIClient.GetRecording(context.Background(), recordingID)
	if err != nil {
		fmt.Printf("Error getting recording %s, err=%s\n", recordingID, err)
	} else {
		fmt.Printf("YAHOO -- Got Recording %s --> %s\n", recordingID, recording)
	}
}
func test2(token string) {
	// calling graphql
	// Create veritone api client
	veritoneAPIConfig := veritoneAPI.APIConfig{
		Token:       token,
		BaseURI:     "https://api.aws-dev.veritone.com/v3/graphql", //"http://localhost:3000/graphql", // Veritone API instance to use (dev/stage/etc.)
		Version:     "",                                            // API version to use
		MaxAttempts: 1,                                             // how many times to call Veritone API for each request until successful response
		Timeout:     "15s",                                         // API call timeout (for example: "3s")
		RetryDelay:  "15s",                                         // time to wait before each retry
	}
	veritoneAPIClient, _ := veritoneAPI.New(veritoneAPIConfig)

	recordingId := "21099164"
	//	testRecording(veritoneAPIClient, recordingId)
	testCreateAsset(veritoneAPIClient, recordingId)
	log.Println("--------------------")
	log.Println("Now check Recording Again")
	//	testRecording(veritoneAPIClient, recordingId)
}
func testCreateAsset(veritoneAPIClient *veritoneAPI.APIClient, recordingID string) {
	ttmlBytes := "HI, BYE, NOMORE HELLOW WORLD.."
	ttmlAsset := veritoneAPI.Asset{
		AssetType:   "transcript",
		ContentType: "text/plain",
		Metadata: map[string]interface{}{
			"fileName": fmt.Sprintf("test-%d.ttml", time.Now().Unix()),
			"source":   fmt.Sprintf("test:%v", time.Now()),
			"size":     len(ttmlBytes),
		},
	}
	resultAsset, err := veritoneAPIClient.CreateAsset(context.Background(), recordingID, strings.NewReader(ttmlBytes), ttmlAsset)
	if err != nil {
		fmt.Printf("Error creating Asset recording %s, err=%s\n", recordingID, err)
	} else {
		fmt.Printf("YAHOO -- Created Asset for Recording %s --> %s\n", recordingID, resultAsset)

	}
}

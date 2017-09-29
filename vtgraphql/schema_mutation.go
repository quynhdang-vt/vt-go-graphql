package vtgraphql

type MutationCreateTDOInput struct {
	StartDateTime int64  `json:"startDateTime"`
	StopDateTime  int64  `json:"stopDateTime"`
	Source        string `json:"source,omitempty"`
	MediaId       string `json:"mediaId, omitempty"`
	Status        string `json:"status, omitempty"`
	ApplicationId string `json:"applicationId, omitempty"`
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
}
type MutationCreateTDO struct {
	Input MutationCreateTDOInput
}
type MutationUpdateTDOInput struct {
	Id            string `json:"id"`
	StartDateTime int64  `json:"startDateTime,omitempty"`
	StopDateTime  int64  `json:"stopDateTime,omitempty"`
	Source        string `json:"source,omitempty"`
	MediaId       string `json:"mediaId, omitempty"`
	Status        string `json:"status, omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
}

type MutationUpdateTDO struct {
//	Input MutationUpdateTDO
}

type MutationDeleteTDO struct {
	Id string `json:"id"`
}

type MutationCreateAssetInput struct {
	ContainerId string       `json:"containerId"`
	ContentType string       `json:"contentType,omitempty"`
	File        UploadedFile `json:"file"`
	Md5Checksum string       `json:"md5Checksum,omitempty"`
	AssetType   string       `json:"assetType,omitempty"`
	Uri         string       `json:"uri,omitempty"`
	Metadata    JSONData     `json:"metadata,omitempty"`
}

type MutationCreateAsset struct {
	Input MutationCreateAssetInput
}

type MutationDeleteAsset struct {
	Id string `json:"id"`
}

type MutationUpdateAssetInput struct {
	Id string `json:id`
}

type MutationUpdateAsset struct {
	Asset MutationUpdateAssetInput `json:"asset"`
}

type MutationCloneRequestInput struct {
	SourceApplicationId      string `json:"sourceApplicationId"`
	DestinationApplicationId string `json:"destinationApplicationId"`
	CloneBlobs               bool   `json:"cloneBlobs,omitempty"`
}

type MutationRequestClone struct {
	Input MutationCloneRequestInput
}

type MutationCreateEngineInput struct {
	IsPublic        bool   `json:"isPublic,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	CategoryId      string `json:"categoryId,omitempty"`
	DeploymentModel int    `json:"deploymentModel,omitempty"`
	Price           int    `json:"price,omitempty"`
}

type MutationCreateEngine struct {
	Input MutationCreateEngineInput
}

type MutationUpdateEngineInput struct {
	Id              string      `json:"id"`
	IsPublic        bool        `json:"isPublic,omitempty"`
	Name            string      `json:"name,omitempty"`
	Description     string      `json:"description,omitempty"`
	CategoryId      string      `json:"categoryId,omitempty"`
	DeploymentModel int         `json:"deploymentModel,omitempty"`
	Price           int         `json:"price,omitempty"`
	State           EngineState `json:"state,omitempty"`
}

type MutationUpdateEngine struct {
	Input MutationUpdateEngineInput
}

type MutationDeleteEngine struct {
	Id string `json:"id"`
}

type MutationCreateBuildInput struct {
	EngineId string `json:"engineId"`
}

type MutationCreateEngineBuild struct {
	Input MutationCreateBuildInput
}

// UpdateBuildInput another multipart!
type MutationUpdateBuildInput struct {
	Id       string            `json:"id"`
	EngineId string            `json:"engineId"`
	Action   BuildUpdateAction `json:"action"`
	File     UploadedFile      `json:"file"`
}

type MutationUpdateEngineBuild struct {
	Input MutationUpdateEngineInput
}

type MutationDeleteEngineBuild struct {
	Id string `json:"id"`
}

type MutationUpdateTaskInput struct {
	Id         string `json:"id"`
	Status     string `json:"status,omitempty"`
	JobId      string `json:"jobId"`
	TaskOutput string `json:"taskOutput,omitempty"`
}

type MutationUpdateTask struct {
	Input MutationUpdateTaskInput
}

type MutationPollTaskInput struct {
	Id          string   `json:"id"`
	JobId       string   `json:"jobId"`
	PollPayload JSONData `json:"pollPayload,omitempty"`
}

type MutationPollTask struct {
	Input MutationPollTaskInput
}

type MutationCreateTaskInput struct {
	TaskType string `json:"taskType,omitempty"`
}

type MutationCreateJobInput struct {
	Status   string                    `json:"status,omitempty"`
	TargetId string                    `json:"targetId, omitempty"`
	Tasks    []MutationCreateTaskInput `json:"tasks,omitempty"`

	Retries int `json:"retries,omitempty"`
}

type MutationCreateJob struct {
	Input MutationCreateJobInput
}

type MutationCancelJob struct {
	Id string `json:"id"`
}

type MutationRetryJob struct {
	Id string `json:"id"`
}

// and we'll do the marshal of this interface to have just
/*
"query":"mutation { input: {json-special-handling since it needs just fieldname:value with fieldname unquoted)  { outputParamType fieldnames } }"
*/

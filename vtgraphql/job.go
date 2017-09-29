package vtgraphql

import (
	"fmt"
)

const (
	TaskStatusRunning  TaskStatus = "running"
	TaskStatusComplete TaskStatus = "complete"
	TaskStatusFailed   TaskStatus = "failed"
)

type Job struct {
	JobID       string `json:"jobId,omitempty"`
	RecordingID string `json:"recordingId,omitempty"`
	Status      string `json:"status,omitempty"`

	CreatedDateTime  int64 `json:"createdDateTime,omitempty"`
	ModifiedDateTime int64 `json:"modifiedDateTime,omitempty"`

	Tasks []Task `json:"tasks,omitempty"`
}

func (this *Job) AddTask(task Task) {
	if this.Tasks == nil {
		this.Tasks = []Task{}
	}
	this.Tasks = append(this.Tasks, task)
}

type Task struct {
	TaskID   string      `json:"taskId,omitempty"`
	Status   TaskStatus  `json:"taskStatus,omitempty"`
	TaskType string      `json:"taskType,omitempty"` // Deprecated. Use EngineID where appropriate
	EngineID string      `json:"engineId,omitempty"`
	Payload  interface{} `json:"taskPayload,omitempty"`
	Output   interface{} `json:"taskOutput,omitempty"`
}

type TaskStatus string

type RecordingJobsResponse struct {
	TotalRecords int   `json:"totalRecords"`
	Records      []Job `json:"records"`
}

type taskPayloadDownloadFile struct {
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	StartDateTime int64                  `json:"startDateTime"`
	FileURI       string                 `json:"fileUri"`
}

func NewDownloadFileTask(fileURI string, startDateTime int64, metadata map[string]interface{}) (Task, error) {
	if fileURI == "" {
		return Task{}, fmt.Errorf("Missing fileURI!")
	}
	if startDateTime < 0 {
		return Task{}, fmt.Errorf("Invalid startDateTime: %d", startDateTime)
	}

	task := Task{
		EngineID: "download-file",
		Payload: taskPayloadDownloadFile{
			FileURI:       fileURI,
			StartDateTime: startDateTime,
			Metadata:      metadata,
		},
	}

	return task, nil
}

type taskPayloadTranscodeFfmpeg struct {
	SetAsPrimary bool   `json:"setAsPrimary"`
	Format       string `json:"format,omitempty"`
	Bitrate      int    `json:"bitrate,omitempty"`
	SampleRate   int    `json:"sampleRate,omitempty"`
	Mono         bool   `json:"mono,omitempty"`
}

func NewTranscodeFFMPEGTask(setAsPrimary bool, format string, bitrate, sampleRate int, mono bool) Task {
	task := Task{
		EngineID: "transcode-ffmpeg",
		Payload: taskPayloadTranscodeFfmpeg{
			SetAsPrimary: setAsPrimary,
			Format:       format,
			Bitrate:      bitrate,
			SampleRate:   sampleRate,
			Mono:         mono,
		},
	}

	return task
}

func NewSimpleTask(engineId string) Task {
	return Task{
		EngineID: engineId,
	}
}

func NewInsertIntoIndexTask() Task {
	return NewSimpleTask("insert-into-index")
}

func NewMentionGenerateTask() Task {
	return NewSimpleTask("mention-generate")
}

func NewTranscribeTask() Task {
	return NewSimpleTask("transcribe")
}

func NewNuanceTask() Task {
	return NewSimpleTask("nuance")
}

func NewGracenoteTask() Task {
	return NewSimpleTask("gracenote")
}

package vtgraphql

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// -----------------
// EngineState enum
type EngineState int16

const (
	Active EngineState = iota
	Disabled
	Pending
	Deleted
	Draft
	Ready
)

func (es EngineState) String() string {
	switch es {
	case Active:
		return "active"
	case Disabled:
		return "disabled"
	case Pending:
		return "pending"
	case Deleted:
		return "deleted"
	case Draft:
		return "draft"
	case Ready:
		return "ready"
	}
	return "active"
}

// Special case for JSONData serialization - it would be key:name with not quoting the key
type JSONData map[string]interface{}

func (this JSONData) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	for k, v := range this {
		vb, err := json.Marshal(v)
		if err == nil {
			buffer.WriteString(fmt.Sprintf("%s:%s, ", k, string(vb)))
		}
	}
	buffer.WriteString("}")

	return buffer.Bytes(), nil
}

/** Muti-part placeholder -- use to set the multipart param */
type UploadedFile struct {
	ParamName  string
	ParamValue string
}

// BuildUpdateAction enum
type BuildUpdateAction int

const (
	Deploy BuildUpdateAction = iota + 50
	Pause
	Unpause
	Approve
	Disapprove
	Invalidate
	Submit

// Upload
// Restart
)

func (ba BuildUpdateAction) String() string {
	switch ba {
	case Deploy:
		return "deploy"
	case Pause:
		return "pause"
	case Unpause:
		return "unpause"
	case Approve:
		return "approve"
	case Disapprove:
		return "disapprove"
	case Invalidate:
		return "invalidate"
	case Submit:
		return "submit"
	}
	return ""
}

// ---------------------
// IGNORE interface Page
// =====================

// Metadata is a beast ..
// "subtype" will need to implment the SpecifyFragment
type Metadata interface {
	Name() string
	SpecifyFrament() (string, error)
}
type CloneAssetIdMap struct {
	OldAssetId string `json:"oldAssetId"`
	NewAssetId string `json:"newAssetId"`
}
type CloneData struct {
	Date       string           `json:"date,omitempty"`
	OriginalId string           `json:"originalId"`
	CloneBlobs bool             `json:"cloneBlobs,omitempty"`
	AssetIdMap FilterAssetIdMapStruct `json:"assetIdMap,omitempty"`
	Name       string           `json:"name"`
}
type FilterAssetIdMapStruct struct {
	Filter CloneAssetIdMap
	values []CloneAssetIdMap
}

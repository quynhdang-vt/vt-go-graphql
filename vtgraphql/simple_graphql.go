package vtgraphql

import (
	"context"

	vtGraphQLAPI "github.com/quynhdang-vt/vt-go-graphql/vtgraphql_api"
	"fmt"
)
/*
func InitializeATemporalDataObjectStruct() *vtgraphQL.TemporalDataObject {
	s := &vtgraphQL.TemporalDataObject{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*s), reflect.ValueOf(s).Elem())
	return s
}

*/
func GraphQLGetRecording(ctx context.Context, apiClient *vtGraphQLAPI.APIClient, recordingID string) (res *vtGraphQLAPI.TemporalDataObject, err error) {
	const GetRecordingRequest = "{temporalDataObject(id: \"%s\") {" +
		"id startDateTime stopDateTime source applicationId    status    createdDateTime    modifiedDateTime " +
			"assets { " +
			"records {id name contentType  assetType  uri  createdDateTime   modifiedDateTime jsondata}}}}"
	output, err := apiClient.Query_temporalDataObject(ctx, fmt.Sprintf(GetRecordingRequest, recordingID))
	if err!=nil {
		return nil, err
	}
	fmt.Printf("Got recording = %s\n",	output.Id)
	return output, nil
}

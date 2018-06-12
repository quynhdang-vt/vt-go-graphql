package main

import (
	"context"
	"encoding/json"
	"fmt"
	veritoneSchema "github.com/quynhdang-vt/vt-go-graphql/vtgraphql_schema"
	vtgraphql "github.com/quynhdang-vt/vt-go-graphql/vtgraphql"
	vtGraphQLAPI "github.com/quynhdang-vt/vt-go-graphql/vtgraphql_api"
	"log"
	"os"
	"flag"
)



func getDefaultOutputDir() string{
	gopath := os.Getenv("GOPATH")
	if len(gopath)>0{
		return gopath+"/src/github.com/quynhdang-vt/vt-go-graphql"
	} else {
		return "/tmp"
	}
}
// Usage -genschema -schemaFile xxxx OR
//  -testRecording -recordingId xxxx
func main() {
	genSchemaPtr := flag.Bool("genschema", false, "Generate schema from /tmp")
	schemaFilePtr := flag.String("schemaFile", "/tmp/schema.json", "Schema file to generate")
	outputRootDir := flag.String("outrootdir", getDefaultOutputDir(), "Root directory to generate the packages")

	testRecordingPtr := flag.Bool("testRecording", false, "Test retrieval recording")
	recordingIdPtr := flag.String("recordingId", "21099164", "Recording id")
	tokenPtr := flag.String("token", "", "Token")

	flag.Parse()

	// check for outputRootDir
	var err error=nil
	if(*genSchemaPtr) {
		err= GenerateTypesFromSchema(*schemaFilePtr, *outputRootDir)
	}
	if (*testRecordingPtr) {
		err = TestRecordingWithGeneratedStub(*recordingIdPtr, *tokenPtr)
	}
	if err != nil {
		fmt.Printf("error = %v\n", err)
	}

}


// SERIOUS BUSINESS HERE... TO KEEP, the REST is just .. junk
func GenerateTypesFromSchema(schemaFile string, outputRootDir string) error {

	outJsonFile := outputRootDir+"/vtgraphql_schema/gen_schema_json.go"
	_, typeMap, err := veritoneSchema.GetSchemaFromFile(schemaFile, &outJsonFile)
	if err != nil {
		return fmt.Errorf("ERRROR? %v", err)
	}

	/** should take the GOPATH!, for now just cheat */
	return veritoneSchema.GenerateGraphQLAllTypes(*typeMap,
		outputRootDir+"/vtgraphql_api/gen_vt_api_types.go",
		outputRootDir+"/vtgraphql_api/gen_vt_api_other_types.go",
		outputRootDir+"/vtgraphql_api/gen_vt_api_impl.go",
		outputRootDir+"/vtgraphql_api/gen_sample_queries.go",
		outputRootDir+"/vtgraphql_api/gen_sample_mutations.go")
}

func TestRecordingWithGeneratedStub(recordingId string, token string) (err error) {
	 err = nil
	veritoneAPIConfig := vtGraphQLAPI.APIConfig{
		Token:       token,
		BaseURI:     "https://api.aws-dev.veritone.com/v3/graphql", //"http://localhost:3000/graphql", // Veritone API instance to use (dev/stage/etc.)
		Version:     "",                                            // API version to use
		MaxAttempts: 1,                                             // how many times to call Veritone API for each request until successful response
		Timeout:     "15s",                                         // API call timeout (for example: "3s")
		RetryDelay:  "15s",                                         // time to wait before each retry
	}
	veritoneAPIClient, _ := vtGraphQLAPI.New(veritoneAPIConfig)
	recording, err := vtgraphql.GraphQLGetRecording(context.Background(), veritoneAPIClient, recordingId )
	if err!=nil {
		log.Fatal(err)
	}
	pretty, _ := json.MarshalIndent(recording, "", "\t")
	fmt.Printf("Got recording yooohooo..%s", pretty)


	return err
}

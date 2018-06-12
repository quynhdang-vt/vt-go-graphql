package vtgraphql_schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

//==========================

/*
 read from a file that was the result of a schema query to the graphql server
 At the minimum, the query is expected to be similar to the query_schema.json
*/
func GetSchemaFromFile(introspectSchemaFile string, genSchemaFileName *string) (*SchemaQueryResponse, *TypeMap, error) {
	// now we can see if we can get the desrialization of the json here
	fileContents, err := ioutil.ReadFile(introspectSchemaFile)
	if err != nil {
		return nil, nil, err
	}
	if genSchemaFileName != nil {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("// GENERATED %s for use with your graphQL generated types\npackage vtgraphql_schema\n\nconst VT_GRAPHQL_SCHEMA = `%s`",
			time.Now(),
			string(fileContents)))
		WriteToFile(*genSchemaFileName, buf.String())
	}
	return getSchemaFromBytes(fileContents)
}

func getSchemaFromBytes(fileContents []byte) (*SchemaQueryResponse, *TypeMap, error) {
	//	globalSchema := InitializeASchemaQueryResponseStruct()
	globalSchema := new(SchemaQueryResponse)
	err := json.Unmarshal(fileContents, globalSchema)

	//	fmt.Printf("GOT IT?? \n%v\n", ss)
	typeMap := processSchema(globalSchema)
	fmt.Printf("Got MAP it? # of regular types=%d, # of interfaces=%d, # of scalar types=%d\n",
		len(typeMap.RegularTypes), len(typeMap.Interfaces), len(typeMap.ScalarTypes))

	return globalSchema, typeMap, err

}

// include this as you can..
func GetDefaultSchema() (*SchemaQueryResponse, *TypeMap, error) {
	return getSchemaFromBytes([]byte(VT_GRAPHQL_SCHEMA))
}

func processSchema(s *SchemaQueryResponse) *TypeMap {
	if s == nil {
		return nil
	}
	typeMap := TypeMap{
		Interfaces:   make(map[string][]string),
		RegularTypes: make(map[string]GraphQLType),
		ScalarTypes:  make([]string, 0, 10),
	}
	x:=s.Data.Schema.Types
	for _, v := range x {
		if strings.HasPrefix(v.Name, "_") {
			continue
		}
		typeMap.RegularTypes[v.Name] = v
		// find interface -- find the "concrete" classes that would delare to be part of interface
		// this is so that we can efficiently generate the fragments for input
		if v.Kind == "INTERFACE" {
			// add to the interfacs
			pp := make([]string, 0, len(v.PossibleTypes))
			for _, p := range v.PossibleTypes {
				//fmt.Printf("%d=%s\n", j, p.Name)
				pp = append(pp, p.Name)
			}
			typeMap.Interfaces[v.Name] = pp
		} else if v.Kind == "SCALAR" {
			typeMap.ScalarTypes = append(typeMap.ScalarTypes, v.Name)
		}
	}

	return &typeMap
}

func mapScalarType(gql_scalar string) string {
	if gql_scalar == "ID" || gql_scalar == "String" {
		return "string"
	}
	if gql_scalar == "Int" {
		return "int"
	}
	if gql_scalar == "Boolean" {
		return "bool"
	}

	return gql_scalar
}

func mapGraphQLScalarType(gql_scalar string) string {
	if gql_scalar == "ID" || gql_scalar == "String" {
		return "string"
	}
	if gql_scalar == "Int" {
		return "int"
	}
	if gql_scalar == "Boolean" {
		return "bool"
	}
	if gql_scalar == "JSON" {
		return "map[string]interface{}"
	}
	return "interface{}"
}

/** tricky tricky tricky
type --> if name ==null, make sure to look at either:
     KIND = LIST --> array of ofType,
	 KIND = NON_NULL --> check its ofType
			--> if its ofType != null --> return that
            --> if its ofType  == null --> then the ofType.Kind == LIST ---> return the ofType's ofType.name array

func mapType(aType *GraphQLSimpleTypeStruct, regularTypes map[string]GraphQLType, usePtr bool) string {

	// Given a type the following is possible
	// type name is empty --> if Kind is List --> []OfType.Name (or map string scalar)
	if aType == nil {
		return "" //ERROR
	}
	isArray := aType.Kind == "LIST"
	typeName := aType.Name
	if len(typeName) == 0 {
		// consider ofType
		if aType.OfType != nil {
			// check to see what type we must use
			if len(aType.OfType.Name) != 0 {
				typeName = aType.OfType.Name
			} else {
				if aType.OfType.OfType != nil &&
					len(aType.OfType.OfType.Name) != 0 {
					typeName = aType.OfType.OfType.Name
				} else {
					return "" // ERROR
				}
			}
			if aType.OfType.Kind == "LIST" {
				isArray = true
			}
		} else {
			return "" // ERROR !!!
		}
	}
	if isArray {
		return "[]" + mapScalarType(typeName)
	}
	if usePtr {
		return "*" + mapScalarType(typeName)
	}
	return mapScalarType(typeName)
}
*/
func getBaseType(aType *GraphQLSimpleTypeStruct, typeMap map[string]GraphQLType, usePtr bool) (res *GraphQLType, typeDesc string, isArray bool, err error) {

	// Given a type the following is possible
	// type name is empty --> if Kind is List --> []OfType.Name (or map string scalar)
	if aType == nil {
		return nil, "", false, fmt.Errorf("Nil type")
	}

	typeName := aType.Name
	isArray = aType.Kind == "LIST"

	if len(typeName) == 0 {

		// consider ofType
		if aType.OfType != nil {
			// check to see what type we must use
			if len(aType.OfType.Name) != 0 {
				typeName = aType.OfType.Name
			} else {
				if aType.OfType.OfType != nil &&
					len(aType.OfType.OfType.Name) != 0 {
					typeName = aType.OfType.OfType.Name
				} else {
					return nil, "", false, fmt.Errorf("Failed to get any type info") // ERROR
				}
			}
			if aType.OfType.Kind == "LIST" {
				isArray = true
			}
		} else {
			return nil, "", false, fmt.Errorf("Failed to get any type info")
		}
	}
	r1, ok := typeMap[typeName]
	if ok {
		if isArray {
			typeDesc = "[]" + mapScalarType(typeName)
		} else if usePtr {
			if r1.Kind == "SCALAR" && typeName == "ID" || typeName == "String" || typeName == "Boolean" || typeName == "Int"{
				typeDesc = mapScalarType(typeName)
			} else {
				typeDesc = "*"  + mapScalarType(typeName)
			}
		} else {
			typeDesc = mapScalarType(typeName)
		}
		return &r1, typeDesc, isArray, nil
	}
	return nil, "", false, fmt.Errorf("Invalid type %s", typeName)
}

func generateAField(v *GraphQLFieldStruct, m map[string]GraphQLType, prefix string) (
	interfaceS string, sampleS string, returnTypeString string) {
	var buf, sampleBuf bytes.Buffer
	buf.WriteString(fmt.Sprintf("    %s_%s (ctx context.Context, requestSpec string /*", prefix, v.Name))
	nArgs := len(v.Args)
	// generate Args
	if nArgs > 0 {

		for _, args := range v.Args {
			_, typeDesc, _, _ := getBaseType(args.Type, m, true)

			buf.WriteString(fmt.Sprintf("%s %s, ", args.Name, typeDesc))

			//TODO output
		}
	}
	buf.WriteString("*/) (output ")
	outputType, typeDesc, _, _ := getBaseType(v.Type, m, true)
	buf.WriteString(typeDesc)
	buf.WriteString(", err error)")

	// Generating sample queries or mutations
	argString, _ := GenerateArgForARequest(MapGenericNameSimpleTypeStructToNameTypeInterface(v.Args), m, "(", ")")
	outputString, _ := GenerateGraphQLTypeResponseStringForARequest(outputType, m, make(map[string]bool), false)
	sampleString := fmt.Sprintf("\nconst Sample%s%sRequest =`\n{\n\t%s %s\n  %s\n`\n", prefix, strings.Title(v.Name), v.Name, argString, outputString)
	//fmt.Println(sampleString)
	sampleBuf.WriteString(sampleString)

	// formulate the output string allocation from typeDesc
	// for struct:  &struct{}
	// for array: make(array,1)
	if strings.HasPrefix(typeDesc,"*"){
		returnTypeString = fmt.Sprintf(OutputStub_struct, typeDesc[1:len(typeDesc)])
	} else if strings.HasPrefix(typeDesc, "[]"){
		returnTypeString = fmt.Sprintf(OutputStub_array, typeDesc)
	}

	return buf.String(), sampleBuf.String(), returnTypeString
}

// from GlobalTypeMap
func GenerateGraphQLInterface(m TypeMap) (interfaceString string, apiImplString string, sampleQueryString string, sampleMutationString string) {
	/*
		Query --> prefix with 'Query_'+fieldName, make sure to camel the name

		Query_fieldName(input args, output args) error
	*/

	var buf ,sampleQueryBuf, sampleMutationBuf, apiBuf  bytes.Buffer
	queryType, foundQuery := m.RegularTypes["Query"]
	mutationType, foundMutation := m.RegularTypes["Mutation"]
	if !foundQuery && !foundMutation {
		log.Fatal("No Query nor Mutation were found in schema")
	}

	buf.WriteString("type VeritoneGraphQL interface {\n")
	for _, v := range queryType.Fields {
		s1, s2, s3 := generateAField( &v, m.RegularTypes, "Query")
		buf.WriteString(fmt.Sprintf("%s\n", s1))
		sampleQueryBuf.WriteString(s2)
		apiBuf.WriteString(fmt.Sprintf(API_Method_stub, s1, s3))
	}
	for _, v := range mutationType.Fields {
		s1, s2, s3 := generateAField( &v, m.RegularTypes, "Mutation")
		buf.WriteString(fmt.Sprintf("%s\n", s1))
		sampleMutationBuf.WriteString(s2)
		apiBuf.WriteString(fmt.Sprintf(API_Method_stub, s1, s3))
	}
	buf.WriteString("}\n")
	formattedBuf, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("Can't format interface? err=%v\n%s\n", err, buf)
	}
	formattedAPIBuf, err := format.Source(apiBuf.Bytes())
	if err != nil {
		log.Fatalf("Can't format interface? err=%v\n%s\n", err, apiBuf)
	}
	formattedSampleQueryBuf, err := format.Source(sampleQueryBuf.Bytes())
	formattedSampleMutationBuf, err := format.Source(sampleMutationBuf.Bytes())
	return string(formattedBuf), string(formattedAPIBuf), string(formattedSampleQueryBuf), string(formattedSampleMutationBuf)
}

/**
generate the types for struct that is not Query, nor Int, String
*/
func GenerateGraphQLTypes(m TypeMap) (structTypes string, otherTypes string, e error) {
	var structTypeBuf bytes.Buffer
	var otherTypeBuf bytes.Buffer
	var mainBuf *bytes.Buffer

	for _, theType := range m.RegularTypes {
		if theType.Name == "Query" || theType.Name == "Mutation" {
			continue
		}
		if theType.Kind == "ENUM" {
			otherTypeBuf.WriteString(GenerateGraphQLEnumInfo(&theType))
			continue
		}
		if theType.Kind == "SCALAR" {
			// want to have the right one?
			otherTypeBuf.WriteString(fmt.Sprintf("// SCALAR %s-->%s\n", theType.Name, mapGraphQLScalarType(theType.Name)))
			otherTypeBuf.WriteString(fmt.Sprintf("type %s %s\n", theType.Name, mapGraphQLScalarType(theType.Name)))
			continue
		}
		if theType.Kind == "INTERFACE" {
			otherTypeBuf.WriteString (fmt.Sprintf("// INTERFACE %s\n", theType.Name))
			mainBuf = &otherTypeBuf
		} else {
			mainBuf = &structTypeBuf
		}

		mainBuf.WriteString(fmt.Sprintf("type %s struct {\n", theType.Name))
		// fields in between here
		fields := theType.Fields
		if theType.Kind == "INPUT_OBJECT" {
			fields = theType.InputFields
		}
		jsonTags := ""
		for _, theField := range fields {
			// make sure to cap the field
			if theField.Type.Kind != "NON-NULL" {
				jsonTags = fmt.Sprintf("json:\"%s,omitempty\"", theField.Name)
			} else {
				jsonTags = fmt.Sprintf("json:\"%s\"", theField.Name)
			}
			baseType, typeDesc, _, err := getBaseType(theField.Type, m.RegularTypes, true)
			if err != nil {
				return "", "", nil
			}
			var buf2 bytes.Buffer
			buf2.WriteString(fmt.Sprintf("graphql:\"%s,%s\"", strings.ToLower(baseType.Kind), baseType.Name))

			// get the type and check to see if we need to get enum Values etc
			mainBuf.WriteString(fmt.Sprintf("%s %s `%s %s`\n", strings.Title(theField.Name), typeDesc, jsonTags, buf2.String()))
		}
		mainBuf.WriteString("}\n")
		if theType.Kind == "INTERFACE" {
			mainBuf.WriteString(GenerateGraphQLInterfaceReflectInfo(&theType))
		}
	}
	formattedBuf, err := format.Source(structTypeBuf.Bytes())
	if err != nil {
		log.Fatalf("Can't generate types?  err=%v\n%s\n", err, structTypeBuf)
	}
	formattedOtherBuf, err := format.Source(otherTypeBuf.Bytes())
	if err != nil {
		log.Fatalf("Can't generate types?  err=%v\n%s\n", err, otherTypeBuf)
	}
	return string(formattedBuf), string (formattedOtherBuf), nil
}

func GenerateGraphQLAllTypes(m TypeMap, outputFileName string, outputOtherFileName string, outputAPIStubFileName string,
	outputQuerySampleFileName string, outputMutationSampleFileName string) error {
	const fileContents = `
// GENERATED by vt-go-graphql %s
package vtgraphql_api

import (
	"context"
    )
%s
%s
`
	const apiStubFileContents = `
// GENERATED by vt-go-graphql %s
// %s
package vtgraphql_api
import (
	"context"
	"reflect"
	vtgraphQLSchema "github.com/quynhdang-vt/vt-go-graphql/vtgraphql_schema"
    )
%s
`

	const otherTypeFileContents = `
// GENERATED by vt-go-graphql %s
// %s
package vtgraphql_api

%s
`
	interfaces, apiStubs, sampleQueries, sampleMutations := GenerateGraphQLInterface(m)
	structTypes, otherTypes, err := GenerateGraphQLTypes(m)
	if err != nil {
		return err
	}

	WriteToFile(outputFileName, fmt.Sprintf(fileContents, time.Now(), interfaces, structTypes))
	WriteToFile(outputAPIStubFileName, fmt.Sprintf(apiStubFileContents, time.Now(), "API Stubs...", apiStubs))
	WriteToFile(outputOtherFileName, fmt.Sprintf(otherTypeFileContents, time.Now(), "ENUM, INTERFACE, SCALAR...", otherTypes))

	// sample queries, mutations
	const sampleFileContents = `
// GENERATED by vt-go-graphql %s
package vtgraphql_api
%s
`
	WriteToFile(outputQuerySampleFileName, fmt.Sprintf(sampleFileContents, time.Now(), sampleQueries))
	WriteToFile(outputMutationSampleFileName, fmt.Sprintf(sampleFileContents, time.Now(), sampleMutations))

	return nil
}

func concatNames(pNames []GenericNameStruct) string {
	nNames := len(pNames)
	if nNames == 0 {
		return ""
	}
	var buf bytes.Buffer
	for i, v := range pNames {
		buf.WriteString(v.Name)
		if i < nNames-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
func GenerateSampleBasedOnType(pType *GraphQLSimpleTypeStruct, regularTypes map[string]GraphQLType) string {
	fType, _, isArray, _ := getBaseType(pType, regularTypes, false)
	if fType.Name == "ID" || fType.Name == "String" {
		return "\"STRING_VALUE\""
	}
	if fType.Name == "Int" {
		return "10234"
	}
	if fType.Name == "Boolean" {
		return "true or false"
	}

	names := concatNames(fType.EnumValues)
	// just return array for now
	if fType.Kind == "ENUM" {
		if isArray {
			return "[" + names + "]"
		}
		return "[" + names + "]"
	}
	if fType.Kind == "INPUT_OBJECT" {
		pInputType, ok := regularTypes[fType.Name]
		if ok {
			// TODO actually, this needs to be treated like args with sample values
			// need to get the fields and feed back to the args.
			// for mutation, InputFields are populated..
			fields := pInputType.Fields
			if len(fields)==0 {
				fields = pInputType.InputFields
			}
			s, _ := GenerateArgForARequest(MapGraphQLFieldStructToNameTypeInterface(fields), regularTypes, "{", "}")
			return s;
		}
	}
	return "EH?"
}
func GenerateArgForARequest(args []NameTypeInterface, regularTypes map[string]GraphQLType, openChar string, closeChar string) (res string, err error) {
	// check for args..
	nArgs := len(args)
	if nArgs == 0 {
		return "", nil
	}
	var buf bytes.Buffer
	buf.WriteString(openChar)
	for i, arg := range args {
		buf.WriteString(fmt.Sprintf("%s: %s", arg.GetName(), GenerateSampleBasedOnType(arg.GetType(), regularTypes)))
		if i < nArgs-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(closeChar)
	return buf.String(), nil
}
func GenerateGraphQLTypeInterfaceWithFragments(pType *GraphQLType, regularTypes map[string]GraphQLType, seenMap map[string]bool, inRecursive bool) (res string, err error) {
    var buf bytes.Buffer
    buf.WriteString("{")
    // get the PossibleTypes
    for _, v:=range(pType.PossibleTypes){
    	aSubType, ok:=regularTypes[v.Name]
    	if (ok){
			s, _:=GenerateGraphQLTypeResponseStringForARequest(&aSubType, regularTypes, seenMap, true)
			buf.WriteString(fmt.Sprintf("...on %s %s", v.Name, s))
		}
	}
  	buf.WriteString("}")
  	return buf.String(), nil
}


// ENUM --> string, and the values are just strings
func GenerateGraphQLEnumInfo (pType *GraphQLType) (res string){
	if pType.Kind=="ENUM" {
		var buf bytes.Buffer

		const t = `
// ENUM %s
type %s string
		`

		const t2 = `const %s_%s %s = "%s"
		`

		buf.WriteString(fmt.Sprintf(t, pType.Name, pType.Name))
		for _, e := range(pType.EnumValues) {
			buf.WriteString(fmt.Sprintf(t2, pType.Name, e.Name, pType.Name, e.Name))
		}
		buf.WriteString(fmt.Sprintf("// END of ENUM %s\n\n", pType.Name))

		//fmt.Println(buf.String())
		return buf.String()
	} else {
		return ""
	}
}

// INTERFACE --> and subinterface
func GenerateGraphQLInterfaceReflectInfo (pType *GraphQLType) (res string) {
	const t1 = `
func (t %s) GetGraphQLInterfaceName() string {
	return "%s"
}
`

	const t2 = `
func (t %s) GetAllGraphQLImplementedTypes() string {
    return "%s"
}
	// END of INTERFACE %s
`
	var buf bytes.Buffer
	var buf2 bytes.Buffer
	if (pType.Kind == "INTERFACE"){
		nSubTypes := len(pType.PossibleTypes)
		if nSubTypes>0{
			for i, x := range(pType.PossibleTypes) {
				buf.WriteString(fmt.Sprintf(t1,
					x.Name, pType.Name))
				buf2.WriteString (x.Name)
				if i<nSubTypes-1 {
					buf2.WriteString(",")
				}
			}
			buf.WriteString(fmt.Sprintf(t2,
					pType.Name, buf2.String(), pType.Name))
		}
	}
	return buf.String()
}
func GenerateGraphQLTypeResponseStringForARequest(pType *GraphQLType, regularTypes map[string]GraphQLType, seenMap map[string]bool, inRecursive bool) (res string, err error) {
	var buf bytes.Buffer
	if pType == nil || pType.Kind == "SCALAR" || pType.Kind == "ENUM" {
		return "", nil
	}


	buf.WriteString("{")
	var fType *GraphQLType
	fTypeStr := ""
	argStr := ""
	if inRecursive && ( pType.Kind == "OBJECT" || pType.Kind == "INPUT_OBJECT" || pType.Kind=="INTERFACE") {
		seenMap[pType.Name] = true
	}


	fields := pType.Fields
	if pType.Kind == "INPUT_OBJECT" {
		fields = pType.InputFields
	}
	for _, f := range fields{
		if len(f.Args) > 0 {
			buf.WriteString(fmt.Sprintf("\n\t%s ", f.Name))
			argStr, err = GenerateArgForARequest(MapGenericNameSimpleTypeStructToNameTypeInterface(f.Args), regularTypes, "(", ")")
			buf.WriteString(argStr)
		} else {
			buf.WriteString(fmt.Sprintf("%s ", f.Name))

		}

		fType, _, _, err = getBaseType(f.Type, regularTypes, false)

		seenIt, ok := seenMap[fType.Name]
		if ok && seenIt {
			// skip it..
			//fmt.Printf("Seen this before %s\n", fType.Name)
		} else {
			fTypeStr, err = GenerateGraphQLTypeResponseStringForARequest(fType, regularTypes, seenMap, true)
			buf.WriteString(fTypeStr)

		}
	}
	buf.WriteString("}\n")
	return buf.String(), err
}

/** GEN API STUB */
const API_Method_stub = `
func (this *APIClient) 	%s {
	%s
	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

`

const OutputStub_struct = `
output = &%s{}
vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())
`

const OutputStub_array = `
output = make(%s, 1)
vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(output), reflect.ValueOf(output).Elem())
`

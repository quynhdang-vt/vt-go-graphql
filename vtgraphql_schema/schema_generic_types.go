package vtgraphql_schema

import (
	"encoding/json"
)

type NameTypeInterface interface{
	GetName() string;
	GetType() *GraphQLSimpleTypeStruct;
}

type GenericNameStruct struct {
	Name string `json:"name"`
}
type GraphQLOfTypeSimpleTypeStruct struct {
	Name   string             `json:"name"`
	Kind   string             `json:"kind"`
	OfType *GenericNameStruct `json:"ofType"`
}

type GenericNameSimpleTypeStruct struct {
	Name string                   `json:"name"`
	Type *GraphQLSimpleTypeStruct `json:"type"`
}
func (t GenericNameSimpleTypeStruct) GetName() string {
	return t.Name;
}
func (t GenericNameSimpleTypeStruct) GetType() *GraphQLSimpleTypeStruct {
	return t.Type;
}
func MapGenericNameSimpleTypeStructToNameTypeInterface(a1 []GenericNameSimpleTypeStruct) []NameTypeInterface {
	a2:=make([]NameTypeInterface, len(a1))
	for i, v:=range(a1) {
		a2[i] = v
	}
	return a2
}
type GraphQLSimpleTypeStruct struct {
	Name   string                         `json:"name"`
	Kind   string                         `json:"kind"`
	OfType *GraphQLOfTypeSimpleTypeStruct `json:"ofType"`
}
type GraphQLFieldStruct struct {
	Name string                        `json:"name"`
	Type *GraphQLSimpleTypeStruct      `json:"type"`
	Args []GenericNameSimpleTypeStruct `json:"args"`
}
func (t GraphQLFieldStruct) GetName() string {
	return t.Name;
}
func (t GraphQLFieldStruct) GetType() *GraphQLSimpleTypeStruct {
	return t.Type;
}
func MapGraphQLFieldStructToNameTypeInterface(a1 []GraphQLFieldStruct) []NameTypeInterface {
	a2:=make([]NameTypeInterface, len(a1))
	for i, v:=range(a1) {
		a2[i] = v
	}
	return a2
}
type GraphQLType struct {
	Name          string               `json:"name"`
	Kind          string               `json:"kind"`
	Fields        []GraphQLFieldStruct `json:"fields,omitempty"`
	Interfaces    []GenericNameStruct  `json:"interfaces"`
	PossibleTypes []GenericNameStruct  `json:"possibleTypes"`
	EnumValues    []GenericNameStruct  `json:"enumValues"`
	InputFields   []GraphQLFieldStruct `json:"inputFields"`
}

type TypesArray struct {
	Types []GraphQLType `json:"types"`
}
type SchemaStruct struct {
	Schema *TypesArray `json:"__schema"`
}
type SchemaQueryResponse struct {
	Data *SchemaStruct `json:"data"`
}

func (e SchemaQueryResponse) String() string {
	pretty, _ := json.MarshalIndent(&e, "", "\t")
	return string(pretty)
}

type TypeMap struct {
	Interfaces   map[string][]string // maintain the interface mapping
	RegularTypes map[string]GraphQLType
	ScalarTypes  []string
}

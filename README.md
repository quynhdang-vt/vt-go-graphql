# vt-go-graphql

This allows the generation of GO binding of a GraphQL schema.

The steps are:
1.  Retrieve schema from the GraphQL server by running `utils/get_schema.sh` script.  This script queries the 
GraphQL server using the query_schema.json by default.  The retrieved output is default to be  `schema.json`.
2.  Generate the structs, samples, API invocations by calling 
`main -genschema -schemaFile SCHEMA.JSON_LOC -outrootdir OUTPUTDIR`.
If GOPATH is defined, default outrootdir is `$GOPATH/src/github.com/quynhdang-vt/vt-go-graphql`
3.  You can then compile the vtgraphql_api and vtgraphql_schema package


## Package descriptions
vtgraphql_api is the main mostly-generated package that could be used to directly access GraphQL. 
See `vtgraphql_api/gen_vt_api_types.go` for the `VeritoneGraphQL` interface which contain
 the _Query_xxx_ or _Mutation_xxx_
as generated from the schema.  Sample of querySpecs to the Query_xx or Mutation_xxx methods could be found in
`vtgraphql_api/gen_sample_queries.go` or `vtgraphql_api/gen_sample_mutations.go`


The package `vtgraphql` contains most used operations at a higher level (e.g. _GetRecording_ vs. _Query_temporalDataObject_)

Some tricky process is the jsondata for input -- make sure to generate the field names WITHOUT quotes otherwise 
GraphQL server will reject the request with syntax errors.

See `vtgraphql/simple_graphql.go` 
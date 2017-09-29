package vtgraphql_api

type IsGraphQLInterfaceImplementation interface {
	GetGraphQLInterfaceName() string
}

type IsGraphQLInterface interface {
	GetAllGraphQLImplementedTypes() string
}

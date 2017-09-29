package vtgraphql

/** The pattern:
create argument structs the Query fields fro
 */
type QueryTemporalDataObjects struct {
	ApplicationId string `json:"applicationId,omitempty"`
	Id            string `json:"id,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty"`
}

type QueryTemporalDataObject struct {
	Id string `json:"id"`
}

type QueryAsset struct {
	Id string `json:"id"`
}

type QueryCloneRequests struct {
	Id            string `json:"id"`
	ApplicationId string `json:"applicationId,omitempty"`
}

type QueryEngines struct {
	Id              string      `json:"id"`
	CategoryId      string      `json:"categoryId,omitempty"`
	State           EngineState `json:"state,omitempty"`
	Owned           string      `json:"owned,omitempty"`
	LibraryRequired bool        `json:"libraryRequired,omitempty"`
	Offset          int         `json:"offset,omitempty"`
	Limit           int         `json:"limit,omitempty"`
}

type QueryEngineCategories struct {
	Id string `json:"id,omitempty"`
}

type QueryJobs struct {
	Id            string   `json:"id"`
	Status        []string `json:"status,omitempty"`
	Offset        int      `json:"offset,omitempty"`
	Limit         int      `json:"limit,omitempty"`
	ApplicationId string   `json:"applicationId,omitempty"`
	TargetId      string   `json:"targetId,omitempty"`
}

type QueryEntityIdentifierTypes struct {
	Id string `json:"id,omitempty"`
}

type QueryLibraryTypes struct {
	Id     string `json:"id,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type QueryLibraries struct {
	Id     string `json:"id,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type QueryApplications struct {
	Id     string `json:"id,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

type QueryOrganizations struct {
	Id     string `json:"id,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}
type QueryOrganization struct {
	Id string `json:"id"`
}
type QueryPermissions struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Offset int    `json:"offset,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}
type QueryUsers struct {
	Id              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
	Offset          int      `json:"offset,omitempty"`
	Limit           int      `json:"limit,omitempty"`
}
type QueryUser struct {
	Id string `json:"id"`
}
type QueryGroups struct {
	Id              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	OrganizationIds []string `json:"organizationIds,omitempty"`
	Offset          int      `json:"offset,omitempty"`
	Limit           int      `json:"limit,omitempty"`
}

// TODO me

type QuerySearchMentions struct {
	Search JSONData `json:"search"`
}

type QuerySearchMedia struct {
	Search JSONData `json:"search"`
}

type QueryFolder struct {
	Id string `json:"id"`
}
type QueryTask struct {
	Id string `json:"id"`
}
type QueryJob struct {
	Id string `json:"id"`
}

// TODO construct query strings for posting by reflecting on the type, and
// generate JSON query such as this:
// query: interface{}

// and we'll do the marshal of this interface to have just
/*
"query":"{ inputParamType(field that has value)  { outputParamType fieldnames } }"
*/
type GenericGraphQLQueryStatement struct {
	InputParams  interface{}
	OutputParams interface{}
}

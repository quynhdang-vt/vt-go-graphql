
// GENERATED by vt-go-graphql 2017-09-28 23:33:00.824056724 -0700 PDT
// API Stubs...
package vtgraphql_api
import (
	"context"
	"reflect"
	vtgraphQLSchema "github.com/quynhdang-vt/vt-go-graphql/vtgraphql_schema"
    )

func (this *APIClient) Query_temporalDataObjects(ctx context.Context, requestSpec string /*applicationId string, id string, offset int, limit int, */) (output *TDOList, err error) {

	output = &TDOList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_temporalDataObject(ctx context.Context, requestSpec string /*id string, */) (output *TemporalDataObject, err error) {

	output = &TemporalDataObject{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_asset(ctx context.Context, requestSpec string /*id string, */) (output *Asset, err error) {

	output = &Asset{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_cloneRequests(ctx context.Context, requestSpec string /*id string, applicationId string, offset int, limit int, */) (output *CloneRequestList, err error) {

	output = &CloneRequestList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_engines(ctx context.Context, requestSpec string /*id string, categoryId string, state []EngineState, owned bool, libraryRequired bool, offset int, limit int, */) (output *EngineList, err error) {

	output = &EngineList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_engineCategories(ctx context.Context, requestSpec string /*id string, */) (output []EngineCategory, err error) {

	output = make([]EngineCategory, 1)
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_jobs(ctx context.Context, requestSpec string /*id string, status []string, offset int, limit int, applicationId string, targetId string, */) (output *JobList, err error) {

	output = &JobList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_entityIdentifierTypes(ctx context.Context, requestSpec string /*id string, offset int, limit int, */) (output *EntityIdentifierTypeList, err error) {

	output = &EntityIdentifierTypeList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_libraryTypes(ctx context.Context, requestSpec string /*id string, offset int, limit int, */) (output *LibraryTypeList, err error) {

	output = &LibraryTypeList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_libraries(ctx context.Context, requestSpec string /*id string, name string, offset int, limit int, */) (output *LibraryList, err error) {

	output = &LibraryList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_library(ctx context.Context, requestSpec string /*id string, */) (output *Library, err error) {

	output = &Library{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_entity(ctx context.Context, requestSpec string /*id string, */) (output *Entity, err error) {

	output = &Entity{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_applications(ctx context.Context, requestSpec string /*id string, status *ApplicationStatus, offset int, limit int, */) (output *ApplicationList, err error) {

	output = &ApplicationList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_organizations(ctx context.Context, requestSpec string /*id string, offset int, limit int, */) (output *OrganizationList, err error) {

	output = &OrganizationList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_organization(ctx context.Context, requestSpec string /*id string, */) (output *Organization, err error) {

	output = &Organization{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_permissions(ctx context.Context, requestSpec string /*id string, name string, offset int, limit int, */) (output *PermissionList, err error) {

	output = &PermissionList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_users(ctx context.Context, requestSpec string /*id string, name string, organizationIds []string, offset int, limit int, */) (output *UserList, err error) {

	output = &UserList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_user(ctx context.Context, requestSpec string /*id string, */) (output *User, err error) {

	output = &User{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_me(ctx context.Context, requestSpec string /**/) (output *User, err error) {

	output = &User{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_groups(ctx context.Context, requestSpec string /*id string, name string, organizationIds []string, offset int, limit int, */) (output *GroupList, err error) {

	output = &GroupList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_searchMentions(ctx context.Context, requestSpec string /*search *JSON, */) (output *SearchResult, err error) {

	output = &SearchResult{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_searchMedia(ctx context.Context, requestSpec string /*search *JSON, */) (output *SearchResult, err error) {

	output = &SearchResult{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_folder(ctx context.Context, requestSpec string /*id string, */) (output *Folder, err error) {

	output = &Folder{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_task(ctx context.Context, requestSpec string /*id string, */) (output *Task, err error) {

	output = &Task{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_job(ctx context.Context, requestSpec string /*id string, */) (output *Job, err error) {

	output = &Job{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_application(ctx context.Context, requestSpec string /*id string, */) (output *Application, err error) {

	output = &Application{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_ingestionConfiguration(ctx context.Context, requestSpec string /*id string, */) (output *IngestionConfiguration, err error) {

	output = &IngestionConfiguration{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_ingestionConfigurations(ctx context.Context, requestSpec string /*id string, offset int, limit int, name string, startDate *DateTime, endDate *DateTime, sources []string, applicationId string, emailAddress string, */) (output *IngestionConfigurationList, err error) {

	output = &IngestionConfigurationList{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Query_graphqlServiceInfo(ctx context.Context, requestSpec string /**/) (output *GraphQLServiceInfo, err error) {

	output = &GraphQLServiceInfo{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createTDO(ctx context.Context, requestSpec string /*input *CreateTDO, */) (output *TemporalDataObject, err error) {

	output = &TemporalDataObject{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateTDO(ctx context.Context, requestSpec string /*input *UpdateTDO, */) (output *TemporalDataObject, err error) {

	output = &TemporalDataObject{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteTDO(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createAsset(ctx context.Context, requestSpec string /*input *CreateAsset, */) (output *Asset, err error) {

	output = &Asset{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteAsset(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateAsset(ctx context.Context, requestSpec string /*asset *UpdateAsset, */) (output *Asset, err error) {

	output = &Asset{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_requestClone(ctx context.Context, requestSpec string /*input *RequestClone, */) (output *CloneRequest, err error) {

	output = &CloneRequest{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createEngine(ctx context.Context, requestSpec string /*input *CreateEngine, */) (output *Engine, err error) {

	output = &Engine{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateEngine(ctx context.Context, requestSpec string /*input *UpdateEngine, */) (output *Engine, err error) {

	output = &Engine{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteEngine(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createEngineBuild(ctx context.Context, requestSpec string /*input *CreateBuild, */) (output *Build, err error) {

	output = &Build{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateEngineBuild(ctx context.Context, requestSpec string /*input *UpdateBuild, */) (output *Build, err error) {

	output = &Build{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteEngineBuild(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateTask(ctx context.Context, requestSpec string /*input *UpdateTask, */) (output *Task, err error) {

	output = &Task{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_pollTask(ctx context.Context, requestSpec string /*input *PollTask, */) (output *Task, err error) {

	output = &Task{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createJob(ctx context.Context, requestSpec string /*input *CreateJob, */) (output *Job, err error) {

	output = &Job{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_cancelJob(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_retryJob(ctx context.Context, requestSpec string /*id string, */) (output *Job, err error) {

	output = &Job{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createApplication(ctx context.Context, requestSpec string /*input *CreateApplication, */) (output *Application, err error) {

	output = &Application{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteApplication(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateApplication(ctx context.Context, requestSpec string /*input *UpdateApplication, */) (output *Application, err error) {

	output = &Application{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createEntityIdentifierType(ctx context.Context, requestSpec string /*input *CreateEntityIdentifierType, */) (output *EntityIdentifierType, err error) {

	output = &EntityIdentifierType{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateEntityIdentifierType(ctx context.Context, requestSpec string /*input *UpdateEntityIdentifierType, */) (output *EntityIdentifierType, err error) {

	output = &EntityIdentifierType{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteEntityIdentifierType(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createLibraryType(ctx context.Context, requestSpec string /*input *CreateLibraryType, */) (output *LibraryType, err error) {

	output = &LibraryType{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateLibraryType(ctx context.Context, requestSpec string /*input *UpdateLibraryType, */) (output *LibraryType, err error) {

	output = &LibraryType{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteLibraryType(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createLibrary(ctx context.Context, requestSpec string /*input *CreateLibrary, */) (output *Library, err error) {

	output = &Library{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateLibrary(ctx context.Context, requestSpec string /*input *UpdateLibrary, */) (output *Library, err error) {

	output = &Library{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteLibrary(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createEntity(ctx context.Context, requestSpec string /*input *CreateEntity, */) (output *Entity, err error) {

	output = &Entity{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateEntity(ctx context.Context, requestSpec string /*input *UpdateEntity, */) (output *Entity, err error) {

	output = &Entity{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteEntity(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createEntityIdentifier(ctx context.Context, requestSpec string /*input *CreateEntityIdentifier, */) (output *EntityIdentifier, err error) {

	output = &EntityIdentifier{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateEntityIdentifier(ctx context.Context, requestSpec string /*input *UpdateEntityIdentifier, */) (output *EntityIdentifier, err error) {

	output = &EntityIdentifier{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteEntityIdentifier(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_applicationWorkflow(ctx context.Context, requestSpec string /*input *ApplicationWorkflow, */) (output *Application, err error) {

	output = &Application{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_createIngestionConfiguration(ctx context.Context, requestSpec string /*input *CreateIngestionConfiguration, */) (output *IngestionConfiguration, err error) {

	output = &IngestionConfiguration{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_updateIngestionConfiguration(ctx context.Context, requestSpec string /*input *UpdateIngestionConfiguration, */) (output *IngestionConfiguration, err error) {

	output = &IngestionConfiguration{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}

func (this *APIClient) Mutation_deleteIngestionConfiguration(ctx context.Context, requestSpec string /*id string, */) (output *DeletePayload, err error) {

	output = &DeletePayload{}
	vtgraphQLSchema.InitializeStruct("  ", "ROOT", reflect.TypeOf(*output), reflect.ValueOf(output).Elem())

	return output, this.doRequestWithJson(ctx, requestSpec, output)
}



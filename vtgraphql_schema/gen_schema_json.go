// GENERATED 2018-02-05 23:48:30.175162255 -0800 PST for use with your graphQL generated types
package vtgraphql_schema

const VT_GRAPHQL_SCHEMA = `{
  "data": {
    "__schema": {
      "types": [
        {
          "name": "Query",
          "fields": [
            {
              "name": "temporalDataObjects",
              "type": {
                "name": "TDOList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "applicationId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "sampleMedia",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "includePublic",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "temporalDataObject",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "asset",
              "type": {
                "name": "Asset",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "widget",
              "type": {
                "name": "Widget",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "collectionId",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "organizationId",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "cloneRequests",
              "type": {
                "name": "CloneRequestList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "applicationId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "engines",
              "type": {
                "name": "EngineList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "categoryId",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "category",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "state",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "EngineState",
                      "kind": "ENUM",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "owned",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "libraryRequired",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "createsTDO",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "engine",
              "type": {
                "name": "Engine",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "engineBuild",
              "type": {
                "name": "Build",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "engineCategories",
              "type": {
                "name": "EngineCategoryList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "engineCategory",
              "type": {
                "name": "EngineCategory",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "jobs",
              "type": {
                "name": "JobList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "status",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "String",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "applicationId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "targetId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "job",
              "type": {
                "name": "Job",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "task",
              "type": {
                "name": "Task",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "entityIdentifierTypes",
              "type": {
                "name": "EntityIdentifierTypeList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "entityIdentifierType",
              "type": {
                "name": "EntityIdentifierType",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "libraryTypes",
              "type": {
                "name": "LibraryTypeList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "libraryType",
              "type": {
                "name": "LibraryType",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "libraries",
              "type": {
                "name": "LibraryList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "type",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "includeOwnedOnly",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "orderBy",
                  "type": {
                    "name": "LibraryOrderBy",
                    "kind": "ENUM",
                    "ofType": null
                  }
                },
                {
                  "name": "orderDirection",
                  "type": {
                    "name": "OrderDirection",
                    "kind": "ENUM",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "library",
              "type": {
                "name": "Library",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "libraryEngineModel",
              "type": {
                "name": "LibraryEngineModel",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "entity",
              "type": {
                "name": "Entity",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "applications",
              "type": {
                "name": "ApplicationList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "status",
                  "type": {
                    "name": "ApplicationStatus",
                    "kind": "ENUM",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "organizations",
              "type": {
                "name": "OrganizationList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "organization",
              "type": {
                "name": "Organization",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "permissions",
              "type": {
                "name": "PermissionList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "users",
              "type": {
                "name": "UserList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "organizationIds",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "user",
              "type": {
                "name": "User",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "me",
              "type": {
                "name": "User",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "groups",
              "type": {
                "name": "GroupList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "organizationIds",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "mention",
              "type": {
                "name": "Mention",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "mentionId",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "userId",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "searchMentions",
              "type": {
                "name": "SearchResult",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "search",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "JSON",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "searchMedia",
              "type": {
                "name": "SearchResult",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "search",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "JSON",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "rootFolders",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Folder",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": [
                {
                  "name": "type",
                  "type": {
                    "name": "RootFolderType",
                    "kind": "ENUM",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "folder",
              "type": {
                "name": "Folder",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "application",
              "type": {
                "name": "Application",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "ingestionConfiguration",
              "type": {
                "name": "IngestionConfiguration",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "ingestionConfigurations",
              "type": {
                "name": "IngestionConfigurationList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "startDate",
                  "type": {
                    "name": "DateTime",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "endDate",
                  "type": {
                    "name": "DateTime",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "sources",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": null,
                      "kind": "NON_NULL",
                      "ofType": {
                        "name": "String"
                      }
                    }
                  }
                },
                {
                  "name": "applicationId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "emailAddress",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "schema",
              "type": {
                "name": "Schema",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "structuredData",
              "type": {
                "name": "StructuredData",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "nielsenTVData",
              "type": {
                "name": "NielsenTVData",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "graphqlServiceInfo",
              "type": {
                "name": "GraphQLServiceInfo",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "getSignedWritableUrl",
              "type": {
                "name": "WritableUrlInfo",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "key",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "myRights",
              "type": {
                "name": "RightsListing",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ID",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Int",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Boolean",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "TDOList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "TemporalDataObject",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Page",
          "fields": [
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "INTERFACE",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": [
            {
              "name": "TDOList"
            },
            {
              "name": "AssetList"
            },
            {
              "name": "TaskList"
            },
            {
              "name": "BuildList"
            },
            {
              "name": "EngineList"
            },
            {
              "name": "EntityIdentifierTypeList"
            },
            {
              "name": "JobList"
            },
            {
              "name": "ApplicationList"
            },
            {
              "name": "PermissionList"
            },
            {
              "name": "UserList"
            },
            {
              "name": "StructuredDataSchemaList"
            },
            {
              "name": "WatchlistList"
            },
            {
              "name": "CollectionList"
            },
            {
              "name": "WidgetList"
            },
            {
              "name": "CloneRequestList"
            },
            {
              "name": "EngineCategoryList"
            },
            {
              "name": "LibraryTypeList"
            },
            {
              "name": "LibraryList"
            },
            {
              "name": "LibraryEngineModelList"
            },
            {
              "name": "EntityList"
            },
            {
              "name": "EntityIdentifierList"
            },
            {
              "name": "OrganizationList"
            },
            {
              "name": "GroupList"
            },
            {
              "name": "IngestionConfigurationList"
            }
          ],
          "enumValues": null
        },
        {
          "name": "TemporalDataObject",
          "fields": [
            {
              "name": "createdDateTime",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mediaId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "metadata",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Metadata",
                  "kind": "INTERFACE",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "assets",
              "type": {
                "name": "AssetList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "type",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": null,
                      "kind": "NON_NULL",
                      "ofType": {
                        "name": "String"
                      }
                    }
                  }
                },
                {
                  "name": "assetType",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": null,
                      "kind": "NON_NULL",
                      "ofType": {
                        "name": "String"
                      }
                    }
                  }
                },
                {
                  "name": "sourceTaskId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "primaryAsset",
              "type": {
                "name": "Asset",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "assetType",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "String",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "security",
              "type": {
                "name": "Security",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "startDateTime",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "DateTime",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "stopDateTime",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "DateTime",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "source",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "applicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "tasks",
              "type": {
                "name": "TaskList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "hasSourceAsset",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "jobs",
              "type": {
                "name": "JobList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "folders",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "Folder"
                  }
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "DateTime",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "String",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Metadata",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "INTERFACE",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": [
            {
              "name": "KVP"
            },
            {
              "name": "CloneData"
            },
            {
              "name": "Program"
            },
            {
              "name": "FileData"
            },
            {
              "name": "StructuredJSONObject"
            },
            {
              "name": "JSONObject"
            }
          ],
          "enumValues": null
        },
        {
          "name": "JSON",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "AssetList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Asset",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Asset",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "contentType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "containerId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "container",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "uri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "signedUri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "type",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "assetType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "details",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsonstring",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": [
                {
                  "name": "indent",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "fileData",
              "type": {
                "name": "AssetFileData",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "sourceData",
              "type": {
                "name": "AssetSourceData",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "transform",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": [
                {
                  "name": "transformFunction",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "TransformFunction",
                      "kind": "ENUM",
                      "ofType": null
                    }
                  }
                }
              ]
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "AssetFileData",
          "fields": [
            {
              "name": "md5sum",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "size",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "originalFileUri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "AssetSourceData",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "taskId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "task",
              "type": {
                "name": "Task",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engineId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engine",
              "type": {
                "name": "Engine",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Task",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "queuedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "completedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "startedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "TaskStatus",
                "kind": "ENUM",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "order",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "isClone",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "applicationId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "targetId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "target",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engineId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engine",
              "type": {
                "name": "Engine",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jobId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "job",
              "type": {
                "name": "Job",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "buildId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "build",
              "type": {
                "name": "Build",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "sourceAsset",
              "type": {
                "name": "Asset",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "sourceAssetId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mediaLengthSec",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "payload",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "output",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "payloadString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "outputString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "log",
              "type": {
                "name": "TaskLog",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "taskPayload",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "taskOutput",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "testTask",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "TaskStatus",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "pending"
            },
            {
              "name": "running"
            },
            {
              "name": "complete"
            },
            {
              "name": "queued"
            },
            {
              "name": "accepted"
            },
            {
              "name": "failed"
            },
            {
              "name": "cancelled"
            }
          ]
        },
        {
          "name": "Engine",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "ownerOrganizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "isPublic",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "logoPath",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "categoryId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "state",
              "type": {
                "name": "EngineState",
                "kind": "ENUM",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "price",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "asset",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "displayName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "validateUri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "executeUri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "applicationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createsTDO",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "website",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "rating",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "libraryRequired",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "deploymentModel",
              "type": {
                "name": "DeploymentModel",
                "kind": "ENUM",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "tasks",
              "type": {
                "name": "TaskList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "status",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "String",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "hasSourceAsset",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "builds",
              "type": {
                "name": "BuildList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "status",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "String",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "fields",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "EngineField"
                  }
                }
              },
              "args": []
            },
            {
              "name": "category",
              "type": {
                "name": "EngineCategory",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EngineState",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "active"
            },
            {
              "name": "disabled"
            },
            {
              "name": "pending"
            },
            {
              "name": "deleted"
            },
            {
              "name": "draft"
            },
            {
              "name": "ready"
            }
          ]
        },
        {
          "name": "DeploymentModel",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "FullyNetworkIsolated"
            },
            {
              "name": "MostlyNetworkIsolated"
            },
            {
              "name": "NonNetworkIsolated"
            },
            {
              "name": "HumanReview"
            }
          ]
        },
        {
          "name": "TaskList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Task",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "BuildList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Build",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Build",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engineId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "price",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "validateUri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "executeUri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "dockerImage",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "runtime",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "version",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "report",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EngineField",
          "fields": [
            {
              "name": "max",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "min",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "step",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "type",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "EngineFieldType",
                  "kind": "ENUM",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "info",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "label",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "options",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "EngineFieldPicklistOption"
                  }
                }
              },
              "args": []
            },
            {
              "name": "defaultValue",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Float",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EngineFieldType",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "Number"
            },
            {
              "name": "Picklist"
            },
            {
              "name": "MultiPicklist"
            },
            {
              "name": "Text"
            }
          ]
        },
        {
          "name": "EngineFieldPicklistOption",
          "fields": [
            {
              "name": "key",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "value",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EngineCategory",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engineIds",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              },
              "args": []
            },
            {
              "name": "totalEngines",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "iconClass",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "editable",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "videoOnly",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engines",
              "type": {
                "name": "EngineList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "libraryEntityIdentifierTypeIds",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              },
              "args": []
            },
            {
              "name": "libraryEntityIdentifierTypes",
              "type": {
                "name": "EntityIdentifierTypeList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EngineList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Engine",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EntityIdentifierTypeList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "EntityIdentifierType"
                  }
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EntityIdentifierType",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "label",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "labelPlural",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "iconClass",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "dataType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "EntityIdentifierDataType",
                  "kind": "ENUM",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EntityIdentifierDataType",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "image"
            },
            {
              "name": "audio"
            },
            {
              "name": "video"
            },
            {
              "name": "text"
            }
          ]
        },
        {
          "name": "Job",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "targetId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "sourceAssetId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "TaskStatus",
                "kind": "ENUM",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "tasks",
              "type": {
                "name": "TaskList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "status",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "String",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "targetId",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "applicationId",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "hasSourceAsset",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "applicationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "target",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "TaskLog",
          "fields": [
            {
              "name": "uri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "text",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "TransformFunction",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "XML2JSON"
            },
            {
              "name": "Transcript2JSON"
            }
          ]
        },
        {
          "name": "Security",
          "fields": [
            {
              "name": "global",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "JobList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Job",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Folder",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "treeObjectId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "ownerId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "parent",
              "type": {
                "name": "Folder",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "subfolders",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "Folder"
                  }
                }
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": "Organization",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "typeId",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "rootFolderTypeId",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "maxDepth",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "orderIndex",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "FolderStatus",
                "kind": "ENUM",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "folderPath",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "Folder"
                  }
                }
              },
              "args": []
            },
            {
              "name": "childTDOs",
              "type": {
                "name": "TDOList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Organization",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "type",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "applications",
              "type": {
                "name": "ApplicationList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "seatLimit",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "OrganizationStatus",
                "kind": "ENUM",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "roles",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Role",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "users",
              "type": {
                "name": "UserList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "blacklist",
              "type": {
                "name": "EngineBlacklist",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "whitelist",
              "type": {
                "name": "EngineWhitelist",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "schemas",
              "type": {
                "name": "StructuredDataSchemaList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "watchlists",
              "type": {
                "name": "WatchlistList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "collections",
              "type": {
                "name": "CollectionList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "rootFolder",
              "type": {
                "name": "Folder",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "type",
                  "type": {
                    "name": "RootFolderType",
                    "kind": "ENUM",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "businessUnit",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ApplicationList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Application",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Application",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "key",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "category",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "iconUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "iconSvg",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "url",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "deploymentModel",
              "type": {
                "name": "DeploymentModel",
                "kind": "ENUM",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "clientSecret",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": [
                {
                  "name": "password",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "oauth2RedirectUrls",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "ApplicationStatus",
                "kind": "ENUM",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "permissionsRequired",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ApplicationStatus",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "active"
            },
            {
              "name": "draft"
            },
            {
              "name": "deleted"
            },
            {
              "name": "pending"
            },
            {
              "name": "rejected"
            },
            {
              "name": "approved"
            },
            {
              "name": "disabled"
            }
          ]
        },
        {
          "name": "OrganizationStatus",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "active"
            },
            {
              "name": "deleted"
            }
          ]
        },
        {
          "name": "Role",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "permissions",
              "type": {
                "name": "PermissionList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "PermissionList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Permission",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Permission",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UserList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "User",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "User",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "permissions",
              "type": {
                "name": "PermissionList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "roles",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "Role"
                  }
                }
              },
              "args": []
            },
            {
              "name": "roleIds",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": "Organization",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "firstName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "lastName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "email",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "acls",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "UserACL"
                  }
                }
              },
              "args": []
            },
            {
              "name": "rootFolder",
              "type": {
                "name": "Folder",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "type",
                  "type": {
                    "name": "RootFolderType",
                    "kind": "ENUM",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "passwordUpdatedDateTime",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "lastLoginDateTime",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UserACL",
          "fields": [
            {
              "name": "applicationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "objectType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "objectId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "access",
              "type": {
                "name": "UserACLAccessRights",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UserACLAccessRights",
          "fields": [
            {
              "name": "owner",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "RootFolderType",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "watchlist"
            },
            {
              "name": "collection"
            },
            {
              "name": "cms"
            }
          ]
        },
        {
          "name": "EngineBlacklist",
          "fields": [
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "engines",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Engine",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "engineCategories",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "EngineCategory",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EngineWhitelist",
          "fields": [
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "engines",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Engine",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "engineCategories",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "EngineCategory",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "StructuredDataSchemaList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "StructuredJSONSchema",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "StructuredJSONSchema",
          "fields": [
            {
              "name": "schema",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "ownerOrganizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Organization",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "WatchlistList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Watchlist",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Watchlist",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Organization",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "watchlistType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "marketName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "trackMyPrograms",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "trackAll",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "agencyId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "advertiserId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "brandId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "campaignId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "marketId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "budget",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mediaSourceTypeId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "startDate",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "stopDate",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CollectionList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Collection",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Collection",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "imageUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "ownerId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Organization",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "orgSharing",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "programCount",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "itemCount",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "typeId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "isActive",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "widgets",
              "type": {
                "name": "WidgetList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "WidgetList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Widget",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Widget",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": "Organization",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "collection",
              "type": {
                "name": "Collection",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "collectionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "displayCollectionName",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "displayTranscription",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "width",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "numberOfMentionsToShow",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "adScript",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "seoTags",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "backgroundColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "borderColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "textColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "FolderStatus",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "active"
            },
            {
              "name": "inactive"
            }
          ]
        },
        {
          "name": "CloneRequestList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "CloneRequest",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CloneRequest",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "sourceApplicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "destinationApplicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "numberOfRecordings",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "numberOfCompletedRecordings",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "percentage",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EngineCategoryList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "EngineCategory"
                  }
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LibraryTypeList",
          "fields": [
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "LibraryType",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LibraryType",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "label",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "iconClass",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "entityIdentifierTypes",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "EntityIdentifierType",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "entityTypeName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "entityTypeNamePlural",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "entityType",
              "type": {
                "name": "EntityType",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EntityType",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "namePlural",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "schema",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "JSON",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LibraryOrderBy",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "id"
            },
            {
              "name": "name"
            },
            {
              "name": "createdDateTime"
            },
            {
              "name": "modifiedDateTime"
            },
            {
              "name": "version"
            }
          ]
        },
        {
          "name": "OrderDirection",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "asc"
            },
            {
              "name": "desc"
            }
          ]
        },
        {
          "name": "LibraryList",
          "fields": [
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Library",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Library",
          "fields": [
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "properties",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Property",
                  "kind": "INTERFACE",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "security",
              "type": {
                "name": "Security",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "applicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "version",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "libraryType",
              "type": {
                "name": "LibraryType",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "libraryTypeId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "coverImageUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engineModels",
              "type": {
                "name": "LibraryEngineModelList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "entities",
              "type": {
                "name": "EntityList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "isPublished",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "identifierTypeId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "name",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "orderBy",
                  "type": {
                    "name": "LibraryEntityOrderBy",
                    "kind": "ENUM",
                    "ofType": null
                  }
                },
                {
                  "name": "orderDirection",
                  "type": {
                    "name": "OrderDirection",
                    "kind": "ENUM",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "collaborators",
              "type": {
                "name": "LibraryCollaboratorList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "collaboratorOrgId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "summary",
              "type": {
                "name": "LibrarySummary",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Property",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "INTERFACE",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": [
            {
              "name": "IntProperty"
            },
            {
              "name": "StringProperty"
            },
            {
              "name": "BooleanProperty"
            },
            {
              "name": "CompoundProperty"
            }
          ],
          "enumValues": null
        },
        {
          "name": "LibraryEngineModelList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "LibraryEngineModel",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LibraryEngineModel",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "engineId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "engine",
              "type": {
                "name": "Engine",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "libraryId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "library",
              "type": {
                "name": "Library",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "libraryVersion",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "trainJobId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "trainStatus",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "LibraryEngineModelTrainStatus",
                  "kind": "ENUM",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "dataUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "contentType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LibraryEngineModelTrainStatus",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "pending"
            },
            {
              "name": "queued"
            },
            {
              "name": "complete"
            },
            {
              "name": "failed"
            },
            {
              "name": "running"
            }
          ]
        },
        {
          "name": "LibraryEntityOrderBy",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "id"
            },
            {
              "name": "name"
            },
            {
              "name": "createdDateTime"
            },
            {
              "name": "modifiedDateTime"
            }
          ]
        },
        {
          "name": "EntityList",
          "fields": [
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Entity",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Entity",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "properties",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Property",
                  "kind": "INTERFACE",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "libraryId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "profileImageUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "identifiers",
              "type": {
                "name": "EntityIdentifierList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "identifierTypeId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "offset",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "limit",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "isPublished",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsonstring",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "summary",
              "type": {
                "name": "EntitySummary",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EntityIdentifierList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "EntityIdentifier",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EntityIdentifier",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "entityId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "identifierType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "EntityIdentifierType",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "identifierTypeId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "title",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "isPriority",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "url",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "contentType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsonstring",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "EntitySummary",
          "fields": [
            {
              "name": "identifierCountsByType",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LibraryCollaboratorList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "LibraryCollaborator",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LibraryCollaborator",
          "fields": [
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": "Organization",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "permissions",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "libraryId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "library",
              "type": {
                "name": "Library",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LibrarySummary",
          "fields": [
            {
              "name": "entityCount",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "unpublishedEntityCount",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "lastTrainedVersion",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "lastTrainedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "OrganizationList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Organization",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "GroupList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Group",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Group",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "applicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "User",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "User",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": "Organization",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Mention",
          "fields": [
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "mediaSourceTypeId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mediaSourceId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "programId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mediaId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "advertiserId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "brandId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "campaignId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "trackingUnitId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mentionStatusId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "complianceStatusId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "spotTypeId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "audienceMarketCount",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "audienceAffiliateCount",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mentionHitCount",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "audience",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mentionRating",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "isMatch",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mentionDate",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "metadata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mentionSnippets",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "MentionSnippets",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "userSnippets",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "MentionUserSnippets",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "adCreative",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "fingerprint",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "cognitiveEngineResults",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "comments",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "MentionComment",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MentionSnippets",
          "fields": [
            {
              "name": "text",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "startTime",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "endTime",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "hits",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "MentionHit",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MentionHit",
          "fields": [
            {
              "name": "queryTerm",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "startTime",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "endTime",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MentionUserSnippets",
          "fields": [
            {
              "name": "text",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "startTime",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "endTime",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "transcriptStartDate",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "transcriptEndDate",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MentionComment",
          "fields": [
            {
              "name": "commentId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "mentionId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "userId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "firstName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "lastName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "userImage",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "commentText",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "SearchResult",
          "fields": [
            {
              "name": "jsondata",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "JSON",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "IngestionConfiguration",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "applicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "type",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "emailAddress",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "job",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "ui",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "IngestionConfigurationList",
          "fields": [
            {
              "name": "offset",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "limit",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "count",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "IngestionConfiguration",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Page"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Schema",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "schema",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "schemaString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": [
                {
                  "name": "indent",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "uiComponent",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "ttlSec",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organizationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "StructuredData",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "externalId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "schemaId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "data",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "dataString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": [
                {
                  "name": "indent",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createdBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "NielsenTVData",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "decorators",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "decoratorsString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": [
                {
                  "name": "indent",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "schemaId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "viewershipType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "uniqueCode",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "networkCode",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "startDateTime",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "endDateTime",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "fileName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "programName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "demographics",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "demographicsString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": [
                {
                  "name": "indent",
                  "type": {
                    "name": "Int",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "modifiedBy",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "createdBy",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "GraphQLServiceInfo",
          "fields": [
            {
              "name": "buildInfo",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "WritableUrlInfo",
          "fields": [
            {
              "name": "bucket",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "key",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "expires",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "url",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "RightsListing",
          "fields": [
            {
              "name": "operations",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": null,
                  "kind": "LIST",
                  "ofType": {
                    "name": null
                  }
                }
              },
              "args": []
            },
            {
              "name": "resources",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Mutation",
          "fields": [
            {
              "name": "createTDO",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateTDO",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateTDO",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateTDO",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteTDO",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "cleanupTDO",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "options",
                  "type": {
                    "name": null,
                    "kind": "LIST",
                    "ofType": {
                      "name": null,
                      "kind": "NON_NULL",
                      "ofType": {
                        "name": "TDOCleanupOption"
                      }
                    }
                  }
                }
              ]
            },
            {
              "name": "createAsset",
              "type": {
                "name": "Asset",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateAsset",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "deleteAsset",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateAsset",
              "type": {
                "name": "Asset",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateAsset",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "requestClone",
              "type": {
                "name": "CloneRequest",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "RequestClone",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createEngine",
              "type": {
                "name": "Engine",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateEngine",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateEngine",
              "type": {
                "name": "Engine",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateEngine",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteEngine",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createEngineBuild",
              "type": {
                "name": "Build",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateBuild",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateEngineBuild",
              "type": {
                "name": "Build",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateBuild",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "deleteEngineBuild",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "DeleteBuild",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateTask",
              "type": {
                "name": "Task",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateTask",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "pollTask",
              "type": {
                "name": "Task",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "PollTask",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createJob",
              "type": {
                "name": "Job",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateJob",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "cancelJob",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "retryJob",
              "type": {
                "name": "Job",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createApplication",
              "type": {
                "name": "Application",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateApplication",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteApplication",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateApplication",
              "type": {
                "name": "Application",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateApplication",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateOrganization",
              "type": {
                "name": "Organization",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateOrganization",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createEntityIdentifierType",
              "type": {
                "name": "EntityIdentifierType",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateEntityIdentifierType",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateEntityIdentifierType",
              "type": {
                "name": "EntityIdentifierType",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateEntityIdentifierType",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createLibraryType",
              "type": {
                "name": "LibraryType",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateLibraryType",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateLibraryType",
              "type": {
                "name": "LibraryType",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateLibraryType",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createLibrary",
              "type": {
                "name": "Library",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateLibrary",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateLibrary",
              "type": {
                "name": "Library",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateLibrary",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "deleteLibrary",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createEntity",
              "type": {
                "name": "Entity",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateEntity",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateEntity",
              "type": {
                "name": "Entity",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateEntity",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "deleteEntity",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createEntityIdentifier",
              "type": {
                "name": "EntityIdentifier",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateEntityIdentifier",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateEntityIdentifier",
              "type": {
                "name": "EntityIdentifier",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateEntityIdentifier",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "deleteEntityIdentifier",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createLibraryEngineModel",
              "type": {
                "name": "LibraryEngineModel",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateLibraryEngineModel",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateLibraryEngineModel",
              "type": {
                "name": "LibraryEngineModel",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateLibraryEngineModel",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "deleteLibraryEngineModel",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "applicationWorkflow",
              "type": {
                "name": "Application",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "ApplicationWorkflow",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createIngestionConfiguration",
              "type": {
                "name": "IngestionConfiguration",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateIngestionConfiguration",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateIngestionConfiguration",
              "type": {
                "name": "IngestionConfiguration",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateIngestionConfiguration",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteIngestionConfiguration",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createWidget",
              "type": {
                "name": "Widget",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateWidget",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateWidget",
              "type": {
                "name": "Widget",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateWidget",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createUser",
              "type": {
                "name": "User",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateUser",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateUser",
              "type": {
                "name": "User",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateUser",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createPasswordUpdateRequest",
              "type": {
                "name": "User",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreatePasswordUpdateRequest",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteUser",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "id",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "ID",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createSchema",
              "type": {
                "name": "Schema",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateSchema",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createStructuredData",
              "type": {
                "name": "StructuredData",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateStructuredData",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateStructuredData",
              "type": {
                "name": "StructuredData",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateStructuredData",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createNielsenTVData",
              "type": {
                "name": "NielsenTVData",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "CreateNielsenTVData",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "updateNielsenTVData",
              "type": {
                "name": "NielsenTVData",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UpdateNielsenTVData",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createCollection",
              "type": {
                "name": "Collection",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateCollection",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateCollection",
              "type": {
                "name": "Collection",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateCollection",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteCollection",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "folderId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "id",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "shareCollection",
              "type": {
                "name": "Share",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "ShareCollection",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "shareMentionFromCollection",
              "type": {
                "name": "Share",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "ShareMentionFromCollection",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createCollectionMention",
              "type": {
                "name": "CollectionMention",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CollectionMentionInput",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteCollectionMention",
              "type": {
                "name": "CollectionMention",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CollectionMentionInput",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createFolder",
              "type": {
                "name": "Folder",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateFolder",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateFolder",
              "type": {
                "name": "Folder",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateFolder",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "moveFolder",
              "type": {
                "name": "Folder",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "MoveFolder",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteFolder",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "DeleteFolder",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createMentionComment",
              "type": {
                "name": "MentionComment",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateMentionComment",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateMentionComment",
              "type": {
                "name": "MentionComment",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateMentionComment",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteMentionComment",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "DeleteMentionComment",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "createMentionRating",
              "type": {
                "name": "MentionRating",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "CreateMentionRating",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "updateMentionRating",
              "type": {
                "name": "MentionRating",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UpdateMentionRating",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "deleteMentionRating",
              "type": {
                "name": "DeletePayload",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "DeleteMentionRating",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "userLogin",
              "type": {
                "name": "LoginInfo",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": "UserLogin",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "userLogout",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": [
                {
                  "name": "token",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "String",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "refreshToken",
              "type": {
                "name": "LoginInfo",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "token",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "String",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "validateToken",
              "type": {
                "name": "LoginInfo",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "token",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "String",
                      "kind": "SCALAR",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "createRootFolders",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Folder",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": [
                {
                  "name": "rootFolderType",
                  "type": {
                    "name": "RootFolderType",
                    "kind": "ENUM",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "bulkUpdateWatchlist",
              "type": {
                "name": "WatchlistList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "filter",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "BulkUpdateWatchlistFilter",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                },
                {
                  "name": "input",
                  "type": {
                    "name": "BulkUpdateWatchlist",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "fileTemporalDataObject",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "FileTemporalDataObject",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "unfileTemporalDataObject",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UnfileTemporalDataObject",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "moveTemporalDataObject",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "MoveTemporalDataObject",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            },
            {
              "name": "uploadEngineResult",
              "type": {
                "name": "Asset",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": [
                {
                  "name": "input",
                  "type": {
                    "name": null,
                    "kind": "NON_NULL",
                    "ofType": {
                      "name": "UploadEngineResult",
                      "kind": "INPUT_OBJECT",
                      "ofType": null
                    }
                  }
                }
              ]
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateTDO",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "startDateTime",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "DateTime",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "stopDateTime",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "DateTime",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "source",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "isPublic",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "parentFolderId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateTDO",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "startDateTime",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "stopDateTime",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "source",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "primaryAsset",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "SetPrimaryAsset"
                  }
                }
              }
            },
            {
              "name": "isPublic",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "SetPrimaryAsset",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "assetType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "DeletePayload",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "message",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "TDOCleanupOption",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "storage"
            },
            {
              "name": "searchIndex"
            },
            {
              "name": "engineResults"
            }
          ]
        },
        {
          "name": "CreateAsset",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "containerId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "contentType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "file",
              "type": {
                "name": "UploadedFile",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "md5sum",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "assetType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "type",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "uri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "fileData",
              "type": {
                "name": "SetAssetFileData",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            },
            {
              "name": "sourceData",
              "type": {
                "name": "SetAssetSourceData",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            },
            {
              "name": "details",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UploadedFile",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "SetAssetFileData",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "md5sum",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "size",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "originalFileUri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "SetAssetSourceData",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "taskId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "engineId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateAsset",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "fileData",
              "type": {
                "name": "SetAssetFileData",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            },
            {
              "name": "sourceData",
              "type": {
                "name": "SetAssetSourceData",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            },
            {
              "name": "details",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "RequestClone",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "sourceApplicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "destinationApplicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "cloneBlobs",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateEngine",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "isPublic",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "categoryId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "deploymentModel",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "DeploymentModel",
                  "kind": "ENUM",
                  "ofType": null
                }
              }
            },
            {
              "name": "price",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "fields",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "CreateEngineField"
                  }
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateEngineField",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "max",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "min",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "step",
              "type": {
                "name": "Float",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "type",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "EngineFieldType",
                  "kind": "ENUM",
                  "ofType": null
                }
              }
            },
            {
              "name": "info",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "label",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "options",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "CreateEngineFieldPicklistOption"
                  }
                }
              }
            },
            {
              "name": "defaultValue",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateEngineFieldPicklistOption",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "key",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "value",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateEngine",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "isPublic",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "categoryId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "deploymentModel",
              "type": {
                "name": "DeploymentModel",
                "kind": "ENUM",
                "ofType": null
              }
            },
            {
              "name": "price",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "state",
              "type": {
                "name": "EngineState",
                "kind": "ENUM",
                "ofType": null
              }
            },
            {
              "name": "fields",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "CreateEngineField"
                  }
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateBuild",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "engineId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateBuild",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "engineId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "action",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "BuildUpdateAction",
                  "kind": "ENUM",
                  "ofType": null
                }
              }
            },
            {
              "name": "dockerImage",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "BuildUpdateAction",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "deploy"
            },
            {
              "name": "pause"
            },
            {
              "name": "unpause"
            },
            {
              "name": "approve"
            },
            {
              "name": "disapprove"
            },
            {
              "name": "invalidate"
            },
            {
              "name": "submit"
            },
            {
              "name": "upload"
            }
          ]
        },
        {
          "name": "DeleteBuild",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "engineId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateTask",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "status",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "TaskStatus",
                  "kind": "ENUM",
                  "ofType": null
                }
              }
            },
            {
              "name": "jobId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "output",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "outputString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "outputJsonKey",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "taskOutput",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "PollTask",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "jobId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "pollPayload",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateJob",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "targetId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "tasks",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "CreateTask"
                  }
                }
              }
            },
            {
              "name": "retries",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateTask",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "taskType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "engineId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "payloadString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "payload",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "isClone",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "buildId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "testTask",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateApplication",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "key",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "category",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "iconUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "iconSvg",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "url",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "oauth2RedirectUrls",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "checkPermissions",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Boolean",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "permissionsRequired",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "deploymentModel",
              "type": {
                "name": "DeploymentModel",
                "kind": "ENUM",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateApplication",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "category",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "iconUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "iconSvg",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "url",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "oauth2RedirectUrls",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "String"
                  }
                }
              }
            },
            {
              "name": "checkPermissions",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "permissionsRequired",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "String"
                  }
                }
              }
            },
            {
              "name": "deploymentModel",
              "type": {
                "name": "DeploymentModel",
                "kind": "ENUM",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateOrganization",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "type",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "seatLimit",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "applications",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              }
            },
            {
              "name": "businessUnit",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "metadata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateEntityIdentifierType",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "label",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "labelPlural",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "iconClass",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "dataType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "EntityIdentifierDataType",
                  "kind": "ENUM",
                  "ofType": null
                }
              }
            },
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateEntityIdentifierType",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "label",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "labelPlural",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "iconClass",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "dataType",
              "type": {
                "name": "EntityIdentifierDataType",
                "kind": "ENUM",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateLibraryType",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "label",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "iconClass",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "entityIdentifierTypeIds",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              }
            },
            {
              "name": "entityType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "CreateEntityType",
                  "kind": "INPUT_OBJECT",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateEntityType",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "namePlural",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "schema",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "JSON",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateLibraryType",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "label",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "iconClass",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "entityIdentifierTypeIds",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              }
            },
            {
              "name": "entityType",
              "type": {
                "name": "UpdateEntityType",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateEntityType",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "namePlural",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "schema",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateLibrary",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "applicationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "organizationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "libraryTypeId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "coverImageUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateLibrary",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "coverImageUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "libraryTypeId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "version",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateEntity",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "libraryId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "profileImageUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsonstring",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "isPublished",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateEntity",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "profileImageUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateEntityIdentifier",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "entityId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "identifierTypeId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "title",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "isPriority",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "url",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsonstring",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "contentType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "file",
              "type": {
                "name": "UploadedFile",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "entityType",
              "type": {
                "name": "CreateEntityType",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateEntityIdentifier",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "title",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "isPriority",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "url",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateLibraryEngineModel",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "engineId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "libraryId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "trainJobId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "trainStatus",
              "type": {
                "name": "LibraryEngineModelTrainStatus",
                "kind": "ENUM",
                "ofType": null
              }
            },
            {
              "name": "dataUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateLibraryEngineModel",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "trainJobId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "trainStatus",
              "type": {
                "name": "LibraryEngineModelTrainStatus",
                "kind": "ENUM",
                "ofType": null
              }
            },
            {
              "name": "dataUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "contentType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "file",
              "type": {
                "name": "UploadedFile",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ApplicationWorkflow",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "action",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ApplicationWorkflowAction",
                  "kind": "ENUM",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ApplicationWorkflowAction",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "submit"
            },
            {
              "name": "approve"
            },
            {
              "name": "reject"
            },
            {
              "name": "deploy"
            },
            {
              "name": "disable"
            },
            {
              "name": "undelete"
            }
          ]
        },
        {
          "name": "CreateIngestionConfiguration",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "applicationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "type",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateIngestionConfiguration",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "type",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateWidget",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "widgetId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "collectionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "organizationId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "folderId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "adScript",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "width",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "numberOfMentionsToShow",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "displayLogo",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "displayCollectionName",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "displayMentionIntro",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "displayTranscription",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "seoTags",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "backgroundColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "borderColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "textColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateWidget",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "widgetId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "organizationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "collectionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "displayCollectionName",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "displayTranscription",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "width",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "numberOfMentionsToShow",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "adScript",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "seoTags",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "backgroundColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "borderColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "textColor",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "displayLogo",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "displayMentionIntro",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateUser",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "requestorId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "password",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "organizationId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "sendNewUserEmail",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "email",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "roleIds",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateUser",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "jsondata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "roleIds",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              }
            },
            {
              "name": "acls",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "UpdateUserACL"
                  }
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateUserACL",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "applicationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "organizationId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "objectType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "objectId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "access",
              "type": {
                "name": "UpdateUserACLAccessRights",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateUserACLAccessRights",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "owner",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreatePasswordUpdateRequest",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "skipPasswordResetEmail",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateSchema",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "schema",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "schemaString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "uiComponent",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "ttlSec",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateStructuredData",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "externalId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "schemaId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "data",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "dataString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateStructuredData",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "externalId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "schemaId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "data",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "dataString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateNielsenTVData",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "decorators",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "decoratorsString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "schemaId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "viewershipType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "uniqueCode",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "networkCode",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "startDateTime",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "endDateTime",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "fileName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "programName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "demographics",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "demographicsString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateNielsenTVData",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "decorators",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "decoratorsString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "schemaId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "viewershipType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "uniqueCode",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "networkCode",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "startDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "endDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "fileName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "programName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "demographics",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "demographicsString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateCollection",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "folderDescription",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "image",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateCollection",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "folderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "folderDescription",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "image",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ShareCollection",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "folderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "shareMessage",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "recipients",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "shareOptions",
              "type": {
                "name": "ShareOptions",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ShareOptions",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "showImage",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showComments",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showRating",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showHeader",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showOrganizationLogo",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "organizationLogoUrl",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showEngineResults",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showHits",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showAffiliateStripdown",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showDownload",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "showDescription",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Share",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "recipients",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "shareMessage",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "shareOptionsJson",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "folderId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mentionId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ShareMentionFromCollection",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "folderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "shareMessage",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "recipients",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "shareOptions",
              "type": {
                "name": "ShareOptions",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CollectionMentionInput",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "folderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CollectionMention",
          "fields": [
            {
              "name": "folderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateFolder",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "description",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "parentId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "rootFolderType",
              "type": {
                "name": "RootFolderType",
                "kind": "ENUM",
                "ofType": null
              }
            },
            {
              "name": "orderIndex",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateFolder",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MoveFolder",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "treeObjectId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "prevParentTreeObjectId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "newParentTreeObjectId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "newOrderIndex",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "prevOrderIndex",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "rootFolderType",
              "type": {
                "name": "RootFolderType",
                "kind": "ENUM",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "DeleteFolder",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "orderIndex",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateMentionComment",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "commentText",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateMentionComment",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "commentId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "commentText",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "DeleteMentionComment",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "commentId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateMentionRating",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "ratingValue",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MentionRating",
          "fields": [
            {
              "name": "ratingId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "mentionId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "userID",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "ratingValue",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "createdDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "modifiedDateTime",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateMentionRating",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "ratingId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "ratingValue",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Int",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "DeleteMentionRating",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "mentionId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "ratingId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UserLogin",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "userName",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "password",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "LoginInfo",
          "fields": [
            {
              "name": "apiToken",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "token",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "lastLoggedIn",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "applications",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "applicationPlatforms",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "ApplicationPlatform",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "groups",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Group",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "hasPassword",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "organization",
              "type": {
                "name": "Organization",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "passwordResetRequired",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "providerId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "providerScreenName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "providerUserId",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "user",
              "type": {
                "name": "User",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "ApplicationPlatform",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "platformType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "platformUrl",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "BulkUpdateWatchlistFilter",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "ids",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "BulkUpdateWatchlist",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "stopDate",
              "type": {
                "name": "DateTime",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "FileTemporalDataObject",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "tdoId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "folderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "orderIndex",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UnfileTemporalDataObject",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "tdoId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "folderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MoveTemporalDataObject",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "tdoId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "oldFolderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "newFolderId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UploadEngineResult",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "taskId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "outputString",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "file",
              "type": {
                "name": "UploadedFile",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "filename",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "assetType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "contentType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "completeTask",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "outputJsonKey",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "__Schema",
          "fields": [
            {
              "name": "types",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": null,
                  "kind": "LIST",
                  "ofType": {
                    "name": null
                  }
                }
              },
              "args": []
            },
            {
              "name": "queryType",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "__Type",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "mutationType",
              "type": {
                "name": "__Type",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "subscriptionType",
              "type": {
                "name": "__Type",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "directives",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": null,
                  "kind": "LIST",
                  "ofType": {
                    "name": null
                  }
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "__Type",
          "fields": [
            {
              "name": "kind",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "__TypeKind",
                  "kind": "ENUM",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "fields",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "__Field"
                  }
                }
              },
              "args": [
                {
                  "name": "includeDeprecated",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "interfaces",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "__Type"
                  }
                }
              },
              "args": []
            },
            {
              "name": "possibleTypes",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "__Type"
                  }
                }
              },
              "args": []
            },
            {
              "name": "enumValues",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "__EnumValue"
                  }
                }
              },
              "args": [
                {
                  "name": "includeDeprecated",
                  "type": {
                    "name": "Boolean",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "inputFields",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "__InputValue"
                  }
                }
              },
              "args": []
            },
            {
              "name": "ofType",
              "type": {
                "name": "__Type",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "__TypeKind",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "SCALAR"
            },
            {
              "name": "OBJECT"
            },
            {
              "name": "INTERFACE"
            },
            {
              "name": "UNION"
            },
            {
              "name": "ENUM"
            },
            {
              "name": "INPUT_OBJECT"
            },
            {
              "name": "LIST"
            },
            {
              "name": "NON_NULL"
            }
          ]
        },
        {
          "name": "__Field",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "args",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": null,
                  "kind": "LIST",
                  "ofType": {
                    "name": null
                  }
                }
              },
              "args": []
            },
            {
              "name": "type",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "__Type",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "isDeprecated",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Boolean",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "deprecationReason",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "__InputValue",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "type",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "__Type",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "defaultValue",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "__EnumValue",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "isDeprecated",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "Boolean",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "deprecationReason",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "__Directive",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "locations",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": null,
                  "kind": "LIST",
                  "ofType": {
                    "name": null
                  }
                }
              },
              "args": []
            },
            {
              "name": "args",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": null,
                  "kind": "LIST",
                  "ofType": {
                    "name": null
                  }
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "__DirectiveLocation",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "QUERY"
            },
            {
              "name": "MUTATION"
            },
            {
              "name": "SUBSCRIPTION"
            },
            {
              "name": "FIELD"
            },
            {
              "name": "FRAGMENT_DEFINITION"
            },
            {
              "name": "FRAGMENT_SPREAD"
            },
            {
              "name": "INLINE_FRAGMENT"
            },
            {
              "name": "SCHEMA"
            },
            {
              "name": "SCALAR"
            },
            {
              "name": "OBJECT"
            },
            {
              "name": "FIELD_DEFINITION"
            },
            {
              "name": "ARGUMENT_DEFINITION"
            },
            {
              "name": "INTERFACE"
            },
            {
              "name": "UNION"
            },
            {
              "name": "ENUM"
            },
            {
              "name": "ENUM_VALUE"
            },
            {
              "name": "INPUT_OBJECT"
            },
            {
              "name": "INPUT_FIELD_DEFINITION"
            }
          ]
        },
        {
          "name": "TokenType",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "API"
            },
            {
              "name": "User"
            }
          ]
        },
        {
          "name": "ScopeRequirement",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "Any"
            },
            {
              "name": "All"
            }
          ]
        },
        {
          "name": "AuthObjectType",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "TemporalDataObject"
            },
            {
              "name": "Job"
            },
            {
              "name": "Task"
            },
            {
              "name": "Folder"
            }
          ]
        },
        {
          "name": "KVP",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "value",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "Property",
                  "kind": "INTERFACE",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Metadata"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "IntProperty",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "value",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Property"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "StringProperty",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "value",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Property"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "BooleanProperty",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "value",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Property"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CompoundProperty",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "value",
              "type": {
                "name": "KVP",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Property"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CloneAssetIdMap",
          "fields": [
            {
              "name": "oldAssetId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "newAssetId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CloneData",
          "fields": [
            {
              "name": "date",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "originalId",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "cloneBlobs",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "assetIdMap",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "CloneAssetIdMap",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": [
                {
                  "name": "oldAssetId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                },
                {
                  "name": "newAssetId",
                  "type": {
                    "name": "ID",
                    "kind": "SCALAR",
                    "ofType": null
                  }
                }
              ]
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Metadata"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Program",
          "fields": [
            {
              "name": "id",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "image",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "liveImage",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Metadata"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "FileData",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "size",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "mimeType",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "fileName",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Metadata"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateTaskStatus",
          "fields": null,
          "kind": "ENUM",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": [
            {
              "name": "running"
            },
            {
              "name": "failed"
            },
            {
              "name": "complete"
            }
          ]
        },
        {
          "name": "UserACLInput",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "roleIds",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": null,
                  "kind": "LIST",
                  "ofType": {
                    "name": null
                  }
                }
              }
            },
            {
              "name": "acls",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "CreateUserACL"
                  }
                }
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateUserACL",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "access",
              "type": {
                "name": "UserACLAccessRightsInput",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UserACLAccessRightsInput",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "owner",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateUserInput",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "access",
              "type": {
                "name": "CreateUserACLAccessRights",
                "kind": "INPUT_OBJECT",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "CreateUserACLAccessRights",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "owner",
              "type": {
                "name": "Boolean",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateOrganizationInput",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "type",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "seatLimit",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "status",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "applications",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": null,
                  "kind": "NON_NULL",
                  "ofType": {
                    "name": "ID"
                  }
                }
              }
            },
            {
              "name": "businessUnit",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "metadata",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "SearchInput",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "offset",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "limit",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "index",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": null,
                  "kind": "LIST",
                  "ofType": {
                    "name": "String"
                  }
                }
              }
            },
            {
              "name": "query",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "JSON",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "select",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "StructuredJSONObject",
          "fields": [
            {
              "name": "data",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "schema",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "StructuredJSONSchema",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Metadata"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "JSONObject",
          "fields": [
            {
              "name": "name",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "String",
                  "kind": "SCALAR",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "data",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [
            {
              "name": "Metadata"
            }
          ],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "UpdateSchema",
          "fields": null,
          "kind": "INPUT_OBJECT",
          "inputFields": [
            {
              "name": "id",
              "type": {
                "name": null,
                "kind": "NON_NULL",
                "ofType": {
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
                }
              }
            },
            {
              "name": "name",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "schema",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "uiComponent",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "ttlSec",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "description",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              }
            }
          ],
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
        }
      ]
    }
  }
}
`
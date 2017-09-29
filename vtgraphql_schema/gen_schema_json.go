// GENERATED 2017-09-28 23:33:00.801199279 -0700 PDT for use with your graphQL generated types
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
              "name": "graphqlServiceInfo",
              "type": {
                "name": "GraphQLServiceInfo",
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
              "name": "CloneRequestList"
            },
            {
              "name": "EngineList"
            },
            {
              "name": "BuildList"
            },
            {
              "name": "JobList"
            },
            {
              "name": "EntityIdentifierTypeList"
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
              "name": "OrganizationList"
            },
            {
              "name": "GroupList"
            },
            {
              "name": "IngestionConfigurationList"
            },
            {
              "name": "MentionCommentList"
            },
            {
              "name": "MentionRatingList"
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
                  "name": "assetType",
                  "type": {
                    "name": "String",
                    "kind": "SCALAR",
                    "ofType": null
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
                  "name": "Int",
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
                  "name": "Int",
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
              "name": "uri",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
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
          "name": "Boolean",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
          "possibleTypes": null,
          "enumValues": null
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
              "name": "engineId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
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
              "name": "buildId",
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
              "name": "mediaLengthSec",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
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
              "name": "createsRecording",
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
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
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
                  "name": "ID",
                  "kind": "SCALAR",
                  "ofType": null
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
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
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
                  "name": "engineModelId",
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
              "name": "libraryId",
              "type": {
                "name": "ID",
                "kind": "SCALAR",
                "ofType": null
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
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
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
              "args": []
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
                  "name": "Role",
                  "kind": "OBJECT",
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
                  "name": "Folder",
                  "kind": "OBJECT",
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
          "name": "DateTime",
          "fields": null,
          "kind": "SCALAR",
          "inputFields": null,
          "interfaces": null,
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
                  "name": "asset",
                  "type": {
                    "name": "UpdateAsset",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
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
                    "name": "UpdateBuild",
                    "kind": "INPUT_OBJECT",
                    "ofType": null
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
              "name": "deleteEntityIdentifierType",
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
              "name": "deleteLibraryType",
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
                  "name": "Int",
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
                  "name": "Int",
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
              "name": "mediaId",
              "type": {
                "name": "ID",
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
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              }
            },
            {
              "name": "stopDateTime",
              "type": {
                "name": "Int",
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
              "name": "mediaId",
              "type": {
                "name": "ID",
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
              "name": "md5Checksum",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
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
                "name": "Int",
                "kind": "SCALAR",
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
                "name": "Int",
                "kind": "SCALAR",
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
            }
          ]
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
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
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
              "name": "taskOutput",
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
                  "name": "CreateTask",
                  "kind": "INPUT_OBJECT",
                  "ofType": null
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
          "name": "MentionList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": "JSON",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "searchToken",
              "type": {
                "name": "String",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "from",
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
              "name": "to",
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
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "Mention",
          "fields": [
            {
              "name": "dataObject",
              "type": {
                "name": "TemporalDataObject",
                "kind": "OBJECT",
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
            },
            {
              "name": "context",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "MentionContext",
                  "kind": "OBJECT",
                  "ofType": null
                }
              },
              "args": []
            },
            {
              "name": "comments",
              "type": {
                "name": "MentionCommentList",
                "kind": "OBJECT",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "ratings",
              "type": {
                "name": "MentionRatingList",
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
          "name": "MentionHit",
          "fields": [
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
              "name": "engineResults",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "EngineResult",
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
          "name": "EngineResult",
          "fields": [
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
              "name": "startTime",
              "type": {
                "name": "Int",
                "kind": "SCALAR",
                "ofType": null
              },
              "args": []
            },
            {
              "name": "stopTime",
              "type": {
                "name": "Int",
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
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MentionContext",
          "fields": [
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
              "name": "engineResults",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "EngineResult",
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
          "name": "MentionCommentList",
          "fields": [
            {
              "name": "records",
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
          "name": "MentionComment",
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
              "name": "text",
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
            }
          ],
          "kind": "OBJECT",
          "inputFields": null,
          "interfaces": [],
          "possibleTypes": null,
          "enumValues": null
        },
        {
          "name": "MentionRatingList",
          "fields": [
            {
              "name": "records",
              "type": {
                "name": null,
                "kind": "LIST",
                "ofType": {
                  "name": "MentionRating",
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
          "name": "MentionRating",
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
              "name": "text",
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
        }
      ]
    }
  }
}
`
// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "microservice for the bonus system of self-service car washes",
    "title": "wash-bonus",
    "version": "1.0.0"
  },
  "paths": {
    "/addTestData": {
      "post": {
        "tags": [
          "Standard"
        ],
        "operationId": "addTestData",
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/healthCheck": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "Standard"
        ],
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "tags": [
          "User"
        ],
        "operationId": "getCurrentUser",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/add": {
      "post": {
        "tags": [
          "User"
        ],
        "operationId": "addUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userAdd"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/list": {
      "post": {
        "tags": [
          "User"
        ],
        "operationId": "listUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/listParams"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "items": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/user"
                  }
                },
                "warnings": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "User"
        ],
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}/delete": {
      "delete": {
        "tags": [
          "User"
        ],
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}/edit": {
      "put": {
        "tags": [
          "User"
        ],
        "operationId": "editUser",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userUpdate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/washServer": {
      "post": {
        "tags": [
          "WashServer"
        ],
        "operationId": "addWashServer",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/washServerAdd"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/washServer/{id}": {
      "get": {
        "tags": [
          "WashServer"
        ],
        "operationId": "getWashServer",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/washServer"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "WashServer"
        ],
        "operationId": "editWashServer",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/washServerUpdate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "WashServer"
        ],
        "operationId": "deleteWashServer",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/washServer/{id}/generate-key": {
      "put": {
        "tags": [
          "WashServer"
        ],
        "operationId": "generateKeyWashServer",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/washServerGenerateKeyResp"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/washServers": {
      "post": {
        "tags": [
          "WashServer"
        ],
        "operationId": "listWashServer",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/listParams"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/listWashServer"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "filter": {
      "type": "object",
      "properties": {
        "ignoreCase": {
          "type": "boolean"
        },
        "operator": {
          "description": "enum ==, !=, \u003c, \u003e, \u003c=, \u003e=, in, !in",
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "filterGroup": {
      "description": "params for grouping filters",
      "type": "object",
      "properties": {
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/filter"
          }
        },
        "key": {
          "type": "string"
        },
        "logicFilter": {
          "description": "Comparing option: true == 'AND', false == 'OR'",
          "type": "boolean"
        }
      }
    },
    "listParams": {
      "description": "params for list method",
      "type": "object",
      "properties": {
        "filterGroups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/filterGroup"
          }
        },
        "limit": {
          "type": "integer",
          "minimum": 1
        },
        "offset": {
          "type": "integer"
        },
        "orderBy": {
          "type": "string",
          "enum": [
            "ASC",
            "DESC"
          ]
        },
        "sortBy": {
          "type": "string"
        }
      }
    },
    "listWashServer": {
      "description": "washServer list object",
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/washServer"
          }
        },
        "warnings": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "user": {
      "description": "user object",
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "firebaseId": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "modifiedAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        }
      }
    },
    "userAdd": {
      "description": "user model for add method",
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        }
      }
    },
    "userUpdate": {
      "description": "user model for edit method",
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        }
      }
    },
    "washServer": {
      "description": "washServer object",
      "type": "object",
      "properties": {
        "createdAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "modifiedAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "name": {
          "type": "string"
        },
        "owner_id": {
          "type": "string"
        },
        "service_key": {
          "type": "string"
        }
      }
    },
    "washServerAdd": {
      "description": "washServer model for add methods",
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "owner_id": {
          "type": "string"
        }
      }
    },
    "washServerGenerateKeyResp": {
      "description": "washServerGenerateKeyResp object",
      "type": "object",
      "properties": {
        "service_key": {
          "type": "string"
        }
      }
    },
    "washServerUpdate": {
      "description": "washServer model for edit methods",
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "owner_id": {
          "type": "string"
        },
        "service_key": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "authKey": {
      "description": "Session token inside Authorization header.",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "authKey": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "microservice for the bonus system of self-service car washes",
    "title": "wash-bonus",
    "version": "1.0.0"
  },
  "paths": {
    "/addTestData": {
      "post": {
        "tags": [
          "Standard"
        ],
        "operationId": "addTestData",
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/healthCheck": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "Standard"
        ],
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "tags": [
          "User"
        ],
        "operationId": "getCurrentUser",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/add": {
      "post": {
        "tags": [
          "User"
        ],
        "operationId": "addUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userAdd"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/list": {
      "post": {
        "tags": [
          "User"
        ],
        "operationId": "listUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/listParams"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "items": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/user"
                  }
                },
                "warnings": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  }
                }
              }
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "User"
        ],
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}/delete": {
      "delete": {
        "tags": [
          "User"
        ],
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}/edit": {
      "put": {
        "tags": [
          "User"
        ],
        "operationId": "editUser",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userUpdate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/washServer": {
      "post": {
        "tags": [
          "WashServer"
        ],
        "operationId": "addWashServer",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/washServerAdd"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Created"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/washServer/{id}": {
      "get": {
        "tags": [
          "WashServer"
        ],
        "operationId": "getWashServer",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/washServer"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "WashServer"
        ],
        "operationId": "editWashServer",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/washServerUpdate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "WashServer"
        ],
        "operationId": "deleteWashServer",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/washServer/{id}/generate-key": {
      "put": {
        "tags": [
          "WashServer"
        ],
        "operationId": "generateKeyWashServer",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/washServerGenerateKeyResp"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/washServers": {
      "post": {
        "tags": [
          "WashServer"
        ],
        "operationId": "listWashServer",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/listParams"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/listWashServer"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "filter": {
      "type": "object",
      "properties": {
        "ignoreCase": {
          "type": "boolean"
        },
        "operator": {
          "description": "enum ==, !=, \u003c, \u003e, \u003c=, \u003e=, in, !in",
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "filterGroup": {
      "description": "params for grouping filters",
      "type": "object",
      "properties": {
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/filter"
          }
        },
        "key": {
          "type": "string"
        },
        "logicFilter": {
          "description": "Comparing option: true == 'AND', false == 'OR'",
          "type": "boolean"
        }
      }
    },
    "listParams": {
      "description": "params for list method",
      "type": "object",
      "properties": {
        "filterGroups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/filterGroup"
          }
        },
        "limit": {
          "type": "integer",
          "minimum": 1
        },
        "offset": {
          "type": "integer",
          "minimum": 0
        },
        "orderBy": {
          "type": "string",
          "enum": [
            "ASC",
            "DESC"
          ]
        },
        "sortBy": {
          "type": "string"
        }
      }
    },
    "listWashServer": {
      "description": "washServer list object",
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/washServer"
          }
        },
        "warnings": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "user": {
      "description": "user object",
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "firebaseId": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "modifiedAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        }
      }
    },
    "userAdd": {
      "description": "user model for add method",
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        }
      }
    },
    "userUpdate": {
      "description": "user model for edit method",
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        }
      }
    },
    "washServer": {
      "description": "washServer object",
      "type": "object",
      "properties": {
        "createdAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "modifiedAt": {
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "name": {
          "type": "string"
        },
        "owner_id": {
          "type": "string"
        },
        "service_key": {
          "type": "string"
        }
      }
    },
    "washServerAdd": {
      "description": "washServer model for add methods",
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "owner_id": {
          "type": "string"
        }
      }
    },
    "washServerGenerateKeyResp": {
      "description": "washServerGenerateKeyResp object",
      "type": "object",
      "properties": {
        "service_key": {
          "type": "string"
        }
      }
    },
    "washServerUpdate": {
      "description": "washServer model for edit methods",
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "owner_id": {
          "type": "string"
        },
        "service_key": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "authKey": {
      "description": "Session token inside Authorization header.",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "authKey": []
    }
  ]
}`))
}

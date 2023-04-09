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
    "/profile": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "user"
        ],
        "operationId": "getProfile",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/profile"
            }
          },
          "404": {
            "description": "Profile not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/profile/balance": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "user"
        ],
        "operationId": "getBalance",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "balance": {
                  "type": "integer"
                }
              }
            }
          },
          "404": {
            "description": "Profile not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/session/{UID}": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "session"
        ],
        "operationId": "getSession",
        "parameters": [
          {
            "type": "string",
            "name": "UID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/session"
            }
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Profile not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/session/{UID}/bonuses": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "session"
        ],
        "operationId": "postSession",
        "parameters": [
          {
            "type": "string",
            "name": "UID",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/bonusCharge"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/bonusCharge"
            }
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Session not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/session/{sessionId}/assign-user": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "session"
        ],
        "operationId": "assignUserToSession",
        "parameters": [
          {
            "type": "string",
            "name": "sessionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Session not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "bonusCharge": {
      "description": "bonus amount for use in session",
      "type": "object",
      "properties": {
        "amount": {
          "type": "integer"
        }
      }
    },
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
    "profile": {
      "description": "user profile",
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "balance": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "session": {
      "description": "session",
      "type": "object",
      "properties": {
        "postBalance": {
          "type": "integer"
        },
        "postID": {
          "type": "integer"
        },
        "washServer": {
          "$ref": "#/definitions/washServer"
        }
      }
    },
    "washServer": {
      "description": "wash server info",
      "type": "object",
      "properties": {
        "Description": {
          "type": "string"
        },
        "Name": {
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
    "/profile": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "user"
        ],
        "operationId": "getProfile",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/profile"
            }
          },
          "404": {
            "description": "Profile not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/profile/balance": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "user"
        ],
        "operationId": "getBalance",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "balance": {
                  "type": "integer"
                }
              }
            }
          },
          "404": {
            "description": "Profile not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/session/{UID}": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "session"
        ],
        "operationId": "getSession",
        "parameters": [
          {
            "type": "string",
            "name": "UID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/session"
            }
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Profile not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/session/{UID}/bonuses": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "session"
        ],
        "operationId": "postSession",
        "parameters": [
          {
            "type": "string",
            "name": "UID",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/bonusCharge"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/bonusCharge"
            }
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Session not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/session/{sessionId}/assign-user": {
      "post": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "session"
        ],
        "operationId": "assignUserToSession",
        "parameters": [
          {
            "type": "string",
            "name": "sessionId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Session not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "bonusCharge": {
      "description": "bonus amount for use in session",
      "type": "object",
      "properties": {
        "amount": {
          "type": "integer"
        }
      }
    },
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
    "profile": {
      "description": "user profile",
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "balance": {
          "type": "integer"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "session": {
      "description": "session",
      "type": "object",
      "properties": {
        "postBalance": {
          "type": "integer"
        },
        "postID": {
          "type": "integer"
        },
        "washServer": {
          "$ref": "#/definitions/washServer"
        }
      }
    },
    "washServer": {
      "description": "wash server info",
      "type": "object",
      "properties": {
        "Description": {
          "type": "string"
        },
        "Name": {
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

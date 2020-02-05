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
    "description": "lakeFS HTTP API",
    "title": "lakeFS API",
    "version": "0.1.0"
  },
  "paths": {
    "/repositories": {
      "get": {
        "summary": "list repositories",
        "operationId": "listRepositories",
        "responses": {
          "200": {
            "description": "repository list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/repository"
              }
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{repositoryId}": {
      "get": {
        "summary": "get repository",
        "operationId": "getRepository",
        "responses": {
          "200": {
            "description": "repository",
            "schema": {
              "$ref": "#/definitions/repository"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "description": "repository not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "summary": "create repository",
        "operationId": "createRepository",
        "parameters": [
          {
            "name": "repository",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/repository_creation"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "repository",
            "schema": {
              "$ref": "#/definitions/repository"
            }
          },
          "400": {
            "description": "validation error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "summary": "delete repository",
        "operationId": "deleteRepository",
        "responses": {
          "204": {
            "description": "repository deleted successfully"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "description": "repository not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/repositories/{repositoryId}/branches": {
      "get": {
        "summary": "list branches",
        "operationId": "listBranches",
        "responses": {
          "200": {
            "description": "branch list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/refspec"
              }
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/repositories/{repositoryId}/branches/{branchId}": {
      "get": {
        "summary": "get branch",
        "operationId": "getBranch",
        "responses": {
          "200": {
            "description": "branch",
            "schema": {
              "$ref": "#/definitions/refspec"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "description": "branch not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "summary": "create branch",
        "operationId": "createBranch",
        "parameters": [
          {
            "name": "branch",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/refspec"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "branch",
            "schema": {
              "$ref": "#/definitions/refspec"
            }
          },
          "400": {
            "description": "validation error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "summary": "delete branch",
        "operationId": "deleteBranch",
        "responses": {
          "204": {
            "description": "branch deleted successfully"
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "description": "branch not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "name": "branchId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/repositories/{repositoryId}/branches/{branchId}/commits": {
      "post": {
        "summary": "create commit",
        "operationId": "commit",
        "parameters": [
          {
            "name": "commit",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/commit_creation"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "commit",
            "schema": {
              "$ref": "#/definitions/commit"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "description": "branch not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "name": "branchId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/repositories/{repositoryId}/commits/{commitId}": {
      "get": {
        "summary": "get commit",
        "operationId": "getCommit",
        "responses": {
          "200": {
            "description": "commit",
            "schema": {
              "$ref": "#/definitions/commit"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "description": "commit not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "name": "commitId",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "commit": {
      "type": "object",
      "properties": {
        "committer": {
          "type": "string"
        },
        "creation_date": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "parents": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "commit_creation": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        },
        "metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "description": "short message explaining the error",
          "type": "string"
        }
      }
    },
    "object": {
      "type": "object",
      "required": [
        "path",
        "type"
      ],
      "properties": {
        "path": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "refspec": {
      "type": "object",
      "required": [
        "id",
        "commit_id"
      ],
      "properties": {
        "commit_id": {
          "type": "string"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "repository": {
      "type": "object",
      "properties": {
        "bucket_name": {
          "type": "string"
        },
        "creation_date": {
          "type": "integer",
          "format": "int64"
        },
        "default_branch": {
          "type": "string",
          "example": "master"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "repository_creation": {
      "type": "object",
      "required": [
        "id",
        "bucket_name"
      ],
      "properties": {
        "bucket_name": {
          "type": "string"
        },
        "default_branch": {
          "type": "string",
          "example": "master"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "basic_auth": {
      "type": "basic"
    }
  },
  "security": [
    {
      "basic_auth": []
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
    "description": "lakeFS HTTP API",
    "title": "lakeFS API",
    "version": "0.1.0"
  },
  "paths": {
    "/repositories": {
      "get": {
        "summary": "list repositories",
        "operationId": "listRepositories",
        "responses": {
          "200": {
            "description": "repository list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/repository"
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{repositoryId}": {
      "get": {
        "summary": "get repository",
        "operationId": "getRepository",
        "responses": {
          "200": {
            "description": "repository",
            "schema": {
              "$ref": "#/definitions/repository"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "repository not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "summary": "create repository",
        "operationId": "createRepository",
        "parameters": [
          {
            "name": "repository",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/repository_creation"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "repository",
            "schema": {
              "$ref": "#/definitions/repository"
            }
          },
          "400": {
            "description": "validation error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "summary": "delete repository",
        "operationId": "deleteRepository",
        "responses": {
          "204": {
            "description": "repository deleted successfully"
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "repository not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/repositories/{repositoryId}/branches": {
      "get": {
        "summary": "list branches",
        "operationId": "listBranches",
        "responses": {
          "200": {
            "description": "branch list",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/refspec"
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/repositories/{repositoryId}/branches/{branchId}": {
      "get": {
        "summary": "get branch",
        "operationId": "getBranch",
        "responses": {
          "200": {
            "description": "branch",
            "schema": {
              "$ref": "#/definitions/refspec"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "branch not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "summary": "create branch",
        "operationId": "createBranch",
        "parameters": [
          {
            "name": "branch",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/refspec"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "branch",
            "schema": {
              "$ref": "#/definitions/refspec"
            }
          },
          "400": {
            "description": "validation error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "summary": "delete branch",
        "operationId": "deleteBranch",
        "responses": {
          "204": {
            "description": "branch deleted successfully"
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "branch not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "name": "branchId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/repositories/{repositoryId}/branches/{branchId}/commits": {
      "post": {
        "summary": "create commit",
        "operationId": "commit",
        "parameters": [
          {
            "name": "commit",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/commit_creation"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "commit",
            "schema": {
              "$ref": "#/definitions/commit"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "branch not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "name": "branchId",
          "in": "path",
          "required": true
        }
      ]
    },
    "/repositories/{repositoryId}/commits/{commitId}": {
      "get": {
        "summary": "get commit",
        "operationId": "getCommit",
        "responses": {
          "200": {
            "description": "commit",
            "schema": {
              "$ref": "#/definitions/commit"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "commit not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "repositoryId",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "name": "commitId",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "commit": {
      "type": "object",
      "properties": {
        "committer": {
          "type": "string"
        },
        "creation_date": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        },
        "parents": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "commit_creation": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        },
        "metadata": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "error": {
      "type": "object",
      "properties": {
        "message": {
          "description": "short message explaining the error",
          "type": "string"
        }
      }
    },
    "object": {
      "type": "object",
      "required": [
        "path",
        "type"
      ],
      "properties": {
        "path": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "refspec": {
      "type": "object",
      "required": [
        "id",
        "commit_id"
      ],
      "properties": {
        "commit_id": {
          "type": "string"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "repository": {
      "type": "object",
      "properties": {
        "bucket_name": {
          "type": "string"
        },
        "creation_date": {
          "type": "integer",
          "format": "int64"
        },
        "default_branch": {
          "type": "string",
          "example": "master"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "repository_creation": {
      "type": "object",
      "required": [
        "id",
        "bucket_name"
      ],
      "properties": {
        "bucket_name": {
          "type": "string"
        },
        "default_branch": {
          "type": "string",
          "example": "master"
        },
        "id": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "basic_auth": {
      "type": "basic"
    }
  },
  "security": [
    {
      "basic_auth": []
    }
  ]
}`))
}

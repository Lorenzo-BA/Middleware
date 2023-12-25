// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Justine Bachelard.",
            "email": "justine.bachelard@ext.uca.fr"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Songs": {
            "get": {
                "description": "Get Songs.",
                "tags": [
                    "Songs"
                ],
                "summary": "Get Songs.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Song"
                            }
                        }
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            }
        },
        "/Songs/{id}": {
            "get": {
                "description": "Get a collection.",
                "tags": [
                    "Songs"
                ],
                "summary": "Get a collection.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Song UUID formatted ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "422": {
                        "description": "Cannot parse id"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Song": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "middleware/example",
	Description:      "API to manage Songs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

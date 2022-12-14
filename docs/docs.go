// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/download/{format}": {
            "get": {
                "description": "Download in json or csv format with filters",
                "tags": [
                    "download"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "download format: json or csv",
                        "name": "format",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "transaction id",
                        "name": "transaction_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "terminal id: n or 1, 2, 3, ..., n",
                        "name": "terminal_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "status: accepted or declined",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "payment type: cash or card",
                        "name": "payment_type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "date post: from yyyy-mm-dd, to yyyy-mm-dd",
                        "name": "date_post",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "payment narrative",
                        "name": "payment_narrative",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/api/upload": {
            "post": {
                "description": "upload csv file, parsing it and saving the parsing results to the database",
                "tags": [
                    "upload"
                ],
                "parameters": [
                    {
                        "type": "file",
                        "description": "csv file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "EVO Fintech",
	Description:      "REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

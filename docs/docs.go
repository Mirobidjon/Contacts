// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/contact": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get All Contacts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contact"
                ],
                "summary": "get all contacts",
                "operationId": "get-all-contacts",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.allContact"
                            }
                        }
                    },
                    "400": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create Contact",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contact"
                ],
                "summary": "create contact",
                "operationId": "create-contact",
                "parameters": [
                    {
                        "description": "profession",
                        "name": "profession",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contact.DefaultContact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/contact.DefaultContact"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    }
                }
            }
        },
        "/api/contact/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get Contact by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contact"
                ],
                "summary": "get contact by id",
                "operationId": "get-contact",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/contact.DefaultContact"
                        }
                    },
                    "400": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Contact",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contact"
                ],
                "summary": "update contact",
                "operationId": "update-contact",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contact.DefaultContact"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/contact.ID"
                        }
                    },
                    "400": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete Contact",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "contact"
                ],
                "summary": "delete contact",
                "operationId": "delete-contact",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/contact.ID"
                        }
                    },
                    "400": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "this is login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "profession",
                        "name": "profession",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.signinInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/contact.Token"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "this is sign up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign up",
                "operationId": "sign-up",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contact.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/contact.ID"
                        }
                    },
                    "500": {
                        "description": "desc",
                        "schema": {
                            "$ref": "#/definitions/handler.errorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "contact.DefaultContact": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "contact.ID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "contact.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "contact.User": {
            "type": "object",
            "required": [
                "name",
                "password",
                "username"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handler.allContact": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "handler.errorMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.signinInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

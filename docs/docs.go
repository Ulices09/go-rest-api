// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/auth/login": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "data",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errors.CustomError"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errors.CustomError"
                        }
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "security": [
                    {
                        "auth-token": []
                    }
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get current user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errors.CustomError"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "tags": [
                    "posts"
                ],
                "summary": "Get posts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "filter attributes",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page size",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.PaginationResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entity.Post"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errors.CustomError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "auth-token": []
                    }
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Crate post",
                "parameters": [
                    {
                        "description": "post",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/posts.CreatePostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Post"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errors.CustomError"
                        }
                    }
                }
            }
        },
        "/posts/{id}": {
            "get": {
                "tags": [
                    "posts"
                ],
                "summary": "Get post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "post id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Post"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errors.CustomError"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "get all users",
                "tags": [
                    "users"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "filters users by email",
                        "name": "filter",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.ListResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entity.User"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errors.CustomError"
                        }
                    }
                }
            },
            "post": {
                "description": "create new user",
                "tags": [
                    "users"
                ],
                "summary": "Crate user",
                "parameters": [
                    {
                        "description": "user",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/errors.CustomError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.ListResult": {
            "type": "object",
            "properties": {
                "data": {}
            }
        },
        "dto.PaginationResult": {
            "type": "object",
            "properties": {
                "data": {},
                "pageSize": {
                    "type": "integer"
                },
                "pages": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "entity.Post": {
            "type": "object",
            "properties": {
                "createddAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/entity.User"
                }
            }
        },
        "entity.Role": {
            "type": "object",
            "properties": {
                "createddAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "createddAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "posts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Post"
                    }
                },
                "role": {
                    "$ref": "#/definitions/entity.Role"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "errors.CustomError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "posts.CreatePostRequest": {
            "type": "object",
            "required": [
                "text",
                "title"
            ],
            "properties": {
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "users.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "roleId"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "roleId": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "auth-token": {
            "type": "apiKey",
            "name": "session-token",
            "in": "cookie"
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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

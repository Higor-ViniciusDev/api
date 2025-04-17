// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "higor souza",
            "url": "https://www.linkedin.com/in/higor-vinicius-de-souza-a44b8416a/",
            "email": "higor.vinicius331@hotmail.com"
        },
        "license": {
            "name": "Full Cycle",
            "url": "http://www.fullcycle.com.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/produtos": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Buscar Todos os produtos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produto"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "pagina",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limite result",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Produto"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handles.Error"
                        }
                    }
                }
            }
        },
        "/produtos/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Cria novo usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produto"
                ],
                "parameters": [
                    {
                        "description": "produto request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateProdutoInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handles.Error"
                        }
                    }
                }
            }
        },
        "/produtos/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Buscar um produto especifico por ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produto"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "produto ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Produto"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handles.Error"
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
                "description": "Altera Produto Existente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produto"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "produto request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateProdutoInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handles.Error"
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
                "description": "Delata um produto",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "produto"
                ],
                "summary": "Delata um produto",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "produto ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handles.Error"
                        }
                    }
                }
            }
        },
        "/usuario/create": {
            "post": {
                "description": "Cria novo usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuario"
                ],
                "parameters": [
                    {
                        "description": "usuario request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUsuarioInpunt"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handles.Error"
                        }
                    }
                }
            }
        },
        "/usuario/generateTolken": {
            "post": {
                "description": "Gerar Novo Tolken de acesso",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuario"
                ],
                "parameters": [
                    {
                        "description": "usuario credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GetJWT"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetJWTOutput"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handles.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handles.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateProdutoInput": {
            "type": "object",
            "properties": {
                "nome": {
                    "type": "string"
                },
                "preco": {
                    "type": "number"
                }
            }
        },
        "dto.CreateUsuarioInpunt": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "senha": {
                    "type": "string"
                }
            }
        },
        "dto.GetJWT": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "senha": {
                    "type": "string"
                }
            }
        },
        "dto.GetJWTOutput": {
            "type": "object",
            "properties": {
                "access_tolken": {
                    "type": "string"
                }
            }
        },
        "entity.Produto": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "preco": {
                    "type": "number"
                }
            }
        },
        "handles.Error": {
            "type": "object",
            "properties": {
                "menssagem": {
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Expert API Com JWT",
	Description:      "Requisição de produto com API e autenticação.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

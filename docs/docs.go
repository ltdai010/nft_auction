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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/collections": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Post Collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Post Collection",
                "operationId": "post-collection",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CollectionReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Collection"
                        }
                    }
                }
            }
        },
        "/collections/query/list": {
            "get": {
                "description": "Get Collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Get Collection",
                "operationId": "query-collection",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "creator_id",
                        "name": "creator_id",
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
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.QueryCollectionRes"
                        }
                    }
                }
            }
        },
        "/collections/{id}": {
            "get": {
                "description": "Get Collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Get Collection",
                "operationId": "get-collection",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Collection"
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
                "description": "Put Collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Put Collection",
                "operationId": "put-collection",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CollectionReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Collection"
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
                "description": "Delete Collection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Collections"
                ],
                "summary": "Delete Collection",
                "operationId": "delete-collection",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/items": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Post Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Post Item",
                "operationId": "post-item",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ItemReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Item"
                        }
                    }
                }
            }
        },
        "/items/action/like/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Like Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Like Item",
                "operationId": "put-item",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/items/query/list": {
            "get": {
                "description": "Get Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Get Item",
                "operationId": "query-items",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "collection_id",
                        "name": "collection_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "liked by",
                        "name": "liked_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "owner_id",
                        "name": "owner_id",
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
                        "description": "page_size",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.QueryItemRes"
                        }
                    }
                }
            }
        },
        "/items/{id}": {
            "get": {
                "description": "Get Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Get Item",
                "operationId": "get-item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Item"
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
                "description": "Put Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Put Item",
                "operationId": "put-item",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ItemReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Item"
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
                "description": "Delete Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Delete Item",
                "operationId": "delete-item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/migrate": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "migrate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Migration"
                ],
                "summary": "migrate",
                "operationId": "migrate",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sales": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Post Sales",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sales"
                ],
                "summary": "Post Sales",
                "operationId": "post-sales",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SaleReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Sale"
                        }
                    }
                }
            }
        },
        "/sales/query/list": {
            "get": {
                "description": "Query Sale",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sales"
                ],
                "summary": "Query Sale",
                "operationId": "query-sale",
                "parameters": [
                    {
                        "type": "string",
                        "description": "creator_id",
                        "name": "item_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.QuerySaleRes"
                        }
                    }
                }
            }
        },
        "/sales/{id}": {
            "get": {
                "description": "Get Sale",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sales"
                ],
                "summary": "Get Sale",
                "operationId": "get-sale",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Sale"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login user",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/profile/{id}": {
            "get": {
                "description": "Get profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get profile",
                "operationId": "get-profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Collection": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Item"
                    }
                },
                "metadata": {
                    "type": "object"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updater_id": {
                    "type": "string"
                }
            }
        },
        "models.CollectionReq": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Item": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "collection": {
                    "$ref": "#/definitions/models.Collection"
                },
                "collection_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "item_id": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updater_id": {
                    "type": "string"
                }
            }
        },
        "models.ItemReq": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "collection_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "item_id": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "models.QueryCollectionRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Collection"
                    }
                },
                "metadata": {}
            }
        },
        "models.QueryItemRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Collection"
                    }
                },
                "metadata": {}
            }
        },
        "models.QuerySaleRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Sale"
                    }
                }
            }
        },
        "models.Sale": {
            "type": "object",
            "properties": {
                "buyer_id": {
                    "type": "string"
                },
                "coin_buy": {
                    "type": "string"
                },
                "coin_buy_address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "decimal": {
                    "type": "integer"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "item": {
                    "$ref": "#/definitions/models.Item"
                },
                "item_id": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updater_id": {
                    "type": "string"
                }
            }
        },
        "models.SaleReq": {
            "type": "object",
            "properties": {
                "coin_buy": {
                    "type": "string"
                },
                "coin_buy_address": {
                    "type": "string"
                },
                "decimal": {
                    "type": "integer"
                },
                "item_id": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "creator_id": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_login": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object"
                },
                "name": {
                    "type": "string"
                },
                "pubkey": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updater_id": {
                    "type": "string"
                }
            }
        },
        "models.UserLoginRequest": {
            "type": "object",
            "properties": {
                "signature": {
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
	Host:             "api.uchain.duckdns.org",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Finan Loyalty API",
	Description:      "This is Finan Loyalty server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/actors": {
            "get": {
                "description": "Get actors with movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actors"
                ],
                "responses": {
                    "200": {
                        "description": "Response body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Actor"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Update actors",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actors"
                ],
                "parameters": [
                    {
                        "description": "movie",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Actor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Get movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Actors"
                ],
                "parameters": [
                    {
                        "description": "movie",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Actor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "description": "Get movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Get movies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "field to sort by",
                        "name": "sortBy",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ASC DESC",
                        "name": "sortType",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Movie"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "description": "Update movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "parameters": [
                    {
                        "description": "movie",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "parameters": [
                    {
                        "description": "movie",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "description": "delete movie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/movies/search-by-actor": {
            "get": {
                "description": "Get movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Get movies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "part of actor's name",
                        "name": "partOfActor",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Movie"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/movies/search-by-title": {
            "get": {
                "description": "Get movies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Get movies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "part of movie title",
                        "name": "partOfTitle",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Response body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Movie"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Actor": {
            "type": "object",
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "movies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Movie"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.AuthResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "model.Movie": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "release_year": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FilmLibrary API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

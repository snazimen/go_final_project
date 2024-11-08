package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "snazimen@yandex.ru",
        "contact": {
            "name": "Danila sushkin",
            "email": "snazimen@yandex.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/signin": {
            "post": {
                "description": "Получение токена по паролю",
                "consumes": [
                    "application/json"
                ],
                "summary": "Получение токена по паролю",
                "parameters": [
                    {
                        "description": "Пароль профиля",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/middleware.bodyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/middleware.getAuthByPassword"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/middleware.errResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/middleware.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/middleware.errResponse"
                        }
                    }
                }
            }
        },
        "/api/task": {
            "get": {
                "description": "Получить задачу",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Получить задачу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор задачи",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TaskResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Редактировать задачу",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Редактировать задачу",
                "parameters": [
                    {
                        "description": "Параметры задачи",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавить новую задачу",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Добавить новую задачу",
                "parameters": [
                    {
                        "description": "Параметры задачи",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Task"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.TaskResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удалить задачу",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Удалить задачу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор задачи",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TaskResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    }
                }
            }
        },
        "/api/task/done": {
            "post": {
                "description": "Выполнить задачу",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Выполнить задачу",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор задачи",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TaskResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    }
                }
            }
        },
        "/api/tasks": {
            "get": {
                "description": "Получить список ближайших задач",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Получить список ближайших задач",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Строка поиска",
                        "name": "search",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TasksResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.errResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "middleware.bodyRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "middleware.errResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "middleware.getAuthByPassword": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "repeat": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.TaskResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.TasksResp": {
            "type": "object",
            "properties": {
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Task"
                    }
                }
            }
        }
    }
}`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:7540",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Пользовательская документация API",
	Description:      "Итоговая работа по курсу \"Go-разработчик с нуля\" (Яндекс Практикум)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

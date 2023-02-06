// Code generated by swaggo/swag. DO NOT EDIT
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
        "/categories": {
            "get": {
                "tags": [
                    "Category"
                ],
                "summary": "Get all categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/modules_storage_internal_handler.dtoCategory"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/category": {
            "post": {
                "tags": [
                    "Category"
                ],
                "summary": "Create new category",
                "parameters": [
                    {
                        "description": "Category",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modules_storage_internal_handler.dtoCreateCategory"
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
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/inventory": {
            "get": {
                "tags": [
                    "Inventory"
                ],
                "summary": "Get all inventory products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/modules_storage_internal_handler.dtoInventory"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Inventory"
                ],
                "summary": "Add product to inventory",
                "parameters": [
                    {
                        "description": "Inventory",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modules_storage_internal_handler.dtoCreateInventory"
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
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/product": {
            "post": {
                "tags": [
                    "Product"
                ],
                "summary": "Create new product",
                "parameters": [
                    {
                        "description": "Product",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modules_storage_internal_handler.dtoCreateProduct"
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
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "tags": [
                    "Product"
                ],
                "summary": "Get product by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/modules_storage_internal_handler.dtoProduct"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "tags": [
                    "Product"
                ],
                "summary": "Get all products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/modules_storage_internal_handler.dtoProduct"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/recipe": {
            "get": {
                "tags": [
                    "Recipe"
                ],
                "summary": "Get all recipes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_modules_storage_internal_service.DtoRecipe"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Recipe"
                ],
                "summary": "Create recipe and add products and instructions",
                "parameters": [
                    {
                        "description": "Recipe",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modules_storage_internal_handler.dtoCreateRecipe"
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
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign in to account",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modules_user_internal_handler.dtoAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Saves token to cookie"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign up into account",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/modules_user_internal_handler.dtoAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        },
        "/upload": {
            "post": {
                "tags": [
                    "Files Uploading"
                ],
                "summary": "Upload file to cloud storage",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/modules_storage_internal_handler.dtoUploadedFile"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "fiber.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_IN-45_INT20H-test-task_modules_storage_internal_service.DtoRecipe": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "cooking_time_minutes": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "instructions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_modules_storage_internal_service.dtoInstruction"
                    }
                },
                "name": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_IN-45_INT20H-test-task_modules_storage_internal_service.dtoProduct"
                    }
                }
            }
        },
        "github_com_IN-45_INT20H-test-task_modules_storage_internal_service.dtoInstruction": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                }
            }
        },
        "github_com_IN-45_INT20H-test-task_modules_storage_internal_service.dtoProduct": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "amount_type": {
                    "type": "string"
                },
                "category_id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "github_com_IN-45_INT20H-test-task_pkg_customerrors.UnauthorizedError": {
            "type": "object"
        },
        "modules_storage_internal_handler.dtoCategory": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "modules_storage_internal_handler.dtoCreateCategory": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "modules_storage_internal_handler.dtoCreateInstruction": {
            "type": "object",
            "required": [
                "priority"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                }
            }
        },
        "modules_storage_internal_handler.dtoCreateInventory": {
            "type": "object",
            "required": [
                "product_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "amount_type": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "modules_storage_internal_handler.dtoCreateProduct": {
            "type": "object",
            "required": [
                "category_id",
                "name"
            ],
            "properties": {
                "amount_type": {
                    "type": "string"
                },
                "category_id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "modules_storage_internal_handler.dtoCreateRecipe": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "cooking_time_minutes": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "instructions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/modules_storage_internal_handler.dtoCreateInstruction"
                    }
                },
                "name": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/modules_storage_internal_handler.dtoCreateRecipesProducts"
                    }
                }
            }
        },
        "modules_storage_internal_handler.dtoCreateRecipesProducts": {
            "type": "object",
            "required": [
                "amount",
                "amount_type",
                "product_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "amount_type": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "modules_storage_internal_handler.dtoInventory": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "amount_type": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string"
                }
            }
        },
        "modules_storage_internal_handler.dtoProduct": {
            "type": "object",
            "properties": {
                "amount_type": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "modules_storage_internal_handler.dtoUploadedFile": {
            "type": "object",
            "properties": {
                "file_url": {
                    "type": "string"
                }
            }
        },
        "modules_user_internal_handler.dtoAuth": {
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
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 6
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Backend API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

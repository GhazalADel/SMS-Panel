// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
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
        "/account/phone-books/": {
            "post": {
                "description": "Create a new phone book entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "Create a phone book entry",
                "parameters": [
                    {
                        "description": "Phone book entry data",
                        "name": "phoneBook",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.PhoneBookRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.PhoneBookResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/account/phone-books/phone-book-numbers": {
            "post": {
                "description": "Create a new phone book number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PhoneBookNumbers"
                ],
                "summary": "Create a new phone book number",
                "parameters": [
                    {
                        "description": "Phone book number object",
                        "name": "phoneBookNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreatePhoneBookNumberRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.PhoneBookNumber"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/account/phone-books/phone-book-numbers/{phoneBookNumberID}": {
            "get": {
                "description": "Get phone book number by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PhoneBookNumbers"
                ],
                "summary": "Get phone book number by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Phone book number ID",
                        "name": "phoneBookNumberID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PhoneBookNumber"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update phone book number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PhoneBookNumbers"
                ],
                "summary": "Update phone book number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Phone book number ID",
                        "name": "phoneBookNumberID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Phone book number object",
                        "name": "phoneBookNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdatePhoneBookNumberRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PhoneBookNumber"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete phone book number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PhoneBookNumbers"
                ],
                "summary": "Delete phone book number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Phone book number ID",
                        "name": "phoneBookNumberID",
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
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/account/phone-books/{phoneBookID}/phone-book-numbers": {
            "get": {
                "description": "Get all phone book numbers for a given PhoneBookID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PhoneBookNumbers"
                ],
                "summary": "Get all phone book numbers for a given PhoneBookID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Phone book ID",
                        "name": "phoneBookID",
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
                                "$ref": "#/definitions/models.PhoneBookNumber"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/account/{accountID}/phone-books/": {
            "get": {
                "description": "Get all phone books for a given account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "Get all phone books",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "accountID",
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
                                "$ref": "#/definitions/handlers.PhoneBookResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/account/{accountID}/phone-books/{phoneBookID}": {
            "get": {
                "description": "Get a phone book by ID for a given account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "Get a phone book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "accountID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Phone Book ID",
                        "name": "phoneBookID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.PhoneBookResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a phone book by ID for a given account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "Update a phone book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "accountID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Phone Book ID",
                        "name": "phoneBookID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Phone Book object",
                        "name": "phoneBook",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.PhoneBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.PhoneBookResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a phone book by ID for a given account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "phonebook"
                ],
                "summary": "Delete a phone book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "accountID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Phone Book ID",
                        "name": "phoneBookID",
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
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/accounts/budget": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the budget amount for the logged-in user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get budget amount",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.BudgetAmountResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts/login": {
            "post": {
                "description": "Login with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.AccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponseRegisterLogin"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponseRegisterLogin"
                        }
                    }
                }
            }
        },
        "/accounts/payment/request": {
            "post": {
                "description": "Zarinpal Payment to add budget to account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "Add budget request",
                "parameters": [
                    {
                        "description": "Payment request details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AmountFee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.RequestResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/accounts/payment/verify": {
            "get": {
                "description": "Verify Zarinpal Payment to add budget to account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payment"
                ],
                "summary": "Verify budget payment and add budget",
                "parameters": [
                    {
                        "description": "Payment verify details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.VerifyResponse"
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
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts/register": {
            "post": {
                "description": "Register a new user with the provided information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.AccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponseRegisterLogin"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponseRegisterLogin"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponseRegisterLogin"
                        }
                    }
                }
            }
        },
        "/sms/phonebooks": {
            "post": {
                "description": "Send sms to phone books numbers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SMS"
                ],
                "summary": "Send sms to phone books numbers",
                "parameters": [
                    {
                        "description": "Phone books sms details.",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SendSMessageToPhoneBooksBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.SendSMSResponse"
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/sms/single-sms": {
            "post": {
                "description": "Sends a single SMS message and saves the result in the SMSMessage table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SMS"
                ],
                "summary": "Send Single SMS",
                "parameters": [
                    {
                        "type": "string",
                        "default": "\"account_token\"",
                        "description": "account_token",
                        "name": "Cookie",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Request body for sending an SMS message",
                        "name": "sendSMSRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SendSMSRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.SendSMSResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponseSingle"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponseSingle"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponseSingle"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.AccountResponse": {
            "type": "object",
            "properties": {
                "Budget": {
                    "type": "integer"
                },
                "ID": {
                    "type": "integer"
                },
                "IsActive": {
                    "type": "boolean"
                },
                "Password": {
                    "type": "string"
                },
                "Token": {
                    "type": "string"
                },
                "UserID": {
                    "type": "integer"
                },
                "Username": {
                    "type": "string"
                }
            }
        },
        "handlers.AmountFee": {
            "type": "object",
            "properties": {
                "fee": {
                    "type": "integer"
                }
            }
        },
        "handlers.BudgetAmountResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                }
            }
        },
        "handlers.CreatePhoneBookNumberRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "phoneBookID": {
                    "type": "integer"
                },
                "prefix": {
                    "type": "string"
                }
            }
        },
        "handlers.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.ErrorResponseRegisterLogin": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "responsecode": {
                    "type": "integer"
                }
            }
        },
        "handlers.ErrorResponseSingle": {
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
        "handlers.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.PhoneBookRequest": {
            "type": "object",
            "required": [
                "accountID",
                "name"
            ],
            "properties": {
                "accountID": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.PhoneBookResponse": {
            "type": "object",
            "properties": {
                "accountID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.RequestResponse": {
            "type": "object",
            "properties": {
                "payment_url": {
                    "type": "string"
                }
            }
        },
        "handlers.SendSMSRequest": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Hello, World!"
                },
                "phone_number": {
                    "type": "string",
                    "example": "1234567890"
                },
                "username": {
                    "type": "string",
                    "example": "johndoe"
                }
            }
        },
        "handlers.SendSMSResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "SMS sent successfully"
                }
            }
        },
        "handlers.SendSMessageToPhoneBooksBody": {
            "type": "object",
            "required": [
                "message",
                "phoneBooks"
            ],
            "properties": {
                "message": {
                    "type": "string"
                },
                "phoneBooks": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "handlers.UpdatePhoneBookNumberRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "prefix": {
                    "type": "string"
                }
            }
        },
        "handlers.UserCreateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "nationalid": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.VerifyResponse": {
            "type": "object",
            "properties": {
                "Authority": {
                    "type": "string"
                },
                "Status": {
                    "type": "string"
                }
            }
        },
        "models.PhoneBook": {
            "type": "object",
            "properties": {
                "accountID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.PhoneBookNumber": {
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
                },
                "phoneBook": {
                    "$ref": "#/definitions/models.PhoneBook"
                },
                "phoneBookID": {
                    "type": "integer"
                },
                "prefix": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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
	Title:            "SMS-PANEL",
	Description:      "Quera SMS-PANEL server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

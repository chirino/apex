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
        "contact": {
            "name": "The Apex Authors",
            "url": "https://github.com/nexodus-io/nexodus/issues"
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
        "/devices": {
            "get": {
                "description": "Lists all devices",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "List Devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Device"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new device",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Add Devices",
                "parameters": [
                    {
                        "description": "Add Device",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddDevice"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/devices/{id}": {
            "get": {
                "description": "Gets a device by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Get Devices",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/peers": {
            "get": {
                "description": "Lists all peers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peers"
                ],
                "summary": "List Peers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Peer"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/peers/{peer_id}": {
            "get": {
                "description": "Gets a peer",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peers"
                ],
                "summary": "Get Peer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Zone ID",
                        "name": "peer_id",
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
                                "$ref": "#/definitions/models.Peer"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Lists all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "List Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Gets a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
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
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Changes the users zone",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Patch User",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PatchUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/zones": {
            "get": {
                "description": "Lists all zones",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Zone"
                ],
                "summary": "List Zones",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Zone"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a named zone with the given CIDR",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Zone"
                ],
                "summary": "Create a Zone",
                "parameters": [
                    {
                        "description": "Add Zone",
                        "name": "zone",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddZone"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Zone"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/zones/{id}": {
            "get": {
                "description": "Gets a Zone by Zone ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Zone"
                ],
                "summary": "Get Zones",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Zone ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Zone"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/zones/{id}/peers": {
            "get": {
                "description": "Lists all peers for this zone",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peers"
                ],
                "summary": "List Peers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Zone ID",
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
                                "$ref": "#/definitions/models.Peer"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/zones/{zone_id}/peers": {
            "post": {
                "description": "Adds a new Peer in this Zone",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peers"
                ],
                "summary": "Add Peer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Zone ID",
                        "name": "zone_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Add Peer",
                        "name": "peer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddPeer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Peer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.Peer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/zones/{zone_id}/peers/{peer_id}": {
            "get": {
                "description": "Gets a peer in a zone by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peers"
                ],
                "summary": "Get Peer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Zone ID",
                        "name": "zone_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Zone ID",
                        "name": "peer_id",
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
                                "$ref": "#/definitions/models.Peer"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddDevice": {
            "type": "object",
            "properties": {
                "public_key": {
                    "type": "string",
                    "example": "rZlVfefGshRxO+r9ethv2pODimKlUeP/bO/S47K3WUk="
                }
            }
        },
        "models.AddPeer": {
            "type": "object",
            "properties": {
                "allowed_ips": {
                    "type": "string",
                    "example": "10.1.1.1"
                },
                "child_prefix": {
                    "type": "string",
                    "example": "172.16.42.0/24"
                },
                "device_id": {
                    "type": "string",
                    "example": "6a6090ad-fa47-4549-a144-02124757ab8f"
                },
                "endpoint_ip": {
                    "type": "string",
                    "example": "10.1.1.1"
                },
                "hub_router": {
                    "type": "boolean"
                },
                "hub_zone": {
                    "type": "boolean"
                },
                "node_address": {
                    "type": "string",
                    "example": "1.2.3.4"
                },
                "zone_prefix": {
                    "type": "string",
                    "example": "10.1.1.0/24"
                }
            }
        },
        "models.AddZone": {
            "type": "object",
            "properties": {
                "cidr": {
                    "type": "string",
                    "example": "172.16.42.0/24"
                },
                "description": {
                    "type": "string",
                    "example": "The Red Zone"
                },
                "hub_zone": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string",
                    "example": "zone-red"
                }
            }
        },
        "models.ApiError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "something bad"
                }
            }
        },
        "models.Device": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "aa22666c-0f57-45cb-a449-16efecc04f2e"
                },
                "peers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "97d5214a-8c51-4772-b492-53de034740c5"
                    ]
                },
                "public_key": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.PatchUser": {
            "type": "object",
            "properties": {
                "zone_id": {
                    "type": "string",
                    "example": "3f51dda6-06d2-4724-bb73-f09ad3501bcc"
                }
            }
        },
        "models.Peer": {
            "type": "object",
            "properties": {
                "allowed_ips": {
                    "type": "string",
                    "example": "10.1.1.1"
                },
                "child_prefix": {
                    "type": "string",
                    "example": "172.16.42.0/24"
                },
                "device_id": {
                    "type": "string",
                    "example": "fde38e78-a4af-4f44-8f5a-d84ef1846a85"
                },
                "endpoint_ip": {
                    "type": "string",
                    "example": "10.1.1.1"
                },
                "hub_router": {
                    "type": "boolean"
                },
                "hub_zone": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string",
                    "example": "aa22666c-0f57-45cb-a449-16efecc04f2e"
                },
                "node_address": {
                    "type": "string",
                    "example": "1.2.3.4"
                },
                "zone_id": {
                    "type": "string",
                    "example": "2b655c5b-cfdd-4550-b7f0-a36a590fc97a"
                },
                "zone_prefix": {
                    "type": "string",
                    "example": "10.1.1.0/24"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "devices": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "4902c991-3dd1-49a6-9f26-d82496c80aff"
                    ]
                },
                "id": {
                    "type": "string",
                    "example": "aa22666c-0f57-45cb-a449-16efecc04f2e"
                },
                "zone_id": {
                    "type": "string",
                    "example": "94deb404-c4eb-4097-b59d-76b024ff7867"
                }
            }
        },
        "models.Zone": {
            "type": "object",
            "properties": {
                "cidr": {
                    "type": "string",
                    "example": "172.16.42.0/24"
                },
                "description": {
                    "type": "string",
                    "example": "The Red Zone"
                },
                "hub_zone": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string",
                    "example": "aa22666c-0f57-45cb-a449-16efecc04f2e"
                },
                "name": {
                    "type": "string",
                    "example": "zone-red"
                },
                "peers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "4902c991-3dd1-49a6-9f26-d82496c80aff"
                    ]
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "/auth/realms/controller/protocol/openid-connect/auth",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "user": " Grants read and write access to resources owned by this user"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Apex API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

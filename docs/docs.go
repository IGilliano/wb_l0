// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "ivv_@mail.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/{id}": {
            "get": {
                "description": "Responds with the order as JSON",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Oder By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Search order by ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/wb_l0.Order"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "wb_l0.Delivery": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "zip": {
                    "type": "string"
                }
            }
        },
        "wb_l0.Item": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "chrt_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nm_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "rid": {
                    "type": "string"
                },
                "sale": {
                    "type": "integer"
                },
                "size": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                }
            }
        },
        "wb_l0.Order": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "string"
                },
                "date_created": {
                    "type": "string"
                },
                "delivery": {
                    "$ref": "#/definitions/wb_l0.Delivery"
                },
                "delivery_service": {
                    "type": "string"
                },
                "entry": {
                    "type": "string"
                },
                "internal_signature": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/wb_l0.Item"
                    }
                },
                "locale": {
                    "type": "string"
                },
                "oof_shard": {
                    "type": "string"
                },
                "order_uid": {
                    "type": "string"
                },
                "payment": {
                    "$ref": "#/definitions/wb_l0.Payment"
                },
                "shardkey": {
                    "type": "string"
                },
                "sm_id": {
                    "type": "integer"
                },
                "track_number": {
                    "type": "string"
                }
            }
        },
        "wb_l0.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "bank": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "custom_fee": {
                    "type": "integer"
                },
                "delivery_cost": {
                    "type": "integer"
                },
                "goods_total": {
                    "type": "integer"
                },
                "payment_dt": {
                    "type": "integer"
                },
                "provider": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                },
                "transaction": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "WB L0",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

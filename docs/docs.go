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
        "/goods/info": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品"
                ],
                "summary": "通过商品ID获取商品信息",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.Goods"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/goods/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品"
                ],
                "summary": "获取所有商品信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/order/cancel": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单"
                ],
                "summary": "取消订单",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CancelOrderRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/order/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单"
                ],
                "summary": "历史订单",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.OrderInfoResponse"
                        }
                    }
                }
            }
        },
        "/order/place": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "订单"
                ],
                "summary": "下单",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.PlaceOrderRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.GetProfileResponse"
                        }
                    }
                }
            }
        },
        "/user/address/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "保存收货地址信息",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UserAddressInfoRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/user/address/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "收货地址列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserAddressInfo"
                        }
                    }
                }
            }
        },
        "/user/address/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "更新收货地址信息",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UserAddressInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/wechat/program/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "微信小程序"
                ],
                "summary": "微信小程序用户登录",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.WechatProgramLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/wechat/qr/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "微信小程序"
                ],
                "summary": "微信小程序登录二维码",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.OrderGoods": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "商品数量",
                    "type": "integer"
                },
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "deleted": {
                    "description": "是否删除",
                    "type": "integer"
                },
                "goodsId": {
                    "description": "商品ID",
                    "type": "integer"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "orderId": {
                    "description": "订单ID",
                    "type": "integer"
                },
                "price": {
                    "description": "商品单价",
                    "type": "integer"
                },
                "updateTime": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "model.UserAddressInfo": {
            "type": "object",
            "properties": {
                "addressDetail": {
                    "description": "详细地址",
                    "type": "string"
                },
                "city": {
                    "description": "市",
                    "type": "string"
                },
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "defaultState": {
                    "description": "是否默认地址",
                    "type": "integer"
                },
                "deleted": {
                    "description": "是否删除",
                    "type": "integer"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "latitude": {
                    "description": "纬度",
                    "type": "integer"
                },
                "longitude": {
                    "description": "经度",
                    "type": "integer"
                },
                "phoneNumber": {
                    "description": "收获人电话",
                    "type": "string"
                },
                "province": {
                    "description": "省",
                    "type": "string"
                },
                "region": {
                    "description": "区",
                    "type": "string"
                },
                "updateTime": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userId": {
                    "description": "用户ID",
                    "type": "integer"
                },
                "userName": {
                    "description": "收获人姓名",
                    "type": "string"
                }
            }
        },
        "v1.CancelOrderRequest": {
            "type": "object",
            "required": [
                "orderId"
            ],
            "properties": {
                "orderId": {
                    "type": "integer"
                }
            }
        },
        "v1.GetProfileResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/v1.GetProfileResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.GetProfileResponseData": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string",
                    "example": "alan"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "v1.Goods": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "v1.OrderGoods": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "商品数量",
                    "type": "integer"
                },
                "goodsId": {
                    "description": "商品ID",
                    "type": "integer"
                },
                "goodsName": {
                    "description": "商品名称",
                    "type": "string"
                },
                "price": {
                    "description": "单价",
                    "type": "integer"
                },
                "total": {
                    "description": "商品金额",
                    "type": "integer"
                }
            }
        },
        "v1.OrderInfoResponse": {
            "type": "object",
            "properties": {
                "addressDetail": {
                    "description": "详细地址",
                    "type": "string"
                },
                "city": {
                    "description": "市",
                    "type": "string"
                },
                "goods": {
                    "description": "商品列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.OrderGoods"
                    }
                },
                "latitude": {
                    "description": "纬度",
                    "type": "integer"
                },
                "longitude": {
                    "description": "经度",
                    "type": "integer"
                },
                "orderId": {
                    "description": "订单ID",
                    "type": "integer"
                },
                "orderState": {
                    "description": "订单状态",
                    "type": "integer"
                },
                "paymentModel": {
                    "description": "支付方式",
                    "type": "integer"
                },
                "phoneNumber": {
                    "description": "收获人电话",
                    "type": "string"
                },
                "placeTime": {
                    "description": "下单时间",
                    "type": "string"
                },
                "province": {
                    "description": "省",
                    "type": "string"
                },
                "region": {
                    "description": "区",
                    "type": "string"
                },
                "remark": {
                    "description": "订单备注",
                    "type": "string"
                },
                "total": {
                    "description": "订单金额",
                    "type": "integer"
                },
                "userId": {
                    "description": "用户ID",
                    "type": "integer"
                },
                "userName": {
                    "description": "收获人姓名",
                    "type": "string"
                }
            }
        },
        "v1.PlaceOrderRequest": {
            "type": "object",
            "required": [
                "addressId",
                "deliveryTime",
                "orderGoods"
            ],
            "properties": {
                "addressId": {
                    "type": "integer"
                },
                "deliveryTime": {
                    "type": "string"
                },
                "orderGoods": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderGoods"
                    }
                },
                "remark": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "v1.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.UserAddressInfoRequest": {
            "type": "object",
            "required": [
                "addressDetail",
                "city",
                "latitude",
                "longitude",
                "phoneNumber",
                "province",
                "region",
                "userName"
            ],
            "properties": {
                "addressDetail": {
                    "description": "详细地址",
                    "type": "string"
                },
                "city": {
                    "description": "市",
                    "type": "string"
                },
                "latitude": {
                    "description": "纬度",
                    "type": "integer"
                },
                "longitude": {
                    "description": "经度",
                    "type": "integer"
                },
                "phoneNumber": {
                    "description": "收获人电话",
                    "type": "string"
                },
                "province": {
                    "description": "省",
                    "type": "string"
                },
                "region": {
                    "description": "区",
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                },
                "userName": {
                    "description": "收获人姓名",
                    "type": "string"
                }
            }
        },
        "v1.WechatProgramLoginRequest": {
            "type": "object",
            "required": [
                "jsCode"
            ],
            "properties": {
                "jsCode": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Nunu Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

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
        "/coin/list": {
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
                    "公共"
                ],
                "summary": "币币信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bar"
                        }
                    }
                }
            }
        },
        "/index/bar/list": {
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
                    "公共"
                ],
                "summary": "时间线",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bar"
                        }
                    }
                }
            }
        },
        "/index/list": {
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
                    "指标"
                ],
                "summary": "指标列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Index"
                        }
                    }
                }
            }
        },
        "/index/test": {
            "post": {
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
                    "指标"
                ],
                "summary": "指标回测",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.IndexRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Index"
                        }
                    }
                }
            }
        },
        "/sms/code": {
            "post": {
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
                    "短信"
                ],
                "summary": "发送短信验证码",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SendSMSCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/strategy/create": {
            "post": {
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
                    "策略"
                ],
                "summary": "创建策略",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.CreateStrategyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/strategy/delete": {
            "post": {
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
                    "策略"
                ],
                "summary": "删除策略",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.StrategyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/strategy/index/delete": {
            "post": {
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
                    "策略"
                ],
                "summary": "删除策略关联指标",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.StrategyIndexRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/strategy/index/list": {
            "post": {
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
                    "策略"
                ],
                "summary": "获取策略关联指标",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.StrategyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/strategy/list": {
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
                    "策略"
                ],
                "summary": "获取当前用户的策略列表",
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
                            "$ref": "#/definitions/model.Strategy"
                        }
                    }
                }
            }
        },
        "/strategy/subscription": {
            "post": {
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
                    "策略"
                ],
                "summary": "更新策略订阅状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.StrategyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/subscription/pay/wallet": {
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
                    "公共"
                ],
                "summary": "获取订阅支付钱包地址",
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
        "/subscription/price": {
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
                    "公共"
                ],
                "summary": "获取订阅价格",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Bar"
                        }
                    }
                }
            }
        },
        "/transaction/recharge/confirm": {
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
                    "交易"
                ],
                "summary": "充值确认",
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
                        "description": "OK"
                    }
                }
            }
        },
        "/update/user": {
            "post": {
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
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UpdateUserInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
        "/verification/sms/code": {
            "post": {
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
                    "短信"
                ],
                "summary": "核对验证码",
                "parameters": [
                    {
                        "description": "params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.SendSMSCodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Bar": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "deleted": {
                    "description": "是否删除",
                    "type": "boolean"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "updateTime": {
                    "description": "更新时间",
                    "type": "string"
                },
                "value": {
                    "description": "bar value",
                    "type": "string"
                }
            }
        },
        "model.Index": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "defaultConfig": {
                    "description": "指标简介",
                    "type": "string"
                },
                "deleted": {
                    "description": "是否删除",
                    "type": "boolean"
                },
                "desc": {
                    "description": "指标简介",
                    "type": "string"
                },
                "icon": {
                    "description": "指标图标",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "name": {
                    "description": "指标名称",
                    "type": "string"
                },
                "updateTime": {
                    "description": "更新时间",
                    "type": "string"
                },
                "warningConfigArray": {
                    "description": "预警配置",
                    "type": "string"
                }
            }
        },
        "model.Strategy": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "deleted": {
                    "description": "是否删除",
                    "type": "boolean"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "instId": {
                    "description": "币种ID",
                    "type": "string"
                },
                "strategyName": {
                    "description": "策略名称",
                    "type": "string"
                },
                "subscriptionState": {
                    "description": "订阅状态: 0 未订阅 1 已订阅",
                    "type": "integer"
                },
                "updateTime": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userId": {
                    "description": "用户ID",
                    "type": "integer"
                }
            }
        },
        "v1.CreateStrategyRequest": {
            "type": "object",
            "required": [
                "InstId",
                "indexList",
                "name"
            ],
            "properties": {
                "InstId": {
                    "type": "string"
                },
                "indexList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v1.IndexRequest"
                    }
                },
                "name": {
                    "type": "string"
                },
                "subscriptionState": {
                    "type": "integer"
                },
                "userId": {
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
        "v1.IndexRequest": {
            "type": "object",
            "required": [
                "InstId",
                "bar",
                "indexConfig",
                "indexId",
                "warningConfig"
            ],
            "properties": {
                "InstId": {
                    "type": "string"
                },
                "bar": {
                    "type": "string"
                },
                "indexConfig": {
                    "type": "string"
                },
                "indexId": {
                    "type": "integer"
                },
                "warningConfig": {
                    "type": "string"
                }
            }
        },
        "v1.SendSMSCodeRequest": {
            "type": "object",
            "required": [
                "code",
                "phoneNumber"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "example": "123456"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "1234567890"
                }
            }
        },
        "v1.StrategyIndexRequest": {
            "type": "object",
            "required": [
                "strategyIndexId"
            ],
            "properties": {
                "strategyIndexId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "v1.StrategyRequest": {
            "type": "object",
            "required": [
                "strategyId"
            ],
            "properties": {
                "strategyId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "v1.UpdateUserInfoRequest": {
            "type": "object",
            "properties": {
                "phoneNumber": {
                    "type": "string"
                },
                "wallet": {
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

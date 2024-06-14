// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/dummy-table": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get dummy table",
                "responses": {
                    "200": {
                        "description": "dummy table rows",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handlers.dummyTable"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/restaurants": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get all restaurants",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/sqlcClient.GetRestaurantsRow"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.dummyTable": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "message": {
                    "type": "string",
                    "example": "example error"
                }
            }
        },
        "sql.NullBool": {
            "type": "object",
            "properties": {
                "bool": {
                    "type": "boolean"
                },
                "valid": {
                    "description": "Valid is true if Bool is not NULL",
                    "type": "boolean"
                }
            }
        },
        "sql.NullFloat64": {
            "type": "object",
            "properties": {
                "float64": {
                    "type": "number"
                },
                "valid": {
                    "description": "Valid is true if Float64 is not NULL",
                    "type": "boolean"
                }
            }
        },
        "sql.NullInt32": {
            "type": "object",
            "properties": {
                "int32": {
                    "type": "integer"
                },
                "valid": {
                    "description": "Valid is true if Int32 is not NULL",
                    "type": "boolean"
                }
            }
        },
        "sql.NullString": {
            "type": "object",
            "properties": {
                "string": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if String is not NULL",
                    "type": "boolean"
                }
            }
        },
        "sqlcClient.GetRestaurantsRow": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "googleRestaurant": {
                    "$ref": "#/definitions/sqlcClient.GoogleRestaurant"
                },
                "location": {
                    "$ref": "#/definitions/sqlcClient.Location"
                },
                "restaurantId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "sqlcClient.GoogleRestaurant": {
            "type": "object",
            "properties": {
                "acceptsCashOnly": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "acceptsCreditCards": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "acceptsDebitCards": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "acceptsNfc": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "avgRating": {
                    "$ref": "#/definitions/sql.NullFloat64"
                },
                "businessStatus": {
                    "$ref": "#/definitions/sqlcClient.NullGooglerestaurantBusinessStatus"
                },
                "description": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "goodForGroups": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "goodForWatchingSports": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "googleRestaurantId": {
                    "type": "integer"
                },
                "googleUrl": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "hasOutdoorSeating": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "hasRestroom": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "placeId": {
                    "type": "string"
                },
                "priceLevel": {
                    "$ref": "#/definitions/sqlcClient.NullGooglerestaurantPriceLevel"
                },
                "servesBeer": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesBreakfast": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesBrunch": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesCocktails": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesCoffee": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesDessert": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesDinner": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesLunch": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesVegetarianFood": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "servesWine": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "supportsCurbsidePickup": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "supportsDelivery": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "supportsDineIn": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "supportsReservations": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "supportsTakeout": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "updatedAt": {
                    "type": "string"
                },
                "website": {
                    "$ref": "#/definitions/sql.NullString"
                },
                "wheelchairAccessibleEntrance": {
                    "$ref": "#/definitions/sql.NullBool"
                },
                "wheelchairAccessibleSeating": {
                    "$ref": "#/definitions/sql.NullBool"
                }
            }
        },
        "sqlcClient.GooglerestaurantBusinessStatus": {
            "type": "string",
            "enum": [
                "Operational",
                "ClosedTemporarily",
                "ClosedPermanently"
            ],
            "x-enum-varnames": [
                "GooglerestaurantBusinessStatusOperational",
                "GooglerestaurantBusinessStatusClosedTemporarily",
                "GooglerestaurantBusinessStatusClosedPermanently"
            ]
        },
        "sqlcClient.GooglerestaurantPriceLevel": {
            "type": "string",
            "enum": [
                "Free",
                "Inexpensive",
                "Moderate",
                "Expensive",
                "VeryExpensive"
            ],
            "x-enum-varnames": [
                "GooglerestaurantPriceLevelFree",
                "GooglerestaurantPriceLevelInexpensive",
                "GooglerestaurantPriceLevelModerate",
                "GooglerestaurantPriceLevelExpensive",
                "GooglerestaurantPriceLevelVeryExpensive"
            ]
        },
        "sqlcClient.Location": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "googleRestaurantId": {
                    "$ref": "#/definitions/sql.NullInt32"
                },
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                },
                "locationId": {
                    "type": "integer"
                },
                "viewportHighLat": {
                    "type": "number"
                },
                "viewportHighLng": {
                    "type": "number"
                },
                "viewportLowLat": {
                    "type": "number"
                },
                "viewportLowLng": {
                    "type": "number"
                }
            }
        },
        "sqlcClient.NullGooglerestaurantBusinessStatus": {
            "type": "object",
            "properties": {
                "googlerestaurantBusinessStatus": {
                    "$ref": "#/definitions/sqlcClient.GooglerestaurantBusinessStatus"
                },
                "valid": {
                    "description": "Valid is true if GooglerestaurantBusinessStatus is not NULL",
                    "type": "boolean"
                }
            }
        },
        "sqlcClient.NullGooglerestaurantPriceLevel": {
            "type": "object",
            "properties": {
                "googlerestaurantPriceLevel": {
                    "$ref": "#/definitions/sqlcClient.GooglerestaurantPriceLevel"
                },
                "valid": {
                    "description": "Valid is true if GooglerestaurantPriceLevel is not NULL",
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Restaurant Flow",
	Description:      "Restaurant reviews for UW Students",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

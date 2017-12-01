// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Flagr is a feature flagging, A/B testing and dynamic configuration microservice",
    "title": "Flagr",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/evaluation": {
      "post": {
        "tags": [
          "evaluation"
        ],
        "operationId": "postEvaluation",
        "parameters": [
          {
            "description": "evalution context",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/evalContext"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "evaluation result",
            "schema": {
              "$ref": "#/definitions/evalResult"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/evaluation/batch": {
      "post": {
        "tags": [
          "evaluation"
        ],
        "operationId": "postEvaluationBatch",
        "parameters": [
          {
            "description": "evalution batch request",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/evaluationBatchRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "evaluation batch result",
            "schema": {
              "$ref": "#/definitions/evaluationBatchResponse"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags": {
      "get": {
        "tags": [
          "flag"
        ],
        "operationId": "findFlags",
        "responses": {
          "200": {
            "description": "list all the flags",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/flag"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "flag"
        ],
        "operationId": "createFlag",
        "parameters": [
          {
            "description": "create a flag",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createFlagRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns the created flag",
            "schema": {
              "$ref": "#/definitions/flag"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}": {
      "get": {
        "tags": [
          "flag"
        ],
        "operationId": "getFlag",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag to get",
            "name": "flagID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns the flag",
            "schema": {
              "$ref": "#/definitions/flag"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "flag"
        ],
        "operationId": "putFlag",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag to get",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "description": "update a flag",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/putFlagRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns the flag",
            "schema": {
              "$ref": "#/definitions/flag"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "flag"
        ],
        "operationId": "deleteFlag",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK deleted"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/enabled": {
      "put": {
        "tags": [
          "flag"
        ],
        "operationId": "setFlagEnabled",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag to get",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "description": "set flag enabled state",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/setFlagEnabledRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns the flag",
            "schema": {
              "$ref": "#/definitions/flag"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/segments": {
      "get": {
        "tags": [
          "segment"
        ],
        "operationId": "findSegments",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag to get",
            "name": "flagID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "segments ordered by rank of the flag",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/segment"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "segment"
        ],
        "operationId": "createSegment",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag to get",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "description": "create a segment under a flag",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createSegmentRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "segment created",
            "schema": {
              "$ref": "#/definitions/segment"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/segments/reorder": {
      "put": {
        "tags": [
          "segment"
        ],
        "operationId": "putSegmentsReorder",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "description": "reorder segments",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/putSegmentReorderRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "segments reordered"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/segments/{segmentID}": {
      "put": {
        "tags": [
          "segment"
        ],
        "operationId": "putSegment",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the segment",
            "name": "segmentID",
            "in": "path",
            "required": true
          },
          {
            "description": "update a segment",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/putSegmentRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "segment updated",
            "schema": {
              "$ref": "#/definitions/segment"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "segment"
        ],
        "operationId": "deleteSegment",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the segment",
            "name": "segmentID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "deleted"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/segments/{segmentID}/constraints": {
      "get": {
        "tags": [
          "constraint"
        ],
        "operationId": "findConstraints",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the segment",
            "name": "segmentID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "constraints under the segment",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/constraint"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "constraint"
        ],
        "operationId": "createConstraint",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the segment",
            "name": "segmentID",
            "in": "path",
            "required": true
          },
          {
            "description": "create a constraint",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createConstraintRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "the constraint created",
            "schema": {
              "$ref": "#/definitions/constraint"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/segments/{segmentID}/constraints/{constraintID}": {
      "put": {
        "tags": [
          "constraint"
        ],
        "operationId": "putConstraint",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the segment",
            "name": "segmentID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the constraint",
            "name": "constraintID",
            "in": "path",
            "required": true
          },
          {
            "description": "create a constraint",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createConstraintRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "constraint just updated",
            "schema": {
              "$ref": "#/definitions/constraint"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "constraint"
        ],
        "operationId": "deleteConstraint",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the segment",
            "name": "segmentID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the constraint",
            "name": "constraintID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "deleted"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/segments/{segmentID}/distributions": {
      "get": {
        "tags": [
          "distribution"
        ],
        "operationId": "findDistributions",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the segment",
            "name": "segmentID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "distribution under the segment",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/distribution"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "description": "replace the distribution with the new setting",
        "tags": [
          "distribution"
        ],
        "operationId": "putDistributions",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the segment",
            "name": "segmentID",
            "in": "path",
            "required": true
          },
          {
            "description": "array of distributions",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/putDistributionsRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "distribution under the segment",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/distribution"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/snapshots": {
      "get": {
        "tags": [
          "flag"
        ],
        "operationId": "getFlagSnapshots",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag to get",
            "name": "flagID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns the flag snapshots",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/flagSnapshot"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/variants": {
      "get": {
        "tags": [
          "variant"
        ],
        "operationId": "findVariants",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "variant ordered by variantID",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/variant"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "variant"
        ],
        "operationId": "createVariant",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "description": "create a variant",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createVariantRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "variant just created",
            "schema": {
              "$ref": "#/definitions/variant"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/flags/{flagID}/variants/{variantID}": {
      "put": {
        "tags": [
          "variant"
        ],
        "operationId": "putVariant",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the variant",
            "name": "variantID",
            "in": "path",
            "required": true
          },
          {
            "description": "update a variant",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/putVariantRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "variant just updated",
            "schema": {
              "$ref": "#/definitions/variant"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "variant"
        ],
        "operationId": "deleteVariant",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the flag",
            "name": "flagID",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "numeric ID of the variant",
            "name": "variantID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "deleted"
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "constraint": {
      "type": "object",
      "required": [
        "property",
        "operator",
        "value"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "readOnly": true
        },
        "operator": {
          "type": "string",
          "minLength": 1,
          "enum": [
            "EQ",
            "NEQ",
            "LT",
            "LTE",
            "GT",
            "GTE",
            "EREG",
            "NEREG",
            "IN",
            "NOTIN",
            "CONTAINS",
            "NOTCONTAINS"
          ]
        },
        "property": {
          "type": "string",
          "minLength": 1
        },
        "value": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "createConstraintRequest": {
      "type": "object",
      "required": [
        "property",
        "operator",
        "value"
      ],
      "properties": {
        "operator": {
          "type": "string",
          "minLength": 1
        },
        "property": {
          "type": "string",
          "minLength": 1
        },
        "value": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "createFlagRequest": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "createSegmentRequest": {
      "type": "object",
      "required": [
        "description",
        "rolloutPercent"
      ],
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        },
        "rolloutPercent": {
          "type": "integer",
          "format": "int64",
          "maximum": 100,
          "minimum": 0
        }
      }
    },
    "createVariantRequest": {
      "type": "object",
      "required": [
        "key"
      ],
      "properties": {
        "attachment": {
          "type": "object"
        },
        "key": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "distribution": {
      "type": "object",
      "required": [
        "percent",
        "variantID",
        "variantKey"
      ],
      "properties": {
        "bitmap": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "readOnly": true
        },
        "percent": {
          "type": "integer",
          "format": "int64",
          "maximum": 100,
          "minimum": 0
        },
        "variantID": {
          "type": "integer",
          "format": "int64",
          "minimum": 1
        },
        "variantKey": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "evalContext": {
      "type": "object",
      "required": [
        "entityID",
        "entityType",
        "flagID"
      ],
      "properties": {
        "enableDebug": {
          "type": "boolean"
        },
        "entityContext": {
          "type": "object"
        },
        "entityID": {
          "type": "string",
          "minLength": 1
        },
        "entityType": {
          "type": "string",
          "minLength": 1
        },
        "flagID": {
          "type": "integer",
          "format": "int64",
          "minimum": 1
        }
      }
    },
    "evalDebugLog": {
      "type": "object",
      "properties": {
        "msg": {
          "type": "string"
        },
        "segmentDebugLogs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/segmentDebugLog"
          }
        }
      }
    },
    "evalResult": {
      "type": "object",
      "required": [
        "flagID",
        "segmentID",
        "variantID",
        "variantKey",
        "variantAttachment",
        "evalContext",
        "timestamp"
      ],
      "properties": {
        "evalContext": {
          "$ref": "#/definitions/evalContext"
        },
        "evalDebugLog": {
          "$ref": "#/definitions/evalDebugLog"
        },
        "flagID": {
          "type": "integer",
          "format": "int64",
          "minimum": 1
        },
        "flagSnapshotID": {
          "type": "integer",
          "format": "int64"
        },
        "segmentID": {
          "type": "integer",
          "format": "int64",
          "minimum": 1
        },
        "timestamp": {
          "type": "string",
          "minLength": 1
        },
        "variantAttachment": {
          "type": "object"
        },
        "variantID": {
          "type": "integer",
          "format": "int64",
          "minimum": 1
        },
        "variantKey": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "evaluationBatchRequest": {
      "type": "object",
      "required": [
        "entities",
        "flagIDs"
      ],
      "properties": {
        "enableDebug": {
          "type": "boolean"
        },
        "entities": {
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/evaluationEntity"
          }
        },
        "flagIDs": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "integer",
            "format": "int64",
            "minimum": 1
          }
        }
      }
    },
    "evaluationBatchResponse": {
      "type": "object",
      "required": [
        "evaluationResults"
      ],
      "properties": {
        "evaluationResults": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/evalResult"
          }
        }
      }
    },
    "evaluationEntity": {
      "type": "object",
      "required": [
        "entityID",
        "entityType"
      ],
      "properties": {
        "entityContext": {
          "type": "object"
        },
        "entityID": {
          "type": "string",
          "minLength": 1
        },
        "entityType": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "flag": {
      "type": "object",
      "required": [
        "description",
        "enabled",
        "dataRecordsEnabled"
      ],
      "properties": {
        "dataRecordsEnabled": {
          "description": "enabled data records will get data logging in the metrics pipeline, for example, kafka.",
          "type": "boolean"
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "enabled": {
          "type": "boolean"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "readOnly": true
        },
        "segments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/segment"
          }
        },
        "variants": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/variant"
          }
        }
      }
    },
    "flagSnapshot": {
      "type": "object",
      "required": [
        "id",
        "flag",
        "updatedAt"
      ],
      "properties": {
        "flag": {
          "$ref": "#/definitions/flag"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "readOnly": true
        },
        "updatedAt": {
          "type": "string",
          "minLength": 1
        },
        "updatedBy": {
          "type": "string"
        }
      }
    },
    "putDistributionsRequest": {
      "type": "object",
      "required": [
        "distributions"
      ],
      "properties": {
        "distributions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/distribution"
          }
        }
      }
    },
    "putFlagRequest": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "dataRecordsEnabled": {
          "description": "enabled data records will get data logging in the metrics pipeline, for example, kafka.",
          "type": "boolean",
          "x-nullable": true
        },
        "description": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "putSegmentReorderRequest": {
      "type": "object",
      "required": [
        "segmentIDs"
      ],
      "properties": {
        "segmentIDs": {
          "type": "array",
          "minItems": 1,
          "items": {
            "type": "integer",
            "format": "int64",
            "minimum": 1
          }
        }
      }
    },
    "putSegmentRequest": {
      "type": "object",
      "required": [
        "description",
        "rolloutPercent"
      ],
      "properties": {
        "description": {
          "type": "string",
          "minLength": 1
        },
        "rolloutPercent": {
          "type": "integer",
          "format": "int64",
          "maximum": 100,
          "minimum": 0
        }
      }
    },
    "putVariantRequest": {
      "type": "object",
      "required": [
        "key",
        "attachment"
      ],
      "properties": {
        "attachment": {
          "type": "object"
        },
        "key": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "segment": {
      "type": "object",
      "required": [
        "description",
        "rank",
        "rolloutPercent"
      ],
      "properties": {
        "constraints": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/constraint"
          }
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "distributions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/distribution"
          }
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "readOnly": true
        },
        "rank": {
          "type": "integer",
          "format": "int64",
          "minimum": 0
        },
        "rolloutPercent": {
          "type": "integer",
          "format": "int64",
          "maximum": 100,
          "minimum": 0
        }
      }
    },
    "segmentDebugLog": {
      "type": "object",
      "properties": {
        "msg": {
          "type": "string"
        },
        "segmentID": {
          "type": "integer",
          "format": "int64",
          "minimum": 1
        }
      }
    },
    "setFlagEnabledRequest": {
      "type": "object",
      "required": [
        "enabled"
      ],
      "properties": {
        "enabled": {
          "type": "boolean"
        }
      }
    },
    "variant": {
      "type": "object",
      "required": [
        "key"
      ],
      "properties": {
        "attachment": {
          "type": "object"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "minimum": 1,
          "readOnly": true
        },
        "key": {
          "type": "string",
          "minLength": 1
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about the flag",
      "name": "flag"
    },
    {
      "description": "Segment defines the audience of the flag, it's the user segmentation",
      "name": "segment"
    },
    {
      "description": "Constraint is the unit of defining a small subset of users",
      "name": "constraint"
    },
    {
      "description": "Distribution is the percent distribution of variants within that segment",
      "name": "distribution"
    },
    {
      "description": "Variants are the possible outcomes of flag evaluation",
      "name": "variant"
    },
    {
      "description": "Evaluation is the process of evaluating a flag given the entity context",
      "name": "evaluation"
    }
  ],
  "x-tagGroups": [
    {
      "name": "Flag Management",
      "tags": [
        "flag",
        "segment",
        "constraint",
        "distribution",
        "variant"
      ]
    },
    {
      "name": "Flag Evaluation",
      "tags": [
        "evaluation"
      ]
    }
  ]
}`))
}

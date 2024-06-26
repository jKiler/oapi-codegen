// Package xorder provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/jKiler/oapi-codegen version v1.0.0-00010101000000-000000000000 DO NOT EDIT.
package xorder

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// DateInterval defines model for DateInterval.
type DateInterval struct {
	Start *openapi_types.Date `json:"start,omitempty"`
	End   *openapi_types.Date `json:"end,omitempty"`
}

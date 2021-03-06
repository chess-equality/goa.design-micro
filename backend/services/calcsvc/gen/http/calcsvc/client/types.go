// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calcsvc HTTP client types
//
// Command:
// $ goa gen goa.design-micro/backend/services/calcsvc/design

package client

import (
	calcsvc "goa.design-micro/backend/services/calcsvc/gen/calcsvc"
)

// MultiplyRequestBody is the type of the "calcsvc" service "multiply" endpoint
// HTTP request body.
type MultiplyRequestBody struct {
	// Multiplicand
	A int `form:"a" json:"a" xml:"a"`
	// Multiplier
	B int `form:"b" json:"b" xml:"b"`
}

// NewMultiplyRequestBody builds the HTTP request body from the payload of the
// "multiply" endpoint of the "calcsvc" service.
func NewMultiplyRequestBody(p *calcsvc.MultiplyPayload) *MultiplyRequestBody {
	body := &MultiplyRequestBody{
		A: p.A,
		B: p.B,
	}
	return body
}

// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the calcsvc service.
//
// Command:
// $ goa gen goa.design-micro/backend/services/calcsvc/design

package server

import (
	"fmt"
)

// AddCalcsvcPath returns the URL path to the calcsvc service add HTTP endpoint.
func AddCalcsvcPath(a int, b int) string {
	return fmt.Sprintf("/add/%v/%v", a, b)
}

// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calcsvc HTTP client CLI support package
//
// Command:
// $ goa gen goa.design-micro/backend/services/calcsvc/design

package client

import (
	"fmt"
	"strconv"

	calcsvc "goa.design-micro/backend/services/calcsvc/gen/calcsvc"
)

// BuildAddPayload builds the payload for the calcsvc add endpoint from CLI
// flags.
func BuildAddPayload(calcsvcAddA string, calcsvcAddB string) (*calcsvc.AddPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcsvcAddA, 10, 64)
		a = int(v)
		if err != nil {
			err = fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcsvcAddB, 10, 64)
		b = int(v)
		if err != nil {
			err = fmt.Errorf("invalid value for b, must be INT")
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &calcsvc.AddPayload{
		A: a,
		B: b,
	}
	return payload, nil
}
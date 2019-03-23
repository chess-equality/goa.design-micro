package calc

import (
	"context"
	"log"

	calcsvc "goa.design-micro/backend/services/calcsvc/gen/calcsvc"
)

// calcsvc service example implementation.
// The example methods log the requests and return zero values.
type calcsvcsrvc struct {
	logger *log.Logger
}

// NewCalcsvc returns the calcsvc service implementation.
func NewCalcsvc(logger *log.Logger) calcsvc.Service {
	return &calcsvcsrvc{logger}
}

// Add implements addition.
// Add returns the sum of attributes a and b of p.
func (s *calcsvcsrvc) Add(ctx context.Context, p *calcsvc.AddPayload) (res int, err error) {
	s.logger.Printf("##### From logger: %s", "calcsvc.add")
	// return
	return p.A + p.B, nil
}

// Multiply implements multiplication.
// Multiply returns the product of attributes a and b of p.
func (s *calcsvcsrvc) Multiply(ctx context.Context, p *calcsvc.MultiplyPayload) (res int, err error) {
	s.logger.Printf("##### From logger: %s", "calcsvc.multiply")
	// return
	return p.A * p.B, nil
}

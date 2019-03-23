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

// Add implements add.
// Add returns the sum of attributes a and b of p.
func (s *calcsvcsrvc) Add(ctx context.Context, p *calcsvc.AddPayload) (res int, err error) {
	s.logger.Print("calcsvc.add")
	// return
	return p.A + p.B, nil
}

package maths

import "context"

// IMathService interface describing the expectation from math services
type IMathService interface {
	Calculate(ctx context.Context, req *Request) float64
}

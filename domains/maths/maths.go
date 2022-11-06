package maths

import "context"

type MathService struct {
	sectorID float64
}

// New returns a new instance of math service
func New(sectorID float64) IMathService {
	return &MathService{
		sectorID: sectorID,
	}
}

// Calculate this takes all the parameters in the request and uses it to
// perform the complex maths needed to serve a response.
// x*SectorID + y*SectorID + z*SectorID + vel
func (ms *MathService) Calculate(ctx context.Context, req *Request) float64 {
	return req.CoordX*ms.sectorID + req.CoordY*ms.sectorID + req.CoordZ*ms.sectorID + req.Velocity
}

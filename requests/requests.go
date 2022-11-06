package requests

import (
	"dns/domains/maths"
	"fmt"
	"strconv"
)

type Request struct {
	CoordX   string `json:"x"`
	CoordY   string `json:"y"`
	CoordZ   string `json:"z"`
	Velocity string `json:"vel"`
}

func (r *Request) ToDomainRequest() (*maths.Request, error) {
	xCoord, err := strconv.ParseFloat(r.CoordX, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid X coordinate value: (%s)", r.CoordX)
	}

	yCoord, err := strconv.ParseFloat(r.CoordY, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid Y coordinate value %s", r.CoordY)
	}

	zCoord, err := strconv.ParseFloat(r.CoordZ, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid Z coordinate value %s", r.CoordZ)
	}

	vel, err := strconv.ParseFloat(r.Velocity, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid velocity value %s", r.Velocity)
	}

	return &maths.Request{
		CoordX:   xCoord,
		CoordY:   yCoord,
		CoordZ:   zCoord,
		Velocity: vel,
	}, nil
}

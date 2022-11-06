package maths

// Request construct for the expected request for performing the calculations.
type Request struct {
	CoordX   float64 `json:"x"`
	CoordY   float64 `json:"y"`
	CoordZ   float64 `json:"z"`
	Velocity float64 `json:"vel"`
}

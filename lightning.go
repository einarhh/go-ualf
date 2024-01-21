package ualf

import (
	"time"
)

type Lightning struct {
	Version          *int      `json:"version,omitempty"`
	Timestamp        time.Time `json:"timestamp"`
	Lat              float64   `json:"lat"`
	Lon              float64   `json:"lon"`
	PeakCurrent      int       `json:"peakcurrent"`
	Multiplicity     int       `json:"multiplicity"`
	Sensors          int       `json:"sensors"`
	DegreesOfFreedom int       `json:"degreesOfFreedom"`
	Angle            float32   `json:"angle"`
	SemiMajorAxis    float32   `json:"semiMajorAxis"`
	SemiMinorAxis    float32   `json:"semiMinorAxis"`
	ChiSquare        float32   `json:"chiSquare"`
	RiseTime         float32   `json:"riseTime"`
	PeakToZeroTime   float32   `json:"peakToZeroTime"`
	MaxRateOfRise    float32   `json:"maxRateOfRise"`
	CloudIndicator   bool      `json:"cloudIndicator"`
	AngleIndicator   bool      `json:"angleIndicator"`
	SignalIndicator  bool      `json:"signalIndicator"`
	TimingIndicator  bool      `json:"timingIndicator"`
}

package ualf

import (
	"testing"
)

const UALFString = "2024 01 20 17 07 34 877390848 63.8099 2.6209 30 0 11 17 118.80 0.50 0.40 7.44 11.8 7.5 3.0 0 1 0 1"
const UALFStringVersion = "0 2024 01 20 17 07 34 877390848 63.8099 2.6209 30 0 11 17 118.80 0.50 0.40 7.44 11.8 7.5 3.0 0 1 0 1"

func TestParseNoVersion(t *testing.T) {
	lightning, err := ParseUALF(UALFString)
	if err != nil {
		t.Fatalf("failed to parse UALF string: %v", err)
	}
	if lightning.Version != nil {
		t.Errorf("expected version 0, got %d", lightning.Version)
	}
	if lightning.Timestamp.Year() != 2024 {
		t.Errorf("expected year 2024, got %d", lightning.Timestamp.Year())
	}
	if lightning.Timestamp.Month() != 1 {
		t.Errorf("expected month 1, got %d", lightning.Timestamp.Month())
	}
	if lightning.Timestamp.Day() != 20 {
		t.Errorf("expected day 20, got %d", lightning.Timestamp.Day())
	}
	if lightning.Timestamp.Hour() != 17 {
		t.Errorf("expected hour 17, got %d", lightning.Timestamp.Hour())
	}
	if lightning.Timestamp.Minute() != 7 {
		t.Errorf("expected minute 7, got %d", lightning.Timestamp.Minute())
	}
	if lightning.Timestamp.Second() != 34 {
		t.Errorf("expected second 34, got %d", lightning.Timestamp.Second())
	}
	if lightning.Timestamp.Nanosecond() != 877390848 {
		t.Errorf("expected nanosecond 877390848, got %d", lightning.Timestamp.Nanosecond())
	}
	if lightning.Lat != 63.8099 {
		t.Errorf("expected latitude 63.8099, got %f", lightning.Lat)
	}
	if lightning.Lon != 2.6209 {
		t.Errorf("expected longitude 2.6209, got %f", lightning.Lon)
	}
	if lightning.PeakCurrent != 30 {
		t.Errorf("expected peak current 30, got %d", lightning.PeakCurrent)
	}
	if lightning.Multiplicity != 0 {
		t.Errorf("expected multiplicity 0, got %d", lightning.Multiplicity)
	}
	if lightning.Sensors != 11 {
		t.Errorf("expected sensors 11, got %d", lightning.Sensors)
	}
	if lightning.DegreesOfFreedom != 17 {
		t.Errorf("expected degrees of freedom 17, got %d", lightning.DegreesOfFreedom)
	}
	if lightning.Angle != 118.80 {
		t.Errorf("expected angle 118.80, got %f", lightning.Angle)
	}
	if lightning.SemiMajorAxis != 0.50 {
		t.Errorf("expected semi-major axis 0.50, got %f", lightning.SemiMajorAxis)
	}
	if lightning.SemiMinorAxis != 0.40 {
		t.Errorf("expected semi-minor axis 0.40, got %f", lightning.SemiMinorAxis)
	}
	if lightning.ChiSquare != 7.44 {
		t.Errorf("expected chi-square 7.44, got %f", lightning.ChiSquare)
	}
	if lightning.RiseTime != 11.8 {
		t.Errorf("expected rise time 11.8, got %f", lightning.RiseTime)
	}
	if lightning.PeakToZeroTime != 7.5 {
		t.Errorf("expected peak-to-zero time 7.5, got %f", lightning.PeakToZeroTime)
	}
	if lightning.MaxRateOfRise != 3.0 {
		t.Errorf("expected max rate of rise 3.0, got %f", lightning.MaxRateOfRise)
	}
	if lightning.CloudIndicator != false {
		t.Errorf("expected cloud indicator false, got %t", lightning.CloudIndicator)
	}
	if lightning.AngleIndicator != true {
		t.Errorf("expected angle indicator true, got %t", lightning.AngleIndicator)
	}
	if lightning.SignalIndicator != false {
		t.Errorf("expected signal indicator false, got %t", lightning.SignalIndicator)
	}
	if lightning.TimingIndicator != true {
		t.Errorf("expected timing indicator true, got %t", lightning.TimingIndicator)
	}
}

func TestParseVersion(t *testing.T) {
	lightning, err := ParseUALF(UALFStringVersion)
	if err != nil {
		t.Fatalf("failed to parse UALF string: %v", err)
	}
	if *lightning.Version != 0 {
		t.Errorf("expected version 0, got %d", lightning.Version)
	}
	if lightning.Timestamp.Year() != 2024 {
		t.Errorf("expected year 2024, got %d", lightning.Timestamp.Year())
	}
	if lightning.Timestamp.Month() != 1 {
		t.Errorf("expected month 1, got %d", lightning.Timestamp.Month())
	}
	if lightning.Timestamp.Day() != 20 {
		t.Errorf("expected day 20, got %d", lightning.Timestamp.Day())
	}
	if lightning.Timestamp.Hour() != 17 {
		t.Errorf("expected hour 17, got %d", lightning.Timestamp.Hour())
	}
	if lightning.Timestamp.Minute() != 7 {
		t.Errorf("expected minute 7, got %d", lightning.Timestamp.Minute())
	}
	if lightning.Timestamp.Second() != 34 {
		t.Errorf("expected second 34, got %d", lightning.Timestamp.Second())
	}
	if lightning.Timestamp.Nanosecond() != 877390848 {
		t.Errorf("expected nanosecond 877390848, got %d", lightning.Timestamp.Nanosecond())
	}
	if lightning.Lat != 63.8099 {
		t.Errorf("expected latitude 63.8099, got %f", lightning.Lat)
	}
	if lightning.Lon != 2.6209 {
		t.Errorf("expected longitude 2.6209, got %f", lightning.Lon)
	}
	if lightning.PeakCurrent != 30 {
		t.Errorf("expected peak current 30, got %d", lightning.PeakCurrent)
	}
	if lightning.Multiplicity != 0 {
		t.Errorf("expected multiplicity 0, got %d", lightning.Multiplicity)
	}
	if lightning.Sensors != 11 {
		t.Errorf("expected sensors 11, got %d", lightning.Sensors)
	}
	if lightning.DegreesOfFreedom != 17 {
		t.Errorf("expected degrees of freedom 17, got %d", lightning.DegreesOfFreedom)
	}
	if lightning.Angle != 118.80 {
		t.Errorf("expected angle 118.80, got %f", lightning.Angle)
	}
	if lightning.SemiMajorAxis != 0.50 {
		t.Errorf("expected semi-major axis 0.50, got %f", lightning.SemiMajorAxis)
	}
	if lightning.SemiMinorAxis != 0.40 {
		t.Errorf("expected semi-minor axis 0.40, got %f", lightning.SemiMinorAxis)
	}
	if lightning.ChiSquare != 7.44 {
		t.Errorf("expected chi-square 7.44, got %f", lightning.ChiSquare)
	}
	if lightning.RiseTime != 11.8 {
		t.Errorf("expected rise time 11.8, got %f", lightning.RiseTime)
	}
	if lightning.PeakToZeroTime != 7.5 {
		t.Errorf("expected peak-to-zero time 7.5, got %f", lightning.PeakToZeroTime)
	}
	if lightning.MaxRateOfRise != 3.0 {
		t.Errorf("expected max rate of rise 3.0, got %f", lightning.MaxRateOfRise)
	}
	if lightning.CloudIndicator != false {
		t.Errorf("expected cloud indicator false, got %t", lightning.CloudIndicator)
	}
	if lightning.AngleIndicator != true {
		t.Errorf("expected angle indicator true, got %t", lightning.AngleIndicator)
	}
	if lightning.SignalIndicator != false {
		t.Errorf("expected signal indicator false, got %t", lightning.SignalIndicator)
	}
	if lightning.TimingIndicator != true {
		t.Errorf("expected timing indicator true, got %t", lightning.TimingIndicator)
	}
}

func TestParseUALFInvalid(t *testing.T) {
	_, err := ParseUALF("Non UALF string")
	if err == nil {
		t.Error("expected error, got nil")
	}
}

package ualf

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseUALF parses a UALF string into a Lightning struct.
func ParseUALF(line string) (Lightning, error) {
	z := strings.Split(line, " ")
	if len(z) < 24 || len(z) > 25 {
		return Lightning{}, fmt.Errorf("expected 24-25 fields, got %d", len(z))
	}

	lightning := Lightning{}
	startIndex := 0

	// Version is optional
	if len(z) == 25 {
		version, err := strconv.Atoi(z[0])
		if err != nil {
			return Lightning{}, fmt.Errorf("failed to convert version to integer: %v", err)
		}
		lightning.Version = &version
		startIndex = 1
	}

	// Parse timestamp
	year, err := strconv.Atoi(z[startIndex])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert year to integer: %v", err)
	}
	month, err := strconv.Atoi(z[startIndex+1])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert month to integer: %v", err)
	}
	day, err := strconv.Atoi(z[startIndex+2])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert day to integer: %v", err)
	}
	hour, err := strconv.Atoi(z[startIndex+3])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert hour to integer: %v", err)
	}
	minute, err := strconv.Atoi(z[startIndex+4])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert minute to integer: %v", err)
	}
	second, err := strconv.Atoi(z[startIndex+5])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert second to integer: %v", err)
	}
	nanosecond, err := strconv.Atoi(z[startIndex+6])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert nanosecond to integer: %v", err)
	}
	lightning.Timestamp = time.Date(year, time.Month(month), day, hour, minute, second, nanosecond, time.UTC)

	lightning.Lat, err = strconv.ParseFloat(z[startIndex+7], 64)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert lat to float: %v", err)
	}

	lightning.Lon, err = strconv.ParseFloat(z[startIndex+8], 64)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert lon to float: %v", err)
	}

	lightning.PeakCurrent, err = strconv.Atoi(z[startIndex+9])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert peakCurrent to integer: %v", err)
	}

	lightning.Multiplicity, err = strconv.Atoi(z[startIndex+10])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert multiplicity to integer: %v", err)
	}

	lightning.Sensors, err = strconv.Atoi(z[startIndex+11])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert sensors to integer: %v", err)
	}

	lightning.DegreesOfFreedom, err = strconv.Atoi(z[startIndex+12])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert degreesOfFreedom to integer: %v", err)
	}

	angle, err := strconv.ParseFloat(z[startIndex+13], 32)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert angle to float: %v", err)
	}
	lightning.Angle = float32(angle)

	semiMajorAxis, err := strconv.ParseFloat(z[startIndex+14], 32)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert semiMajorAxis to float: %v", err)
	}
	lightning.SemiMajorAxis = float32(semiMajorAxis)

	semiMinorAxis, err := strconv.ParseFloat(z[startIndex+15], 32)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert semiMinorAxis to float: %v", err)
	}
	lightning.SemiMinorAxis = float32(semiMinorAxis)

	chiSquare, err := strconv.ParseFloat(z[startIndex+16], 32)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert chiSquare to float: %v", err)
	}
	lightning.ChiSquare = float32(chiSquare)

	riseTime, err := strconv.ParseFloat(z[startIndex+17], 32)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert riseTime to float: %v", err)
	}
	lightning.RiseTime = float32(riseTime)

	peakToZeroTime, err := strconv.ParseFloat(z[startIndex+18], 32)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert peakToZeroTime to float: %v", err)
	}
	lightning.PeakToZeroTime = float32(peakToZeroTime)

	maxRateOfRise, err := strconv.ParseFloat(z[startIndex+19], 32)
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert maxRateOfRise to float: %v", err)
	}
	lightning.MaxRateOfRise = float32(maxRateOfRise)

	lightning.CloudIndicator, err = strconv.ParseBool(z[startIndex+20])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert cloudIndicator to bool: %v", err)
	}

	lightning.AngleIndicator, err = strconv.ParseBool(z[startIndex+21])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert angleIndicator to bool: %v", err)
	}

	lightning.SignalIndicator, err = strconv.ParseBool(z[startIndex+22])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert signalIndicator to bool: %v", err)
	}

	lightning.TimingIndicator, err = strconv.ParseBool(z[startIndex+23])
	if err != nil {
		return Lightning{}, fmt.Errorf("failed to convert timingIndicator to bool: %v", err)
	}

	return lightning, nil
}

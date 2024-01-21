# UALF Parser for Go

## Overview

UALF Parser is a Go library designed to parse lightning data formatted in the Universal ASCII Lightning Format (UALF). It provides an efficient and user-friendly way to convert UALF strings into structured `Lightning` structs in Go, making it easier for developers to integrate lightning data into their applications.

## Installation

To install the UALF Parser, use the `go get` command:

```bash
go get github.com/einarhh/go-ualf
```

## Usage

To use the UALF parser, import it into your Go project and call the ParseUALF method with a UALF-formatted string.

Example:

```go
package main

import (
    "fmt"
    "github.com/einarhh/go-ualf"
)

func main() {
    ualfData := "2024 01 20 17 07 34 877390848 63.8099 2.6209 30 0 11 17 118.80 0.50 0.40 7.44 11.8 7.5 3.0 0 1 0 1"
    lightningData, err := ualf.ParseUALF(ualfData)
    if err != nil {
        fmt.Println("Error parsing UALF data:", err)
        return
    }

    fmt.Printf("Parsed Lightning Data: %+v\n", lightningData)
}

```

## The Lightning Struct

The Lightning struct is the core data structure used by UALF Parser.
It includes fields corresponding to the [UALF specification](docs/UALF.md).

The Version field, while not specified in the UALF standard, is included in the data obtained from frost.met.no. In this parser, we treat Version as an optional field. It is omitted if the string's first field represents the year.

```go
type Lightning struct {
	Version          int       `json:"version"`
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
```

## Acknowledgments

-   SMHI Open Data API Docs - Lightning Archive
-   The Norwegian Meteorological institute

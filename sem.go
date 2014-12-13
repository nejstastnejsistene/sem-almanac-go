// http://www.navcen.uscg.gov/?pageName=gpsSem
package sem

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	numGlobalFields       = 4
	numFieldsPerSatellite = 14
)

var MissingFields = errors.New("missing almanac fields")

type Almanac struct {
	NumRecords int
	Title      string
	WeekNumber int
	TOA        int
	Records    []SatelliteRecord
}

type SatelliteRecord struct {
	PRN                        int
	SVN                        int
	AverageURA                 int
	Eccentricity               float64
	InclinationOffset          float64
	RateOfRightAscension       float64
	SqrtOfSemiMajorAxis        float64
	LongitudeOfOrbitalPlane    float64
	ArgumentOfPerigee          float64
	MeanAnomaly                float64
	ZerothOrderClockCorrection float64
	FirstOrderClockCorrection  float64
	SatelliteHealth            int
	SatelliteConfiguration     int
}

func Unmarshal(buf []byte) (al *Almanac, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	fields := strings.Fields(string(buf))
	if len(fields) < numGlobalFields {
		err = MissingFields
		return
	}

	al = new(Almanac)
	al.NumRecords = parseInt(fields[0])
	al.Title = fields[1]
	al.WeekNumber = parseInt(fields[2])
	al.TOA = parseInt(fields[3])

	if len(fields) != numGlobalFields+al.NumRecords*numFieldsPerSatellite {
		err = MissingFields
		return
	}

	al.Records = make([]SatelliteRecord, al.NumRecords)
	for i := 0; i < al.NumRecords; i++ {
		offset := numGlobalFields + i*numFieldsPerSatellite
		al.Records[i].PRN = parseInt(fields[offset+0])
		al.Records[i].SVN = parseInt(fields[offset+1])
		al.Records[i].AverageURA = parseInt(fields[offset+2])
		al.Records[i].Eccentricity = parseFloat(fields[offset+3])
		al.Records[i].InclinationOffset = parseFloat(fields[offset+4])
		al.Records[i].RateOfRightAscension = parseFloat(fields[offset+5])
		al.Records[i].SqrtOfSemiMajorAxis = parseFloat(fields[offset+6])
		al.Records[i].LongitudeOfOrbitalPlane = parseFloat(fields[offset+7])
		al.Records[i].ArgumentOfPerigee = parseFloat(fields[offset+8])
		al.Records[i].MeanAnomaly = parseFloat(fields[offset+9])
		al.Records[i].ZerothOrderClockCorrection = parseFloat(fields[offset+10])
		al.Records[i].FirstOrderClockCorrection = parseFloat(fields[offset+11])
		al.Records[i].SatelliteHealth = parseInt(fields[offset+12])
		al.Records[i].SatelliteConfiguration = parseInt(fields[offset+13])
	}

	return
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parseFloat(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return n
}

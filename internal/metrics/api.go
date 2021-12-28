package metrics

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type Metric float64

type Data struct {
	Height string `json:"height"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
	Sex    string  `json:"sex"`
	Level  float64 `json:"level"`
}


func (d *Data) String() string {
	out, _ := json.MarshalIndent(d, "", "  ")
	return string(out)
}


const (
	CMsInInch  float64 = 2.54
	KGsInPound float64 = 0.45359237
)

const (
	WeightPounds    = "lbs"
	WeightKilograms = "kg"
)

const (
	HeightInches = "in"
	HeightCm     = "cm"
)

func InchesToCM(inches float64) float64 {
	return inches * CMsInInch
}

func CMToInches(cm float64) float64 {
	return cm / CMsInInch
}

func PoundsToKG(lbs float64) float64 {
	return lbs * KGsInPound
}

func KGToPounds(kg float64) float64 {
	return kg / KGsInPound
}

// ParseHeight reads from string input, looking for
// metric markers `cm` or `in`. The metric is stripped off
// and the resulting number is parsed as a float, and converted
// to cm if needed. If no metric designation is found, the
// number is assumed to be in cm.
func ParseHeight(height string) (float64, error) {
	if strings.Contains(height, HeightInches) {
		h, err := strconv.ParseFloat(
			strings.ReplaceAll(height, HeightInches, ""),
			64,
		)
		if err != nil {
			return 0, err
		}

		return InchesToCM(h), nil
	}

	if strings.Contains(height, HeightCm) {
		return strconv.ParseFloat(
			strings.ReplaceAll(height, HeightCm, ""),
			64,
		)
	}

	if h, err := strconv.ParseFloat(height, 64); err == nil {
		return h, nil
	}

	return 0, errors.New("invalid height")
}

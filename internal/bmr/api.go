package bmr

type Sex int

const (
	Male   Sex = 5
	Female Sex = -161
)

type Workload float64

const (
	Sedintary       Workload = 1.2
	OneTwoPerWeek   Workload = 1.4
	TwoThreePerWeek Workload = 1.5
	FourFivePerWeek Workload = 1.6
	SixSevenPerWeek Workload = 1.7
	TwoPerDay       Workload = 1.8
)

func (w Workload) AsFloat() float64 {
	return float64(w)
}

func MifflinStJeor(weight, height, age float64, s Sex) float64 {
	return (10 * weight) + (6.25 * height) - (5 * age) + float64(s)
}

func KatchMcArdle(weight, bodyFat float64) float64 {
	leanBodyMass := weight * ((100 - bodyFat) / 100)
	return 370 + (21.6 * leanBodyMass)
}

func TDEE(bmr float64, work Workload) float64 {
	return bmr * work.AsFloat()
}

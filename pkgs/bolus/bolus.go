package bolus

import "strconv"

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	idealGlucose         = (lowerGlucose + upperGlucose) / 2
)

type Glucometr struct {
	bolus   float64
	glucose float64
	bUnit   float64
	sensiti float64 // Чувствительность к инсулину
	carb    float64 // Углеводный коэффициент (ед / 1хе)
}

func (g *Glucometr) bolusForFood() {
	g.bolus = g.carb * g.bUnit
}

func (g *Glucometr) bolusForCorrect() {
	g.bolus = (g.glucose - idealGlucose) / g.sensiti
}

func (g *Glucometr) fullBolus() {
	g.bolusForFood()
	tempBolus := g.bolus
	g.bolusForCorrect()
	g.bolus += tempBolus
}

func (g *Glucometr) setBolus() {
	g.coeffs()
	switch {
	case g.glucose == 0:
		g.bolusForFood()
	case g.bUnit == 0:
		g.bolusForCorrect()
	default:
		g.fullBolus()
	}
}

func SetGlucometr(glucose, bUnit string) float64 {
	g := Glucometr{}
	g.glucose, _ = strconv.ParseFloat(glucose, 64)
	g.bUnit, _ = strconv.ParseFloat(bUnit, 64)
	g.setBolus()
	return g.bolus
}
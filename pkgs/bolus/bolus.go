package bolus

import (
	"fmt"
	"strconv"
)

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	idealGlucose         = (lowerGlucose + upperGlucose) / 2
)

type glucometr struct {
	bolus   float64
	glucose float64
	bUnit   float64
	sensiti float64 // Чувствительность к инсулину
	carb    float64 // Углеводный коэффициент (ед / 1хе)
}

func (g *glucometr) bolusForFood() {
	g.bolus = g.carb * g.bUnit
}

func (g *glucometr) bolusForCorrect() {
	g.bolus = (g.glucose - idealGlucose) / g.sensiti
}

func (g *glucometr) fullBolus() {
	g.bolusForFood()
	tempBolus := g.bolus
	g.bolusForCorrect()
	g.bolus += tempBolus
}

func (g *glucometr) setBolus() {
	g.coeffs()
	switch {
	case g.glucose == 0:
		g.bolusForFood()
	case g.bUnit == 0:
		g.bolusForCorrect()
	default:
		g.fullBolus()
	}
	if g.bolus < 0 {
		g.bolus = 0
	}
	strBolus := fmt.Sprintf("%.1f", g.bolus)
	g.bolus, _ = strconv.ParseFloat(strBolus, 64)
}


func SetGlucometr(glucose, bUnit float64) float64 {
	g := glucometr{}
	g.glucose = glucose
	g.bUnit = bUnit
	g.setBolus()
	return g.bolus
}

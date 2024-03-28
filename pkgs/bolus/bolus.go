package bolus

import (
	"fmt"
	"strconv"

	"github.com/Starc151/dia.v2/pkgs/ydb"
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

func (g *glucometr) insert() {
	glucometrParams := make(map[string]float64)
	glucometrParams["glucose"] = g.glucose
	glucometrParams["bUnit"] = g.bUnit
	glucometrParams["bolus"] = g.bolus

	connectDB := ydb.Connected{}
	connectDB.Insert(glucometrParams)
}

func SetGlucometr(glucose, bUnit string) (string, error) {
	g := glucometr{}
	g.glucose, _ = strconv.ParseFloat(glucose, 64)
	g.bUnit, _ = strconv.ParseFloat(bUnit, 64)
	g.setBolus()
	g.insert()
	return fmt.Sprintf("Bolus: %.1f", g.bolus), nil
}

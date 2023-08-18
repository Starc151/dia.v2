package bolus

const (
	lowerGlucose float64 = 7.0
	upperGlucose float64 = 9.0
	idealGlucose = (lowerGlucose + upperGlucose) / 2
)

type Glucometr struct {
	bolus float64
	glucose float64
	bUnit float64
	sensiti float64 // Чувствительность к инсулину
	carb float64  // Углеводный коэффициент (ед / 1хе)
}

func (g *Glucometr) bolusForFood() {
	g.bolus = g.carb * g.bUnit
}

func (g *Glucometr) bolusForCorrect() {
	g.bolus = (g.glucose - idealGlucose) / g.sensiti
}

func (g *Glucometr) fullBolus() {
	g.bolusForFood()
	bolus := g.bolus
	g.bolusForCorrect()
	g.bolus += bolus
}

func (g *Glucometr) Bolus(glucose, bUnit float64) float64 {
	g.coeffs()
	g.glucose = glucose
	g.bUnit = bUnit
	switch {
	case glucose == 0:
		g.bolusForFood()
	case bUnit == 0:
		g.bolusForCorrect()
	default:
		g.fullBolus()
	}
	return g.bolus
}
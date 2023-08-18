package bolus

import "time"

func (g *Glucometr) coeffs() {
	timeNow := time.Now().Format("15:04")
	if "00:00" <= timeNow && timeNow <= "10:30"{
		g.sensiti  = 4.0 
		g.carb = 1.25
	} else if "10:31" <= timeNow && timeNow <= "13:00"{
		g.sensiti  = 3.5
		g.carb = 0.8
	} else if "13:01" <= timeNow && timeNow <= "18:00"{
		g.sensiti  = 3.5
		g.carb = 1
	} else if "18:01" <= timeNow && timeNow <= "23:59"{
		g.sensiti  = 3.0
		g.carb = 1
	}
}

// compilation for Android
// fyne package -os android -appID ru.dia.android -icon img/icon.png

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Starc151/dia.v2/pkgs/apk"
)

func main() {
	app := app.New()
	icon, _ := fyne.LoadResourceFromPath("img/icon.png")
	app.SetIcon(icon)

	apk.NewGlucometr(app)
	app.Run()
}
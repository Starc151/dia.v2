// compilation for Android
// fyne package -os android -appID ru.dia.android -icon img/icon.png

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Dia")
	icon, _ := fyne.LoadResourceFromPath("img/icon.png")
	app.SetIcon(icon)
	
	window.ShowAndRun()
}
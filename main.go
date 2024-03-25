// compilation for Android
// fyne package -os android -icon img/icon.png -appID ru.dia.android

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Starc151/dia.v2/pkgs/apk"
)
// func init() {
// 	loc, _ := time.LoadLocation("Europe/Moscow")
// 	time.Local = loc
// }

func main() {
	app := app.New()
	icon, _ := fyne.LoadResourceFromPath("img/icon.png")
	app.SetIcon(icon)

	apk.NewGlucometr(app)
	app.Run()
}
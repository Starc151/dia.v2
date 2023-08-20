package apk

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Apk struct {
	bolus *widget.Label
	
}

func (a *Apk) te() {
	i := widget.NewEntry()
	i.Keyboard()
}

func (a *Apk) LoadApk() *fyne.Container {
	a.bolus = &widget.Label{}
	cont := container.NewVBox(
		container.NewHBox(
			a.digitalBtn(7),
			a.digitalBtn(8),
			a.digitalBtn(9),
		),
		container.NewHBox(
			a.digitalBtn(4),
			a.digitalBtn(5),
			a.digitalBtn(6),
		),
		container.NewHBox(
			a.digitalBtn(1),
			a.digitalBtn(2),
			a.digitalBtn(3),
		),
		container.NewHBox(
			a.addBtn("."),
			a.digitalBtn(0),
			a.addBtn("<"),
		),
		container.NewHBox(
			a.addBtn("C"),
			a.addBtn("ok"),
		),
	)
	return cont
}
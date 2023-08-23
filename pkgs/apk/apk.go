package apk

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Apk struct {
	entry  *canvas.Text
	btn *widget.Button
	window  fyne.Window
	glucometr struct {
		glucose *canvas.Text
		bUint *canvas.Text
		bolus *canvas.Text
	}
}

func (a *Apk) character(str string) {
	// a.display(c.equation + string(char))
}

func (a *Apk) digitBtn(number int) *widget.Button {
	str := strconv.Itoa(number)
	return a.addBtn(str, func() {
		a.character(str)
	})
}

func (a *Apk) charBtn(str string) *widget.Button {
	return a.addBtn(str, func() {
		a.character(str)
	})
}

func (a *Apk) addBtn(text string, action func()) *widget.Button {
	a.btn = &widget.Button{}
	a.btn.Text = text
	a.btn.OnTapped = action
	return a.btn
}

func setAlign (position string) int {
	positionNum := 1 
	switch position {
	case "left":
		positionNum = 0
	case "right":
		positionNum = 2
	}
	return positionNum
}

func (a *Apk) setCanvasText(field *canvas.Text, text, position string) *canvas.Text {
	field = &canvas.Text{Color: color.White, TextSize: 25}
	field.Text = text
	field.Alignment = fyne.TextAlign(setAlign(position))
	return field
}

func (a *Apk) loadApk() {
	border := widget.NewLabel("  ")

	tabs := container.NewAppTabs(
		container.NewTabItem("BOLUS",
			container.NewBorder(nil, nil, border, border,
				container.NewGridWithColumns(1,
					container.NewGridWithColumns(2,
						a.input("Entry:", "left"),
						a.input("_", "right")),
					container.NewGridWithColumns(2,
						a.input("Glucose:", "left"),
						a.input("0", "right")),
					container.NewGridWithColumns(2,
						a.input("B Unit:", "left"),
						a.input("0", "right")),
					container.NewGridWithColumns(1,
						a.input("BOLUS: 3", "central")),
					container.NewGridWithColumns(3,
						a.charBtn("GL"),
						a.charBtn("BU"),
						a.charBtn("OK")),
					container.NewGridWithColumns(3,
						a.digitBtn(7),
						a.digitBtn(8),
						a.digitBtn(9)),
					container.NewGridWithColumns(3,
						a.digitBtn(4),
						a.digitBtn(5),
						a.digitBtn(6)),
					container.NewGridWithColumns(3,
						a.digitBtn(1),
						a.digitBtn(2),
						a.digitBtn(3)),
					container.NewGridWithColumns(3,
						a.charBtn("<"),
						a.digitBtn(0),
						a.charBtn(".,")),
				),
			),
		),
	)

	a.window.SetContent(tabs)
}

func NewGlucometr(app fyne.App) {
	a := Apk{}
	a.window = app.NewWindow("Dia")
	a.glucometr.bUint = a.setCanvasText()
	a.glucometr.glucose = &canvas.Text{}
	a.glucometr.bolus = &canvas.Text{}

	a.loadApk()
	a.window.Show()
}
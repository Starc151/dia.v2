package apk

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Apk struct {
	entry     *canvas.Text
	btn       *widget.Button
	window    fyne.Window
	glucometr struct {
		glucose *canvas.Text
		bUint   *canvas.Text
		bolus   *canvas.Text
	}
}

func (a *Apk) digitBtn(number int) *widget.Button {
	str := strconv.Itoa(number)
	return a.addBtn(str, func() {
		a.setEntry(str)
	})
}

func (a *Apk) setGlucometr(text string, field *canvas.Text) {
	field.Text = text
	field.Refresh()
	a.entry.Text = "_"
	a.entry.Refresh()
}

func (a *Apk) addBtn(text string, action func()) *widget.Button {
	a.btn = &widget.Button{}
	a.btn.Text = text
	a.btn.OnTapped = action
	return a.btn
}

func (a *Apk) setEntry(text string) {
	if a.entry.Text == "_" {
		a.entry.Text = ""
	}
	if len(a.entry.Text) < 4 {
		a.entry.Text += text
	}
	a.entry.Refresh()
}

func (a *Apk) clear() {
	a.entry.Text = "_"
	a.glucometr.glucose.Text = "0"
	a.glucometr.bUint.Text = "0"
	a.glucometr.bolus.Text = ""
	a.entry.Refresh()
	a.glucometr.glucose.Refresh()
	a.glucometr.bUint.Refresh()
	a.glucometr.bolus.Refresh()
}

func (a *Apk) backSpace() {
	len := len(a.entry.Text)
	if len == 1 {
		a.entry.Text = "_"
		a.entry.Refresh()
		return
	}
	a.entry.Text = a.entry.Text[:len-1]
	a.entry.Refresh()
}

func setAlign(position string) int {
	positionNum := 1
	switch position {
	case "left":
		positionNum = 0
	case "right":
		positionNum = 2
	}
	return positionNum
}

func (a *Apk) setCanvasText(text, position string) *canvas.Text {
	canvasText := &canvas.Text{Color: color.White, TextSize: 25}
	canvasText.Text = text
	canvasText.Alignment = fyne.TextAlign(setAlign(position))
	return canvasText
}

func (a *Apk) getBolus() {
	a.glucometr.bolus.Text = "Bolus OK"
}

func NewGlucometr(app fyne.App) {
	a := Apk{}
	a.window = app.NewWindow("Dia")
	a.entry = a.setCanvasText("_", "right")
	a.glucometr.bUint = a.setCanvasText("0", "right")
	a.glucometr.glucose = a.setCanvasText("0", "right")
	a.glucometr.bolus = a.setCanvasText("", "centr")

	a.loadApk()
	a.window.Show()
}

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

func (a *Apk) digitBtn(number int) *widget.Button {
	str := strconv.Itoa(number)
	return a.addBtn(str, func() {
		a.setEntry(str)
	})
}

func (a *Apk) charBtn(text string) {
	a.glucometr.glucose.Text = text
	a.glucometr.glucose.Refresh()
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
	a.entry.Refresh()
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

func (a *Apk) setCanvasText(text, position string) *canvas.Text {
	canvasText := &canvas.Text{Color: color.White, TextSize: 25}
	canvasText.Text = text
	canvasText.Alignment = fyne.TextAlign(setAlign(position))
	return canvasText
}

func (a *Apk) loadApk() {
	border := widget.NewLabel("  ")

	tabs := container.NewAppTabs(
		container.NewTabItem("BOLUS",
			container.NewBorder(nil, nil, border, border,
				container.NewGridWithColumns(1,
					container.NewGridWithColumns(2,
						a.setCanvasText("Entry:", "left"),
						a.entry),
					container.NewGridWithColumns(2,
						a.setCanvasText("Glucose:", "left"),
						a.glucometr.glucose),
					container.NewGridWithColumns(2,
						a.setCanvasText("B Unit:", "left"),
						a.glucometr.bUint),
					container.NewGridWithColumns(1,
						a.glucometr.bolus),
					container.NewGridWithColumns(3,
						a.addBtn("GL", func() {a.charBtn(a.glucometr.glucose.Text)}),
						a.addBtn("GL", func() {a.charBtn(a.glucometr.glucose.Text)}),
						a.addBtn("C", func() {a.clear()})),
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
						a.addBtn("<", func() {a.backSpace()}),
						a.digitBtn(0),
						a.addBtn(".", func() {a.setEntry(".")})),						
					// container.NewGridWithColumns(1,
					// 	a.charBtn("GET BOLUS")),
				),
			),
		),
	)

	a.window.SetContent(tabs)
}

func NewGlucometr(app fyne.App) {
	a := Apk{}
	a.window = app.NewWindow("Dia")
	a.entry = a.setCanvasText("_", "right")
	a.glucometr.bUint = a.setCanvasText("0", "right")
	a.glucometr.glucose = a.setCanvasText("0", "right")
	a.glucometr.bolus = a.setCanvasText("", "right")


	a.loadApk()
	a.window.Show()
}
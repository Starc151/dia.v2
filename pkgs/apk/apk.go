package apk

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Starc151/dia.v2/pkgs/bolus"
	"github.com/Starc151/dia.v2/pkgs/ydb"
)

type glucometr struct {
	entry *canvas.Text
	btn   *widget.Button

	glucose *canvas.Text
	bUnit   *canvas.Text
	bolus   *canvas.Text
	history [][]string

	window fyne.Window
	err    error
}

func (g *glucometr) digitBtn(number int) *widget.Button {
	str := strconv.Itoa(number)
	return g.addBtn(str, func() {
		g.setEntry(str)
	})
}

func (g *glucometr) setGlucometr(text string, field *canvas.Text) {
	if text == "_" {
		text = "0"
	}
	for range text {
		text = strings.ReplaceAll(text, "..", ".")
	}
	if strings.HasPrefix(text, ".") {
		text = "0" + text
	}
	
	text = strings.TrimSuffix(text, ".")
	
	field.Text = text
	field.Refresh()
	g.entry.Text = "_"
	g.entry.Refresh()
}

func (g *glucometr) addBtn(text string, action func()) *widget.Button {
	g.btn = &widget.Button{}
	g.btn.Text = text
	g.btn.OnTapped = action
	return g.btn
}

func (g *glucometr) setEntry(text string) {
	if g.entry.Text == "_" {
		g.entry.Text = ""
	}
	if len(g.entry.Text) < 4 {
		g.entry.Text += text
	}
	g.entry.Refresh()
}

func (g *glucometr) clear(btn... *widget.Button) {
	g.entry.Text = "_"
	g.glucose.Text = "0"
	g.bUnit.Text = "0"
	g.bolus.Text = "0"
	g.entry.Refresh()
	g.glucose.Refresh()
	g.bUnit.Refresh()
	g.bolus.Refresh()
	for i := 0; i < len(btn); i++ {
		btn[i].Enable()
	}
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

func (g *glucometr) setCanvasText(text, position string) *canvas.Text {
	canvasText := &canvas.Text{Color: color.White, TextSize: 25}
	canvasText.Text = text
	canvasText.Alignment = fyne.TextAlign(setAlign(position))
	return canvasText
}

func (g *glucometr) getBolus() {
	glucose := strToFl64(g.glucose.Text)
	bUnit := strToFl64(g.bUnit.Text)
	bolus := bolus.SetGlucometr(glucose, bUnit)

	g.bolus.Text = fmt.Sprintf("%.1f", bolus)
	g.bolus.Refresh()
}

func (g *glucometr) save() {
	glucometrParams := make(map[string]float64, 3)
	glucometrParams["glucose"], _ = strconv.ParseFloat(g.glucose.Text, 64)
	glucometrParams["bUnit"], _ = strconv.ParseFloat(g.bUnit.Text, 64)
	glucometrParams["bolus"], _ = strconv.ParseFloat(g.bolus.Text, 64)

	g.err = ydb.Insert(glucometrParams)
	g.history, g.err = ydb.SelectAll()
	if g.err != nil {
		dialog.ShowError(g.err, g.window)
	}
}

func NewGlucometr(app fyne.App) {
	g := glucometr{}
	g.history, g.err = ydb.SelectAll()
	app.Settings().SetTheme(&myTheme{})
	g.window = app.NewWindow("Dia")
	g.entry = g.setCanvasText("_", "right")
	g.bUnit = g.setCanvasText("0", "right")
	g.glucose = g.setCanvasText("0", "right")
	g.bolus = g.setCanvasText("0", "right")

	g.loadApk()
	g.window.Show()
}

func strToFl64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func (g *glucometr) checkNullParam() bool {
	if g.glucose.Text == "0" && g.bUnit.Text == "0" && g.bolus.Text == "0"{
		return false
	}
	return true
}
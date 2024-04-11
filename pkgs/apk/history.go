package apk

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (g *glucometr) lastHistory() *fyne.Container {
	date := strToCNT(g.history[0][0])
	date.Alignment = fyne.TextAlignCenter
	resBox := container.NewVBox(date)

	for i := 0; i < 4 && i < len(g.history); i++ {
		tempAssay  := strToCNT(fmt.Sprintf("%s", g.history[i][1:3]))
		resBox.Add(tempAssay)
		tempAssay = strToCNT(fmt.Sprintf("%s", g.history[i][3:]))
		resBox.Add(tempAssay)
		resBox.Add(strToCNT(""))
		if g.history[i][0] != g.history[i+1][0] {
			date = strToCNT(g.history[i+1][0])
			date.Alignment = fyne.TextAlignCenter
			resBox.Add(date)
		}
	}
	return resBox
}

func strToCNT(str string) *canvas.Text {
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	cStr := canvas.NewText(str, nil)
	return cStr
}
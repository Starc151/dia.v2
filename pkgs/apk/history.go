package apk

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (g *glucometr) lastHistory() *fyne.Container {
	resBox := container.NewVBox()
	curentDate := ""
	for i := 0; i < 4 && i < len(g.history); i++ {
		date := strToCNT(g.history[i][0])
		date.Alignment = fyne.TextAlignCenter
		if curentDate != date.Text {
			resBox.Add(date)
			curentDate = date.Text
		}
		tempAssay  := strToCNT(fmt.Sprintf("%s", g.history[i][1:3]))
		resBox.Add(tempAssay)
		tempAssay = strToCNT(fmt.Sprintf("%s", g.history[i][3:]))
		resBox.Add(tempAssay)
		resBox.Add(strToCNT(""))
	}
	return resBox
}

func strToCNT(str string) *canvas.Text {
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	cStr := canvas.NewText(str, nil)
	return cStr
}
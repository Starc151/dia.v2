package apk

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (g *glucometr) lastHystory() *fyne.Container {
	fDate := strToCNT(g.hystory[0][0])
	fDate.Alignment = fyne.TextAlignCenter
	tempAssay  := strToCNT(fmt.Sprintf("%s", g.hystory[0][1:3]))
	resBox := container.NewVBox(fDate, tempAssay)
	tempAssay = strToCNT(fmt.Sprintf("%s", g.hystory[0][3:]))
	resBox.Add(tempAssay)
	resBox.Add(strToCNT(""))
	resBox.Add(strToCNT("sdgfdsgsfdg"))

	

	// for i := 1; i <= 10; i++ {
	// 	if g.hystory[i][0] == g.hystory[i-1][0]{
	// 		resBox.Add(strToCNT(fmt.Sprintf("%s", g.hystory[i][1:])))
	// 	}
	// 	resBox.Add(canvas.NewText(g.hystory[i][0], nil))
	// }
	return resBox
}

func strToCNT(str string) *canvas.Text {
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	cStr := canvas.NewText(str, nil)
	return cStr
}
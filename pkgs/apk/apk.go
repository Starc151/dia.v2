package apk

import (
	"strconv"

	"fyne.io/fyne/v2/widget"
)

func (a *Apk) addBtn(text string) *widget.Button {
	btn := widget.NewButton(text, func() {
		a.bolus.SetText(text)
	})
	return btn
}

func (a *Apk) digitalBtn(num int) *widget.Button {
	str := strconv.Itoa(num)
	return a.addBtn(str)
}

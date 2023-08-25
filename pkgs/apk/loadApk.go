package apk

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
						a.addBtn("GL", func() { a.setGlucometr(a.entry.Text, a.glucometr.glucose) }),
						a.addBtn("BU", func() { a.setGlucometr(a.entry.Text, a.glucometr.bUint) }),
						a.addBtn("C", func() { a.clear() })),
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
						a.addBtn("<", func() { a.backSpace() }),
						a.digitBtn(0),
						a.addBtn(".", func() { a.setEntry(".") })),
					container.NewGridWithColumns(1,
						a.addBtn("GET BOLUS", func() {a.getBolus()})),
				),
			),
		),
	)

	a.window.SetContent(tabs)
}

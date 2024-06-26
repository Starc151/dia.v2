package apk

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (g *glucometr) loadApk() {
	historyCont := g.lastHistory()
	historyTab := container.NewTabItem("HISTORY", historyCont)
	getBolusBtn := g.addBtn("GET BOLUS", nil)
	getBolusBtn.OnTapped = func() {
		g.getBolus(getBolusBtn)
		historyCont.RemoveAll()
		historyCont.Add(g.lastHistory())
		historyTab.Content.Refresh()
	}
	border := widget.NewLabel("  ")

	tabs := container.NewAppTabs(
		container.NewTabItem("BOLUS",
			container.NewBorder(nil, nil, border, border,
				container.NewGridWithColumns(1,
					container.NewGridWithColumns(2,
						g.setCanvasText("Entry:", "left"),
						g.entry),
					container.NewGridWithColumns(2,
						g.setCanvasText("Glucose:", "left"),
						g.glucose),
					container.NewGridWithColumns(2,
						g.setCanvasText("B Unit:", "left"),
						g.bUnit),
					container.NewGridWithColumns(1,
						g.bolus),
					container.NewGridWithColumns(3,
						g.addBtn("GL", func() {g.setGlucometr(g.entry.Text, g.glucose) }),
						g.addBtn("BU", func() {g.setGlucometr(g.entry.Text, g.bUnit) }),
						g.addBtn("C", func() {g.clear(getBolusBtn) })),
					container.NewGridWithColumns(3,
						g.digitBtn(7),
						g.digitBtn(8),
						g.digitBtn(9)),
					container.NewGridWithColumns(3,
						g.digitBtn(4),
						g.digitBtn(5),
						g.digitBtn(6)),
					container.NewGridWithColumns(3,
						g.digitBtn(1),
						g.digitBtn(2),
						g.digitBtn(3)),
					container.NewGridWithColumns(3,
						g.addBtn("<", func() { g.backSpace() }),
						g.digitBtn(0),
						g.addBtn(".", func() { g.setEntry(".") })),
					container.NewGridWithColumns(1,
						getBolusBtn),
				),
			),
		),
		historyTab,
	)
	g.window.SetContent(tabs)
}

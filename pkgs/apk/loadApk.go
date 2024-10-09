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
		if !g.checkNullParam(){
			return
		} else {
			g.getBolus()
		}
	}

	saveBtn := g.addBtn("SAVE", nil)
	saveBtn.OnTapped = func() {
		if !g.checkNullParam(){
			return
		} else {
			g.save()
			historyCont.RemoveAll()
			historyCont.Add(g.lastHistory())
			historyTab.Content.Refresh()
			saveBtn.Disable()
			getBolusBtn.Disable()
		}
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
					container.NewGridWithColumns(2,
						g.setCanvasText("Bolus:", "left"),
						g.bolus),
					container.NewGridWithColumns(3,
						g.addBtn("GL", func() {g.setGlucometr(g.entry.Text, g.glucose) }),
						g.addBtn("XE", func() {g.setGlucometr(g.entry.Text, g.bUnit) }),
						g.addBtn("BO", func() {g.setGlucometr(g.entry.Text, g.bolus) }),),
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
						g.addBtn("C", func() {g.clear(getBolusBtn, saveBtn) }),
						g.digitBtn(0),
						g.addBtn(".", func() { g.setEntry(".") })),
					container.NewGridWithColumns(2,
						getBolusBtn,
						saveBtn),
				),
			),
		),
		historyTab,
	)
	g.window.SetContent(tabs)
}

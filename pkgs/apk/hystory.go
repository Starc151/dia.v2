package apk

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/Starc151/dia.v2/pkgs/ydb"
)

func getHystory() *fyne.Container {
	res, _ := ydb.NewConnect("SELECT", nil)
	return container.NewVBox()
}
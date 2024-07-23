package app

import (
	. "SMS/app/dialog"
	"bytes"

	"github.com/lxn/walk"
	// . "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*MultiPageMainWindow
}

// 切换页面同时更新Title
func (mw *MyMainWindow) UpdateTitle(prefix string) {
	var buf bytes.Buffer

	if prefix != "" {
		buf.WriteString(prefix)
		buf.WriteString(" - ")
	}

	buf.WriteString("Walk Multiple Pages Example")

	mw.SetTitle(buf.String())
}

//about弹窗
func (mw *MyMainWindow) aboutAction_Triggered() {
	walk.MsgBox(mw,
		"About Walk Multiple Pages Example",
		"An example that demonstrates a main window that supports multiple pages.",
		walk.MsgBoxOK|walk.MsgBoxIconInformation)
}

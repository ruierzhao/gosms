package dialog

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type SettingDialog struct {
	*walk.MainWindow
}

func NewSettingDialog(owner walk.Form) {
	var settingdlg *walk.Dialog
	var cancelPB *walk.PushButton
	Dialog{
		AssignTo: &settingdlg,
		Title:    "软件设置",
		MinSize:  Size{450, 400},
		Layout:   VBox{},
		Children: []Widget{
			Label{
				Text:    "shzhi",
				MinSize: Size{40, 20},
			},
			Label{
				Text:    "ccccc",
				MinSize: Size{40, 20},
			},
			TextEdit{
				Text: "textEdit",
				OnKeyUp: func(key walk.Key) {
					fmt.Println("按下了", key)
				},
			},
			// Layout:HSpacer{},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						AssignTo: &cancelPB,
						Text:     "取消",
						OnClicked: func() {
							fmt.Println("test priontln")
							settingdlg.Cancel()
						},
					},
					PushButton{

						Text: "确定",
						OnClicked: func() {
							fmt.Println(settingdlg.Bounds())
						},
					},
				},
			},
		},
	}.Run(owner)
}

func (sd *SettingDialog) Hide() {

	fmt.Println("SettingDialog-- ", sd.Bounds(), sd.Children())
	sd.Close()
}

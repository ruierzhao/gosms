package dialog

import (
	"SMS/app/util"
	"fmt"
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type SendMsg struct {
	*walk.Composite
}

func NewSendMsg(parent walk.Container) (Page, error) {
	var (
		mobanMsg, rizhi, haoma *walk.TextEdit
		haomaNumberLabel       *walk.Label
		haomaNumber            int
	)

	sendMsg := new(SendMsg)
	if err := (Composite{
		AssignTo:  &sendMsg.Composite,
		Name:      "sendMsg",
		Layout:    HBox{},
		Alignment: AlignHCenterVNear,
		Children: []Widget{
			// 读取电话号码
			Composite{
				Layout: VBox{},
				Children: []Widget{
					// 限制textedit大小
					Composite{
						Layout: VBox{MarginsZero: true, SpacingZero: true},
						Children: []Widget{
							TextEdit{
								AssignTo:    &haoma,
								Text:        "15288144694",
								ToolTipText: "输入电话号码【一行一个】或者拖入文件",
								Enabled:     true,
								Font:        Font{PointSize: 15},
								// Alignment:   AlignHNearVNear,
								Column:   5,
								Row:      5,
								ReadOnly: false,
								VScroll:  true,
								OnBoundsChanged: func() {
								},
								OnTextChanged: func() {
									// fmt.Println("bounds->", haoma.Bounds())
									// haomaNumberLabel.SetText("写入" + haomaNumber + "个号码")
								},
								OnKeyUp: func(key walk.Key) {
									if key == walk.KeyOEMPlus|walk.KeyControl {
										// control + 放大字体
										newfont, err := walk.NewFont("", haoma.Font().PointSize()+1, 0)
										if err != nil {
											fmt.Println(err)
											return
										}
										haoma.SetFont(newfont)
									}
									// control - 减小字体
									if key == walk.KeyOEMMinus|walk.KeyControl {
										newfont, err := walk.NewFont("", haoma.Font().PointSize()-1, 0)
										if err != nil {
											fmt.Println(err)
											return
										}
										haoma.SetFont(newfont)
									}
									fmt.Println("Up--按下了", key)
								},

								OnSizeChanged: func() {
								},
							},
						},
					},
					// 后面写
					Label{
						AssignTo: &haomaNumberLabel,
						Text:     "写入" + strconv.Itoa(haomaNumber) + "个号码",
					},
					Composite{
						Layout: HBox{},
						Children: []Widget{
							PushButton{
								AssignTo:    nil,
								Text:        "导入txt",
								Name:        "txtButton",
								ToolTipText: "点击选择txt文件",
								MinSize:     Size{Height: 40},
								OnClicked: func() {
									dlg := new(walk.FileDialog)
									tlog := util.NewTasklog(rizhi)
									phoneNumber := util.NewTasklog(haoma)

									dlg.Title = "请选择txt文件"
									dlg.Filter = "*.txt"
									dlg.ShowOpenMultiple(parent.Form())

									var results []string = handleFilePath(dlg.FilePaths)
									var contents []string

									for _, v := range results {
										tlog.AppendLogText("\n读入文件")
										tlog.AppendLogText(v)
										tlog.AppendLogText("中的内容")
										contentsi := readPaths(v)

										contents = append(contents, contentsi...)
									}

									for _, v := range contents {
										phoneNumber.AppendLogText(v)
									}
								},
							},
							PushButton{
								AssignTo: nil,
								Text:     "去重复",
								MinSize:  Size{Height: 40},
								// OnClicked: quChFu,
								OnClicked: func() {
								},
							},
							PushButton{
								AssignTo:    nil,
								Text:        "导入Excel",
								MinSize:     Size{Height: 40},
								Name:        "ExcelButton",
								ToolTipText: "点击选择Excel文件",
								OnKeyPress: func(key walk.Key) {
								},
								OnSizeChanged: func() {
								},
								OnClicked: func() {
								},
							},
						},
					},
				},
			},
			// 选择短信模板
			Composite{
				Layout: VBox{},
				Children: []Widget{
					HSplitter{
						Children: []Widget{
							Composite{
								MinSize:    Size{Height: 31, Width: 300},
								DataBinder: DataBinder{
									// AssignTo:   &operateEngineDb,
									// DataSource: operateFrom,
								},
								Layout: Grid{Columns: 5},
								Children: []Widget{
									Label{
										Text: "选择模板：",
									},
									ComboBox{
										// AssignTo:      &engineOptionsBox,
										Value:         Bind("EngineId", SelRequired{}),
										BindingMember: "Id",
										DisplayMember: "Name",
										// Model:         Engine.GetEngineOptionsSelects(),
										OnCurrentIndexChanged: func() {

										},
									},
									PushButton{
										Text:    "选择",
										MaxSize: Size{Width: 40},
										OnClicked: func() {

										},
									},
								},
							},
						},
					},
					HSplitter{
						Children: []Widget{
							TextEdit{
								AssignTo:  &mobanMsg,
								Enabled:   true,
								ReadOnly:  false,
								TextColor: walk.RGB(136, 136, 136),
								Text:      "请选择短信模板...",
								Row:       8,
								Column:    8,
								OnTextChanged: func() {
									// fmt.Println("bounds->", mobanMsg.Bounds(), "bounds.Width->", mobanMsg.Bounds().Width)
								},
							},
						},
					},

					HSplitter{
						Children: []Widget{
							PushButton{
								Text:    "取消",
								MinSize: Size{Height: 40},
								OnClicked: func() {
									fmt.Println("test priontln")
								},
							},
							PushButton{
								Text:    "确定",
								MinSize: Size{Height: 40},
								OnClicked: func() {
									util.LogToFile("单击了确定！", "", -1)
								},
							},
						},
					},
					VSplitter{
						Children: []Widget{
							TextEdit{
								AssignTo:  &rizhi,
								Enabled:   true,
								TextColor: walk.RGB(136, 136, 136),
								Text:      "日志打印...",
								Font: Font{
									PointSize: 12,
								},
								Column:  5,
								Row:     5,
								VScroll: true,
								OnTextChanged: func() {
								},
							},
						},
					},
				},
			},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}
	return sendMsg, nil
}

package dialog

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type FooPage struct {
	*walk.Composite
}
type BarPage struct {
	*walk.Composite
}
type BazPage struct {
	*walk.Composite
}

func NewBazPage(parent walk.Container) (Page, error) {
	p := new(BazPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "bazPage",
		Layout:   HBox{},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					HSpacer{},
					ToolBar{
						Items: []MenuItem{
							Menu{
								Text:  "打开",
								Image: "./data/img/open.png",
								Items: []MenuItem{
									Action{
										// AssignTo: &dakaiAction,
										Image: "./data/img/media.png",
										Text:  "媒体文件",
										OnTriggered: func() {
										},
									},
								},
							},
							Menu{
								Text: "quedi",
							},
						},
					},
				},
			},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}

func NewBarPage(parent walk.Container) (Page, error) {
	p := new(BarPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "barPage",
		Layout:   HBox{},
		Children: []Widget{
			HSpacer{},
			Label{Text: "I'm the Bar page"},
			HSpacer{},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}

func NewFooPage(parent walk.Container) (Page, error) {
	p := new(FooPage)

	if err := (Composite{
		AssignTo: &p.Composite,
		Name:     "fooPage",
		Layout:   HBox{},
		Children: []Widget{

			Label{Text: "I'm the Foo page"},
			Label{Text: "I'm the Foo page"},

			Label{Text: "I'm the Foo page"},
			Label{Text: "I'm the Foo page"},
			HSpacer{},
		},
	}).Create(NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}

	return p, nil
}

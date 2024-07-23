package main

import (
	. "SMS/app"
	. "SMS/app/dialog"
	"SMS/app/util"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

//应用版本号
const APP_VERSION = "0.1.1"

var AppRootDir string
var mw *MyMainWindow

var (
	outputSrtChecked *walk.CheckBox
	outputLrcChecked *walk.CheckBox
	outputTxtChecked *walk.CheckBox

	globalFilterChecked  *walk.CheckBox
	definedFilterChecked *walk.CheckBox
)

func init() {
	//主 Window
	mw = new(MyMainWindow)

	AppRootDir = util.AppRoorDir

}

func main() {
	// var taskFiles = new(TaskHandleFile)
	util.LogToFile("--1|2|5--", "<ruier>", 19)
	var logText *walk.TextEdit

	// var operateEngineDb *walk.DataBinder
	// var operateTranslateEngineDb *walk.DataBinder
	// var operateTranslateDb *walk.DataBinder
	// var operateDb *walk.DataBinder
	// var operateFilter *walk.DataBinder

	// var operateFrom = new(OperateFrom)

	// var startBtn *walk.PushButton          //生成字幕Btn
	// var startTranslateBtn *walk.PushButton //字幕翻译Btn
	// var engineOptionsBox *walk.ComboBox
	// var translateEngineOptionsBox *walk.ComboBox
	// var dropFilesEdit *walk.TextEdit

	// var appSetings = Setings.GetCacheAppSetingsData()
	// var appFilter = Filter.GetCacheAppFilterData()

	//初始化展示配置
	// operateFrom.Init(appSetings)

	//日志
	var tasklog = util.NewTasklog(logText)

	//字幕生成应用
	// var videosrt = NewApp(AppRootDir)
	//注册日志事件
	// videosrt.SetLogHandler(func(s string, video string) {
	// 	baseName := tool.GetFileBaseName(video)
	// 	strs := strings.Join([]string{"【", baseName, "】", s}, "")
	// 	//追加日志
	// 	tasklog.AppendLogText(strs)
	// })
	//字幕输出目录
	// videosrt.SetSrtDir(appSetings.SrtFileDir)
	//注册[字幕生成]多任务
	// var multitask = NewVideoMultitask(appSetings.MaxConcurrency)

	//字幕翻译应用
	// var srtTranslateApp = NewSrtTranslateApp(AppRootDir)
	//注册日志回调事件
	// srtTranslateApp.SetLogHandler(func(s string, file string) {
	// 	baseName := tool.GetFileBaseName(file)
	// 	strs := strings.Join([]string{"【", baseName, "】", s}, "")
	// 	//追加日志
	// 	tasklog.AppendLogText(strs)
	// })
	//文件输出目录
	// srtTranslateApp.SetSrtDir(appSetings.SrtFileDir)
	//注册[字幕翻译]多任务
	// var srtTranslateMultitask = NewTranslateMultitask(appSetings.MaxConcurrency)
	// var dakaiAction *walk.Action
	cfg := &MultiPageMainWindowConfig{
		Name:    "mainWindow",
		MinSize: Size{600, 400},
		MaxSize: Size{600, 600},
		MenuItems: []MenuItem{
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "About",
						OnTriggered: func() {},
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text:        "About",
						OnTriggered: func() {},
					},
				},
			},
		},
		OnCurrentPageChanged: func() {
			mw.UpdateTitle(mw.CurrentPageTitle())
		},
	}

	mpmw, err := NewMultiPageMainWindow(cfg)
	if err != nil {
		panic(err)
	}

	mw.MultiPageMainWindow = mpmw

	mw.UpdateTitle(mw.CurrentPageTitle())

	//更新
	tasklog.SetTextEdit(logText)

	//尝试校验新版本
	go func() {
		if CheckNewVersionMessage {

			appV := new(AppVersion)
			vtag := appV.GetAppVersion()
			if vtag != "" && CompareVersion(vtag, APP_VERSION) == 1 {
				_ = appV.ShowVersionNotifyInfo(vtag, mw)
			}
		}
	}()
	mw.Run()
}

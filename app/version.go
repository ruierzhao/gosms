package app

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lxn/walk"
)

const (
	APPVERSIOM_URL = "https://gitee.com/641453620/video-srt-windows/tags"
)

var (
	CheckNewVersionMessage = false // 新版本信息检测默认开启
)

type AppVersion struct {
}

//版本号比较
func CompareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	var limit int

	if len(v1) > len(v2) {
		limit = len(v1)
	} else {
		limit = len(v2)
	}

	for {
		if len(v1) >= limit {
			break
		}
		v1 = append(v1, "0")
	}

	for {
		if len(v2) >= limit {
			break
		}
		v2 = append(v2, "0")
	}

	for i := 0; i < limit; i++ {
		num1, _ := strconv.Atoi(v1[i])
		num2, _ := strconv.Atoi(v2[i])
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}
	return 0
}

// 码云查询新版本
func (app *AppVersion) GetAppVersion() string {
	// timeout := time.Duration(3 * time.Second)
	client := &http.Client{}
	//请求码云 获取版本信息
	res, err := client.Get(APPVERSIOM_URL)
	if err != nil {
		return ""
	}
	//解析网页
	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	vs := ""
	dom.Find("#git-tags-container .releases-tags-wrap .releases-tag-content .tag-list .tag-item").Each(func(i int, s *goquery.Selection) {
		if vs != "" {
			return
		}
		tag, is := s.Find(".tag-name a").Attr("title")
		if is && tag != "" {
			vs = strings.TrimSpace(tag)
		}
	})
	return vs
}

//显示更新提醒
func (v *AppVersion) ShowVersionNotifyInfo(version string, own *MyMainWindow) error {
	mw, err := walk.NewMainWindow()
	if err != nil {
		return err
	}
	ni, err := walk.NewNotifyIcon(mw)
	if err != nil {
		return err
	}

	defer func() {
		time.Sleep(time.Second * 15)
		_ = ni.Dispose()
	}()

	if err := ni.SetVisible(true); err != nil {
		return err
	}
	if err := ni.ShowMessage("更新提醒", "检测到AliYun SMS的新版本（v"+version+"），请及时下载更新哦"); err != nil {
		return err
	}
	return nil
}

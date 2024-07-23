package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 初始化Logger
var (
	AppRoorDir string
	LogToFile  func(logTxt interface{}, prefix string, flag int)
)

func init() {
	AppRoorDir = GetAppRootDir()
	LogToFile = InitLogger("log.log")
}

func GetAppRootDir() string {
	if rootDir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	} else {
		return WinDir(rootDir)
	}
}

func InitLogger(fileName string) func(logTxt interface{}, prefix string, flag int) {
	return func(logTxt interface{}, prefix string, flag int) {
		if prefix == "" {
			prefix = "< ruier >"
		}
		if flag == -1 {
			flag = (1 | 2 | 5)
		}
		_logFile := AppRoorDir + "/data/" + fileName
		logFile, err := os.OpenFile(_logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Println("open "+_logFile+" failed, err:", err)
			return
		}
		log.New(logFile, prefix, flag).Println(" >>", logTxt)
	}
}

//Windows下Dir路径转换
func WinDir(dir string) string {
	return strings.Replace(dir, "\\", "/", -1)
}

package dialog

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/lxn/walk"
)

//处理选择路径集
func handleFilePath(filpaths []string) (results []string) {
	for _, v := range filpaths {
		v = strings.Replace(v, "\\", "/", -1)
		results = append(results, v)
	}
	return
}

//读取路径下文件内容
func readPaths(result string) []string {
	var cs []string
	vF, _ := os.Open(result)
	defer vF.Close() //关闭bug

	br := bufio.NewReader(vF)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		cs = append(cs, string(line))
	}
	return cs
}

func quChFu(haoma *walk.TextEdit) {
	haomaArr := readPaths(haoma.Text())
	var haomaMap = make(map[string]bool)
	for _, v := range haomaArr {
		haomaMap[v] = true
	}
}

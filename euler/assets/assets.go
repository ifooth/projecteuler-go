package assets

import (
	"embed"
	"strings"
)

//go:embed dat
var f embed.FS

func GetAsset(name string) ([]byte, error) {
	return f.ReadFile("dat/" + name)
}

// GetTrimString : 获取文件内容，去掉掉空行和换行
func GetTrimString(name string) (string, error) {
	asset, err := GetAsset(name)
	if err != nil {
		return "", err
	}
	var dat string
	for _, v := range strings.Split(string(asset), "\n") {
		dat += strings.TrimSpace(v)
	}
	return dat, nil
}

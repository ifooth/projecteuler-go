package assets

import (
	"embed"
	"strconv"
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

func GetMatrixInt(name string) ([][]int64, error) {
	asset, err := GetAsset(name)
	if err != nil {
		return nil, err
	}
	matrix := [][]int64{}
	for _, v := range strings.Split(string(asset), "\n") {
		array := []int64{}
		for _, value := range strings.Split(strings.TrimSpace(v), " ") {
			a, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			array = append(array, int64(a))
		}
		matrix = append(matrix, array)
	}
	return matrix, nil
}

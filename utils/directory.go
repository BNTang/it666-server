package utils

import (
	"os"
)

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

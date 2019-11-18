package main

import "strings"

//验证图片格式
func IsImage(path string) bool{
	path = strings.ToLower(path)
	paths := strings.Split(path, ".")
	if len(paths) < 2{
		return false
	}
	fileType := paths[len(paths) - 1]
	switch fileType {
	case "jpg":
	case "jpeg":
	case "png":
	default:
		return false
	}
	return true
}

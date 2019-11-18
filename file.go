package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var(
	Sep = string(filepath.Separator)
)

//获取文件列表
// 按照文件名，倒叙排列返回
//param src string 查询的文件夹路径
//param filters []string 仅保留的文件，文件夹除外
//param isSrc bool 返回是否为文件路径
//return []string,error 文件列表,错误
func getFileList(src string, filters []string, isSrc bool) ([]string, error) {
	//初始化
	var fs []string
	//读取目录
	dir, err := ioutil.ReadDir(src)
	if err != nil {
		return nil, err
	}
	//遍历目录文件
	for _, v := range dir {
		var appendSrc string
		if isSrc == true {
			appendSrc = src + Sep + v.Name()
		} else {
			appendSrc = v.Name()
		}
		if v.IsDir() == true || len(filters) < 1 {
			fs = append(fs, appendSrc)
			continue
		}
		names := strings.Split(v.Name(), ".")
		if len(names) == 1 {
			fs = append(fs, appendSrc)
			continue
		}
		t := names[len(names)-1]
		for _, filterValue := range filters {
			if t != filterValue {
				continue
			}
			fs = append(fs, appendSrc)
		}
	}
	//对数组进行倒叙排序
	sort.Sort(sort.Reverse(sort.StringSlice(fs)))
	//返回
	return fs, nil
}

//判断是否为文件夹
//param src string 文件夹路径
//return bool 是否为文件夹
func isFolder(src string) bool {
	info, err := os.Stat(src)
	return err == nil && info.IsDir()
}

//判断是否为文件
//param src string 文件路径
//return bool 是否为文件
func isFile(src string) bool {
	info, err := os.Stat(src)
	return err == nil && !info.IsDir()
}

//读取文件
//param src string 文件路径
//return []byte,error 文件数据,错误
func loadFile(src string) ([]byte, error) {
	fd, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	c, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return c, nil
}
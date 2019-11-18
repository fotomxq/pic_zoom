package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type configType struct {
	Width int
	Height int
}

var(
	configData configType
)

//图像快速处理模块和主程序
// 该模块可以将指定目录下的所有图片文件，自动压缩到约定的比例范围内
func main(){
	fmt.Println("begin...")
	configFileData, err := loadFile("./config.json")
	if err != nil{
		fmt.Println("cannot load config, " + err.Error())
		time.Sleep(time.Second * 10)
		return
	}
	if err := json.Unmarshal(configFileData, &configData); err != nil{
		fmt.Println("cannot load config, " + err.Error())
		time.Sleep(time.Second * 10)
		return
	}
	if err := run("./"); err != nil{
		fmt.Println("cannot run, " + err.Error())
		time.Sleep(time.Second * 10)
		return
	}
	fmt.Println("finish.")
	time.Sleep(time.Second * 10)
}

func run(path string) error {
	//无限制遍历本目录下的所有图片文件，并自动覆盖式修改
	fileList, err := getFileList(path, []string{}, true)
	if err != nil{
		return err
	}
	for _, v := range fileList{
		if isFolder(v) {
			fmt.Println("find new dir: " + v)
			if err := run(v); err != nil{
				return err
			}
			continue
		}
		if IsImage(v) {
			fmt.Println("zoom new image: " + v)
			if err := Zoom(v, v, configData.Width, configData.Height); err != nil{
				return err
			}
			continue
		}
		fmt.Println("find new file, but not image or dir, file path: " + v)
	}
	return nil
}
package main

import (
	"errors"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
)

//压缩程序
//params path string 目标文件
//params destPath string 新的文件路径
//params destMaxWidth float64 最大宽度
//params destMaxHeight float64 最大高度
func Zoom(path string, destPath string, destMaxWidth int, destMaxHeight int) error {
	if IsImage(path) == false{
		return errors.New("path file is not image")
	}
	if IsImage(destPath) == false{
		return errors.New("dest path file is not image")
	}
	if destMaxWidth < 1 || destMaxHeight < 1{
		return errors.New("dest image size is error")
	}
	srcImage, err := imaging.Open(path)
	if err != nil{
		return errors.New("cannot open file, " + err.Error())
	}
	destImage1 := imaging.Fit(srcImage, destMaxWidth, destMaxHeight, imaging.Lanczos)
	destImage2 := imaging.New(destImage1.Bounds().Size().X, destImage1.Bounds().Size().Y, color.NRGBA{0, 0, 0, 0})
	destImage2 = imaging.Paste(destImage2, destImage1, image.Pt(0, 0))
	if err := imaging.Save(destImage2, destPath); err != nil{
		return errors.New("cannot save new file, " + err.Error())
	}
	return nil
}

set GOOS=windows
set GOARCH=amd64
go build -o ./../pic_zoom_win64.exe ./../

set GOOS=linux
set GOARCH=amd64
go build -o ./../pic_zoom_linux64.exe ./../
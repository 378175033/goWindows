package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
	"fmt"
	"runtime"
)

var (
	inTE1, outTE1, inTE2, outTE2 *walk.TextEdit
	m = map[int]string{1: "A", 2: "B", 3: "C"}
	info string
)

func init() { //代码包初始化函数
	fmt.Printf("Map: %v\n", m)
	info = fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
	fmt.Println(info)
	test()
}

func test()  {

}

func main() {

	MainWindow{
		Title:   "SCREAM",
		MinSize: Size{600, 400},
		Layout: VBox{
			Margins: Margins{10, 10, 10, 10},
		},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					HSplitter{
						Children: []Widget{
							TextEdit{AssignTo: &inTE1},
							TextEdit{AssignTo: &outTE1, ReadOnly: true, Text: "Upper"},
						},
					},
					HSplitter{
						Children: []Widget{
							TextEdit{AssignTo: &inTE2},
							TextEdit{AssignTo: &outTE2, ReadOnly: true, Text: "Lower"},
						},
					},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE1.SetText(strings.ToUpper(inTE1.Text()))
					outTE2.SetText(strings.ToLower(inTE2.Text()))
				},
			},
		},
	}.Run()
}
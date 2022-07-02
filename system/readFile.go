package system

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var fileName = "/Users/jonny/IMPORTANT/learnCode/hi_go/assets/a.text"

// 读取文件并显示所有内容
func ReadFileAllContent() {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}

// 读取文件并逐行显示所有内容
func ReadFileLineContent() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	line := bufio.NewReader(file)

	for {
		content, _, err := line.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(content) + "\n")
	}
}

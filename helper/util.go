package helper

import (
	"errors"
	"fmt"
	"regexp"
)


// 获取文件名或者截取路径
func GetFileName(path,needle string) (fileName string, err error) {
	re := regexp.MustCompile(needle)
	match := re.FindIndex([]byte(path))
	fmt.Println(match)
	if len(match) == 0 {
		fmt.Println("没有匹配的ptah，文件路径有问题")
		return "",errors.New("没有匹配的ptah，文件路径有问题")
	}
	content := path[match[1] : len(path)]
	fmt.Println(content)

	return content,nil
}

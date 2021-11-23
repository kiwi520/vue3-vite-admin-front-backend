package helper

import (
	"errors"
	"fmt"
	"net"
	"os"
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


//获取本机IP
func GetLocalIP() []string {
	var ipStr []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces error:", err.Error())
		return ipStr
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//获取IPv6
					/*if ipnet.IP.To16() != nil {
					    fmt.Println(ipnet.IP.String())
					    ipStr = append(ipStr, ipnet.IP.String())

					}*/
					//获取IPv4
					if ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())
						ipStr = append(ipStr, ipnet.IP.String())

					}
				}
			}
		}
	}
	return ipStr

}

func GetHttps() string {
	hps:= os.Getenv("Open_Https")
	if hps == "on" {
		return "https"
	} else {
		return "http"
	}
}

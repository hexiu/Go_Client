package main

import (
	"fmt"
	"net"
	"net/http"
	// "net/url"
	"os"
	"strconv"
	"time"
)

const (
	host    string = "192.168.0.105"        //远端服务器的 主机域名 Or IP
	port    string = ":8080"                //远端服务器的ip
	path    string = "/SignIn/DealUser.jsp" //远端服务器的URI
	logPath string = "./log/"               //日志记录
)

var false_count int64 = 0

func main() {
	logging()
	for true {
		//向服务器发送信息
		err := send_message()
		if err {
			time.Sleep(time.Second * 10)
		} else {
			false_count++
			fmt.Println("false count:", false_count)
			if false_count >= 10 {
				time.Sleep(60 * time.Second)
			}
			continue
		}
	}
}

//向服务器发送信息
func send_message() bool {
	// 获取mac地址，macAddress 是一个切片类型
	macAddress := getMacAddress()
	fmt.Println(macAddress)

	var client http.Client
	fmt.Println(macAddress[1])
	_, err := client.Get("http://" + host + port + path + "?mac=" + macAddress[1])
	if err != nil {
		return false
	}
	// fmt.Println(mess)
	// fmt.Println(mess.Proto)
	return true
}

// 获取Mac地址,返回一个切片
func getMacAddress() (macAdress []string) {
	interfaces, _ := net.Interfaces()
	macAdress = make([]string, 3)
	index := 0
	for _, inter := range interfaces {
		mac := inter.HardwareAddr
		macAdress[index] = fmt.Sprintf("%s", mac)
		index++
	}
	return macAdress
}

func logging() {
	initLog()

}

//初始化日志相关信息
func initLog() {

	file, err := os.OpenFile(logPath, 0, os.ModePerm)
	if err != nil {
		err := createDir()
		if !err {
			fmt.Printf("Log Dir is Exist!\n")
			return
		}
	}
	defer file.Close()

	logname := GetLogName()
	logFile, err := os.OpenFile(logPath+"/"+logname, 0, os.ModePerm)
	if err != nil {
		err1 := createFile()
		if !err1 {
			return
		}

	}
	defer logFile.Close()
}

// 创建当天的Log文件
func createFile() bool {
	logname := GetLogName()
	fmt.Println(logname)
	logFile, err := os.Create(logPath + "/" + logname)
	if err != nil {
		fmt.Println("Create Log File " + logname + " error ")
		return false
	}
	defer logFile.Close()
	return true
}

//创建日志目录
func createDir() bool {
	file, err := os.OpenFile(logPath, 0, os.ModePerm)
	if err != nil {
		os.Mkdir("./log", os.ModePerm)
		return true
	}
	defer file.Close()
	return false
}

//获取字符串的数据名
func GetLogName() (name string) {
	year := strconv.Itoa(time.Now().Year())
	month := time.Now().Month().String()
	day := strconv.Itoa(time.Now().Day())

	return "logging" + "_" + year + "_" + month + "_" + day + ".log"

}

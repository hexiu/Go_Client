package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	host    string = "222.24.24.102" //远端服务器的 主机域名 Or IP
	port    string = ":8080"         //远端服务器的ip
	path    string = "/logs"         //远端服务器的URI
	logPath string = "./log/"        //日志记录
	Time    int    = 60
)

var logname = GetLogName()
var false_count int = 0

func main() {
	logging()
	writeLog("test")
	writeLog("test1")
	writeLog("test2")
	for true {
		//向服务器发送信息
		err := send_message()
		if err {
			time.Sleep(time.Second * time.Duration(Time))
		} else {
			false_count++
			data := strconv.Itoa(false_count)
			writeLog("Connect false count:" + data)
			// fmt.Println("false count:" + data)
			if false_count >= 10 {
				time.Sleep(5 * time.Duration(Time) * time.Second)
			}
			continue
		}
	}
}

//向服务器发送信息
func send_message() bool {
	// 获取mac地址，macAddress 是一个切片类型
	macAddress := getMacAddress()
	// fmt.Println(macAddress)

	wrinteInit("My MacAddress : " + macAddress[1])

	var client http.Client
	// fmt.Println(macAddress[0])
	_, err := client.Get("http://" + host + port + path + "?mac=" + macAddress[1])
	if err != nil {
		return false
	} else {
		writeLog("Send Success " + time.Now().String())
	}
	// fmt.Println(mess)
	// fmt.Println(mess.Proto)
	return true
}

// 获取Mac地址,返回一个切片
func getMacAddress() (macAdress []string) {
	interfaces, _ := net.Interfaces()
	macAdress = make([]string, 5)
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
		wrinteInit("Create Log Dir Success !")
		if !err {
			writeLog("Log Dir is Exist!\n")
			return
		}
	}
	defer file.Close()

	logname := GetLogName()
	logFile, err := os.OpenFile(logPath+"/"+logname, 0, os.ModePerm)
	if err != nil {
		err1 := createFile()
		wrinteInit("Create Log File" + logname + "success")
		if !err1 {
			return
		}

	}
	defer logFile.Close()
}

// 创建当天的Log文件
func createFile() bool {
	logname := GetLogName()
	// fmt.Println(logname)
	logFile, err := os.Create(logPath + "/" + logname)
	if err != nil {
		writeLog("Create Log File " + logname + " error ")
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

//写入启动信息
func wrinteInit(mess string) {
	file, err := os.Create("init.log")
	if err != nil {
		return
	}
	file.WriteString(time.Now().String() + "  " + mess + " \n")
	file.Close()
}

//写入日志
func writeLog(mess string) {
	filename := logPath + logname
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	// fmt.Println(time.Now().String() + "  " + mess + " \n")
	//data := time.Now().String() + "  " + mess + " \n"
	file.WriteString(time.Now().String() + "  " + mess + " \n")
	defer file.Close()
}

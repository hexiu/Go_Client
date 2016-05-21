package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	// "strings"
	"io/ioutil"
	"time"
)

const (
	//	host    string = "127.0.0.1" //远端服务器的 主机域名 Or IP
	host      string = "ims.smartxupt.com" //远端服务器的 主机域名 Or IP
	port      string = ""                  //远端服务器的ip
	path      string = "/sign_action.php"  //远端服务器的URI
	logPath   string = "./log/"            //日志记录
	Time      int    = 20
	macSelect int    = 0
)

var logname = GetLogName()
var false_count int = 0

func main() {
	//logging()
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

	wrinteInit("My MacAddress : " + macAddress[0] + macAddress[1])

	var client http.Client
	// fmt.Println(macAddress[0])
	mac := macAddress[macSelect][0:2] + macAddress[macSelect][3:5] + macAddress[macSelect][6:8] + macAddress[macSelect][9:11] + macAddress[macSelect][12:14] + macAddress[macSelect][15:]
	// fmt.Println(mac)
	conn, err := client.Get("http://" + host + port + path + "?agent=" + mac)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	body, err := ioutil.ReadAll(conn.Body)
	if err != nil {
		fmt.Println("Get Server message error! ", err)
	}

	os.Create("message.txt")

	file, err := os.OpenFile("message.txt", os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Write message error !", err)
	}
	file.WriteString(string(body))

	defer file.Close()

	defer conn.Body.Close()

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
		Rmac := fmt.Sprintf("%s", mac)
		// macAdress[index] = Rmac
		if len(Rmac) == 17 {
			// fmt.Println(Rmac[:2])
			if string(Rmac[:2]) != "00" && string(Rmac[:2]) != "01" {
				macAdress[index] = Rmac
				index++
				if index == 2 {
					break
				}
			} else {
				continue
			}

		}
	}
	// fmt.Println(macAdress)

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

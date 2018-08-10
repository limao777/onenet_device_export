package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/limao777/onenet_device_export/config"
	"github.com/limao777/onenet_device_export/structs"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var xlsx *excelize.File
var curl_thread int
var all_dev_counter int = 2
var lock_co *sync.Mutex

var curl_chan chan bool

/**
初始化配置和锁
*/
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	lock_co = new(sync.Mutex)
}

/**
向excel写入设备数据
*/
func infoToExcel(data structs.BackDevDevices) {
	lock_co.Lock()
	defer lock_co.Unlock()
	xlsx.SetCellValue("Export", string('A')+strconv.Itoa(all_dev_counter), data.Id)
	if data.Private {
		xlsx.SetCellValue("Export", string('B')+strconv.Itoa(all_dev_counter), "true")
	} else {
		xlsx.SetCellValue("Export", string('B')+strconv.Itoa(all_dev_counter), "false")
	}
	xlsx.SetCellValue("Export", string('C')+strconv.Itoa(all_dev_counter), data.Title)
	xlsx.SetCellValue("Export", string('D')+strconv.Itoa(all_dev_counter), data.Desc)
	tmp, _ := json.Marshal(data.Tags)
	xlsx.SetCellValue("Export", string('E')+strconv.Itoa(all_dev_counter), tmp)
	xlsx.SetCellValue("Export", string('F')+strconv.Itoa(all_dev_counter), data.Url)
	xlsx.SetCellValue("Export", string('G')+strconv.Itoa(all_dev_counter), data.Isdn)
	tmp, _ = json.Marshal(data.Location)
	if string(tmp) != "null" {
		xlsx.SetCellValue("Export", string('H')+strconv.Itoa(all_dev_counter), tmp)
	}
	xlsx.SetCellValue("Export", string('I')+strconv.Itoa(all_dev_counter), data.Protocol)

	tmp, _ = json.Marshal(data.Route_to)
	if string(tmp) != "null" {
		xlsx.SetCellValue("Export", string('J')+strconv.Itoa(all_dev_counter), tmp)
	}
	tmp, _ = json.Marshal(data.Auth_info)
	if string(tmp) != "null" {
		xlsx.SetCellValue("Export", string('K')+strconv.Itoa(all_dev_counter), tmp)
	}
	xlsx.SetCellValue("Export", string('L')+strconv.Itoa(all_dev_counter), data.Active_code)
	xlsx.SetCellValue("Export", string('M')+strconv.Itoa(all_dev_counter), data.Interval)
	tmp, _ = json.Marshal(data.Other)
	if string(tmp) != "null" {
		xlsx.SetCellValue("Export", string('N')+strconv.Itoa(all_dev_counter), tmp)
	}
	tmp, _ = json.Marshal(data.Key)
	if string(tmp) != "null" {
		xlsx.SetCellValue("Export", string('O')+strconv.Itoa(all_dev_counter), tmp)
	}
	xlsx.SetCellValue("Export", string('P')+strconv.Itoa(all_dev_counter), data.Create_time)

	all_dev_counter++
	return
}

/**
主函数
配置文件检查
excel初步工作
多路数据请求
*/
func main() {

	uri := "http://api.heclouds.com/devices?per_page=100"

	api_key := config.Get("onenet", "apiKey")
	if api_key == "" {
		fmt.Println("get config api-key error(配置文件获取api-key错误)")
		for {
		}
	}

	curl_thread_str := config.Get("app", "goroutine")
	if curl_thread_str == "" {
		fmt.Println("get config goroutine error(配置文件获取goroutine错误)")
		for {
		}
	}
	var err error
	curl_thread, err = strconv.Atoi(curl_thread_str)
	if err != nil {
		fmt.Println("get config format goroutine error(配置文件获取goroutine格式错误)")
		for {
		}
	}
	if curl_thread < 1 {
		curl_thread = 1
	}
	if curl_thread > 100 {
		curl_thread = 100
	}
	api_key = strings.Trim(api_key, " ")

	search_key_words := config.Get("search", "key_words")
	if search_key_words != "" {
		uri = uri + "&key_words=" + search_key_words
	}
	search_online := config.Get("search", "online")
	if search_online != "" {
		uri = uri + "&online=" + search_online
	}
	search_auth_info := config.Get("search", "auth_info")
	if search_auth_info != "" {
		uri = uri + "&auth_info=" + search_auth_info
	}
	search_tag := config.Get("search", "tag")
	if search_tag != "" {
		uri = uri + "&tag=" + search_tag
	}
	search_private := config.Get("search", "private")
	if search_private != "" {
		uri = uri + "&private=" + search_private
	}
	search_begin := config.Get("search", "begin")
	if search_begin != "" {
		uri = uri + "&begin=" + search_begin
	}
	search_end := config.Get("search", "end")
	if search_end != "" {
		uri = uri + "&end=" + search_end
	}
	
//	fmt.Println(uri)

	xlsx = excelize.NewFile()
	index := xlsx.NewSheet("Export")
	xlsx.SetActiveSheet(index)
	xlsx.SetCellValue("Export", "A1", "id")
	xlsx.SetCellValue("Export", "B1", "private")
	xlsx.SetCellValue("Export", "C1", "title")
	xlsx.SetCellValue("Export", "D1", "desc")
	xlsx.SetCellValue("Export", "E1", "tags")
	xlsx.SetCellValue("Export", "F1", "url")
	xlsx.SetCellValue("Export", "G1", "idsn")
	xlsx.SetCellValue("Export", "H1", "location")
	xlsx.SetCellValue("Export", "I1", "protocol")
	xlsx.SetCellValue("Export", "J1", "route_to")
	xlsx.SetCellValue("Export", "K1", "auth_info")
	xlsx.SetCellValue("Export", "L1", "active_code")
	xlsx.SetCellValue("Export", "M1", "interval")
	xlsx.SetCellValue("Export", "N1", "other")
	xlsx.SetCellValue("Export", "O1", "key")
	xlsx.SetCellValue("Export", "P1", "create_time")

	fmt.Println("start getting device info(开始获取设备数据...)")

	curl_chan = make(chan bool, 1)

	for i := 1; i <= curl_thread; i++ {
		go do_curl(uri, i, api_key)
	}

	for i := 1; i <= curl_thread; i++ {
		<-curl_chan
	}

	err = xlsx.SaveAs("./ExportOneNETDevice.xlsx")
	if err != nil {
		fmt.Println("write to excel error(写入excel发生错误)", err)
	}

	fmt.Println("end getting device info， please quit this application(设备数据获取完毕，可以退出该程序了...)")

	select {}
}

func do_curl(uri string, i int, api_key string) {
	for j := i; j <= 10000; j += curl_thread {
		//API限制最大10000

		client := &http.Client{}
		var req *http.Request
		var err error

		curl_uri := uri + "&page=" + strconv.Itoa(j)
		
		req, err = http.NewRequest("GET", curl_uri, nil)
		req.Header.Set("api-key", api_key)

		if err != nil {
			//TODO handle error
			req = nil
		}

		resp, err := client.Do(req)
		if err != nil {
			//TODO handle error
			resp = nil
		}

		if resp != nil && req != nil && client != nil {

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//TODO handle error
				resp.Body.Close()
			}
			//		fmt.Println(string(body))
			back_ret := structs.BackDevRet{}
			err = json.Unmarshal([]byte(body), &back_ret)
			if err == nil {
				if len(back_ret.Data.Devices) > 0 {
					for _, v := range back_ret.Data.Devices {
						infoToExcel(v)
					}

				} else {
					break
				}
			} else {
				fmt.Println("[err] deal json error", i)
			}
			resp.Body.Close()
			resp = nil
			req = nil
			client = nil
		} else {
			fmt.Printf("[err] page %d get data error\r\n", i)
		}
		fmt.Printf("channel %d accumplished(通道%d获取数据完成)\r\n", j, j)
	}
	curl_chan <- true
	return
}

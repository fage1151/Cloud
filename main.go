package main

import (
	// "bytes"
	"fmt"
	// "os"
	// "strings"
	"net/http"
	"bytes"
	"sort"
	"strings"
	"io/ioutil"
	"time"
	"strconv"
	"encoding/base64"
	"os"
	"log"
	"net/url"
)

func main() {

	//打印
	/*var deviceid = "123457";
	var devicesecret = "jnxiaer7";
	var printdata = "您好，欢迎使用中午打印机";
	var appid = "8000000";
	var appsecret = "*****";
	rpc := rpc{appid, appsecret, http.Client{}}

	printer := Printer{deviceid, devicesecret, rpc};
	res := printer.print(printdata);
	*/
	//设置声音
	/*
		var deviceid = "123457";
		var devicesecret = "jnxiaer7";
		var appid = "8000000";
		var appsecret = "*****";
		rpc := rpc{appid, appsecret, http.Client{}}

		printer := Printer{deviceid, devicesecret, rpc};


		res := printer.set_sound("3");*/
	//设置logo
	/*var deviceid = "123457";
	var devicesecret = "jnxiaer7";
	var appid = "8000000";
	var appsecret = "*****";
	rpc := rpc{appid, appsecret, http.Client{}}

	printer := Printer{deviceid, devicesecret, rpc};

	res := printer.set_logo("/home/heqian/elind/openapi/SDK/Go/images/logo.png");
	*/

	//清空打印队列
	/*var deviceid = "123457";
	var devicesecret = "jnxiaer7";
	var appid = "8000000";
	var appsecret = "*****";
	rpc := rpc{appid, appsecret, http.Client{}}
	printer := Printer{deviceid, devicesecret, rpc};
	res := printer.empty_print_queue();
	*/
	//获取打印状态
	/*	var deviceid = "123457";
		var devicesecret = "jnxiaer7";
		var appid = "8000000";
		var appsecret = "*****";
		rpc := rpc{appid, appsecret, http.Client{}}
		printer := Printer{deviceid, devicesecret, rpc};
		res := printer.get_print_status("34233183");
		*/

	//获取打印机状态

	var deviceid = "123457";
	var devicesecret = "jnxiaer7";
	var appid = "8000000";
	var appsecret = "*****";
	rpc := rpc{appid, appsecret, http.Client{}}
	printer := Printer{deviceid, devicesecret, rpc};
	res := printer.get_status();
	fmt.Println(res)

}

type Printer struct {
	deviceid     string
	devicesecret string
	rpc          rpc
}

func (c *Printer) print(printdata string) interface{} {

	var data = make(map[string]string); /*创建集合 */

	data["printdata"] = printdata;
	return c.Post(data, "");
}

func (c *Printer) get_print_status(id string) interface{} {

	var data = make(map[string]string); /*创建集合 */
	data["dataid"] = id;
	return c.Get(data, "/printstatus");
}

func (c *Printer) get_status() interface{} {

	var data = make(map[string]string); /*创建集合 */

	return c.Get(data, "/status");
}

func (c *Printer) set_sound(sound string) interface{} {

	var data = make(map[string]string); /*创建集合 */
	data["sound"] = sound;
	return c.Post(data, "/sound");
}

func (c *Printer) set_logo(logo string) interface{} {

	var data = make(map[string]string); /*创建集合 */

	data["logodata"] = c.base_64(logo);
	return c.Post(data, "/logo");
}

func (c *Printer) empty_print_queue() interface{} {

	var data = make(map[string]string); /*创建集合 */

	return c.Post(data, "/emptyprintqueue");
}

func (c *Printer) base_64(filepath string) string {

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	encodeString := base64.StdEncoding.EncodeToString(data)

	return encodeString;
}

func (c *Printer) Post(data map[string]string, url string) interface{} {
	data["deviceid"] = c.deviceid;
	data["devicesecret"] = c.devicesecret;
	data["appid"] = c.rpc.appid;
	data["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10);
	return c.rpc.post(data, "http://api.zhongwuyun.com"+url);
}

func (c *Printer) Get(data map[string]string, url string) interface{} {
	data["deviceid"] = c.deviceid;
	data["devicesecret"] = c.devicesecret;
	data["appid"] = c.rpc.appid;
	data["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10);
	return c.rpc.get(data, "http://api.zhongwuyun.com"+url);
}

func (c *Printer) set_args(deviceid string, devicesecret string) {

	c.deviceid = deviceid;
	c.devicesecret = devicesecret;
}

type rpc struct {
	appid     string
	appsecret string
	rpc       http.Client
}

func (c *rpc) post(param map[string]string, url string) interface{} {

	var urlstring = c.MakeParams(param)

	var contentType = "application/x-www-form-urlencoded";
	resp, err := c.rpc.Post(url, contentType, strings.NewReader(urlstring));
	defer resp.Body.Close()

	if err != nil {

		//fmt.Println(err.Error());

		return false;
	}

	if resp.StatusCode == 200 {

		body, _ := ioutil.ReadAll(resp.Body)

		return string(body)

	}

	return false;
}

func (c *rpc) get(param map[string]string, url string) interface{} {

	var urlstring = c.MakeParams(param)

	resp, err := c.rpc.Get(url + "?" + urlstring);
	defer resp.Body.Close()

	if err != nil {

		fmt.Println(err.Error());
		return false;
	}
	if resp.StatusCode == 200 {

		body, _ := ioutil.ReadAll(resp.Body)

		return string(body)

	}

	return false;
}

func (c *rpc) MakeParams(params map[string]string) (params_str string) {

	var p string

	var keys []string

	b := bytes.Buffer{}

	for k, _ := range params {

		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {

		b.WriteString(v)

		b.WriteString("=" + url.QueryEscape(params[v]) + "&")

	}

	p = b.String()

	p = strings.TrimRight(p, "&")

	return p;
}

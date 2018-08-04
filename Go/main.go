package main

import (
	// "bytes"
	"fmt"
	// "os"
	// "strings"
	"net/http"
)

func main() {
	//client := &http.Client{}
	//向服务端发送get请求
	//request, _ := http.NewRequest("GET", "http://192.168.1.35:9091/?publicKey=&privateKey=&info=sad&message=", nil)
	//response, _ := http.PostForm("http://192.168.1.35:9091", url.Values{"publicKey": {"-----BEGIN+PUBLIC+KEY-----%0D%0AMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY%2BP4e3cAtmv%0D%0AppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7%2B5C0Dgacd%0D%0AwYWd%2F7PeCELyEipZJL07Vro7Ate8Bfjya%2BwltGK9%2BXNUIHiumUKULW4KDx21%2B1NL%0D%0AAUeJ6PeW%2BDAkmJWF6QIDAQAB%0D%0A-----END+PUBLIC+KEY-----"}, "privateKey": {""}, "info": {"sad"}, "message": {""}})
	//接收服务端返回给客户端的信息
	// response, _ := client.Do(request)
	// if response.StatusCode == 200 {
	// 	str, _ := ioutil.ReadAll(response.Body)
	// 	bodystr := string(str)
	// 	fmt.Println(bodystr)
	// }

	//post请求
	/*postValues := url.Values{}
	postValues.Add("publicKey", "")
	postValues.Add("privateKey", `----nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----`)
	postValues.Add("info", "")
	postValues.Add("message", `DkCxcs0z6Z03uHWOHOASf2xen+7oNoSad+KG2ss0hkE79211GlgjepmMFRW4zLiF51pVYHHOBFDYYJrnokq5d0ceKYY6ONzbBYKCJMzD7guN3qMYf48Cl9g0bDVb1oMbuN2PstzORe800Q72moQaHVRPiqh7VZ6NCXnkLrtnY64=`)
*/
	var deviceid = "123457";

	var devicesecret = "jnxiaer7";

	var printdata = "您好，欢迎使用中午打印机";

	var appid = "8000000";

	var appsecret = "*****";

	printer = new ()
	rpc := rpc{appid,appsecret,http.Client{}}

	res := Printer{deviceid,devicesecret,rpc}.set_args(deviceid, devicesecret).print(printdata);

	fmt.Println(string(res))

}

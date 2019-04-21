package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	httpUrlRoot := "https://localhost:8081"
	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		fmt.Println("Loadx509keypair err:", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(httpUrlRoot + "/login")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//upload file
	file, err := os.OpenFile("client.csr", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fileBuffer := &bytes.Buffer{}
	writer := multipart.NewWriter(fileBuffer)

	part, err := writer.CreateFormFile("upload_file", "/temp/"+file.Name())
	if err != nil {
		panic(err)
	}

	//### 应该有更好的方法。这一步相当于把数据加载到内存中，文件小了还可以文件大了就不行了。
	_, err = io.Copy(part, file) //把意思好似是把这个发送过去
	//###
	if err != nil {
		panic(err)
	}
	writer.Close()

	request, err := http.NewRequest("POST", httpUrlRoot+"/upload", fileBuffer)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	uploadResp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer uploadResp.Body.Close()
	buff, err := ioutil.ReadAll(uploadResp.Body)

	fmt.Println(string(buff))


	deployJson := make(map[string]interface{})
	deployJson["jsonrpc"] = 2.0
	deployJson["method"] = "query"

}

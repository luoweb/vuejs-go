package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	_ "strings"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!\n")
	// w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080") //允许访问所有域
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             // 返回数据格式是json
	r.ParseForm()

	fmt.Println("method:", r.Method) //获取请求的方法

	switch r.URL.Path {
	case "/login":
		fmt.Println("method:", r.Method) //获取请求的方法
		if r.Method == "GET" {
			t, _ := template.ParseFiles("login.html")
			log.Println(t.Execute(w, nil))
		} else {
			defer r.Body.Close()
			con, _ := ioutil.ReadAll(r.Body) //获取post的数据
			fmt.Println(string(con))

			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(string(con)), &dat); err == nil {
				fmt.Println(dat)
				fmt.Println(dat["username"])
			} else {
				fmt.Println(err)
			}
		}
	case "/upload":
		fmt.Println("method:", r.Method) //获取请求的方法
		if r.Method == "GET" {
			t, _ := template.ParseFiles("helloword.html")
			log.Println(t.Execute(w, nil))
		} else {
			defer r.Body.Close()
			con, _ := ioutil.ReadAll(r.Body) //获取post的数据
			fmt.Println(string(con))
			
			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(string(con)), &dat); err == nil {
				fmt.Println(dat)
				fmt.Println(dat["username"])
			} else {
				fmt.Println(err)
			}
		}
	case "/deploy":
		fmt.Println("method:", r.Method) //获取请求的方法
		if r.Method == "GET" {
			t, _ := template.ParseFiles("helloword.html")
			log.Println(t.Execute(w, nil))
		} else {
			defer r.Body.Close()
			con, _ := ioutil.ReadAll(r.Body) //获取post的数据
			fmt.Println(string(con))

			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(string(con)), &dat); err == nil {
				fmt.Println(dat)
				fmt.Println(dat["username"])
			} else {
				fmt.Println(err)
			}
		}
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", r.URL)
	}

}

func main() {
	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8081",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err = s.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS err:", err)
	}
}

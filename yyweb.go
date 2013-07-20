package main

import (
	"fmt"
	"html/template"
	/*"io"*/
	"io/ioutil"
	"log"
	"net/http"
	/*"path/filepath"*/
	"bytes"
	"gomyadmin"
	"strconv"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	filename := r.URL.Path
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	namef := []rune(filename)
	namefb := bytes.Buffer{}
	for _, t := range namef {
		//fmt.Println(t)
		namefb.WriteRune(t)
	}
	b, err := ioutil.ReadFile("." + namefb.String())
	if err != nil {
		fmt.Fprintf(w, "Hello astaxie")
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, string(b))
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		me := 1
		t.Execute(w, me)
	} else {
		var p pp.Peopel
		for k, v := range r.Form {
			val := v[0]
			if k == "Name" {

				p.Name = val
			}
			if k == "Height" {

				h, _ := strconv.ParseFloat(val, 32)
				p.Height = float32(h)
			}
			if k == "Weight" {

				w, _ := strconv.ParseFloat(val, 32)
				p.Weight = float32(w)
			}
			if k == "Age" {

				age, _ := strconv.Atoi(val)
				p.Age = age
			}
			if k == "sex" {

				if val == "0" {
					p.Sex = pp.M
				}
				if val == "1" {
					p.Sex = pp.W
				}
			}
		}
		var fp pp.Peopel
		fp.Name = p.Name
		po, err := fp.FindUser()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Fprintf(w, err.Error())
		}

		fmt.Println(po)
		if po.Id > 0 {
			p.Id = po.Id
			b, err := p.Update()
			if err != nil {
				fmt.Fprintf(w, "更新失败"+err.Error())
			}
			if b {
				fmt.Fprintf(w, "更新的id:%v", p.Id)
			}
		} else {
			id, err := p.Save()
			if err != nil {
				fmt.Fprintf(w, "保存失败"+err.Error())
			}
			fmt.Fprintf(w, "保存的id:%v", id)
		}
		fmt.Fprintf(w, "%v", p)

		fmt.Println("执行完成")
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

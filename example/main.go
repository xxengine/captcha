// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

// +build ignore

package main

import (
	"github.com/orivil/captcha"
	"log"
	"net/http"
)

func main() {
	var addr = ":8886"
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		code := request.URL.Query().Get("code")
		img := captcha.NewImage("1123", []byte(code), 280, 80)
		_, err := img.WriteTo(writer)
		if err != nil {
			panic(err)
		}
	})
	log.Printf("open url: http://localhost%s?code=BACDL123 and try to change the 'code' parameter\n", addr)
	err := http.ListenAndServe(addr, nil)
	log.Println(err)
}

package main

import "net/http"

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hellow World?"))
	})
	http.ListenAndServe(":2000", nil) //只写端口，默认localhost
}
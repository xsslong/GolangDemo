package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`这他妈也可以`))
	})
	http.ListenAndServe(":3000", nil)
}

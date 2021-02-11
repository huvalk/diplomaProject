package main

import (
	"fmt"
	"github.com/kataras/golog"
	"net/http"
	"os"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		golog.Debug("Knok-knok")

		_, err := w.Write([]byte("I am flying, Jack"))

		if err != nil {
			golog.Error("Knok failed: ", err)
		}
	})

	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", 8080),
		os.Getenv("SSL_PATH") + "fullchain.pem",
		os.Getenv("SSL_PATH") + "privkey.pem",
		nil)

	if err != nil {
		golog.Error("Server haha failed: ", err)
	}
}
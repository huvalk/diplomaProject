package main

import (
	"flag"
	"fmt"
	"github.com/kataras/golog"
	"net/http"
)


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		golog.Debug("Knok-knok")

		_, err := w.Write([]byte("I am flying, Jack"))

		if err != nil {
			golog.Error("Knok failed: ", err)
		}
	})

	var (
		env  = flag.String("env", "dev", "enables debug mode")
		port = flag.Uint64("p", 8080, "port")
	)

	if *env == "dev" {
		golog.SetLevel(golog.DebugLevel.String())
		golog.Debug("Debug")
	}

	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", port),
		"/etc/letsencrypt/fullchain.pem",
		"/etc/letsencrypt/privkey.pem",
		nil)

	if err != nil {
		golog.Error("Server haha failed: ", err)
	}
}
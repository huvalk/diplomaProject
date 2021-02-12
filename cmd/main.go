package main

import (
	"flag"
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

	var (
		env  = flag.String("env", "dev", "enables debug mode")
		sslPath  = flag.String("ssl_path", "/etc/letsencrypt/", "ssl cert path")
		port = flag.Uint64("p", 8080, "port")
	)

	if *env == "dev" {
		golog.SetLevel("debug")
		golog.Debug("Debug")
	}

	golog.Debug(os.Getenv("SSL_PATH"))

	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", *port),
		*sslPath + "fullchain.pem",
		*sslPath + "privkey.pem",
		nil)

	if err != nil {
		golog.Error("Server haha failed: ", err)
	}
}
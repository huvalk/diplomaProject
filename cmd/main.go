package main

import (
	"diplomaProject/application/server"
	"diplomaProject/pkg/infrastructure"
	"flag"
	"github.com/labstack/echo"
	"log"
)


func main() {
	e := echo.New()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conn, err := infrastructure.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	serv := server.NewServer(e, conn)
	//sslPath := flag.String("ssl_path", "/etc/letsencrypt", "ssl cert path")
	flag.Parse()

	log.Fatal(serv.ListenAndServe())
	//log.Fatal(serv.ListenAndServeTLS(*sslPath))
}

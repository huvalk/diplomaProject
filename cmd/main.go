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

	//
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	golog.Debug("Knok-knok")
	//
	//	_, err := w.Write([]byte("I am flying, Jack"))
	//
	//	if err != nil {
	//		golog.Error("Knok failed: ", err)
	//	}
	//})
	//
	//var (
	//	env     = flag.String("env", "test", "enables debug mode")
	//	//sslPath = flag.String("ssl_path", "/etc/letsencrypt", "ssl cert path")
	//	//port    = flag.Uint64("p", 8080, "port")
	//)
	//flag.Parse()
	//
	//if *env != "prod" {
	//	golog.SetLevel("debug")
	//	golog.Debug("Debug")
	//}
	//
	////err := http.ListenAndServeTLS(fmt.Sprintf(":%d", *port),
	////	*sslPath+"/fullchain.pem",
	////	*sslPath+"/privkey.pem",
	////	nil)
	//
	//err := http.ListenAndServe(":8080", nil)
	//
	//if err != nil {
	//	golog.Error("Server haha failed: ", err)
	//}
}

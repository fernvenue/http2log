package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	path string
	key  string
	file string
	addr string
	port string
)

func initFlag() {
	flag.StringVar(&path, "u", "/http2log", "URL path of log interface;")
	flag.StringVar(&key, "k", os.Getenv("KEY"), "HTTP header 'X-API-KEY' content;")
	flag.StringVar(&file, "f", "/var/log/http2log.log", "Path to log file;")
	flag.StringVar(&addr, "a", "127.0.0.1", "Bind address;")
	flag.StringVar(&port, "p", "8283", "Bind port;")
	flag.Parse()
}

func main() {
	initFlag()

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("X-API-KEY")
        if auth != key {
            w.WriteHeader(http.StatusForbidden)
            fmt.Fprintf(w, http.StatusText(http.StatusForbidden)+".\n")
			log.Printf("Authentication failed, rejected.")
            return
        }
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		b := string(body) + "\n"
		defer func() {
            _ = f.Close()
        }()
		if _, err = f.WriteString(b); err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Succeed.\n")
		log.Printf("Authentication succeeded, written.")
	})

	log.Fatal(http.ListenAndServe(net.JoinHostPort(addr, port), nil))
}

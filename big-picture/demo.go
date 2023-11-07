package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	//comment out to run the cli demo
	logCli()

	//Comment out to run WebServer demo
	//startWebServer()
}

func startWebServer() {
	//It really takes this few lines???
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web Services are easy with Go!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	http.ListenAndServe(":3000", nil)
}

func logCli() {
	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	bufReader := bufio.NewReader(f)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}

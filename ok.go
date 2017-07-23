package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	portArg := flag.Int("p", 8080, "port to listen on")

	flag.Parse()

	statusCode := 200
	if flag.NArg() > 0 {
		code, err := strconv.ParseInt(flag.Arg(0), 0, 16)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		statusCode = int(code)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
	})

	address := fmt.Sprintf(":%d", *portArg)
	fmt.Println(http.ListenAndServe(address, nil))
}

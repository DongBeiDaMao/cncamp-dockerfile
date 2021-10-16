package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"fmt"
	"time"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	//return status 200
	w.WriteHeader(http.StatusOK)
	logResponse(http.StatusOK, r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	//copy request to response
	for k, v := range r.Header {
		w.Header().Add(k, getString(v))
	}
	//VERSION_ENV
	v_env := os.Getenv("VERSION")
	w.Header().Add("VERSION", v_env)
	//set status code
	w.WriteHeader(statusCode)
	//log record
	logResponse(statusCode, r)
}

//transfer response value from []string to string
func getString(strArr []string) string {
	var str = ""
	for i, v := range strArr {
		str += v
		if(i != len(strArr) - 1) {
			str += ", "
		}
	}
	return str
}

//return ip address
func getClientIP(addr string) string {
	index := 0
	for i, v := range(addr) {
		if(v == ':') {
			index = i
			break
		}
	}
	return addr[0 : index]
}

func logResponse(status int, r *http.Request) {
	clientIP := getClientIP(r.RemoteAddr)
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf(t + "  From " + clientIP + " status: " + strconv.Itoa(status) + "\n")
}
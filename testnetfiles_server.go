package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

type TestnetfilesServer struct {
	port uint16
}

func NewTestnetfilesServer(port uint16) *TestnetfilesServer {
	return &TestnetfilesServer{port}
}

func (tfs *TestnetfilesServer) Port() uint16 {
	return tfs.port
}

func (tfs *TestnetfilesServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		//io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")

	}
}

func (tfs *TestnetfilesServer) Run() {

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(tfs.Port())), nil))
}

func init() {
	log.SetPrefix("TestnetFilesServer: ")
}

func main() {
	port := flag.Uint("port", 9000, "TCP Port Number for Testnetfiles Server")
	flag.Parse()
	app := NewTestnetfilesServer(uint16(*port))
	app.Run()
}

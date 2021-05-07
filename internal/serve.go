package internal

import (
	"net"
	"net/http"

	"github.com/ka1i/wisper/pkg/utils"
)

func Serve(dir string) *net.TCPAddr {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	server := &http.Server{
		Handler: mux,
	}
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		utils.Fatal(err)
	}

	go server.Serve(ln)

	return ln.Addr().(*net.TCPAddr)
}

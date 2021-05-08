package internal

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ka1i/wisper/pkg/utils"
)

var (
	pwd  = utils.GetHome()
	mime = make(map[string]string)
)

func init() {
	//Content-Type mime
	mime[".svg"] = "image/svg+xml"
	mime[".ico"] = "image/x-icon"
	mime[".html"] = "text/html; charset=UTF-8"
	mime[".css"] = "text/css; charset=UTF-8"
	mime[".js"] = "application/javascript; charset=UTF-8"
}

func Serve() *net.TCPAddr {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", serverHandler{}))

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

type serverHandler struct{}

func (h serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	rq, err := urlParser(r.URL) //r.RequestURI
	if err != nil {
		w.WriteHeader(404)
		log.Println(err)
		return
	}
	//fmt.Println(rq)
	file, _ := os.Open(rq)
	if k, ok := mime[filepath.Ext(file.Name())]; ok {
		w.Header().Set("Content-Type", k)
	}
	io.Copy(w, file)
	defer file.Close()
}

func urlParser(url *url.URL) (string, error) {
	var rq string

	fileinfo, err := os.Lstat(filepath.Join(pwd, url.Path))
	if err != nil {
		return "", err
	}
	if fileinfo.IsDir() {
		rq = filepath.Join(pwd, "/index.html")
		_, err := os.Lstat(rq)
		if err != nil {
			return "", err
		}
	} else {
		rq = filepath.Join(pwd, url.Path)
	}
	return rq, nil
}

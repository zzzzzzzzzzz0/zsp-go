package zsp2

import (
	"net"
	"net/http"
	"os"
)

type Serve___ struct {
	*http.Server
	port string
}

func (this *Serve___) Close__() {
	os.Exit(0)
}

func New_serve__(addr string, handler http.Handler) (srv *Serve___, err error) {
	srv = &Serve___{Server:&http.Server{Addr: addr, Handler: handler}}

	//err = srv.ListenAndServe()
	var l net.Listener
	l, srv.port, err = Listen__(addr, true)
	if err != nil {
		return
	}
	err = srv.Serve(l)
	return
}

func Listen__(addr string, use bool) (l net.Listener, port string, err error) {
	/*if addr == "" {
		addr = ":http"
	}*/
	l, err = net.Listen("tcp", addr)
	if !use {
		defer l.Close()
	}
	if err != nil {
		return
	}

	s := l.Addr().String()
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		if r >= '0' && r <= '9' {
		} else {
			port = s[i + 1:]
			break
		}
	}
	return
}
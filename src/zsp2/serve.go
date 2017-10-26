package zsp2

import (
	"net"
	. "net/http"
	"strings"
	"util4"
)

type Serve___ struct {
	*Server
}

func (this *Serve___) Close__() {
	util4.Exit__(0)
}

func New_serve__(thiz *Zsp___) (srv *Serve___, err error) {
	srv = &Serve___{Server:&Server{Addr: thiz.addr, Handler: thiz}}

	//err = srv.ListenAndServe()

	var l net.Listener
	l, _, err = Listen__(thiz.addr, true)
	if err != nil {
		return
	}

	ip1 := "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, a := range addrs {
			if ipn, ok := a.(*net.IPNet); ok {
				ip := ipn.IP
				if(!ip.IsLoopback()) {
					ip1 = ip.String()
					break
				}
			}
		}
	}
	addr := strings.Replace(l.Addr().String(), "[::]", ip1, -1)
	if addr[0] == ':' {
		addr = ip1 + addr
	}
	thiz.addr = addr
	println("addr " + thiz.addr)
	if thiz.hou_code != "" {
		thiz.zs__(thiz.hou_code, false)
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

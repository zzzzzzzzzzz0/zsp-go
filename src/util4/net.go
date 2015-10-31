package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"net/url"
	"net/http"
	"io/ioutil"
)

func Net__(qv *Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}),
ret__ func(...interface{})) (no_use bool, goto1 *Goto___) {
	switch k {
	case "urlencode":
		for _, s1 := range s {
			s2, ok := s__(s1); if !ok {return}
			ret__(url.QueryEscape(s2))
		}
		return
	case "抓网页":
		if buzu__(1) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		resp, err := http.Get(s0)
		if err != nil {
			ret__("", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ret__("", err)
			return
		}
		ret__(string(body))
		return
	}
	no_use = true
	return
}
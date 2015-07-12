package zsp2

import (
	"net/url"
	"net/http"
	"strings"
	"strconv"
	"regexp"
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"util4"
)

func (this *Zsp___) i__(qv *Qv___,
s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}),
ret__ func(...interface{}), c *util4.Chan___, goto1__ func(*Goto___),
tag string, s ...interface{}) {
	var (
		no_use bool
		goto1 *Goto___
	)
	loop:
	for {
		no_use, goto1 = util4.Util__(qv, tag, s, s__, err__, buzu__, buzhichi__, ret__)
		if !no_use {
			break
		}
		no_use = util4.Chan__(tag, s, s__, buzu__, ret__)
		if !no_use {
			break
		}
		no_use, goto1 = util4.Str__(qv, tag, s, s__, err__, buzu__, buzhichi__, ret__)
		if !no_use {
			break
		}
		no_use, goto1 = util4.File__(qv, tag, s, s__, err__, buzu__, buzhichi__, this.can_stat__, ret__, c)
		if !no_use {
			break
		}
		no_use, goto1 = util4.Regexp__(qv, tag, s, s__, err__, buzu__, buzhichi__, ret__)
		if !no_use {
			break
		}
		no_use, goto1 = util4.Os__(qv, tag, s, s__, err__, buzu__, buzhichi__, ret__)
		if !no_use {
			break
		}
		no_use, goto1 = util4.Net__(qv, tag, s, s__, err__, buzu__, buzhichi__, ret__)
		if !no_use {
			break
		}
		no_use, goto1 = util4.Time__(qv, tag, s, s__, err__, buzu__, buzhichi__, ret__)
		if !no_use {
			break
		}
		no_use, goto1 = this.clpars.Z__(qv, tag, s, s__, err__, ret__)
		if !no_use {
			break
		}
	
		switch tag {
		case "参":
			if data, ok := qv.Not_my.(*data___); ok {
				r := data.r
				if len(s) == 0 {
					ret, _ := url.QueryUnescape(r.URL.RawQuery)
					ret__(ret)
					break loop
				}
				if !data.r_is_parse {
					r.ParseForm()
					data.r_is_parse = true
				}
				name, ok := s__(s[0]); if !ok {break loop}
				ret__(r.Form.Get(name))
			}
			break loop
		case "url":
			if data, ok := qv.Not_my.(*data___); ok {
				ret__(data.r.URL.String())
			}
			break loop
		case "uri":
			if data, ok := qv.Not_my.(*data___); ok {
				ret__(data.r.RequestURI)
			}
			break loop
		case "ip":
			if data, ok := qv.Not_my.(*data___); ok {
				ret := data.r.RemoteAddr
				i := strings.Index(ret, ":")
				if i >= 0 {
					ret__(ret[0:i])
					ret__(ret[i + 1:])
				} else {
					ret__(ret)
				}
			}
			break loop
		case "端口号":
			if this.serve != nil {
				ret__(this.serve.port)
			}
			break loop
		case "空闲端口号":
			_, port, err := Listen__(":0", false)
			if err == nil {
				ret__(port)
			}
			break loop
		case "转向":
			if data, ok := qv.Not_my.(*data___); ok {
				for _, s1 := range s {
					s2, ok := s__(s1); if !ok {break loop}
					http.ServeFile(*data.w, data.r, s2)
				}
			}
			break loop
		case "加根路径":
			for _, s1 := range s {
				s2, ok := s__(s1); if !ok {break loop}
				this.known_path.Add__(s2)
			}
			break loop
		case "伪装":
			if len(s) % 2 != 0 {
				buzu__(-1)
			}
			for i := 0; i < len(s); {
				v, ok := s__(s[i]); if !ok {break loop}
				i++
				k, ok := s__(s[i]); if !ok {break loop}
				i++
				add := true
				for k1, _ := range this.weizhuang {
					if k == k1.re.String() {
						add = false
						break
					}
				}
				if add {
					re, err := regexp.Compile(k)
					if err != nil {
						err__(err)
						break loop
					}
					this.weizhuang[&weizhuang___{re}] = v
				}
			}
			break loop
		case "服务关闭":
			this.serve.Close__()
			break loop
		}
		buzhichi__("")
		break
	}
	if c != nil {
		c.Close_if__()
	}
	goto1__(goto1)
}

func (this *Zsp___) I__(qv *Qv___, s ...interface{}) (goto1 *Goto___, err1 *Errinfo___, ret1 []interface{}) {
	if len(s) == 0 {
		return
	}
	err__ := func(s ...interface{}) {
		err1 = New_errinfo__()
		for _, i := range s {
			err1.Add__(i)
		}
		err1.Add__(Errs_.Fail)
	}
	s__ := func(i interface{}) (s2 string, ok bool) {
		s2, ok = i.(string)
		if !ok {
			//O__("%T", i)
			err__("s转换")
		}
		return
	}
	tag, ok := s__(s[0]); if !ok {return}
	var c *util4.Chan___
	ret__ := func(s ...interface{}) {
		for _, i := range s {
			if err, ok := i.(error); ok {
				ret1 = append(ret1, err)
				continue
			}
			ret1 = append(ret1, i)
		}
	}
	buzu__ := func(i int) bool {
		if i == -1 || len(s) < i {
			err__(tag + " 不足")
			return true
		}
		return false
	}
	if tag == "背后" {
		if buzu__(2) {
			return
		}
		s1, ok := s__(s[1]); if !ok {return}
		n, err := strconv.Atoi(s1)
		if err != nil {
			err__(err)
			return
		}
		c = util4.New_chan__(n)
		s = s[2:]
		tag, ok = s__(s[0]); if !ok {return}
		ret__(c)
	}
	buzhichi__ := func(s ...interface{}) {
		var s1 string
		for _, i := range s {
			s2, ok := s__(i)
			if ok {
				s1 += " " + s2
			}
		}
		err__(tag + s1 + " 不支持")
	}
	goto1__ := func(g *Goto___) {
		goto1 = g
	}
	s = s[1:]
	if c != nil {
		go this.i__(qv, s__, err__, buzu__, buzhichi__, ret__, c, goto1__, tag, s...)
	} else {
		this.i__(qv, s__, err__, buzu__, buzhichi__, ret__, c, goto1__, tag, s...)
	}
	return
}

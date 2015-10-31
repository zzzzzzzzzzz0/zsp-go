package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"regexp"
)

func Regexp__(qv *Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}),
ret__ func(...interface{})) (no_use bool, goto1 *Goto___) {
	re__ := func(i1, i int) *regexp.Regexp {
		if i >= 0 && buzu__(i) {
			return nil
		}
		s1, ok := s__(s[i1]); if !ok {return nil}
		re, err := regexp.Compile(s1)
		if err != nil {
			ret__("", err)
			return nil
		}
		return re
	}
	switch k {
	case "正则配":
		s0, ok := s__(s[0]); if !ok {return}
		for i := 1; i < len(s); i++ {
			re := re__(i, -1)
			if re == nil {
				return
			}
			if re.MatchString(s0) {
				ret__("1")
				return
			}
		}
		return
	case "正则格":
		re := re__(1, 2)
		if re == nil {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		s2 := "$1"
		if len(s) > 2 {
			s2, ok = s__(s[2]); if !ok {return}
		}
		ret__(re.ReplaceAllString(s0, s2))
		return
	case "正则迭":
		re := re__(1, 3)
		if re == nil {
			return
		}
		code, ok := s__(s[2]); if !ok {return}
		var (
			err1 *Errinfo___
			buf *Buf___
			kw *Keyword___
			ret string
		)
		s0, ok := s__(s[0]); if !ok {return}
		for _, s1 := range re.FindAllStringSubmatch(s0, -1) {
			buf, goto1, err1 = Zs2__(code, qv, s1[1:]...)
			if err1 != nil {
				err__(err1)
				return
			}
			ret += buf.S__()
			kw, goto1 = Goto1__(goto1)
			if kw == Kws_.Continue {
				continue
			}
			if kw == Kws_.Break {
				break
			}
			if goto1 != nil {
				break
			}
		}
		ret__(ret)
		return
	case "正则代":
		re := re__(1, 3)
		if re == nil {
			return
		}
		code, ok := s__(s[2]); if !ok {return}
		var (
			err1 *Errinfo___
			buf *Buf___
			kw *Keyword___
			ret string
		)
		s0, ok := s__(s[0]); if !ok {return}
		for i, s1 := range re.FindAllString(s0, -1) {
			buf, goto1, err1 = Zs2__(code, qv, s1, Itoa__(i + 1))
			if err1 != nil {
				err__(err1)
				break
			}
			ret += buf.S__()
			kw, goto1 = Goto1__(goto1)
			if kw == Kws_.Continue {
				continue
			}
			if kw == Kws_.Break {
				break
			}
			if goto1 != nil {
				break
			}
		}
		ret__(ret)
		return
	}
	no_use = true
	return
}

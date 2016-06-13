package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"strconv"
	"math/rand"
	"time"
)

func array_oper(f1__ func(a *Strings___, s1 string), f2__ func(s2 string) int,
s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{})) (ok2 bool) {
	if buzu__(3) {
		return
	}
	a, ok := s[1].(*Strings___)
	if !ok {
		err__("")
		return
	}
	len1 := len(s) - 1
	s1, ok := s__(s[len1]); if !ok {return}
	f1__(a, s1)
	for i := 2; i < len1; i++ {
		s2, ok := s__(s[i]); if !ok {return}
		switch(f2__(s2)) {
		case -1:
			return
		case 1:
			buzhichi__(s[0], s2)
			return
		}
	}
	ok2 = true
	return
}

func Util__(qv *Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}), can_stat__ func(string) bool,
ret__ func(...interface{}), c *Chan___) (no_use bool, goto1 *Goto___) {
	switch k {
	case "随机数":
		//rand.Seed(time.Now().UnixNano())
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		max := 100
		min := 0
		l := len(s)
		var err error
		if l >= 1 {
			s0, ok := s__(s[0]); if !ok {return}
			if s0 != "" {
				max, err = strconv.Atoi(s0)
				if err != nil {
					err__(err)
					return
				}
			}
		}
		if l >= 2 {
			s1, ok := s__(s[1]); if !ok {return}
			if s1 != "" {
				min, err = strconv.Atoi(s1)
				if err != nil {
					err__(err)
					return
				}
			}
		}
		if max < min {
			err__(k + "大小")
			return
		}
		ret__(Itoa__(r.Intn(max - min) + min))
		return
	case "迭代":
		if buzu__(1) {
			return
		}
		var (
			err1 *Errinfo___
			buf *Buf___
			kw *Keyword___
			ret string
		)
		s0, ok := s__(s[0]); if !ok {return}
		for i := 1; i < len(s); i++ {
			si, ok := s__(s[i]); if !ok {return}
			buf, goto1, err1 = Zs2__(s0, qv, si, Itoa__(i))
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
	case "数组":
		if buzu__(1) {
			return
		}
		switch(s[0]) {
		case "新建":
			i := 1
			if len(s) >= 2 {
				s1, ok := s__(s[1]); if !ok {return}
				var err error
				i, err = strconv.Atoi(s1)
				if err != nil {
					err__(err)
					return
				}
			}
			for ; i > 0; i-- {
				ret__(New_strings__())
			}
		case "增至", "不重复增至":
			var (
				a *Strings___
				s1 string
				f2__ func(string) int
			)
			if s[0] == "增至" {
				f2__ = func(s3 string) int {
					a.Add__(s3)
					return 0
				}
			} else {
				f2__ = func(s3 string) int {
					if !a.Find__(func (s2 string) bool {return s2 == s3}) {
						a.Add__(s3)
					}
					return 0
				}
			}
			if(!array_oper(func(a2 *Strings___, s2 string) {a = a2; s1 = s2}, f2__, s, s__, err__, buzu__, buzhichi__)) {
				return
			}
			f2__(s1)
		case "遍历":
			var (
				a *Strings___
				code string
				desc bool
			)
			if(!array_oper(
			func(a2 *Strings___, s2 string) {a = a2; code = s2},
			func(s3 string) int {
				switch(s3) {
				case "逆序":
					desc = true
				default:
					return 1
				}
				return 0
			}, s, s__, err__, buzu__, buzhichi__)) {
				return
			}
			var (
				err1 *Errinfo___
				buf *Buf___
				kw *Keyword___
				ret string
				i int
			)
			if desc {
				i = a.Len__() - 1
			} else {
				i = 0
			}
			for i >= 0 && i < a.Len__() {
				buf, goto1, err1 = Zs2__(code, qv, a.A[i], Itoa__(i + 1))
				if err1 != nil {
					err__(err1)
					break
				}
				if desc {
					i--
				} else {
					i++
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
		case "搜索":
			var (
				a *Strings___
				s1 string
			)
			if(!array_oper(
			func(a2 *Strings___, s2 string) {a = a2; s1 = s2},
			func(string) int {return 1},
			s, s__, err__, buzu__, buzhichi__)) {
				return
			}
			for i, s2 := range a.A {
				if s2 == s1 {
					ret__(Itoa__(i + 1))
					return
				}
			}
		case "长度", "得", "设":
			a, ok := s[1].(*Strings___)
			if !ok {
				err__("")
				return
			}
			switch s[0] {
			case "长度":
				ret__(Itoa__(a.Len__()))
			case "得":
				for i := 2; i < len(s); i++ {
					si, ok := s__(s[i]); if !ok {return}
					i2, err := strconv.Atoi(si)
					if err != nil {
						err__(err)
						return
					}
					ret__(a.A[i2 - 1])
				}
			}
		default:
			buzhichi__(s[0])
		}
		return
	case "遍历变量":
		if buzu__(1) {
			return
		}
		code, ok := s__(s[0]); if !ok {return}
		var (
			err1 *Errinfo___
			buf *Buf___
		)
		fa__ := func (a *Strings___) (s string) {
			a.Find__(func (s2 string) bool {
				s += Kws_.Kaifangkuohao.String() + s2 + Kws_.Bifangkuohao.String()
				return false
			})
			return
		}
		fb__ := func (b bool) string {
			if(b) {
				return "1"
			} else {
				return "0"
			}
		}
		qv2 := qv
		var (
			ret string
			i int
		)
		for {
			if qv2 == nil {
				break
			}
			if qv2.Vars.Find__(func(v *Var___) bool {
				kw := v.Kw.String()
				tm := time.Unix(v.Timestamp, 0)
				timestamp := tm.Format("050415-02012006")
				annota_val := fa__(v.Annota_val)
				is_lock := fb__(v.Is_lock)
				i2 := Itoa__(i)
				switch v.Kw {
				case Kws_.Def:
					buf, goto1, err1 = Zs2__(code, qv,
						timestamp, kw, is_lock, i2, v.Name, v.Val.S, annota_val,
						fb__(v.Is_no_arg), fb__(v.Is_through), Itoa__(v.Qian_arg))
				default:
					buf, goto1, err1 = Zs2__(code, qv, timestamp, kw, is_lock, i2, v.Name, v.Val.S, annota_val)
				}
				if err1 != nil {
					err__(err1)
					return true
				}
				if goto1 != nil {
					return true
				}
				ret += buf.S__()
				return false
			}) {
				break
			}
			qv2 = qv2.Up
			i++
		}
		ret__(ret)
		return
	}
	no_use = true
	return
}
package util4

import (
	"github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"strconv"
	"math/rand"
	"time"
)

func Util__(qv *zhscript.Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}),
ret__ func(...interface{})) (no_use bool, goto1 *zhscript.Goto___) {
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
		ret__(strconv.Itoa(r.Intn(max - min) + min))
		return
	case "迭代":
		var (
			err1 *zhscript.Errinfo___
			buf *zhscript.Buf___
			kw *zhscript.Keyword___
			ret string
		)
		s0, ok := s__(s[0]); if !ok {return}
		for i := 1; i < len(s); i++ {
			si, ok := s__(s[i]); if !ok {return}
			buf, goto1, err1 = Zs2__(s0, qv, si)
			if err1 != nil {
				err__(err1)
				break
			}
			ret += buf.S__()
			kw, goto1 = Goto1__(goto1)
			if kw == zhscript.Kws_.Continue {
				continue
			}
			if kw == zhscript.Kws_.Break {
				break
			}
			if goto1 != nil {
				break
			}
		}
		ret__(ret)
		return
	case "遍历变量":
		if buzu__(1) {
			return
		}
		var err1 *zhscript.Errinfo___
		var buf *zhscript.Buf___
		code, ok := s__(s[0]); if !ok {return}
		fa := func (a *zhscript.Strings___) (s string) {
			a.Find__(func (s2 string) bool {
				s += zhscript.Kws_.Kaifangkuohao.String() + s2 + zhscript.Kws_.Bifangkuohao.String()
				return false
			})
			return
		}
		fb := func (b bool) string {
			if(b) {
				return "1"
			} else {
				return "0"
			}
		}
		var ret string
		qv2 := qv
		var i int
		for {
			if qv2 == nil {
				break
			}
			if qv2.Vars.Ls.Find__(func(e *zhscript.Em___) bool {
				v := zhscript.Var__(e)
				buf, goto1, err1 = Zs2__(code, qv,
					v.Name, v.Val.S,
					fa(v.Annota_val),
					fb(v.Is_lock),
					v.Kw.String(),
					strconv.Itoa(i))
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
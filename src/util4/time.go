package util4

import (
	"github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"time"
	"strconv"
)

func Time__(qv *zhscript.Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}),
ret__ func(...interface{})) (no_use bool, goto1 *zhscript.Goto___) {
	switch k {
	case "时间":
		if len(s) < 1 {
			ret__(time.Now().String())
		} else {
			s0, ok := s__(s[0]); if !ok {return}
			ret__(time.Now().Format(s0))
		}
		return
	case "等待":
		if buzu__(1) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		d := time.Millisecond
		for {
			if Ends__(s0, "s") {
				d = time.Second
			} else if Ends__(s0, "m") {
				d = time.Minute
			} else if Ends__(s0, "h") {
				d = time.Hour
			} else {
				break
			}
			s0 = s0[0:len(s0) - 1]
			break
		}
		n, err := strconv.Atoi(s0)
		if err != nil {
			err__(err)
		}
		time.Sleep(time.Duration(n) * d)
		return
	}
	no_use = true
	return
}
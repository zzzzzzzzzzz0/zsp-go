package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"time"
	"strconv"
	"errors"
)

func get_time_dur__(i interface{}) (d time.Duration, err error) {
	s0, ok := i.(string)
	if !ok {
		err = errors.New("s转换")
		return
	}

	d = time.Millisecond
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
	var n int
	n, err = strconv.Atoi(s0)
	if err != nil {
		return
	}
	d *= time.Duration(n)
	return
}

func timer1__(d time.Duration, code string, one bool, qv *Qv___, c *Chan___) {
	timer1 := time.NewTicker(d)
	defer timer1.Stop()
	var (
		err1 *Errinfo___
		buf *Buf___
		kw *Keyword___
		goto1 *Goto___
	)
	for {
		select {
		case <-timer1.C:
			if c != nil && c.is_close {
				return
			}
			if qv != nil {
				buf, goto1, err1 = Zs2__(code, qv)
				if err1 != nil {
					if c != nil {
						c.Err__(err1)
					} else {
						println(err1)
					}
					return
				}
				s2 := buf.S__()
				if c != nil {
					select {
					case c.o <- s2:
					case <- c.x:
						return
					}
				}
				kw, goto1 = Goto1__(goto1)
				if goto1 != nil || kw == Kws_.Break {
					if c != nil {
						c.Goto__(goto1)
					}
					return
				}
			} else {
				println(code)
				select {
				case ret, ok := <- c.o:
					if ok {
						println(ret)
					}
				case g, ok := <- c.goto1:
					if ok {
						println(g)
					}
				case err, ok := <- c.err:
					if ok {
						println(err)
					}
				case <- c.x:
					c.close__()
					return
				}
			}
			if one {
				if c != nil {
					c.x__()
				}
				return
			}
		}
	}
}
func Time__(qv *Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}),
ret__ func(...interface{}), c *Chan___) (no_use bool, goto1 *Goto___) {
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
		d, err := get_time_dur__(s[0])
		if err != nil {
			err__(err)
			return
		}
		time.Sleep(d)
		return
	case "定时器", "一次定时器":
		if buzu__(2) {
			return
		}
		d, err := get_time_dur__(s[0])
		if err != nil {
			err__(err)
			return
		}
		code, ok := s__(s[1]); if !ok {return}
		if c != nil {
			c.use = true
		}
		timer1__(d, code, k == "一次定时器", qv, c)
		return
	}
	no_use = true
	return
}
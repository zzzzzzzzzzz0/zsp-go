package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"strings"
	"strconv"
)

func Ends__(s0, s1 string) bool {
	return strings.HasSuffix(s0, s1)
}
func Starts__(s0, s1 string) bool {
	return strings.HasPrefix(s0, s1)
}

func Itoa__(i int) string {
	return strconv.Itoa(i)
}

func Str__(qv *Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}), can_stat__ func(string) bool,
ret__ func(...interface{}), c *Chan___) (no_use bool, goto1 *Goto___) {
	switch k {
	case "尾匹配":
		if buzu__(2) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		s1, ok := s__(s[1]); if !ok {return}
		if Ends__(s0, s1) {
			ret__("1")
		}
		return
	case "头匹配":
		if buzu__(2) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		s1, ok := s__(s[1]); if !ok {return}
		if Starts__(s0, s1) {
			ret__("1")
		}
		return
	case "子串":
		if buzu__(2) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		s1, ok := s__(s[1]); if !ok {return}
		start, err := strconv.Atoi(s1)
		if err != nil {
			err__(err)
			return
		}
		r := []rune(s0)
		l := len(r)
		end := l
		if len(s) > 2 {
			s1, ok = s__(s[2]); if !ok {return}
			end, err = strconv.Atoi(s1)
			if err != nil {
				err__(err)
				return
			}
			if end < 0 {
				end += l
			}
		}
		if start < 0 {
			start += l
			if start < 0 {
				start = 0
			}
		}
		var ret string
		for i := start; i < l && i < end; i++ {
			ret += string(r[i])
		}
		ret__(ret)
		return
	case "大写", "小写":
		if buzu__(1) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		if k == "大写" {
			ret__(strings.ToUpper(s0))
		} else {
			ret__(strings.ToLower(s0))
		}
		return
	case "省略":
		if buzu__(1) {
			return
		}
		filename, ok := s__(s[0]); if !ok {return}
		var (
			rr [][]rune
			r1, r2 []rune
			s1 string
		)
		i1 := strings.LastIndex(filename, "/")
		i2 := strings.LastIndex(filename, ".")
		if i2 > i1 {
			s1 = filename[0:i2]
		} else {
			s1 = filename
		}
		index := []rune("index")
		rr_add__ := func() {
			if r1 != nil {
				if _, ok := Startswith__(r1, index, 0); ok {
				} else {
					rr = append(rr, r1)
				}
				r1 = nil
			}
		}
		for _, r := range s1 {
			switch r {
			case '/':
				rr_add__()
				continue
			}
			r1 = append(r1, r)
		}
		rr_add__()
		for1:
		for i, r3 := range rr {
			if len(r3) == 0 {
				continue
			}
			for i1 := 0; i1 < i; i1++ {
				r4 := rr[i1]
				len4 := len(r4)
				if len4 == 0 {
					continue
				}
				/*if _, ok := zhscript.Startswith__(r3, r4, 0); ok {
					if len4 == len(r3) {
						continue for1
					}
					r3 = r3[len4:]
					break
				}*/
				i2 := 0
				len3 := len(r3)
				for ; i2 < len4 && i2 < len3; i2++ {
					if r4[i2] != r3[i2] {
						break
					}
				}
				if i2 > 0 {
					if i2 == len3 {
						continue for1
					}
					r3 = r3[i2:]
				}
			}
			if r2 != nil {
				r2 = append(r2, '-')
			}
			r2 = append(r2, r3...)
		}
		ret__(string(r2))
		return
	case "符号换":
		if buzu__(2) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		var start rune
		{
			s1, ok := s__(s[1]); if !ok {return}
			start = []rune(s1)[0]
		}
		var (
			ret2 string
			i int
		)
		l := len(s)
		for _, r := range s0 {
			i = int(r - start) + 2
			if i >= 2 && i < l {
				si, ok := s__(s[i]); if !ok {return}
				ret2 += si
			} else {
				ret2 += string(r)
			}
		}
		ret__(ret2)
		return
	case "数换":
		if buzu__(2) {
			return
		}

		s0, ok := s__(s[0]); if !ok {return}
		i, err := strconv.ParseInt(s0, 10, 0)
		if err != nil {
			err__(err)
			return
		}

		s1, ok := s__(s[1]); if !ok {return}
		r := []rune(s1)
		jinzhi := len(r)
		if jinzhi < 2 {
			err__(Itoa__(jinzhi) + "进制")
			return
		}

		var ret string
		jinzhi1 := int64(jinzhi)
		for i > 0 {
			ret = string(r[i % jinzhi1]) + ret
			i /= jinzhi1
		}
		if ret == "" {
			ret = string(r[0])
		}
		ret__(ret)
		return
	}
	no_use = true
	return
}
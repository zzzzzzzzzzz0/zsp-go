package util4

import (
	"strings"
	"strconv"
)

func file_okname(buzu__ func(int) bool, s []interface{}, s__ func(interface{}) (string, bool), ret__ func(...interface{})) {
	if buzu__(2) {
		return
	}
	filename, ok := s__(s[0]); if !ok {return}
	if Ends__(filename, "/") {
		filename = filename[0:len(filename) - 1]
	}
	i := strings.LastIndex(filename, "/")
	if i >= 0 {
		filename = filename[i + 1:]
	}
	if filename == "" {
		return
	}
	dir, ok := s__(s[1]); if !ok {return}
	dir = Dir__(dir)
	file := dir + filename
	var (
		ext string
		minw int
	)
	i = -1
	for {
		if !Exist_file__(file) {
			break
		}
		if i == -1 {
			s1 := File_main_name__(filename)
			if s1 != "" {
				if filename != s1 {
					ext = filename[len(s1):]
				}
				filename = s1
			}
			i = 0
			l := len(filename)
			i2 := l - 1
			var b bool
			for ; i2 >= 0; i2-- {
				d := filename[i2]
				if !('0' <= d && d <= '9') {
					b = true
					break
				}
			}
			if b || i2 == -1 {
				i2++
				minw = l - i2
				i, _ = strconv.Atoi(filename[i2:])
				filename = filename[0:i2]
			}
		}

		i++
		a := strconv.Itoa(i)
		for len(a) < minw {
			a = "0" + a
		}

		file = dir + filename + a + ext
	}
	ret__(file)
}
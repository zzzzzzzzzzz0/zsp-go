package util4

import (
	"github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"path/filepath"
	"os"
	"strings"
	"strconv"
	"io"
	"io/ioutil"
	"bufio"
)

func Dir__(dir string) string {
	dir = filepath.Clean(dir)
	if !Ends__(dir, "/") {
		dir += "/"
	}
	return dir
}

func Get_dir__(s string) string {
	return filepath.Dir(s)
}

func Exist_file__(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func File_main_name__(s string) (s1 string) {
	s1 = filepath.Base(s)
	i1 := strings.LastIndex(s1, ".")
	if i1 >= 0 {
		s1 = s1[0:i1]
	}
	return
}

func File__(qv *zhscript.Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}), can_stat__ func(string) bool,
ret__ func(...interface{}), c *Chan___) (no_use bool, goto1 *zhscript.Goto___) {
	switch k {
	case "目录名":
		if buzu__(1) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		ret__(Get_dir__(s0))
		return
	case "文件":
		if buzu__(1) {
			return
		}
		filename, ok := s__(s[0]); if !ok {return}
		var (
			fi os.FileInfo
			err error
		)
		stat := func() bool {
			if fi == nil {
				fi, err = os.Lstat(filename)
			}
			return !os.IsNotExist(err)
		}
		for i := 1; i < len(s); i++ {
			switch s[i] {
			case "在":
				if stat() {
					ret__("1")
				}
			case "类型":
				if stat() {
					if (fi.Mode() & os.ModeSymlink) != 0 {
						ret__("l")
					} else if fi.IsDir() {
						ret__("d")
					} else if fi.Mode().IsRegular() {
						ret__("f")
					} else {
						ret__("?")
					}
				}
			case "目录":
				ret__(Get_dir__(filename))
			case "名":
				ret__(filepath.Base(filename))
			case "副名":
				ret__(filepath.Ext(filename))
			case "主名":
				ret__(File_main_name__(filename))
			case "绝对":
				s1, err := filepath.Abs(filename)
				if err != nil {
					err__(err)
					return
				}
				ret__(s1)
			case "链接":
				s1, err := os.Readlink(filename)
				if err == nil {
					ret__(s1)
				}
			case "净":
				ret__(filepath.Clean(filename))
			case "大小":
				if stat() {
					ret__(strconv.FormatInt(fi.Size(), 10))
				}
			case "权限":
				if stat() {
					ret__(fi.Mode().String())
				}
			case "修改时间":
				if stat() {
					len1 := len(s) - 1
					if i == len1 {
						ret__(fi.ModTime().String())
					} else {
						if s1, ok := s__(s[len1]); ok {
							ret__(fi.ModTime().Format(s1))
						}
					}
					return
				}
			case "读":
				var (
					begin, end int64
					use_begin, use_end, head bool
					err error
				)
				for i++; i < len(s); i++ {
					var b rune
					switch s[i] {
					case "始", "终":
						si, _ := s__(s[i])
						b = []rune(si)[0]
					case "头":
						head = true
						continue
					}
					if b == 0 {
						buzhichi__(s[i])
						return
					}
					i++
					if buzu__(i + 1) {
						return
					}
					si, ok := s__(s[i]); if !ok {return}
					switch b {
					case '始':
						begin, err = strconv.ParseInt(si, 10, 0)
						if err != nil {
							err__(err)
							return
						}
						use_begin = true
					case '终':
						end, err = strconv.ParseInt(si, 10, 0)
						if err != nil {
							err__(err)
							return
						}
						use_end = true
					}
				}

				if use_begin || use_end || head {
					f, err := os.Open(filename)
					if err != nil {
						err__(err)
						return
					}
					defer f.Close()

					if head {
						br := bufio.NewReader(f)
						ret, err := br.ReadBytes('\n')
						if err != nil {
							if err != io.EOF {
								err__(err)
							}
							return
						}
						ret__(string(ret[:len(ret) - 1]))
						return
					}

					var (
						buf = make([]byte, 32)
						pos = begin
						ret string
					)
					for {
						n, err := f.ReadAt(buf, pos)
						if err == io.EOF {
							break
						}
						if err != nil {
							err__(err)
							return
						}
						if use_end && pos + int64(n) > end {
							n = int(end - pos + 1)
						}
						if n <= 0 {
							break
						}
						ret += string(buf[:n])
						pos += int64(n)
						if use_end && pos > end {
							break
						}
					}
					ret__(ret)
					return
				}

				ret, err := ioutil.ReadFile(filename)
				if err != nil {
					err__(err)
					return
				}
				ret__(string(ret))
				return
			case "写":
				if buzu__(i + 2) {
					return
				}
				len1 := len(s) - 1
				s1, ok := s__(s[len1]); if !ok {return}
				
				flags := os.O_CREATE | os.O_WRONLY
				for i++; i < len1; i++ {
					switch s[i] {
					case "追加":
						flags |= os.O_APPEND
					default:
						buzhichi__(filename, s[i])
						return
					}
				}
				if flags & os.O_APPEND == 0 {
					os.Remove(filename)
				}
				f, err := os.OpenFile(filename, flags, 0660)
				if err != nil {
					err__(err)
					return
				}
				defer f.Close()
				_, err = f.WriteString(s1)
				if err != nil {
					err__(err)
					return
				}
				return
			default:
				buzhichi__(s[i])
				return
			}
		}
		return
	case "遍历目录":
		if buzu__(2) {
			return
		}
		dir, ok := s__(s[0]); if !ok {return}
		dir = Dir__(dir)
		if !can_stat__(dir) {
			ret__("", dir + " 拒绝")
			return
		}
		fi, err := os.Stat(dir)
		if os.IsNotExist(err) {
			ret__("", dir + " 不存在")
			return
		}
		if !fi.IsDir() {
			ret__("", dir + " 不是目录")
			return
		}
		
		var inc_subdir, ret_dir, only_dir, only_root_file bool
		for i1 := 1; i1 < len(s) - 1; i1++ {
			switch s[i1] {
			case "含子目录":
				inc_subdir = true
			case "返回目录":
				ret_dir = true
			case "仅目录":
				only_dir = true
			case "仅根文件":
				only_root_file = true
			default:
				buzhichi__(s[i1])
				return
			}
		}

		var (
			err1 *zhscript.Errinfo___
			buf *zhscript.Buf___
			kw *zhscript.Keyword___
			ret string
		)
		code, ok := s__(s[len(s) - 1]); if !ok {return}

		err2 := filepath.Walk(dir, func(path string, fi2 os.FileInfo, err error) error {
			if fi2 == nil {
				return err
			}
			if fi2.IsDir() {
				if path == dir {
					return nil
				}
				if !ret_dir {
					if !inc_subdir {
						return filepath.SkipDir
					}
					return nil
				}
			} else {
				if only_dir {
					return nil
				}
			}
			if Starts__(path, dir) {
				path = path[len(dir):]
			}
			if fi2.IsDir() {
				if ret_dir {
					if !inc_subdir && strings.Contains(path, "/") {
						return filepath.SkipDir
					}
					path = Dir__(path)
				}
			} else {
				if only_root_file && strings.Contains(path, "/") {
					return nil
				}
			}

			buf, goto1, err1 = Zs2__(code, qv, path)
			if err1 != nil {
				return err1
			}
			s2 := buf.S__()
			if c != nil {
				if s2 != "" {
					select {
					case c.o <- s2:
					case <- c.x:
						return nil
					}
				}
			} else {
				ret += s2
			}
			kw, goto1 = Goto1__(goto1)
			return nil
		})
		if c != nil {
			c.o <- ""
		}
		ret__(ret)
		if err2 != nil {
			ret__(dir + " " + err2.Error())
			return
		}
		return
	case "合适文件名":
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
		return
	}
	no_use = true
	return
}
package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"path/filepath"
	"os"
	"strings"
	"strconv"
	"io"
	"io/ioutil"
	"bufio"
	"errors"
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

func Exist_file2__(path string) (ok, is_dir, is_symlink bool) {
	fi, err := os.Lstat(path)
	ok = err == nil
	if ok {
		is_dir = fi.IsDir()
		is_symlink = (fi.Mode() & os.ModeSymlink) != 0
		if is_symlink {
			path2, err2 := os.Readlink(path)
			if err2 == nil {
				if !Starts__(path2, "/") {
					path2 = path + "/" + path2
				}
				ok2, _, is_dir2 := Exist_file2__(path2)
				if ok2 {
					is_dir = is_dir2
				}
			}
		}
	}
	return
}

func File_main_name__(s string) (s1 string) {
	s1 = filepath.Base(s)
	i1 := strings.LastIndex(s1, ".")
	if i1 >= 0 {
		s1 = s1[0:i1]
	}
	return
}

func File__(qv *Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}), can_stat__ func(string) bool,
ret__ func(...interface{}), c *Chan___) (no_use bool, goto1 *Goto___) {
	switch k {
	case "文件":
		if buzu__(1) {
			return
		}
		filename, ok := s__(s[0]); if !ok {return}
		var (
			fi os.FileInfo
			err error
		)
		stat__ := func() bool {
			if fi == nil {
				fi, err = os.Lstat(filename)
			}
			return !os.IsNotExist(err)
		}
		for i := 1; i < len(s); i++ {
			switch s[i] {
			case "在":
				if stat__() {
					ret__("1")
				}
			case "类型":
				if stat__() {
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
				if stat__() {
					ret__(strconv.FormatInt(fi.Size(), 10))
				}
			case "权限":
				if stat__() {
					ret__(fi.Mode().String())
				}
			case "修改时间":
				if stat__() {
					if i == len(s) - 1 {
						ret__(fi.ModTime().String())
					} else {
						i++
						s1, ok := s__(s[i]); if !ok {return}
						ret__(fi.ModTime().Format(s1))
					}
				}
			case "读":
				var (
					begin, end int64
					use_begin, use_end, head bool
					err error
				)
				for i++; i < len(s); i++ {
					var b string
					switch s[i] {
					case "始", "终":
						b, _ = s__(s[i])
					case "头":
						head = true
						continue
					default:
						buzhichi__(s[i])
						return
					}
					i++
					if buzu__(i + 1) {
						return
					}
					si, ok := s__(s[i]); if !ok {return}
					switch b {
					case "始":
						begin, err = strconv.ParseInt(si, 10, 0)
						if err != nil {
							err__(err)
							return
						}
						use_begin = true
					case "终":
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
				i_old := i
				for i++; i < len1; i++ {
					switch s[i] {
					case "追加":
						flags |= os.O_APPEND
					default:
						buzhichi__(filename, s[i_old], s[i])
						return
					}
				}
				if flags & os.O_APPEND == 0 {
					os.Remove(filename)
				}
				err := os.MkdirAll(Get_dir__(filename), os.ModeDir | 0750)
				if err != nil {
					err__(err)
					return
				}
				f, err := os.OpenFile(filename, flags, 0640)
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
	case "目录名":
		if buzu__(1) {
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		ret__(Get_dir__(s0))
		return
	case "遍历目录":
		goto1 = file_for_dir(buzu__, s, s__, can_stat__, ret__, buzhichi__, qv, c)
		return
	case "合适文件名":
		file_okname(buzu__, s, s__, ret__)
		return
	}
	no_use = true
	return
}

var file_walk_break_ = errors.New("")

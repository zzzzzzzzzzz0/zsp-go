package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"os"
	"path/filepath"
	"strings"
)

func file_for_dir(buzu__ func(int) bool, s []interface{}, s__ func(interface{}) (string, bool),
can_stat__ func(string) bool, ret__ func(...interface{}), buzhichi__ func(...interface{}),
qv *Qv___, c *Chan___) (goto1 *Goto___) {
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
		err1 *Errinfo___
		buf *Buf___
		kw *Keyword___
		ret string
	)
	code, ok := s__(s[len(s) - 1]); if !ok {return}

	if c != nil {
		c.use = true
	}
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
			select {
			case c.o <- s2:
			case <- c.x:
				return file_walk_break_
			}
		} else {
			ret += s2
		}
		kw, goto1 = Goto1__(goto1)
		if goto1 != nil || kw == Kws_.Break {
			if c != nil {
				c.Goto__(goto1)
			}
			return file_walk_break_
		}
		return nil
	})
	ret__(ret)
	switch err2 {
		case nil, file_walk_break_:
		default:
		ret2 := dir + " " + err2.Error()
		if c != nil {
			c.o <- ret2
		} else {
			ret__(ret2)
		}
	}
	if c != nil {
		if err2 != file_walk_break_ {
			c.x__()
		}
	}
	return
}
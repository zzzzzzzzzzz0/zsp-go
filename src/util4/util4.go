package util4

import (
	"strings"
	"strconv"
	"os"
	"path/filepath"
	"os/exec"
	"regexp"
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
)

func Zs__(src string, src_is_file bool, up_qv *Qv___, r interface{}, s ...string) (buf *Buf___, goto1 *Goto___, err *Errinfo___) {
	var args Args___
	if src_is_file {
		args.Src_file__(src)
	} else {
		args.Src_code__(src)
	}
	args.Add__(s...)
	buf = New_buf__()
	var qv *Qv___
	qv, err = New_qv__(&args, up_qv)
	if err == nil {
		qv.Not_my = r
		goto1, err = qv.Z__(0, buf)
	}
	return
}

func Dir__(dir string) string {
	if !Ends__(dir, "/") {
		dir += "/"
	}
	return dir
}

func Ends__(s0, s1 string) bool {
	return strings.HasSuffix(s0, s1)
}
func Starts__(s0, s1 string) bool {
	return strings.HasPrefix(s0, s1)
}

func Util__(qv *Qv___, k string, s []string, err__ func(s string), buzu__ func(i int) bool) (no_use bool, ret, ret_err string, goto1 *Goto___) {
	switch k {
		case "得环境变量":
		if buzu__(1) {
			return
		}
		ret = os.Getenv(s[0])
		return

		case "尾匹配":
		if buzu__(2) {
			return
		}
		if Ends__(s[0], s[1]) {
			ret = "1"
		}
		return
		case "头匹配":
		if buzu__(2) {
			return
		}
		if Starts__(s[0], s[1]) {
			ret = "1"
		}
		return

		case "正则配":
		if buzu__(2) {
			return
		}
		for i := 1; i < len(s); i++ {
			re, err := regexp.Compile(s[i])
			if err != nil {
				ret_err = err.Error()
				return
			}
			ret += re.FindString(s[0])
		}
		return

		case "得目录名":
		if buzu__(1) {
			return
		}
		ret = filepath.Dir(s[0])
		return

		case "遍历目录":
		if buzu__(2) {
			return
		}
		dir := Dir__(s[0])
		fi, err := os.Stat(dir)
		if os.IsNotExist(err) {
			ret_err = (dir + " 不存在")
			return
		}
		if !fi.IsDir() {
			ret_err = (dir + " 不是目录")
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
				err__(s[i1] + " 不识别选项")
				return
			}
		}
	
		var err1 *Errinfo___
		var buf *Buf___
		code := s[len(s) - 1]
		
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
			path = path[len(dir):]
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
	
			buf, goto1, err1 = Zs__(code, false, qv, qv.Not_my, path)
			if err1 != nil {
				return err1
			}
			if goto1 != nil && goto1.Kw == Kws_.Continue {
				goto1 = nil
			}
			ret += buf.String()
			return nil
		})
		if err2 != nil {
			err__(dir + " " + err2.Error())
			return
		}
		return

		case "遍历变量":
		if buzu__(1) {
			return
		}
		var err1 *Errinfo___
		var buf *Buf___
		code := s[0]
		fa := func (a *List___) (s string) {
			a.Find__(func (e *Em___) bool {
				s += Kws_.Kaifangkuohao.String() + e.String() + Kws_.Bifangkuohao.String()
				return false
			});
			return
		};
		fb := func (b bool) string {
			if(b) {
				return "1"
			} else {
				return "0"
			}
		}
		For_var__(qv, func(v *Var___, i int) bool {
			buf, goto1, err1 = Zs__(code, false, qv, qv.Not_my,
				v.Name, v.Val,
				fa(v.Annota), fa(v.Annota_val),
				fb(v.Is_lock), fb(v.Is_no_arg),
				v.Kw.String(), strconv.Itoa(i))
			if err1 != nil {
				err__(err1.Error())
				return true
			}
			if goto1 != nil {
				return true
			}
			return Bool__(buf.String())
		});
		return
		
		case "重定向输出":
		if buzu__(1) {
			return
		}
		cmd := exec.Command(s[0], s[1:]...)
		buf, err := cmd.Output()
		if err != nil {
			ret_err = err.Error()
		}
		ret = string(buf)
		return
	}
	no_use = true
	return
}
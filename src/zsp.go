package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
    "strconv"
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
)

var s_begin_ = Kws_.Begin_yuanyang.String()
var s_end_ = Kws_.End_yuanyang.String() + Kws_.Juhao.String()
var r_ = strings.NewReplacer("<%", s_end_, "%>", s_begin_)
func content_convert__(b []byte) []byte {
	s := string(b)
	s = r_.Replace(s)
	s = s_begin_ + s + s_end_
	return []byte(s)
}

func dir__(dir string) string {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	return dir
}

type zsp___ struct {
	main_qv *Qv___
	z *Zhscript___
	addr, index string
}

//r.URL.RawQuery
func (this *zsp___) I__(qv *Qv___, s ...string) (goto1 *Goto___, err1 *Errinfo___, ret string) {
	if len(s) == 0 {
		return
	}
	tag := s[0]
	err__ := func(s string) {
		err1 = New_errinfo__(s, Errs_.Fail)
	}
	buzu__ := func(i int) bool {
		if len(s) < i {
			err__(tag + " 不足")
			return true
		}
		return false
	}
	switch tag {
	case "得参":
		if buzu__(2) {
	        return
		}
		if r, ok := qv.Not_my.(*http.Request); ok {
			r.ParseForm()
			ret = r.Form.Get(s[1])
		}
		return
	case "得环境变量":
		if buzu__(2) {
	        return
		}
		ret = os.Getenv(s[1])
		return
	case "尾匹配":
		if buzu__(3) {
	        return
		}
		if strings.HasSuffix(s[1], s[2]) {
			ret = "1"
		}
		return
	case "遍历目录":
		if buzu__(3) {
	        return
		}
		dir := dir__(s[1])
	    fi, err := os.Stat(dir)
	    if os.IsNotExist(err) {
	    	err__(dir + " 不存在")
	        return
	    }
	    if !fi.IsDir() {
	    	err__(dir + " 不是目录")
	        return
	    }
    	code := s[2]
    	start := -1
    	length := 0
    	if len(s) > 4 {
    		start, _ = strconv.Atoi(s[3])
    		length, _ = strconv.Atoi(s[4])
    	}
    	i2 := 0
		err2 := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			path = path[len(dir):]
			if path == "" {
				return nil
			}

			i2++
			if start >= 0 && length >= 0 {
				if i2 < start || i2 >= start + length {
					return nil
				}
			}

			var buf *Buf___
			buf, goto1, err1 = this.zz__(code, false, qv, qv.Not_my, path)
			if err1 == nil {
				ret += buf.String()
			}
			return nil
		})
		if err2 != nil {
			err__(s[1] + " " + err2.Error())
	        return
        }
    	if len(s) > 5 {
			var buf *Buf___
			buf, goto1, err1 = this.zz__(s[5], false, qv, qv.Not_my, strconv.Itoa(i2))
			if err1 == nil {
				ret += buf.String()
			}
    	}
	    return
	case "配置":
		this.parse__(s[1:])
		return
	}
	err__(tag + " 不支持")
	return
}

func (this *zsp___) zz__(src string, src_is_file bool, up_qv *Qv___, r interface{}, s ...string) (buf *Buf___, goto1 *Goto___, err *Errinfo___) {
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

func (this zsp___) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Path
	if strings.HasSuffix(src, "/") {
		src += this.index
	}
	if strings.HasSuffix(src, ".zsp") {
		buf, _, err := this.zz__(src, true, this.main_qv, r)
		if err != nil {
			fmt.Fprint(w, buf, err)
			return
		}
		fmt.Fprint(w, buf)
		return
	}
	if src2, ok := Get_path__(src); ok {
		http.ServeFile(w, r, src2)
		return
	}
	http.NotFound(w, r)
}

func (this *zsp___) parse__(a []string) (err *Errinfo___) {
	tag := []string {"-r", "--root", "-a", "--addr", "-i", "--index",}
	it := -1
	for _, s := range a {
		for i1, s1 := range tag {
			if s1 == s {
				it = i1
				break
			}
		}
		if it < 0 {
			err = New_errinfo__(s, Errs_.Exist)
			break
		}
		switch it {
		case 0, 1:
			Known_path_add__(dir__(s))
		case 2, 3:
			this.addr = s
		case 4, 5:
			this.index = s
		}
	}
	return
}

func (this *zsp___) z__() {
	z, err := New__(os.Args, content_convert__, this)
	this.z = z
	if err == nil {
		this.index = "index.zsp"
		this.addr = "127.0.0.1:4000"
		err = this.parse__(z.Args.A)
		if err == nil {
			this.main_qv, err = z.New_main_qv__(nil)
			if err == nil {
				err = this.main_qv.Set_var__("让我",
				Kws_.Call.String() + "I__" +
				Kws_.Dunhao.String() +
				Kws_.Kaidanyinhao.String() +
				Kws_.Args.String() +
				Kws_.Bidanyinhao.String() +
				Kws_.Juhao.String(),
				true, Kws_.Def)
			}
		}
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	if z.Args.Src_type == Src_is_file_ {
		buf, _, err := this.zz__(z.Args.Src, true, this.main_qv, nil)
		fmt.Print(buf.String())
		if err != nil {
			fmt.Println()
			fmt.Println(err)
		}
	} else {
		err2 := http.ListenAndServe(this.addr, this)
		fmt.Println(err2)
	}
}

func main() {
	(&zsp___{}).z__()
}

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"util4"
)

var s_begin_ = Kws_.Juhao.String() + Kws_.Begin_yuanyang.String()
var s_end_ = Kws_.End_yuanyang.String() + Kws_.Juhao.String()
var r_ = strings.NewReplacer("<%", s_end_, "%>", s_begin_)
func content_convert__(b []byte, src string) []byte {
	if !util4.Ends__(src, ".zsp") {
		return b
	}
	s := string(b)
	s = r_.Replace(s)
	s = s_begin_ + s + s_end_
	return []byte(s)
}

type zsp___ struct {
	main_qv *Qv___
	z *Zhscript___
	addr, index string
}

type data___ struct {
	r *http.Request
	r_is_parse bool
}

//r.URL.RawQuery
func (this *zsp___) I__(qv *Qv___, s ...string) (goto1 *Goto___, err1 *Errinfo___, ret, ret_err string) {
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
	s = s[1:]

	var no_use bool
	no_use, ret, ret_err, goto1 = util4.Util__(qv, tag, s, err__, buzu__)
	if !no_use {
		return
	}

	switch tag {
	case "得参":
		data := qv.Not_my.(*data___)
		r := data.r
		if len(s) == 0 {
			ret, _ = url.QueryUnescape(r.URL.RawQuery)
			return
		}
		if !data.r_is_parse {
			r.ParseForm()
			data.r_is_parse = true
		}
		ret = r.Form.Get(s[0])
		return
	case "配置":
		this.parse__(s, false)
		return
	}
	err__(tag + " 不支持")
	return
}

func (this zsp___) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Path
	if util4.Ends__(src, "/") {
		src += this.index
	}
	if util4.Ends__(src, ".zsp") {
		buf, _, err := util4.Zs__(src, true, this.main_qv, &data___{r:r})
		fmt.Fprint(w, buf.S__())
		if err != nil {
			fmt.Fprint(w, err)
			fmt.Println(err)
		}
		return
	}
	if src2, ok := Get_path__(src); ok {
		http.ServeFile(w, r, src2)
		return
	}
	http.NotFound(w, r)
}

func (this *zsp___) parse__(a []string, shou1 bool) (shou []string, err *Errinfo___) {
	tag := []string {"-r", "--root", "-a", "--addr", "-i", "--index",}
	it := -1
	for1:
	for i := 0; i < len(a); i++ {
		s := a[i]
		for i1, s1 := range tag {
			if s1 == s {
				it = i1
				continue for1
			}
		}
		if it < 0 {
			if shou1 {
				shou = append(shou, s)
				continue
			} else {
				err = New_errinfo__(s, Errs_.Exist)
				break
			}
		}
		switch it {
		case 0, 1:
			Known_path_add__(util4.Dir__(s))
		case 2, 3:
			this.addr = s
			it = -1
		case 4, 5:
			this.index = s
			it = -1
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
		var a []string
		a, err = this.parse__(z.Args.A, true)
		if err == nil {
			var args Args___
			args.Add__(a...)
			this.main_qv, err = z.New_main_qv__(&args)
			if err == nil {
				name := New_buf__()
				name.WriteString("让我")
				val := New_buf__()
				val.WriteString(Kws_.Call.String() + "I__" +
				Kws_.Dunhao.String() +
				Kws_.Kaidanyinhao.String() +
				Kws_.Args.String() +
				Kws_.Bidanyinhao.String() +
				Kws_.Juhao.String())
				err = this.main_qv.Set_var__(name, val, true, false, Kws_.Def)
			}
		}
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	if z.Args.Src_type == Src_is_file_ {
		buf, _, err := util4.Zs__(z.Args.Src, true, this.main_qv, nil)
		fmt.Print(buf.String())
		if err != nil {
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

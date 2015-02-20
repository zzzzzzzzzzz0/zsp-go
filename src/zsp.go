package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"flag"
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
	}
	err__(tag + " 不支持")
	return
}

func (this zsp___) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Path
	if util4.Ends__(src, "/") {
		src += this.index
	}
	if src2, ok := Get_path__(src); ok {
		if util4.Ends__(src, ".zsp") {
			buf, _, err := util4.Zs__(src2, true, this.main_qv, &data___{r:r})
			fmt.Fprint(w, buf.S__())
			if err != nil {
				fmt.Fprint(w, err)
				fmt.Println(err)
			}
		} else {
			http.ServeFile(w, r, src2)
		}
		return
	}
	http.NotFound(w, r)
}

func (this *zsp___) z__() {
	z, err := New__(os.Args, content_convert__, this)
	this.z = z
	if err == nil {
		var root string
		flag.StringVar(&root, "r", ".", "root")
		flag.StringVar(&this.addr, "a", ":4000", "侦听地址")
		flag.StringVar(&this.index, "i", "index.zsp", "索引页")
	    flag.CommandLine.Parse(z.Args.A)
		Known_path_add__(util4.Dir__(root))

		var args Args___
		println(flag.Args())
		args.Add__(flag.Args()...)
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

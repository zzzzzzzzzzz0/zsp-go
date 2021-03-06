package zsp2

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"regexp"
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"util4"
	"clpars4"
)

const _zsp_, _zs_ = ".zsp", ".zs"

var s_begin_ = Kws_.Juhao.String() + Kws_.Begin_yuanyang.String()
var s_end_ = Kws_.End_yuanyang.String() + Kws_.Juhao.String()
var r_ = strings.NewReplacer("<%", s_end_, "%>", s_begin_)
func content_convert__(b []byte, src string) []byte {
	if !util4.Ends__(src, _zsp_) {
		return b
	}
	s := string(b)
	s = r_.Replace(s)
	s = s_begin_ + s + s_end_
	return []byte(s)
}

type Zsp___ struct {
	main_qv *Qv___
	z *Zhscript___
	pgrname, finalpgrname string
	addr, index string
	known_path *Strings___
	weizhuang map[*weizhuang___]string
	clpars *clpars4.C___
	serve *Serve___
	hou_code string
}

type weizhuang___ struct {
	re *regexp.Regexp
}

type data___ struct {
	r *http.Request
	r_is_parse bool
	w *http.ResponseWriter
}

func (this *Zsp___) weizhuang__(src string) string {
	for k, v := range this.weizhuang {
		re := k.re
		if re.MatchString(src) {
			return v
		}
	}
	return src
}

func (this *Zsp___) known_path_add__(s string) {
	for util4.Ends__(s, "/") {
		s = s[0:len(s) - 1]
	}
	if !this.known_path.Find__(func(s2 string)bool {return s == s2}) {
		this.known_path.Add__(s)
	}
}

func (this *Zsp___) can_stat__(src string) bool {
	for _, s := range this.known_path.A {
		if util4.Starts__(src, s) {
			return true
		}
	}
	return false
}

func (this *Zsp___) get_path__(src string) (src2 string, ok, is_dir, is_symlink bool) {
	ok, is_dir, is_symlink = util4.Exist_file2__(src)
	if ok && this.can_stat__(src) {
		src2 = src
		return
	}
	for _, s := range this.known_path.A {
		src2 = s + "/" + src
		ok, is_dir, is_symlink = util4.Exist_file2__(src2)
		if ok {
			return
		}
	}
	return
}

func (this Zsp___) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	src := r.URL.Path
	if util4.Ends__(src, "/") {
		src += this.index
	}
	src = this.weizhuang__(src)
	if src2, ok, is_dir, is_symlink := this.get_path__(src); ok {
		if util4.Ends__(src, _zsp_) {
			buf, goto1, err := util4.Zs__(src2, true, "", this.main_qv, &data___{r:r, w:&w})
			fmt.Fprint(w, buf.S__())
			if goto1 != nil {
				switch goto1.Kw {
				case Kws_.Quit, Kws_.Return:
				default:
					err = New_errinfo__(goto1.S, goto1.Kw, Errs_.Keyword, Kws_.For)
				}
			}
			if err != nil {
				fmt.Fprint(w, err)
				util4.Errln__(err)
			}
			return
		}
		if is_symlink && is_dir {
			http.Redirect(w, r, src2 + "/", http.StatusMovedPermanently)
			return
		}
		http.ServeFile(w, r, src2)
		return
	}
	http.NotFound(w, r)
}

func (this *Zsp___) set_main_qv_var__(name2, val2 string) {
	name := New_buf__()
	name.WriteString(name2)
	val := New_buf__()
	val.WriteString(val2)
	this.main_qv.Set_var__(name, val, nil, Kws_.Set)
}

func (this *Zsp___) zs__(src string, src_is_file bool) {
	buf, goto1, err := util4.Zs__(src, src_is_file, "",
		this.main_qv, nil, this.main_qv.Args.A__()...)
	fmt.Print(buf.S__())
	if err != nil {
		util4.Errln__(err)
		util4.Exit__(251)
		return
	}
	if goto1 != nil {
		switch goto1.Kw {
		case Kws_.Quit, Kws_.Return:
			if goto1.S == "" {
				util4.Exit__(0)
				return
			}
			i, err2 := strconv.Atoi(goto1.S)
			if err2 == nil {
				util4.Exit__(i)
				return
			}
		}
		util4.Errgotoln__(goto1)
		util4.Exit__(252)
		return
	}
}

func (this *Zsp___) Z__() {
	var is_serve bool
	{
		ss := New_strings__()
		var root string
		err := this.flag__(&root, &is_serve, ss)
		if err == nil {
			this.z, err = New__(ss.A, content_convert__, this)
			if err == nil {
				var args Args___
				for _, v := range this.z.Args.A {
					args.Add2__(v)
				}
				this.main_qv, err = this.z.New_main_qv__(&args)
				if err == nil {
					name := New_buf__()
					name.WriteString("我的")
					val := New_buf__()
					val.WriteString(Kws_.Call.String() + "I__" +
					Kws_.Dunhao.String() +
					Kws_.Kaidanyinhao.String() +
					Kws_.Args.String() +
					Kws_.Bidanyinhao.String() +
					Kws_.Juhao.String())
					set2 := new(Var___)
					set2.Is_lock = true
					set2.Is_through = true
					err = this.main_qv.Set_var__(name, val, set2, Kws_.Def)
				}
			}
		}
		if err != nil {
			util4.Errln__(err)
			util4.Exit__(250)
			return
		}

		this.main_qv.Annota.Add__("主")
		this.known_path = New_strings__()
		this.weizhuang = make(map[*weizhuang___]string)
		this.clpars = clpars4.New__()

		if this.z.Args.Src_type == Src_is_file_ {
			if !util4.Starts__(root, "/") {
				dir := util4.Get_dir__(this.z.Args.Src)
				if root == "." {
					root = dir
				} else {
					root = dir + "/" + root
				}
			}
		}
		if this.z.Args.Src_type == Src_is_file_ && !util4.Ends__(this.z.Args.Src, _zsp_) {
			is_serve = false
		}
		if is_serve {
			this.known_path_add__(root)
		} else {
			this.known_path_add__("/")
		}
		root = util4.Dir__(root)
		Known_path_add__(root)
		this.set_main_qv_var__("根", root)

		if this.z.Args.Src_type == Src_is_file_ {
			this.zs__(this.z.Args.Src, true)
		}
	}
	if is_serve {
		//err2 := http.ListenAndServe(this.addr, this)
		var err2 error
		this.serve, err2 = New_serve__(this)
		if err2 != nil {
			util4.Errln__(err2)
			util4.Exit__(255)
			return
		}
	}
}

package zsp2

import (
	"flag"
	"os"
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"util4"
)

func (this *Zsp___) flag2__(a []string, ss *Strings___) *Errinfo___ {
	all_is := false
	for i := 0; i < len(a); i++ {
		s := a[i]
		if util4.Ends__(s, " " + Shebang_flag_) {
			err := this.flag2__(Fields__(s), ss)
			if err != nil {
				return err
			}
			continue
		}
		if s == Shebang_flag_ {
			continue
		}
		if s == "-h" || util4.Starts__(s, "--help") {
			ss.Add__(s)
			continue
		}
		if s == "----" {
			all_is = !all_is
		}
		if !all_is {
			if s == "-zsp-help" {
				flag.PrintDefaults()
				util4.Exit__(250)
			}
			if s != "" && s[0] == '-' {
				name := s[1:]

				has_value := false
				value := ""
				for i := 1; i < len(name); i++ { // equals cannot be first
					if name[i] == '=' {
						value = name[i+1:]
						has_value = true
						name = name[0:i]
						break
					}
				}
				if zsp_clpars_ && !has_value {
					i1 := i + 1
					if i1 < len(a) {
						s1 := a[i1]
						if s1 != "" && s1[0] == '-' {
						} else {
							i++
							value = a[i]
							s += "=" + value
						}
					}
				}
	
				f := flag.Lookup(name)
				if f != nil {
					if err := f.Value.Set(value); err != nil {
						return New_errinfo__("-" + f.Name, Errs_.Case, value)
					}
					continue
				}
			}
		}
		ss.Add__(s)
	}
	return nil
}

var zsp_clpars_ bool

func (this *Zsp___) flag__(root *string, is_serve *bool, ss *Strings___) (err *Errinfo___) {
	flag.StringVar(root, "zsp-root", ".", "根")
	flag.StringVar(&this.addr, "zsp-addr", ":0", "侦听地址")
	flag.StringVar(&this.index, "zsp-index", "", "索引页")
	flag.BoolVar(is_serve, "zsp-serve", true, "做为服务")
	
	flag.BoolVar(&zsp_clpars_, "zsp-clpars", true, "")

	//flag.CommandLine.Parse(z.Args.A__())
	//args.Add__(flag_.Args()...)

	a := os.Args

	this.pgrname = a[0]
	this.finalpgrname = util4.Finalfilename__(this.pgrname)
	Known_path_add__(this.finalpgrname)
	ss.Add__(this.finalpgrname)

	start := this.pgrname
	for {
		s := start
		for {
			if util4.Ends__(s, "_") {
				s = s[0:len(s) - 1]
			} else {
				break
			}
		}
		start = s + _zsp_
		if util4.Exist_file__(start) {
			break
		}
		start = s + _zs_
		if util4.Exist_file__(start) {
			break
		}
		start = ""
		break
	}
	if start != "" {
		ss.Add__(start)
	}
	err = this.flag2__(a[1:], ss)

	ht := "http://"
	if util4.Starts__(this.addr, ht) {
		this.addr = this.addr[len(ht):]
	}
	return
}
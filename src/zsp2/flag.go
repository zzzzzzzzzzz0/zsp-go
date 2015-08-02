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
		if !all_is {
			if s == "----" {
				all_is = true
				continue
			}
			if util4.Ends__(s, " " + Shebang_flag_) {
				err := this.flag2__(Fields__(s), ss)
				if err != nil {
					return err
				}
				continue
			}
			if s == "-zsp-help" {
				flag.PrintDefaults()
				os.Exit(250)
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
	
				f := flag.Lookup(name)
				if f != nil {
					if !has_value {
						i++
						if i >= len(a) {
							return New_errinfo__("-" + f.Name, Errs_.Case)
						}
						value = a[i]
					}
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

func (this *Zsp___) flag__(root *string, is_serve *bool, ss *Strings___) *Errinfo___ {
	flag.StringVar(root, "r", ".", "根")
	flag.StringVar(&this.addr, "a", ":0", "侦听地址")
	flag.StringVar(&this.index, "i", "index.zsp", "索引页")
	flag.BoolVar(is_serve, "s", true, "做为服务")

	//flag.CommandLine.Parse(z.Args.A__())
	//args.Add__(flag_.Args()...)

	a := os.Args
	ss.Add__(a[0])
	return this.flag2__(a[1:], ss)
}
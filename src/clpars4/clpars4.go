package clpars4

import (
	"github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"regexp"
	"util4"
	"strconv"
)

type C___ struct {
	items *zhscript.List___
}

type item___ struct {
	tag, help, code, typ string
	re *regexp.Regexp
}

func (this *C___) help__() (ret string) {
	this.items.Find__(func(e *zhscript.Em___)bool {
		item := e.Value.(*item___)
		if item.tag != "" {
			ret += item.tag + "\t"
		}
		ret += item.help + "\n"
		return false
	})
	return
}

func (this *C___) Z__(qv *zhscript.Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), ret__ func(...interface{})) (no_use bool, goto1 *zhscript.Goto___) {
	switch k {
	case "命令行加回调":
		for i := 0; i < len(s); i++ {
			tag, ok := s__(s[i]); if !ok {return}
			var typ string
			if util4.Ends__(tag, "...") {
				tag = tag[0:len(tag) - 3]
				typ = "..."
			}

			var item *item___
			has := this.items.Find__(func(e *zhscript.Em___)bool {
				item = e.Value.(*item___)
				return item.tag == tag
			})

			if !has {
				item = &item___{}
				item.tag = tag
				item.typ = typ
				switch tag {
				case "":
					tag = "(.*)"
				default:
					tag = "^" + tag + "$"
				}
				re, err := regexp.Compile(tag)
				if err != nil {
					err__(err)
					return
				}
				item.re = re
			}

			i++
			if i >= len(s) {
				err__(item.tag + " 缺帮助")
				return
			}
			item.help, ok = s__(s[i]); if !ok {return}

			i++
			if i >= len(s) {
				err__(item.tag + " 缺代码")
				return
			}
			item.code, ok = s__(s[i]); if !ok {return}
			if item.code == "" {
				item.code = item.help
			}

			if !has {
				this.items.PushBack(item)
			}
		}
		return
	case "命令行解析":
		var (
			err1 *zhscript.Errinfo___
			kw *zhscript.Keyword___
			i int
			ss1 []string
			item1 *item___
		)
		zs__ := func(item *item___, ss []string) {
			_, goto1, err1 = util4.Zs3__(item.code, item.tag, qv, ss...)
			if err1 != nil {
				return
			}
			kw, goto1 = util4.Goto1__(goto1)
		}
		for i1, s1 := range s {
			s2, ok := s__(s1); if !ok {return}
			if(this.items.Find__(func(e *zhscript.Em___)bool {
				item := e.Value.(*item___)
				var fa bool
				for _, ss := range item.re.FindAllStringSubmatch(s2, -1) {
					fa = true
					if item.typ == "..." {
						ss1 = ss[1:]
						item1 = item
						return true
					}
					if item.tag == "" {
						i++
						ss = append(ss, strconv.Itoa(i))
					}
					zs__(item, ss[1:])
					if err1 != nil {
						return true
					}
					if kw == zhscript.Kws_.Continue {
						continue
					}
					if kw == zhscript.Kws_.Break || goto1 != nil {
						return true
					}
				}
				return fa
			})) {
				if item1 != nil {
					for {
						i1++
						if i1 >= len(s) {
							break
						}
						s2, ok = s__(s[i1]); if !ok {return}
						ss1 = append(ss1, s2)
					}
					zs__(item1, ss1)
				}
				if err1 != nil {
					err__(err1)
					break
				}
				if goto1 != nil || kw == zhscript.Kws_.Break {
					break
				}
				if item1 != nil {
					break
				}
				continue
			}
			switch s2 {
			case "-h", "--help":
				print(this.help__())
				util4.Exit__(100)
				return
			}
			err__("无法" + k, s2)
			return
		}
		return
	case "命令行帮助":
		ret__(this.help__())
		return
	}
	no_use = true
	return
}

func New__() *C___ {
	v := &C___{new(zhscript.List___)}
	return v
}
package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
)

func Zs__(src string, src_is_file bool, src2 string, up_qv *Qv___, r interface{},
s ...string) (buf *Buf___, goto1 *Goto___, err *Errinfo___) {
	var args Args___
	if src_is_file {
		args.Src_file__(src)
	} else {
		args.Src_code__(src)
	}
	args.Src2 = src2
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

func Zs2__(src string, qv *Qv___, s ...string) (buf *Buf___, goto1 *Goto___, err *Errinfo___) {
	return Zs3__(src, "", qv, s...)
}

func Zs3__(src, src2 string, qv *Qv___, s ...string) (buf *Buf___, goto1 *Goto___, err *Errinfo___) {
	return Zs__(src, false, src2, qv.Up, qv.Not_my, s...)
}

func Goto1__(goto1 *Goto___) (kw *Keyword___, goto2 *Goto___) {
	if goto1 != nil {
		if goto1.Kw == Kws_.Continue && goto1.S == "" {
			kw = goto1.Kw
			return
		}
		if goto1.Kw == Kws_.Break && goto1.S == "" {
			kw = goto1.Kw
			return
		}
		goto2 = goto1
	}
	return
}

type Zs_writer___ struct {
	code string
	qv *Qv___
	s string
}

func(this *Zs_writer___) Write(p []byte) (n int, err error) {
	s := string(p)
	for _, b := range s {
		switch b {
		case '\n', '\r':
			_, goto1, err1 := Zs2__(this.code, this.qv, this.s, string(b))
			if goto1 != nil {
				Errgotoln__(goto1)
			}
			if err1 != nil {
				Errln__(err1)
			}
			this.s = ""
		default:
			this.s += string(b)
		}
	}
	return len(p), nil
}

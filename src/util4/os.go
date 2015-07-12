package util4

import (
	"github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"os"
	"os/exec"
	"io"
	"strconv"
	"fmt"
)

func Errln__(v interface{}) {
	//fmt.Println(v)
	fmt.Fprintf(os.Stderr, "%v\n", v)
}

func Errgotoln__(g *zhscript.Goto___) {
	Errln__(zhscript.New_errinfo__("无法到达目标", g.S, g.Kw))
}

type zs_cmd___ struct {
	i io.WriteCloser
	o, e *Zs_writer___
}

func cmd__(s []interface{}, i1 int, s__ func(interface{}) (string, bool)) *exec.Cmd {
	name, ok := s__(s[i1]); if !ok {return nil}
	args := []string{}
	for i := i1 + 1; i < len(s); i++ {
		si, ok := s__(s[i]); if !ok {return nil}
		args = append(args, si)
	}
	return exec.Command(name, args...)
}

func Os__(qv *zhscript.Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}),
ret__ func(...interface{})) (no_use bool, goto1 *zhscript.Goto___) {
	switch k {
	case "环境变量":
		if buzu__(1) {
			return
		}
		name, ok := s__(s[0]); if !ok {return}
		ret__(os.Getenv(name))
		return
	case "程序名":
		ret__(os.Args[0])
		return
	case "重定向输出":
		if buzu__(1) {
			return
		}
		cmd := cmd__(s, 0, s__)
		if cmd == nil {
			return
		}
		buf, err := cmd.Output()
		ret__(string(buf))
		if err != nil {
			ret__(err)
		}
		return
	case "启动被动者":
		if buzu__(3) {
			return
		}
		cmd := cmd__(s, 2, s__)
		if cmd == nil {
			return
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			err__(err)
			return
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			err__(err)
			return
		}
		stdin, err := cmd.StdinPipe()
		if err != nil {
			err__(err)
			return
		}
		err = cmd.Start()
		if err != nil {
			err__(err)
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		s1, ok := s__(s[1]); if !ok {return}
		zc := &zs_cmd___{stdin,
			&Zs_writer___{code:s0, qv:qv},
			&Zs_writer___{code:s1, qv:qv}}
		go io.Copy(zc.o/*os.Stdout*/, stdout) 
		go io.Copy(zc.e/*os.Stderr*/, stderr)
		ret__(zc, strconv.Itoa(cmd.Process.Pid)) 
		return
	case "命令被动者":
		if buzu__(2) {
			return
		}
		zc, ok := s[1].(*zs_cmd___); if !ok {
			zhscript.O__("%v",s[1])
			err__(k)
			return
		}
		s0, ok := s__(s[0]); if !ok {return}
		zc.i.Write([]byte(s0))
		return
	}
	no_use = true
	return
}
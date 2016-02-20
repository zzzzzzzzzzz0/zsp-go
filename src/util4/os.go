package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"os"
	"os/exec"
	"io"
	"path/filepath"
	"strconv"
	"fmt"
	"bytes"
)

func Errln__(v interface{}) {
	//fmt.Println(v)
	fmt.Fprintf(os.Stderr, "%v\n", v)
}

func Errgotoln__(g *Goto___) {
	Errln__(New_errinfo__("无法到达目标", g.S, g.Kw))
}

type zs_cmd___ struct {
	i io.WriteCloser
	o, e *Zs_writer___
}

func cmd__(s []interface{}, i1 int, s__ func(interface{}) (string, bool)) *exec.Cmd {
	var name string
	args := []string{}
	for i := i1; i < len(s); i++ {
		si, ok := s__(s[i]); if !ok {return nil}
		for _, s2 := range Fields__(si) {
			if name == "" {
				name = s2
				continue
			}
			args = append(args, s2)
		}
	}
	return exec.Command(name, args...)
}

func Finalpgrname__(ret string) string {
	wd__ := func(ret string) string {
		if !Starts__(ret, "/") {
			dir, err := os.Getwd()
			if err == nil {
				ret = dir + "/" + ret
			}
		}
		return ret
	}
	var ln string
	ln__ := func() bool {
		var err error
		ln, err = os.Readlink(ret)
		return err == nil
	}
	for {
		ret = wd__(ret)
		for {
			dir := filepath.Dir(ret)
			dirln, err := os.Readlink(dir)
			if err != nil {
				break
			}
			is_ln := ln__()
			if !Starts__(dirln, "/") && !is_ln {
				if Starts__(dirln, "..") {
					dirln = wd__(dirln)
				} else {
					dirln = dir + "/" + dirln
				}
			}
			ret = dirln + "/" + filepath.Base(ret)
			if is_ln {
				ret = wd__(ret)
			}
		}
		if !ln__() {
			break
		}
		if !Starts__(ln, "/") {
			ln = filepath.Dir(ret) + "/" + ln
		}
		ret = ln
	}
	return filepath.Clean(ret)
}

func Os__(qv *Qv___, k string, s []interface{}, s__ func(interface{}) (string, bool),
err__ func(...interface{}), buzu__ func(int) bool, buzhichi__ func(...interface{}), can_stat__ func(string) bool,
ret__ func(...interface{}), c *Chan___) (no_use bool, goto1 *Goto___) {
	switch k {
	case "环境变量":
		if buzu__(1) {
			return
		}
		name, ok := s__(s[0]); if !ok {return}
		ret__(os.Getenv(name))
		return
	case "改变目录":
		if buzu__(1) {
			return
		}
		name, ok := s__(s[0]); if !ok {return}
		err := os.Chdir(name)
		if err != nil {
			ret__(err)
		}
		return
	case "当前目录":
		dir, err := os.Getwd()
		if err != nil {
			err__(err)
			return
		}
		ret__(dir)
		return
	case "重定向输出":
		if buzu__(1) {
			return
		}
		cmd := cmd__(s, 0, s__)
		if cmd == nil {
			return
		}
		var buf bytes.Buffer
		cmd.Stdout = &buf
		cmd.Stderr = &buf
		err := cmd.Run()
		ret__(string(buf.Bytes()))
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
			O__("%v",s[1])
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
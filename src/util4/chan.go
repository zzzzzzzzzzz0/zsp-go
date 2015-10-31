package util4

import (
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
)

type Chan___ struct {
	o chan string
	x chan bool
	goto1 chan *Goto___
	err chan error
	use, has_x, is_x, is_close bool
}

func New_chan__(n int) *Chan___ {
	c := &Chan___{o:make(chan string, n), x:make(chan bool),
		goto1:make(chan *Goto___), err:make(chan error)}
	return c
}

func (this *Chan___) close__() {
	if !this.is_close {
		this.is_close = true
		close(this.o)
		close(this.goto1)
		close(this.err)
	}
}

func (this *Chan___) Close_if__() {
	if !this.use {
		this.x__()
		this.close__()
	}
}

func (this *Chan___) x__() {
	if !this.is_x {
		this.is_x = true
		close(this.x)
	}
}

func (this *Chan___) Goto__(g *Goto___) {
	if !this.is_close {
		this.goto1 <- g
	}
}
func (this *Chan___) Err__(err error) {
	if !this.is_close {
		this.err <- err
	}
}

func Chan__(k string, s []interface{}, s__ func(interface{})(string, bool),
err__ func(...interface{}), buzu__ func(int) bool,
ret__ func(...interface{})) (no_use bool, goto1 *Goto___) {
	switch k {
	case "信道":
		if buzu__(1) {
			return
		}
		c, ok := s[0].(*Chan___)
		if !ok {
			err__("非" + k)
			return
		}
		if c.is_close {
			return
		}
		if c.has_x && len(c.o) == 0 {
			c.close__()
			ret__("", "1")
			return
		}
		select {
		case ret, ok := <- c.o:
			if ok {
				ret__(ret)
			}
		case g, ok := <- c.goto1:
			if ok {
				goto1 = g
			}
		case err, ok := <- c.err:
			if ok {
				err__(err)
			}
		case <- c.x:
			c.has_x = true
		}
		return
	case "信道关闭":
		if buzu__(1) {
			return
		}
		c, ok := s[0].(*Chan___)
		if !ok {
			return
		}
		c.x__()
		return
	}
	no_use = true
	return
}
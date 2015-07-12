package util4

import (

)

type Chan___ struct {
	o chan string
	x chan bool
}

func New_chan__(n int) *Chan___ {
	c := &Chan___{make(chan string, n), make(chan bool)}
	return c
}

func (this *Chan___) Close__() {
	close(this.x)
	close(this.o)
}

func (this *Chan___) Close_if__() {
	if len(this.o) == 0 {
		this.Close__()
	}
}

func Chan__(k string, s []interface{}, s__ func(interface{})(string, bool),
buzu__ func(int) bool, ret__ func(...interface{})) (no_use bool) {
	switch k {
	case "信道":
		if buzu__(1) {
			return
		}
		c, ok := s[0].(*Chan___)
		if !ok {
			return
		}
		ret, ok := <- c.o
		if !ok {
			return
		}
		if ret == "" {
			c.Close__()
		}
		ret__(ret)
		return
	case "信道关闭":
		if buzu__(1) {
			return
		}
		c, ok := s[0].(*Chan___)
		if !ok {
			return
		}
		c.Close__()
		return
	}
	no_use = true
	return
}
package util4

import (
    "testing"
    "time"
	. "github.com/zzzzzzzzzzz0/zhscript-go/zhscript"
	"os"
)

func Test_timer1__(t *testing.T) {
	var qv *Qv___

	z, _ := New__(os.Args, nil, nil)
	qv, _ = z.New_main_qv__(&z.Args)
	O_ansi_ = true
	O_tree_ = true

	c := New_chan__(1)
	go timer1__(2 * time.Second, `
赋予丑【顶】以算术先如果存在丑【顶】那么‘丑’+了1。
乙‘丑’
`, false, qv, c)
	timer1__(1 * time.Second, "甲", false, nil, c)
}


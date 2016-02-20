#!../bin/zsp --。
定义“运行”【主】以下代码
	赋予1以‘参数栈’。
	显示####“ ”‘1’换行。
	赋予2、3、pid【上】以执行‘1’。
	显示[‘2’,"‘3’",‘pid’]换行。
上代码。

运行下原样msgbox 1 22 &
echo $!上原样。

我的等待、3s。
赋予ppid以‘pid’。
赋予cid以算术‘pid’+1。

运行。
赋予pidz以‘pid’。

运行“ps h --ppid ‘ppid’ -o pid”。
运行“pidof -o ‘ppid’ gtkmmsh”。
运行“pgrep -P ‘ppid’ gtkmmsh”。
运行“k1 "msgbox 1 22" -x”。

我的等待、100。
显示是‘cid’吗？换行。
循环先算术‘pidz’-‘ppid’-1了次先
	显示先算术‘ppid’+‘次’了,。
了。
显示换行。

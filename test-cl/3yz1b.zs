#!../bin/zsp --。
（三英战吕布）

定义打得过【倒挂3】、我1、我2、我3、敌以下代码
	赋予力以合战力‘我1’、‘我2’、‘我3’。
	如果‘力’大于‘敌’那么1。
上代码。

定义合战力、我1、我2、我3以下代码
	赋予力、错以算术(‘我1’/‘我1负担’) + ‘我2’ + ‘我3’、1。
	如果‘错’那么先
		显示‘错’换行。
		结束。
	了。
	‘力’。
上代码。

定义 \o/ 以下代码
	定义示、名以下代码
		‘名’英雄(有‘‘名’’之强)
	上代码。
	显示先示刘了“、”先示关了“、”先示张了“撕布！”换行。
上代码。
定义 /..\ 以下代码
	定义示、名以下代码
		渣‘名’(仅‘‘名’’)
	上代码。
	显示先示刘了“、”先示关了“、”先示张了“被铩羽止飞……”换行。
上代码。

赋予刘、关、张、吕以10、104、100、205。

赋予我1负担以‘参数1’。
如果‘我1负担’等于“”那么赋予我1负担以2。
显示备备的拖累——‘我1负担’换行换行。

如果‘刘’、‘关’、‘张’打得过‘吕’那么 \o/ 否则 /..\。

显示先合战力‘刘’、‘关’、‘张’了“ 比 ‘吕’”换行。

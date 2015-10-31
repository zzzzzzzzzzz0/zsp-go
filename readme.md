基于 zhscript-go 的中文语法的动态网页。不是一种通常用于站点的动态网页，而是期望用于嵌入式设备，服务器端直接解释，不似 jsp、php 般需要再编译，这样便可以由客户端发送代码，最大灵活性地调用设备的功能，由于是中文语法，也便能让小白用户自己来定制。

测试方法：

./bin/zsp -a :4000 -i ""

在浏览器里访问 127.0.0.1:4000，点击其中 [helloworld.zsp](https://github.com/zzzzzzzzzzz0/zsp-go/blob/master/helloworld.zsp) 的话，将会显示：

hello, woorrrllllddddd.

已有应用：

[跨目录的基于浏览器的图片管理器](https://github.com/zzzzzzzzzzz0/zsp-go-imgmgr)
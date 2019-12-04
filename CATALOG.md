# Go Web场景实战(第一篇)：极致性能篇

## 第一章:纯数据库篇

### [第1讲:开张课、技术栈介绍、课程内容说明 试听](http://www.jtthink.com/course/play/2143)
* 开张课，说明下本课程的技术栈和开发内容

### [第2讲:上线准备:表模型、API设计、取出商品列表](http://www.jtthink.com/course/play/2144)
* 讲解下我们使用到的表。（课件中有近2万左右的商品测试数据，供大家玩耍）

### [第3讲:上线准备:三层构架代码封装(上)](http://www.jtthink.com/course/play/2145)
* 如果纯按照官方文档写代码就太low了。我们把gokit中的三层代码架构迁移到gin中。做个深度使用

### [第4讲:上线准备:三层构架代码封装(下) ---三部曲 试听](http://www.jtthink.com/course/play/2146)
* 上节课我们讲到了三层架构的原理，这节课我们用代码实现出来

### [第5讲:练习课：利用三层架构实现商品详细API](http://www.jtthink.com/course/play/2157)
* 基于上节课内容，我们做个商品详细API。作为练习

### [第6讲:本机压测: apache ab 基本操作、压测商品列表API](http://www.jtthink.com/course/play/2158)
* 做个开端。下节课我们正式来压测和优化

### [第7讲:补充课时：日志保存、error统一拦截、](http://www.jtthink.com/course/play/2159)
* 今天是个补充课时，简单说明下日志保存和gin的错误拦截（不使用中间件）

### [第8讲:把程序和mysql部署到Linux(docker)中、测试运行 试听](http://www.jtthink.com/course/play/2170)
* 为了更过瘾的测试，我们把系统部署到Linux中(gin和mysql)

### [第9讲:使用配置文件、第一次本地压测(ab):最基本的参数调整](http://www.jtthink.com/course/play/2171)
* 今天我们来进行最基本的第一次本地压测,并调整一个最基本的系统参数

### [第10讲:开启mysql日志(docker)、初步设置连接池](http://www.jtthink.com/course/play/2172)
* 今天演示下超小连接池，抗高并发

### [第11讲:（补充课时）MySQL5.7缓存的设置、提高查询效率](http://www.jtthink.com/course/play/2191)
* 补充课时。可选学，由于网友要求，补充下。mysql缓存本身比较鸡肋。

### [第12讲:使用新工具代替传统ab压测](http://www.jtthink.com/course/play/2192)
* 传统ab工具并不是特别好用，因此今天介绍个go做的压测工具，作为ab的替代品

### [第13讲:增加一个API:商品元数据表、记录点击量 试听](http://www.jtthink.com/course/play/2193)
* 为了后面的压测，我们增加一个API

### [第14讲:代码扩展：显示商品详细API时，同时显示meta信息](http://www.jtthink.com/course/play/2194)
* 由于我们使用了三层架构的代码。因此今天说明下怎么扩展

### [第15讲:nginx+gin+mysql 的docker环境部署](http://www.jtthink.com/course/play/2195)
* 实际环境我们不可能把gin服务裸对外。一般都会挡一层网关或反代。后面的压测也要基于此，因此我们今天讲下基本的部署
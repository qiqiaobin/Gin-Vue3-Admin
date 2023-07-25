# log
基于zap封装的log库

logrus 是 golang 一款非常优秀的日志框架, 其优点非常明显:

 - 优雅的代码框架设计
 - 使用简单
 - 组件化的开发思路
 - 灵活的输出方式

但是, 性能终究是忍痛舍弃 logrus 的“阿喀琉斯之踵”
目前 golang 日志库的大众选择主要集中在: logrus, zap, zerolog. zap 和 zerolog 的性能都是优秀的, 但是从用法习惯上我更倾向于 zap.

### 简单介绍 Zap 的使用
Zap 提供三种不同方式的输出(以 Info为 例)

```go
log.Info("hello zap") // {"level":"info","ts":1576423173.016333,"caller":"test_zap/main.go:28","msg":"hello zap"}
log.Infof("hello %s", "zap") // {"level":"info","ts":1576423203.056074,"caller":"test_zap/main.go:29","msg":"hello zap"}
log.Infow("hello zap", "field1", "value1") //{"level":"info","ts":1576423203.0560799,"caller":"test_zap/main.go:30","msg":"hello zap","field1":"value1"}
```

如果我们对 logrus 的 key-value 理论比较在意的话, 使用 zap infow 可以完美解决

 - Zap 使用起来不便利的地方
 - Zap 使用上不能像 logrus 那样开箱即用
 - 使用者需要自己去组装相关函数
 - Zap 同样不提供日志切割的功能, 但是想添加上这个功能没有 logrus 那样便利

基于这些问题, 我封装了一套开箱即用的日志组件: https://github.com/yizhidaozuihou/log

打造 Zap 开箱即用日志组件
提供的功能:

 - 提供的功能: 像 logrus 一样, 全局的 Debug, Info ... 函数
 - 日志分割功能. 默认文件大小1024M，自动压缩, 最大有3个文件备份，备份保存时间7天, 不会打印日志被调用的文文件名和位置
 - 日志默认会被分成五类文件：debug、info、warn、error、panic 都会打印在xxx.log. xxx.log.Request输出 request log 的地方(如果有需要的话)

ES的写入有两种方式一种是逐个文档写入（index），另一种是多个文档批量写入（bulk）。对于这两种写入方式，ES都会将其转换为bulk写入。本节，我们就以bulk写入为例，根据代码执行主线来分析ES写入的流程。


> https://cloud.tencent.com/developer/article/1361160

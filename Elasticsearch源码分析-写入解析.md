ES的写入有两种方式一种是逐个文档写入（index），另一种是多个文档批量写入（bulk）。对于这两种写入方式，ES都会将其转换为bulk写入。本节，我们就以bulk写入为例，根据代码执行主线来分析ES写入的流程。


> https://cloud.tencent.com/developer/article/1361160

Elasticsearch查询解析
> https://cloud.tencent.com/developer/article/1154904


2.3 DFS_QUERY_THEN_FETCH
       这类查询用于解决ES在多分片、少数据量的场景下计算相关度不准确的问题：以TF/IDF算法为例，ES在计算相关度时仅考虑单个分片内的IDF，可能导致查询结果中，类似的文档因为在不同分片而相关度大为不同的问题。此时可以使用此类查询，在QUERY_THEN_FETCH之前再增加一轮任务调度，用于计算分布式的IDF。但通常情况下，局部和全局IDF的差异会随着索引里文档数的增多渐渐消失，在真实世界的数据量下，这个问题几乎没有影响，没有必要使用此类查询增加一轮任务调度的开销。


ES提供用户Transport和Rest两种接口：用户可以通过ES官方提供的Transport Client访问ES集群，这种接口使用的协议与ES集群内部节点间的通讯协议一致；也可以使用简单易用的Rest接口，直接发送Http请求访问ES集群，由ES完成Rest请求到Transport请求的转换。考虑Rest接口的易用性，以及Rest层极低的额外开销，建议用户直接使用Rest接口。

上述即为查询入口的处理流程，它对任何Rest请求都适用。实际上，除了自带的Rest请求外，ES提供强大的扩展能力，**用户可以通过自定义插件实现自己的请求及处理逻辑。**此外，ES还支持自定义过滤器Filter，在实际进行Transport层处理前进行统一的预处理工作。



所以，在不需要通过_id字段去重、update的使用场景中，写入不指定_id可以提升写入速率。腾讯云CES技术团队的测试结果显示，无_id的数据写入性能可能比有_id的高出近一倍，实际损耗和具体测试场景相关。

8. 使用routing
对于数据量较大的index，一般会配置多个shard来分摊压力。这种场景下，一个查询会同时搜索所有的shard，然后再将各个shard的结果合并后，返回给用户。对于高并发的小查询场景，每个分片通常仅抓取极少量数据，此时查询过程中的调度开销远大于实际读取数据的开销，且查询速度取决于最慢的一个分片。开启routing功能后，ES会将routing相同的数据写入到同一个分片中（也可以是多个，由index.routing_partition_size参数控制）。如果查询时指定routing，那么ES只会查询routing指向的那个分片，可显著降低调度开销，提升查询效率。

9. 为string类型的字段选取合适的存储方式
存为text类型的字段（string字段默认类型为text）： 做分词后存储倒排索引，支持全文检索，可以通过下面几个参数优化其存储方式：
		norms：用于在搜索时计算该doc的_score（代表这条数据与搜索条件的相关度），如果不需要评分，可以将其关闭。
		index_options：控制倒排索引中包括哪些信息（docs、freqs、positions、offsets）。对于不太注重_score/highlighting的使用场景，可以设为 docs来降低内存/磁盘资源消耗。
		fields: 用于添加子字段。对于有sort和聚合查询需求的场景，可以添加一个keyword子字段以支持这两种功能。

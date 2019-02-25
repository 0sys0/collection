### long和integer

### cutoff_frequency 
> https://www.elastic.co/guide/cn/elasticsearch/guide/current/common-terms.html

### match_phrase的坑
ES中文分词器之精确短语匹配（解决了match_phrase匹配不全的问题）
> https://www.jianshu.com/p/51a756e85eb6
主要是分词引起的，常常和预想的结果不同

### 死磕 Elasticsearch 方法论
> https://yq.aliyun.com/articles/679317

### aggregation
Metrics Aggregations
> https://my.oschina.net/bingzhong/blog/1861105

Bucket Aggregations
> https://my.oschina.net/bingzhong/blog/1917915

Pipeline Aggregations
> https://my.oschina.net/bingzhong/blog/1975879

### Significant Terms Aggregation

### suggester
> Elasticsearch Suggester详解
> https://elasticsearch.cn/article/142

那么类似的功能在Elasticsearch里如何实现呢？ 答案就在Suggesters API。 Suggesters基本的运作原理是将输入的文本分解为token，然后在索引的字典里查找相似的term并返回。 根据使用场景的不同，Elasticsearch里设计了4种类别的Suggester，分别是:
- Term Suggester
- Phrase Suggester
- Completion Suggester
- Context Suggester

### collapse

> Elasticsearch 5.x 字段折叠的使用
> https://elasticsearch.cn/article/132

### Search After
在使用Elasticsearch的过程中，一般的分页需求我们可以使用form和size的方式实现，但是这种分页方式在深度分页的场景下应该是要避免使用的。深度分页会随着请求的页次增加，所消耗的内存和时间的增长也是成比例的增加
对于深度分页，es推荐使用 scroll 接口，详情请查看《Elasticsearch普通分页from&size VS scroll滚动分页》。注意，scroll接口不适合用在实时搜索的场景里。
从es 5.0版本开始，es提供了新的参数 search_after 来解决这个问题，search_after 提供了一个活的游标来拉取从上次返回的最后一个请求开始拉取下一页的数据。
> http://www.piaoyi.org/database/Elasticsearch-Search_After.html


### store属性和_source字段
> https://blog.csdn.net/jingkyks/article/details/41785887

哪些情形下需要显式的指定store属性呢？大多数情况并不是必须的。从_source中获取值是快速而且高效的。如果你的文档长度很长，存储_source或者从_source中获取field的代价很大，你可以显式的将某些field的store属性设置为yes。缺点如上边所说：假设你存储了10个field，而如果想获取这10个field的值，则需要多次的io，如果从_source中获取则只需要一次，而且_source是被压缩过的。

### es父文档子文档
> https://blog.csdn.net/hereiskxm/article/details/44937455

在使用搜索“has-child”搜索父文档时，一般情况只返回子文档符合条件的父文档。用 Inner-hits 则可以把父子文档同时返回——既返回父文档，也返回匹配has-child条件的子文档，相当于在父子之间join了。

这是一个相当有用的特性，假设我们使用父文档存储邮件内容，子文档存储每个邮件拥有者的信息以及对于此用户这封邮件的状态。搜索某个账户的邮件列表时，我们希望搜索到邮件内容和邮件状态，可以设想假如没有Inner-hits，我们必须得分两次查询，因为邮件内容和邮件状态分别存放在父文档和子文档中。而有了Inner_hits属性后，我们可以使用一次查询完成。

> Elasticsearch 学习之Search API inner hits
> https://blog.csdn.net/yiyiholic/article/details/81634393

用inner_hits可以把父子文档同时返回——既返回,不加inner_hits只返回一个type里的数据。inner_hits默认只查询3条数据，可以自定义设置from 和size。

> elasticsearch 关联查询
> http://www.cnblogs.com/double-yuan/p/9798103.html

### es字段
ctx._source 代表当前将被更新的源文档。

### es5.x中文文档
> http://cwiki.apachecn.org/pages/viewpage.action?pageId=9405298


### alias索引别名
> https://www.cnblogs.com/jajian/p/10152681.html

索引别名就像一个快捷方式或软连接，可以指向一个或多个索引，也可以给任何一个需要索引名的API来使用，而且别名不能与索引同名。

别名带给我们极大的灵活性，允许我们做下面这些：

- 在运行的集群中可以无缝的从一个索引切换到另一个索引。
- 给多个索引分组。
- 给索引的一个子集创建视图

别名的开销很小，应该广泛使用

### string类型与index字段

> https://www.elastic.co/guide/en/elasticsearch/reference/5.2/string.html

2.x中创建的索引只有string，没有text和keyword。在5.x中创建的索引中string类型会被自动转换为text或keyword。保留该字段是为了向下兼容。

> https://www.elastic.co/guide/en/elasticsearch/reference/5.2/mapping-index.html

index表示是否可查询。对于string类型index为analyzed表示该字段为fulltext，index为not_analyzed表示该字段为keyword，no表示不可查询。对于其他类型true表示可以查询，false表示不可查询

//所以说就是旧版本的坑

es的版本是5.2.2

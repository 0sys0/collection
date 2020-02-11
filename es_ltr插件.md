> https://elasticsearch-learning-to-rank.readthedocs.io/en/latest/index.html

实际上，您永远都不想以sltr这种方式使用查询。为什么？该模型对索引中的每个结果执行。这些模型占用大量CPU。您将使用上面的查询快速使您的Elasticsearch集群爬行。

通常，您将在基线相关性查询的前N个执行模型。您可以使用Elasticsearch的内置rescore功能进行此操作：

```
POST tmdb/_search
{
    "query": {
        "match": {
            "_all": "rambo"
        }
    },
    "rescore": {
        "window_size": 1000,
        "query": {
            "rescore_query": {
                "sltr": {
                    "params": {
                        "keywords": "rambo"
                    },
                    "model": "my_model"
                }
            }
        }
    }
}
```

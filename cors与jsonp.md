跨域资源共享 CORS 详解
> http://www.ruanyifeng.com/blog/2016/04/cors.html

jsonp 跨域原理分析
> https://segmentfault.com/a/1190000009773724

```
<html>
<head>
    <title>跨域测试</title>
    <script src="js/jquery-1.7.2.js"></script>
    <script>

        $(document).ready(function () {

            $("#btn").click(function () {

                $.ajax({
                    url: "http://localhost:9090/student",
                    type: "GET",
                    dataType: "jsonp", //指定服务器返回的数据类型
                    success: function (data) {
                        var result = JSON.stringify(data); //json对象转成字符串
                        $("#text").val(result);
                    }
                });

            });

        });
    </script>
</head>
<body>
    <input id="btn" type="button" value="跨域获取数据" />
    <textarea id="text" style="width: 400px; height: 100px;"></textarea>

</body>
</html>
```

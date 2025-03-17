## todo
- 重写一个工具，解决生成乱码的问题

## sql2pb

https://github.com/Mikaelemmmm/sql2pb

### 安装
```
go install github.com/Mikaelemmmm/sql2pb@latest
```

### 命令示例

```
sql2pb -go_package ./pb -host localhost -package pb -password 123456 -port 3306 -schema lottery -service_name lotter -user root > lottery.proto
```

## 解决报错

如果提示如下图所示的错误：根据提示我们可以看出来是编码问题。

![](https://files.mdnice.com/user/36414/da1991a5-38c9-49e8-a13c-29f80c3733ab.png)

我们重置一下编码即可解决：

点击utf-8
![](https://files.mdnice.com/user/36414/2e00391a-f1a8-4c03-ba96-43733a87c963.png)

然后convert
![](https://files.mdnice.com/user/36414/9fda69d3-4f0f-479e-a81d-13b7ff430b1c.png)

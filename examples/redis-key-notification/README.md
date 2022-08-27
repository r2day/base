# redis 监控事件发生

默认redis不开启，所以需要通过如下方式开启

```
config set notify-keyspace-events Kx
```
或者直接配置在 `redis.con`中


这里我们只关心过期expire事件，所以使用`Kx`

## 其他更多内容参考

https://xiaorui.cc/archives/4123
http://redisdoc.com/topic/notification.html
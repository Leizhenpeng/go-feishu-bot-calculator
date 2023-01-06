# go-feishu-bot-calculator
> 给飞书机器人添加计算器功能

![img.png](doc/img.png)

## 关于反向代理

由于需要订阅事件，所以需要在公网上部署，

如果你的服务器没有公网 IP，可以使用反向代理的方式

飞书的服务器在国内对ngrok的访问速度很慢，所以推荐使用一些国内的反向代理服务商

- [cpolar](https://dashboard.cpolar.com/)
- [natapp](https://natapp.cn/)

### 部署
```bash
go run main.go

nohup cpolar http 8080 -log=stdout &
```
### 查看服务器状态
https://dashboard.cpolar.com/status

### 查询并kill进程
```bash
ps -ef | grep cpolar

kill -9 PID 
``` 


## 关于飞书机器人
 
[事件订阅的文档](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM?lang=zh-CN#2eb3504a)
[事件订阅的SDK](https://github.com/larksuite/oapi-sdk-go#%E5%A4%84%E7%90%86%E6%B6%88%E6%81%AF%E4%BA%8B%E4%BB%B6%E5%9B%9E%E8%B0%83)
[所有SDK](https://github.com/larksuite/oapi-sdk-go)
[消息类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json)
[发送消息文档](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create)
应用审核: https://fork-way.feishu.cn/admin/appCenter/audit (需要飞书企业版)


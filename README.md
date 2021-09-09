# feishu-sender(webhook)

Nightingale的理念，是将告警事件扔到redis里就不管了，接下来由各种sender来读取redis里的事件并发送，毕竟发送报警的方式太多了，适配起来比较费劲，希望社区同仁能够共建。

这里提供一个飞书机器人的sender，参考了https://github.com/n9e/wechat-sender，https://github.com/n9e/feishu-sender


飞书拉群，并添加群机器人，得到webhook，配置自定义关键词安全设置，需要在夜莺项目monapi.yaml设置里的notify添加im告警,
例如：
```
notify:
  p1: ["im"]
  p2: ["im"]
  p3: ["im"]
```

## compile

```bash
go 1.15以上需要交叉编译
env GOOS=linux GOARCH=amd64 go build
编译后得到二进制文件
feishu-sender
```

如上编译完就可以拿到二进制了。

## configuration

直接修改etc/feishu-sender.yml即可

## pack

编译完成之后可以打个包扔到线上去跑，将二进制和配置文件打包即可：

```bash
tar zcvf feishu-sender.tar.gz feishu-sender etc/feishu-sender.yml etc/feishu.tpl
```

## test

配置etc/feishu-sender.yml，相关配置修改好，我们先来测试一下是否好使.
目录结构
feishu-sender
etc/feishu.tpl
etc/feishu-sender.yml
./feishu-sender启动测试
## run

如果测试发送没问题，扔到线上跑吧，使用systemd或者supervisor之类的托管起来，systemd的配置实例：


```
$ cat feishu-sender.service
[Unit]
Description=Nightingale feishu sender
After=network-online.target
Wants=network-online.target

[Service]
User=root
Group=root

Type=simple
ExecStart=/data/service/n9e/feishu-sender
WorkingDirectory=/data/service/n9e

Restart=always
RestartSec=1
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
```

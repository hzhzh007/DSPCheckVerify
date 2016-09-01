# What this?
[![GoDoc](https://godoc.org/github.com/hzhzh007/DSPCheckVerify?status.svg)](https://godoc.org/github.com/hzhzh007/DSPCheckVerify) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)]()

对接芒果tv pmp 的dsp 协议自我检查程序



# How to Use

bin  文件夹下包含各个平台的二进制，基本使用如下：

```
./bin/check_mac -h
这是芒果tv pmp 对接dsp 使用的api 校验工具,各家需要对接mgtv的dsp 通过此工具先完成api 格式的连调，然后在每天17:00 －18:00 可以找芒果技术对接实际的
Usage of ./bin/check_mac:
  -a string
        dsp 使用的通信url (default "http://127.0.0.1/mgtv/bid")
  -t string
        检查类型, pc_front: pc 前贴片, pc_pause: pc 暂停 (default "pc_front")
  -v    打印详细
```

# 测试素材

## 尺寸

pc贴片 848*480
app贴片 512*288
时长 15s
暂停 480*270

## 素材

pc贴片
http://mp4.res.hunantv.com/pmp/27/1467959430737.flv
app贴片
http://res.hunantv.com//mediafiles/wiad_creative/213/1460434626.mp4
暂停
http://res.hunantv.com/mediafiles/wiad_creative/49/1460450426.jpg

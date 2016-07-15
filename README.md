# What this?
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

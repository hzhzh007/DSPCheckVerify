package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	CheckPcFront = "pc_front"
	CheckAppFont = "app_front"
)

var (
	checkType string
	apiUrl    string
	verbose   bool
)

func init() {
	flag.StringVar(&checkType, "t", "pc_front", "检查类型, pc_front: pc 前贴片, pc_pause: pc 暂停")
	flag.StringVar(&apiUrl, "a", "http://127.0.0.1/mgtv/bid", "dsp 使用的通信url")
	flag.BoolVar(&verbose, "v", false, "打印详细")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr,
			`这是芒果tv pmp 对接dsp 使用的api 校验工具,各家需要对接mgtv的dsp 通过此工具先完成api 格式的连调，然后在每天17:00 －18:00 可以找芒果技术对接实际的`)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	var data string
	switch checkType {
	case CheckPcFront:
		data = PC_FRONT
	case CheckAppFont:
		data = APP_FRONT
	}
	result, err := Request(data)
	if err != nil {
		log.Fatal(err)
	}
	err = check(data, result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("test ok")
	}
}

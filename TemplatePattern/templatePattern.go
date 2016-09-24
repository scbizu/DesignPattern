package main

import "fmt"

//HummerModel 悍马模型
type HummerModel interface {
	start()
	stop()
	alarm()
	engineBoom()
	// run()
}

//HummerH1Model 悍马类
type HummerH1Model struct {
}

func (hm *HummerH1Model) alarm() {
	fmt.Println("悍马H1鸣笛...")
}

func (hm *HummerH1Model) engineBoom() {
	fmt.Println("悍马H1引擎声音是这样的...")
}

func (hm *HummerH1Model) start() {
	fmt.Println("悍马H1发动...")
}

func (hm *HummerH1Model) stop() {
	fmt.Println("悍马H1停车...")
}

//调用入口...
func run(hm HummerModel) {
	//发动..
	hm.start()
	//引擎开始轰鸣
	hm.engineBoom()
	//看到一条狗挡路 就开始按喇叭 (喵喵喵?)
	hm.alarm()
	//停车...
	hm.stop()
}

func main() {
	hummer := new(HummerH1Model)
	run(hummer)
}

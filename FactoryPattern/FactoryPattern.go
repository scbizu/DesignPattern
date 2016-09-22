package main

import "fmt"

type human interface {
	//人种
	getcolor()
	//说话
	talk()
}

//BlackHuman 黑人
type BlackHuman struct {
}

//YellowHuman 黄人
type YellowHuman struct {
}

//WhiteHuman 白人
type WhiteHuman struct {
}

//HumanFactory 造人工厂
type HumanFactory struct {
}

func (bh BlackHuman) getcolor() {
	fmt.Println("黑色人种的皮肤颜色是黑色的.")
}

func (bh BlackHuman) talk() {
	fmt.Println("黑人会说话,一般人听不懂.")
}

func (yh YellowHuman) getcolor() {
	fmt.Println("黄色人种的皮肤颜色是黄色的.")
}

func (yh YellowHuman) talk() {
	fmt.Println("黄色人种会说话,一般说的都是双字节.")
}

func (wh WhiteHuman) getcolor() {
	fmt.Println("白色人种的皮肤颜色是白色的.")
}
func (wh WhiteHuman) talk() {
	fmt.Println("白色人种会说话,一般都是单字节.")
}

func (hf *HumanFactory) humanFactory(race string) human {
	switch race {
	case "black":
		return new(BlackHuman)
	case "yellow":
		return new(YellowHuman)
	case "white":
		return new(WhiteHuman)
	default:
		return nil
	}
}

//造人计划
func main() {
	var factory HumanFactory
	blackhuman := factory.humanFactory("black")
	fmt.Println("--造出的第一批是黑人--")
	blackhuman.getcolor()
	blackhuman.talk()
	yellowhuman := factory.humanFactory("yellow")
	fmt.Println("--造出的第二批是黄人--")
	yellowhuman.getcolor()
	yellowhuman.talk()
	whitehuman := factory.humanFactory("white")
	fmt.Println("--造出的第三批是白人--")
	whitehuman.getcolor()
	whitehuman.talk()
}

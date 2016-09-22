package main

import "fmt"

//Human ...
type Human interface {
	getcolor()
	talk()
	// getsex()
}

//SexHuman ..
type SexHuman interface {
	Human
	getSex()
}

//HumanFactory ...
type HumanFactory interface {
	createYellowHuman()
	// createWhiteHuman()
	// createBlackHuman()
}

//BlackHuman ...
type BlackHuman struct {
}

//YellowHuman ...
type YellowHuman struct {
}

//WhiteHuman ...
type WhiteHuman struct {
}

//MaleYellowHuman ...
type MaleYellowHuman struct {
	YellowHuman
}

//FemaleYellowHuman ...
type FemaleYellowHuman struct {
	YellowHuman
}

//FemaleFactory ...
type FemaleFactory struct {
}

//MaleFactory ...
type MaleFactory struct {
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
func (myh MaleYellowHuman) getsex() {
	fmt.Println("黄人男性")
}
func (fmyh FemaleYellowHuman) getsex() {
	fmt.Println("黄人女性")
}
func (fmf FemaleFactory) createYellowHuman() *FemaleYellowHuman {
	return new(FemaleYellowHuman)
}
func (mf MaleFactory) createYellowHuman() *MaleYellowHuman {
	return new(MaleYellowHuman)
}
func main() {
	var malefac MaleFactory
	myh := malefac.createYellowHuman()
	myh.getcolor()
	myh.talk()
	myh.getsex()
	var femalefac FemaleFactory
	fyh := femalefac.createYellowHuman()
	fyh.getcolor()
	fyh.talk()
	fyh.getsex()
}

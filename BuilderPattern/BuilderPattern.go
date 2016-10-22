package main

//CarModel ...
type CarModel interface {
	strat()
	stop()
	alarm()
	engineBoom()
	run()
}

//CarBuilder ...
type CarBuilder interface {
	//set builder sequence
	setSequence(Sequence []string)
	//get CarModel
	getCarmodel() CarModel
}

//Director ...
type Director struct {
	Sequence    []string
	Benzbuilder BenzBuilder
	BMWbuilder  BMWBuilder
}

//getAbenzModel 获取一个奔驰模型
func (director Director) getAbenzModel() BenzModel {
	sequence := []string{"start", "stop"}
	director.Benzbuilder.setSequence(sequence)
	return director.Benzbuilder.BenzModel
}

//getABMWModel 获取一个宝马模型
func (director Director) getABMWModel() BMWModel {
	sequence := []string{"start", "engineBoom", "stop"}
	director.BMWbuilder.setSequence(sequence)
	return director.BMWbuilder.BMWModel
}

//BMWModel ...
type BMWModel struct {
	Sequence []string
}

//BMWBuilder ...
type BMWBuilder struct {
	BMWModel
}

//BenzModel ...
type BenzModel struct {
	Sequence []string
}

//BenzBuilder ...
type BenzBuilder struct {
	BenzModel
}

func (bb *BMWBuilder) setSequence(sequence []string) {
	bb.BMWModel.Sequence = sequence
}

func (bb *BMWBuilder) getCarmodel() BMWModel {
	return bb.BMWModel
}

func (cb *BenzBuilder) setSequence(sequence []string) {
	cb.BenzModel.Sequence = sequence
}

func (cb *BenzBuilder) getCarmodel() BenzModel {
	return cb.BenzModel
}

func (bb *BMWBuilder) run() {
	for _, v := range bz.Sequence {
		switch v {
		case "start":
			bz.strat()
		case "stop":
			bz.stop()
		case "alarm":
			bz.alarm()
		case "engineBoom":
			bz.engineBoom()
		}
	}
}

func (bb *BMWBuilder) strat() {
	println("宝马车跑起来是这个样子的...")
}
func (bb *BMWBuilder) stop() {
	println("宝马车应该这样停车.....")
}
func (bb *BMWBuilder) alarm() {
	println("宝马车的喇叭声是这个样子的...")
}
func (bb *BMWBuilder) engineBoom() {
	println("宝马车的引擎是这个声音的....")
}

func (bz *BenzModel) run() {
	for _, v := range bz.Sequence {
		switch v {
		case "start":
			bz.strat()
		case "stop":
			bz.stop()
		case "alarm":
			bz.alarm()
		case "engineBoom":
			bz.engineBoom()
		}
	}
}

func (bz *BenzModel) strat() {
	println("奔驰车跑起来是这个样子的...")
}
func (bz *BenzModel) stop() {
	println("奔驰车应该这样停车.....")
}
func (bz *BenzModel) alarm() {
	println("奔驰车的喇叭声是这个样子的...")
}
func (bz *BenzModel) engineBoom() {
	println("奔驰车的引擎是这个声音的....")
}

func main() {
	var benzM BenzModel
	director := new(Director)
	benzM = director.getAbenzModel()
	// benzM = benzBuilder.getCarmodel()
	benzM.run()
}

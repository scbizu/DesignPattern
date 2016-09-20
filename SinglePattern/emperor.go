package emperor

import (
	"fmt"
	"sync"
)

//Emperor 皇帝类
type Emperor struct {
}

//Getinstance 单例模式
func Getinstance() *Emperor {
	var emperor *Emperor
	fn := func() {
		//new By emperor itself
		emperor = &Emperor{}
	}
	var once sync.Once
	// Do Just Once
	once.Do(fn)
	return emperor
}

//Say 皇帝发话
func (em Emperor) Say() {
	fmt.Println("我是皇帝某某某....")
}

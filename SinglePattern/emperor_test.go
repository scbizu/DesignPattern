package emperor

import "testing"

func TestEmperor(t *testing.T) {
	for i := 0; i < 3; i++ {
		emperor := Getinstance()
		emperor.Say()
	}
	t.Log("End")
}

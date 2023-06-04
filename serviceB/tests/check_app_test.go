package tests

import (
	"kafka-go/serviceB/utils"
	"math/rand"
	"testing"
	"time"
)

func TestRandState(t *testing.T) {
	state := "НЕ УСПЕШНО"
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(2)
	if num == 1 {
		state = "УСПЕШНО"
	}
	stateFunc := utils.RandState()
	if stateFunc != state {
		t.Errorf("Incorrect result. Expect %s, got %s",
			state, stateFunc)
	}
}

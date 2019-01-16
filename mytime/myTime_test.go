package mytime

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T)  {
	currTime := time.Now()
	fmt.Println(currTime)

	nextTime := currTime.Add(1 * time.Minute)
	fmt.Println(nextTime)

	gap := nextTime.Sub(currTime)
	fmt.Println(gap)
}

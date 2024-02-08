package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Weekday()
	fmt.Println(now)
}

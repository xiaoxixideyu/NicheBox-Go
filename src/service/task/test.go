package main

import (
	"fmt"
	"time"
)

func main() {
	parse, _ := time.Parse(time.DateOnly, "2024-01-17")
	now, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
	after := now.After(parse)
	fmt.Printf("after:%v", after)
}

package main

import "github.com/WhiCu/mangazeya/cmd"

func main() {
	cmd.Execute()
}

// package main

// import (
// 	"fmt"

// 	"github.com/WhiCu/mangazeya/pkg/chart"
// )

// func clearScreen() {
// 	// ANSI-код очистки экрана и возврата курсора в 0,0
// 	fmt.Print("\033[H\033[2J")
// }

// func main() {
// 	crt := chart.NewChart[uint64](10, 100, 25)
// 	fmt.Println(crt.View())
// 	// values := []uint64{0, 0, 83, 329, 479, 55, 402, 7610, 0, 55, 901, 901, 901, 854, 438, 823, 456, 845, 1235, 5000, 1274, 234, 2768, 235, 347, 6231}
// 	// for i := range values {
// 	// 	clearScreen()
// 	// 	crt.Add(values[i], values[len(values)-1-i])
// 	// 	// fmt.Println(crt.Memory(), crt.Pos())
// 	// 	fmt.Println(crt.View())
// 	// 	time.Sleep(1 * time.Second)
// 	// }

// }

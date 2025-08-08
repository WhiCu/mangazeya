// package main

// import "github.com/WhiCu/mangazeya/cmd"

// func main() {
// 	cmd.Execute()
// }

package main

import (
	"fmt"

	"github.com/WhiCu/mangazeya/pkg/chart"
)

func main() {
	crt := chart.NewChart[uint64](10, 10, 10)
	values := []uint64{0, 0, 83, 329, 479, 55, 402, 7610, 0, 55, 901, 901, 901, 1235, 5000}
	for _, v := range values {
		crt.Add(v)
	}
	fmt.Println(crt.View())

}

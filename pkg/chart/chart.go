package chart

import (
	// "fmt"

	// "fmt"

	"github.com/guptarohit/asciigraph"
	"golang.org/x/exp/constraints"
)

// type node interface {
// 	Measure() int
// }

type Chart[T constraints.Integer] struct {
	// размеры
	height int
	width  int
	// буфер
	size    int
	memory  [][]T // кольцевой буфер
	pos     int
	legends []string
	// агрегаты
	sum     T // суммарное значение (int64 для безопасных арифметик)
	average T
	max     T
	min     T
}

func NewChart[T constraints.Integer](size int, width, height int) *Chart[T] {

	return &Chart[T]{
		height: height,
		width:  width,

		size: size,
		pos:  0,
	}
}

func (c *Chart[T]) Add(values ...T) {
	c.pos = (c.pos + 1) % c.size
	// fmt.Println("values", values)
	// fmt.Println("c.memory", c.memory)
	// fmt.Println("c.pos", c.pos)
	for i, v := range values {
		if i >= len(c.memory) {
			c.memory = append(c.memory, make([]T, c.size))
		}
		c.memory[i][c.pos] = v
		c.sum = c.sum + v
	}

	// c.sum = c.sum - c.memory[(c.pos+1)%len(c.memory)] + value

	// c.average = c.sum / T(len(c.memory))

	c.max = Max(c.memory)
	// c.min = min(c.min, value)

}

func (c *Chart[T]) AddLegend(legends ...string) {
	if c.legends == nil {
		c.legends = make([]string, 0, len(legends))
	}
	c.legends = append(c.legends, legends...)

}
func Max[T constraints.Ordered](t [][]T) T {
	max := t[0][0]
	for i := range t {
		for j := range t[i] {
			if t[i][j] > max {
				max = t[i][j]
			}
		}
	}
	return max
}

// TEST
func (c *Chart[T]) Memory() [][]T {
	return c.memory
}
func (c *Chart[T]) Pos() int {
	return c.pos
}

func (c *Chart[T]) View() string {
	// fmt.Println("c.memory", c.memory)
	graph := make([][]float64, len(c.memory))
	for i := range c.memory {
		graph[i] = make([]float64, 0, len(c.memory[i]))
		// fmt.Println("c.memory[i]", c.memory[i])
		// fmt.Println("len(c.memory[i])", len(c.memory[i]))
		for j := range c.memory[i] {
			// fmt.Println("c.memory[i][(c.pos-j+c.size)%c.size]", c.memory[i][(c.pos-j+c.size)%c.size])
			// fmt.Println("(c.pos-j+c.size)%c.size", (c.pos-j+c.size)%c.size)
			graph[i] = append(graph[i], float64(c.memory[i][(c.pos-j+c.size)%c.size]))
		}
	}

	// fmt.Println("graph", graph)
	legend := defeaultLegends
	if len(c.legends) > 0 {
		legend = append(c.legends, defeaultLegends...)
	}
	// fmt.Println("legend", legend)
	return asciigraph.PlotMany(
		graph,
		asciigraph.Height(c.height),
		asciigraph.Width(c.width),
		asciigraph.LowerBound(0),
		asciigraph.UpperBound(float64(c.max)),
		asciigraph.SeriesLegends(legend[:len(c.memory)]...),
		asciigraph.SeriesColors(defeaultColor[:len(c.memory)]...),
	)
}

var defeaultLegends = []string{"A", "B", "C", "D", "E", "F", "G"}

var defeaultColor = []asciigraph.AnsiColor{
	asciigraph.Green,
	asciigraph.Red,
	asciigraph.Blue,
	asciigraph.Yellow,
	asciigraph.Magenta,
	asciigraph.Cyan,
	asciigraph.YellowGreen,
}

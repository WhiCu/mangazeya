package chart

import (
	"fmt"

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
	memory []T // кольцевой буфер
	pos    int

	// агрегаты
	sum     T // суммарное значение (int64 для безопасных арифметик)
	average T
	max     T
	min     T
	delta   T // signed difference last->value
}

func NewChart[T constraints.Integer](size int, width, height int) *Chart[T] {
	return &Chart[T]{
		height: height,
		width:  width,
		memory: make([]T, size),
		pos:    0,
	}
}

func (c *Chart[T]) Add(value T) {

	c.sum = c.sum - c.memory[(c.pos+1)%len(c.memory)] + value

	c.average = c.sum / T(len(c.memory))

	c.max = max(c.max, value)
	c.min = min(c.min, value)

	c.delta = c.max / T(len(c.memory))

	c.pos = (c.pos + 1) % len(c.memory)
	c.memory[c.pos] = value

}

func (c *Chart[T]) View() string {
	graph := func() [][]float64 {

		main := make([]float64, 0, len(c.memory))
		average := make([]float64, 0, len(c.memory))

		for i := range c.memory {
			main = append(main, float64(c.memory[i]))
			average = append(average, float64(c.average))
		}
		fmt.Println("main", main, "len", len(main))
		fmt.Println("average", average, "len", len(average))
		return [][]float64{
			average,
			main,
		}
	}()
	return asciigraph.PlotMany(
		graph,
		asciigraph.Height(c.height),
		asciigraph.Width(c.width),
		asciigraph.Precision(0),
		asciigraph.LowerBound(0),
		asciigraph.UpperBound(float64(c.max)),
	)
}

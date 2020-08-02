package gameoflife

import (
	"errors"
	"fmt"
)

//Block is a matrix of x * y size
type Block struct {
	m [][]int
	x int
	y int
}

//NewBlock return a Block, width on the horizontal line(x axis), height on vertical line(y axis)
//TODO 4 bit a byte
func NewBlock(width, height int, s []byte) (Block, error) {
	size := width * height
	if size != len(s) {
		return Block{}, errors.New("invalid length")
	}
	b := Block{}

	b.m = make([][]int, height)
	for i := range b.m {
		b.m[i] = make([]int, width)
	}

	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b.m[i][j] = int(s[count])
			count++
		}
	}
	b.x = width
	b.y = height
	return b, nil
}

func (b *Block) String() string {
	var str string
	for i := 0; i < b.x; i++ {
		str += fmt.Sprintf("%v\n", b.m[i])
	}
	return str
}

//Next moves to next status of matrix
func (b *Block) Next() {
	for i := 0; i < len(b.m); i++ {
		for j := 0; j < len(b.m[i]); j++ {
			b.Update(i, j)
		}
	}
}

//Update status of a pixel in matrix
func (b *Block) Update(x, y int) {
	count := b.countAround(x, y)
	if count == 2 {
		return
	}
	if count == 3 {
		b.m[x][y] = 1
		return
	}
	b.m[x][y] = 0
}

//count alive pixel around
func (b *Block) countAround(x, y int) int {
	a := make([][2]int, 8)
	var val, count int
	a[0] = [2]int{x + 1, y}
	a[1] = [2]int{x - 1, y}
	a[2] = [2]int{x, y + 1}
	a[3] = [2]int{x, y - 1}
	a[4] = [2]int{x + 1, y + 1}
	a[5] = [2]int{x - 1, y - 1}
	a[6] = [2]int{x + 1, y - 1}
	a[7] = [2]int{x - 1, y + 1}
	for i := 0; i < 8; i++ {
		val = b.get(a[i][0], a[i][1])
		if val == 1 {
			count++
		}
	}
	return count

}

//get status of a pixel, if it dose not exit just return -1
func (b *Block) get(x, y int) int {
	if x < 0 || x >= b.x {
		return -1
	}
	if y < 0 || y >= b.y {
		return -1
	}
	return b.m[x][y]
}

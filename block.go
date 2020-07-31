package existenz

import (
	"fmt"
	"errors"
)

type block struct {
	m [][]int
	x int
	y int
} 

//returm a block with matrix and initialize it
//TODO 4 bit a byte
func newBlock(m, n int, s []byte) (block, error) {
	size := m * n
	if size != len(s) {
		return block{}, errors.New("invalid length!")
	} 
	b := block{}

	b.m = make([][]int, m)
	for i := range b.m {
		b.m[i] = make([]int, n)
	}
	
	count := 0
	for i:=0; i<m; i++ {
		for j:=0;j<n;j++ {
			b.m[i][j] = int(s[count])
			count ++ 
		}
	}
	b.x = m
	b.y = n
	return b, nil
}

func (b *block) String() string {
	var str string
	for i := 0; i< b.x; i++ {
	str += fmt.Sprintf("%v\n", b.m[i])
	}
	return str
}


//move to next status of matrix
func (b *block) next() {
	for i:= 0 ; i < len(b.m); i++ {
		for j:=0; j < len(b.m[i]); j++ {
			b.update(i,j)
		}
	}
}

//update status of a pixel in matrix
func (b *block) update(x, y int) {
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
func (b *block) countAround(x, y int) int {
	a := make([][2]int, 8)
	var val, count int
	a[0] = [2]int{x+1,y}
	a[1] = [2]int{x-1,y}
	a[2] = [2]int{x,y+1}
	a[3] = [2]int{x,y-1}
	a[4] = [2]int{x+1, y+1}
	a[5] = [2]int{x-1, y-1}
	a[6] = [2]int{x+1, y-1}
	a[7] = [2]int{x-1, y+1}
	for i:= 0; i<8; i++ {
		val = b.get(a[i][0], a[i][1]) 
		if val == 1 {
			count++
		}
	} 
	return count
		
}

//get status of a pixel, if it dose not exit just return -1
func (b *block) get(x, y int) int {
	if x < 0 || x >= b.x {
		return -1
	}
	if y < 0 || y >= b.y {
		return -1
	}
	return b.m[x][y]
}

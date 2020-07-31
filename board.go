package gameoflife

import (
	"image"
	"image/color"
	"image/jpeg"
	"strconv"
	"os"
	"log"
	"fmt"
)

const (
	Framename = "generation"
	Ext = ".jpg"
	MaxTransitions = 1024
)

var (
	pixel_alive = color.RGBA{0x94, 0xda, 0x67, 0xff}
	pixel_dead = color.RGBA{0, 0, 0, 0xff}
)

type board struct {
	image.RGBA
	b *block
	unit int
}

func NewBoard(b *block, unit int) *board {
	rect := image.Rect(0, 0, b.x * unit, b.y * unit)
	br := board{
		*image.NewRGBA(rect), 
		b, 
		unit,
	}
	br.visualize()
	return &br
}

//fill a block area of board with a color 
func (b *board) setBlock( x, y int, colour color.Color) {
	length := b.unit
	for i:=x*length; i<(x+1)*length; i++ {
		for j:=y*length; j<(y+1)*length; j++{
			b.Set(i, j, colour)
		}
	}
}

//visualize current status matrix  
func (b *board) visualize() {
	for w :=0; w < b.b.x; w++ {
		for l :=0; l < b.b.y; l++ {
			if b.b.m[w][l] == 1 {
				b.setBlock(w, l, pixel_alive)
			} else {
				b.setBlock(w, l, pixel_dead)
			}
		}
	} 
}

//update the board to next frame 
func (b *board) update() {
	b.b.next()
	b.visualize()
}

//write board to a image file
func (b *board) writeImg(name string) {
	image := b.SubImage(image.Rect(0, 0, b.Bounds().Dx(), b.Bounds().Dy()))
	out, err := os.Create(name)
	defer out.Close()
	if err != nil {
		log.Println(err)
	}
	op := new(jpeg.Options)
	op.Quality = 50
	jpeg.Encode(out, image, op)
}

//board transitions for n times
func (b *board) transition(n int) {
	if n > MaxTransitions { 
		n = MaxTransitions
	}
	for i := 0; i < n; i++ {
		b.writeImg(Framename + strconv.Itoa(i) + Ext)
		b.update()
		fmt.Println(b.b)
	}
} 
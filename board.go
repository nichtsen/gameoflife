package gameoflife

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strconv"
)

const (
	Framename      = "generation"
	Ext            = ".jpg"
	MaxTransitions = 1024
)

var (
	pixel_alive = color.RGBA{0x94, 0xda, 0x67, 0xff}
	pixel_dead  = color.RGBA{0, 0, 0, 0xff}
)

//Board can visualize its underlying Block
type Board struct {
	image.RGBA
	b    *Block
	unit int
}

//NewBoard return a new Board for visualizing its underlying Block
func NewBoard(b *Block, unit int) *Board {
	rect := image.Rect(0, 0, b.x*unit, b.y*unit)
	br := Board{
		*image.NewRGBA(rect),
		b,
		unit,
	}
	br.visualize()
	return &br
}

//Fill a Block area of Board with a color
func (b *Board) setBlock(x, y int, colour color.Color) {
	length := b.unit
	for i := x * length; i < (x+1)*length; i++ {
		for j := y * length; j < (y+1)*length; j++ {
			b.Set(i, j, colour)
		}
	}
}

//visualize current status matrix
func (b *Board) visualize() {
	for w := 0; w < b.b.y; w++ {
		for l := 0; l < b.b.x; l++ {
			if b.b.m[w][l] == 1 {
				b.setBlock(l, w, pixel_alive)
			} else {
				b.setBlock(l, w, pixel_dead)
			}
		}
	}
}

//Update the Board to next frame
func (b *Board) Update() {
	b.b.Next()
	b.visualize()
}

//WriteImg writes the Board to a image file
func (b *Board) WriteImg(name string) {
	image := b.SubImage(image.Rect(0, 0, b.Bounds().Dx(), b.Bounds().Dy()))
	out, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	op := new(jpeg.Options)
	op.Quality = 50
	jpeg.Encode(out, image, op)
}

//Transition for n times
func (b *Board) Transition(n int) {
	if n > MaxTransitions {
		n = MaxTransitions
	}
	for i := 0; i < n; i++ {
		b.WriteImg(Framename + strconv.Itoa(i) + Ext)
		b.Update()
		fmt.Println(b.b)
	}
}

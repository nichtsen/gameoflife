# gameoflife
John Conway's  Game of Life


## example

```
package main

import(
	gl "github.com/nichtsen/gameoflife" 
)

func main() {
	//board length, width,
	//and block length by pixels
	width, len, unit := 5, 5, 100
	//1 for alive 0 for dead
	b, err := gl.NewBlock(
		width,
		len,
		[]byte{
		0,1,0,0,1,
		0,1,1,1,0,
		0,1,1,0,0,
		0,1,0,0,0,
		1,0,0,1,0},
	)
	if err != nil {
		panic(err)
	}
	br := gl.NewBoard(&b, unit)
	// transition for 10 times
	br.Transition(10)
}
```
![generation0](img/generation0.jpg)
![generation1](img/generation1.jpg)
![generation2](img/generation2.jpg)
![generation3](img/generation3.jpg)
![generation4](img/generation4.jpg)
![generation5](img/generation5.jpg)
![generation6](img/generation6.jpg)
![generation7](img/generation7.jpg)
![generation8](img/generation8.jpg)
![generation9](img/generation9.jpg)

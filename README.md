# gameoflife
John Conway's  Game of Life


## example

```
package main

import(
	"github.com/nichtsen/gameoflife"
)

func main() {
	//board length, width,
	//and block length by pixels
	width, len, unit := 5, 5, 100
	//1 for alive 0 for dead
	b, err := newBlock(
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
	br := newBoard(&b, unit)
	// transition for 10 times
	br.transition(10)
}
```
![](img/generation0.jpg)
![](img/generation1.jpg)
![](img/generation2.jpg)
![](img/generation3.jpg)
![](img/generation4.jpg)
![](img/generation5.jpg)
![](img/generation6.jpg)
![](img/generation7.jpg)
![](img/generation8.jpg)
![](img/generation9.jpg)

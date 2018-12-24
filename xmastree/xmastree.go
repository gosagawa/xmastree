package xmastree

import (
	"fmt"
	"time"
)

const (
	// EscapeRed display red color
	EscapeRed = "\033[31m"

	// EscapeGreen  display green color
	EscapeGreen = "\033[32m"

	// EscapeYellow display Yellow color
	EscapeYellow = "\033[33m"

	// EscapeReset  reset color
	EscapeReset = "\033[0m"
)

//Tree is Christmas tree
type Tree struct {
	MaxWidth    int // Mas display Width
	LeafHeight  int // Leaf height
	StemWidth   int // stem width
	StemHeight  int // stem height
	TotalHeight int // total height:
}

//Display the tree
func Display(size int, speed int) {

	maxWidth := 0
	maxHeight := size*2 + 9

	for i := 0; i < size; i++ {
		if i != 0 {
			fmt.Printf("\033[%vA", maxHeight)
		}
		tree := newTree(i*2+5, i+1, 2, maxHeight)
		tree.display()
		time.Sleep(time.Duration(speed) * time.Millisecond)
		maxWidth = tree.MaxWidth
	}

	fmt.Print("\033[1A")
	fmt.Print(EscapeYellow)

	spaceLength := (maxWidth - 20) / 2
	echoString(" ", spaceLength)
	fmt.Println("Merry, Christmas !!!")
	echoString(" ", spaceLength)
	fmt.Println(EscapeReset)
}

func newTree(leafHeight int, stemWidth int, stemHeight int, totalHeight int) *Tree {
	tree := Tree{
		LeafHeight:  leafHeight,
		StemWidth:   stemWidth,
		StemHeight:  stemHeight,
		TotalHeight: totalHeight,
	}

	tree.MaxWidth = tree.LeafHeight*2 + 1
	return &tree
}

func (t *Tree) display() {

	echoString("\r\n", t.TotalHeight-t.LeafHeight-t.StemHeight-3)

	fmt.Print(EscapeGreen)
	for i := 0; i < t.LeafHeight; i++ {
		leafAmount := i*2 + 1 - (i/5)*4
		spaceLength := (t.MaxWidth - leafAmount) / 2
		echoString(" ", spaceLength)
		echoString("¥", leafAmount)
		echoString(" ", spaceLength)
		fmt.Println("")
	}

	fmt.Print(EscapeRed)
	for i := 0; i < t.StemHeight; i++ {
		spaceLength := (t.MaxWidth - t.StemWidth - 2) / 2
		echoString(" ", spaceLength)
		echoString("[", 1)
		echoString(" ", t.StemWidth)
		echoString("]", 1)
		echoString(" ", spaceLength)
		fmt.Println("")
	}
	spaceLength := (t.MaxWidth - t.StemWidth - 2) / 2
	echoString(" ", spaceLength)
	echoString("`", t.StemWidth+2)
	echoString(" ", spaceLength)
	fmt.Println("")

	echoString("\r\n", 2)

}

func echoString(str string, amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print(str)
	}
}

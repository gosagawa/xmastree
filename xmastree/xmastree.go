package xmastree

import (
	"fmt"
	"os"
	"time"
)

const (
	EscapeRed    = "\033[31m"
	EscapeGreen  = "\033[32m"
	EscapeYellow = "\033[33m"
	EscapeReset  = "\033[0m"
)

type Tree struct {
	MaxWidth    int
	LeafHeight  int
	StemWidth   int
	StemHeight  int
	TotalHeight int
}

func Display() {

	maxWidth := 0

	displayStep := 15
	maxHeight := displayStep*2 + 9

	for i := 0; i < displayStep; i++ {
		if i != 0 {
			fmt.Fprint(os.Stdout, fmt.Sprintf("\033[%vA", maxHeight))
		}
		tree := newTree(i*2+5, i+1, 2, maxHeight)
		tree.display()
		time.Sleep(50 * time.Millisecond)
		maxWidth = tree.MaxWidth
	}

	fmt.Fprint(os.Stdout, "\033[1A")
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

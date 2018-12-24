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
	for i := 0; i < 8; i++ {
		if i != 0 {
			fmt.Fprint(os.Stdout, "\033[25A")
		}
		tree := newTree(i*2+5, i+1, 2, 25)
		tree.display()
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Print(EscapeYellow)
	fmt.Println("Merry, Christmas !!!")
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

	fmt.Print(EscapeGreen)
	for i := 0; i < t.LeafHeight; i++ {
		leafAmount := i*2 + 1
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

	echoString("\r\n", t.TotalHeight-t.LeafHeight-t.StemHeight-1)
}

func echoString(str string, amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print(str)
	}
}

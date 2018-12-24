package xmastree

import "fmt"

const (
	EscapeRed    = "\033[31m"
	EscapeGreen  = "\033[32m"
	EscapeYellow = "\033[33m"
	EscapeReset  = "\033[0m"
)

type Tree struct {
	MaxWidth   int
	LeafHeight int
	StemWidth  int
	StemHeight int
}

func Display() {
	tree := newTree(5, 1, 2)
	tree.display()

}

func newTree(leafHeight int, stemWidth int, stemHeight int) *Tree {
	tree := Tree{
		LeafHeight: leafHeight,
		StemWidth:  stemWidth,
		StemHeight: stemHeight,
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

	fmt.Print(EscapeYellow)
	fmt.Println("Merry, Christmas !!!")
	fmt.Println(EscapeReset)
}

func echoString(str string, amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print(str)
	}
}

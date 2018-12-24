package xmastree

import "fmt"

const (
	EscapeRed    = "\033[31m"
	EscapeGreen  = "\033[32m"
	EscapeYellow = "\033[33m"
	EscapeReset  = "\033[0m"
)

func Display() {
	fmt.Print(EscapeGreen)
	fmt.Println("    ¥    ")
	fmt.Println("   ¥¥¥   ")
	fmt.Println("  ¥¥¥¥¥  ")
	fmt.Println(" ¥¥¥¥¥¥¥ ")
	fmt.Print(EscapeRed)
	fmt.Println("   [ ]   ")
	fmt.Println("   [ ]   ")
	fmt.Println("   ```   ")
	fmt.Print(EscapeYellow)
	fmt.Println("Merry, Christmas !!!")
	fmt.Println(EscapeReset)

}

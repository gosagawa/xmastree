package xmastree

import "fmt"

const (
	EscapeRed   = "\033[31m"
	EscapeGreen = "\033[32m"
	EscapeBlue  = "\033[33m"
	EscapeReset = "\033[0m"
)

func Display() {
	fmt.Println(EscapeGreen)
	fmt.Println("    ¥    ")
	fmt.Println("   ¥¥¥   ")
	fmt.Println("  ¥¥¥¥¥  ")
	fmt.Println(" ¥¥¥¥¥¥¥ ")
	fmt.Println(EscapeRed)
	fmt.Println("   [ ]   ")
	fmt.Println("   [ ]   ")
	fmt.Println("   ```   ")
	fmt.Println(EscapeBlue)
	fmt.Println("Merry, Christmas !!!")
	fmt.Println(EscapeReset)

}

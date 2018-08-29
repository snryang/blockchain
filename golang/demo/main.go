package main
import (
	"fmt"
	"sort"
	"bufio"
	"os"
)

func main(){
	GuessingGame()
}

func GuessingGame() {
	stdin := bufio.NewReader(os.Stdin)
	var s string
    fmt.Printf("Pick an integer from 0 to 100.\n")
    answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Fscan(stdin, &s)
		//fmt.Scanf("%s", &s)
		fmt.Println(s)
        return s[0] == 'y'
    })
    fmt.Printf("Your number is %d.\n", answer)
}

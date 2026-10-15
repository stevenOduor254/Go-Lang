package main
import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var num int
	conter := 1


	for{
		randomInt := rand.IntN(100)
		fmt.Printf("Enter your guesse Between [0 - 100]: ")
		fmt.Scanln(&num)
		check_if_compare(num, randomInt)
	}
	conter++
}

func check_if_compare(userGuess int, randomGen int) {
	if(userGuess == randomGen){
		fmt.Printf("You have won!!!!")
	}else{
		fmt.Printf("Your Guess: %v\tMachine Guess: %v\n", userGuess, randomGen)
	}
}

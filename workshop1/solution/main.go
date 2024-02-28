package main

import (
	"fmt"
	"solution/guesser"
	"strconv"
	//strconv是"String Conversion"的缩写，是Go标准库中的一个包，
	//提供了一些字符串与基本数据类型之间相互转换的函数，例如将字符串转换为整数或浮点数，
	//以及将整数或浮点数转换为字符串。通过导入"strconv"包，
	//你可以在你的Go程序中使用strconv包中提供的各种函数来进行字符串和基本数据类型之间的转换操作。
)

func main() {
	/***************************************
	 * Set up boundaries and instances.
	 **************************************/
	bounds := guesser.Bounds{
		Min: 4,
		Max: 15,
	}

	g1 := &guesser.Next_guesser{
		Bounds:  bounds,
		N_tried: 0,
	}

	g2 := &guesser.Rand_guesser{
		Bounds: bounds,
	}

	/***************************************
	 * Play the guessing game.
	 **************************************/
	turn := 0
	correct_val := 11
	var curr guesser.IGuess

	fmt.Println("Correct value: " + strconv.Itoa(correct_val))

	for {
		if turn%2 == 0 {
			curr = g1
		} else {
			curr = g2
		}

		the_guess := curr.Guess()
		fmt.Printf("%s guessed: %d\n", curr.Name(), the_guess)

		if the_guess == correct_val {
			fmt.Println(curr.Name() + " won!")
			break
		}

		turn++
	}
}

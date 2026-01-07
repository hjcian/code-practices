package firstfactorial

func FirstFactorial(num int) int {
	if num <= 1 {
		return 1
	}
	// code goes here
	return FirstFactorial(num-1) * num
}

// do not modify below here, readline is our function
// that properly reads in the input for you
// func main() {

// 	fmt.Println(FirstFactorial(readline()))

// }

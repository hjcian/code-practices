package nextgreaterelementwithpositionoffset

/*
 * Complete the 'findNextGreaterElementsWithDistance' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY readings as parameter.
 */

func findNextGreaterElementsWithDistance(readings []int32) [][]int32 {
	// Write your code here
	if len(readings) == 0 {
		return [][]int32{}
	}
	if len(readings) == 1 {
		return [][]int32{{-1, -1}}
	}

	res := make([][]int32, len(readings))
	for i := range res {
		// init.
		res[i] = []int32{-1, -1}
	}

	stack := make([]int, 0)
	stack = append(stack, 0)
	for i := 1; i < len(readings); i++ {
		if readings[i] <= readings[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 {
			if readings[i] <= readings[stack[len(stack)-1]] {
				break
			}
			idx := stack[len(stack)-1]
			distance := i - idx
			res[idx] = []int32{readings[i], int32(distance)}
			stack = stack[:len(stack)-1] // pop up the last ele.
		}
		stack = append(stack, i)
	}
	return res
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     readingsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var readings []int32

//     for i := 0; i < int(readingsCount); i++ {
//         readingsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         readingsItem := int32(readingsItemTemp)
//         readings = append(readings, readingsItem)
//     }

//     result := findNextGreaterElementsWithDistance(readings)

//     for i, rowItem := range result {
//         for j, colItem := range rowItem {
//             fmt.Printf("%d", colItem)

//             if j != len(rowItem) - 1 {
//                 fmt.Printf(" ")
//             }
//         }

//         if i != len(result) - 1 {
//             fmt.Printf("\n")
//         }
//     }

//     fmt.Printf("\n")
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }

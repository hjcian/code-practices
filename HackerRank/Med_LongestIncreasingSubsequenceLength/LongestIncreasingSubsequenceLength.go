package longestincreasingsubsequencelength

/*
 * Complete the 'computeLongestIncreasingSubsequenceLength' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY quality
 */

func computeLongestIncreasingSubsequenceLength(n int32, quality []int32) int32 {
	// Write your code here
	if n == 0 || n == 1 {
		return n
	}

	tails := make([]int32, 0)
	tails = append(tails, quality[0])
	//fmt.Println("tails:", tails)
	for i := 1; i < len(quality); i++ {
		if quality[i] > tails[len(tails)-1] {
			tails = append(tails, quality[i])
			//fmt.Println("tails:", tails)
			continue
		}
		// find the index to replace
		left := 0
		right := len(tails) - 1
		for left < right {
			mid := (right + left) / 2
			if tails[mid] < quality[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		//fmt.Println(".   replace:", right, "with:", quality[i])
		tails[right] = quality[i]
		//fmt.Println(".   tails:", tails)
	}
	return int32(len(tails))
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     n := int32(nTemp)

//     qualityCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var quality []int32

//     for i := 0; i < int(qualityCount); i++ {
//         qualityItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         qualityItem := int32(qualityItemTemp)
//         quality = append(quality, qualityItem)
//     }

//     result := computeLongestIncreasingSubsequenceLength(n, quality)

//     fmt.Printf("%d\n", result)
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

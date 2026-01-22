package longestarithmeticsubsequencewithgivendifference

/*
 * Complete the 'findLongestArithmeticProgression' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY arr
 *  2. INTEGER k
 */

/*
Question:
- Given an array of integers and a positive integer k, find the length of the longest arithmetic progression with common difference k. Ignore duplicates.
*/

import "slices"

//	func findLongestArithmeticProgression(arr []int32, k int32) int32 {
//		dp := make(map[int32]int32)
//		var maxLen int32 = 0
//		for _, num := range arr {
//			if v, ok := dp[num-k]; ok {
//				dp[num] = v + 1
//			} else {
//				dp[num] = 1
//			}
//			if dp[num] > maxLen {
//				maxLen = dp[num]
//			}
//		}
//		return maxLen
//	}
func findLongestArithmeticProgression(arr []int32, k int32) int32 {
	// Write your code here
	if len(arr) == 0 || len(arr) == 1 {
		return int32(len(arr))
	}
	slices.Sort(arr)
	//fmt.Println("Sorted:", arr)
	record := make(map[int32]bool, 0)
	for _, num := range arr {
		record[num] = false
	}
	var longLen int32 = 1
	for i := 0; i < len(arr); i++ {
		used, exists := record[arr[i]+k]
		//fmt.Println("now:", arr[i], "look for:", arr[i]+k, "exists:", exists, "used:", used)
		if !exists {
			// arr[i] is not a start, next
			continue
		}
		if used {
			// means this is used for previous start
			continue
		}
		// arr[i] is a start,
		var length int32 = 1
		num := arr[i]
		for {
			_, exists := record[num+k]
			if !exists {
				break
			}
			//fmt.Println(".   touch:", num+k)
			length++
			record[num+k] = true
			num = num + k
		}
		if length > longLen {
			longLen = int32(length)
		}
	}
	return longLen
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var arr []int32

//     for i := 0; i < int(arrCount); i++ {
//         arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         arrItem := int32(arrItemTemp)
//         arr = append(arr, arrItem)
//     }

//     kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     k := int32(kTemp)

//     result := findLongestArithmeticProgression(arr, k)

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

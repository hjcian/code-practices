package subarrayswithgivensumandboundedmaximum

/*
 * Complete the 'countSubarraysWithSumAndMaxAtMost' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY nums
 *  2. LONG_INTEGER k
 *  3. LONG_INTEGER M
 */

func countSubarraysWithSumAndMaxAtMost(nums []int32, k int64, M int64) int64 {
	// Write your code here
	counts := make(map[int64]int)
	counts[0] = 1
	var sum int64 = 0

	total := 0
	for _, num := range nums {
		if int64(num) > M {
			// this value is kind of a delimiter
			counts = make(map[int64]int)
			counts[0] = 1
			sum = 0
			continue
		}
		sum += int64(num) // sum[i] = sum[i-1] + num[i]
		target := sum - k
		//fmt.Println("sum[", i, "]=", sum, " target:", target)
		if numOfStart, ok := counts[target]; ok {
			total += numOfStart
			//fmt.Println(".   hit, plus:", numOfStart)
		}
		counts[sum]++ // this is a starter
		//fmt.Println(".   start", sum, "#", counts[sum])
	}
	return int64(total)
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     numsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var nums []int32

//     for i := 0; i < int(numsCount); i++ {
//         numsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         numsItem := int32(numsItemTemp)
//         nums = append(nums, numsItem)
//     }

//     k, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     M, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     result := countSubarraysWithSumAndMaxAtMost(nums, k, M)

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

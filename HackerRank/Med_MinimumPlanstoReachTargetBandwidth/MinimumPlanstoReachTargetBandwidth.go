package minimumplanstoreachtargetbandwidth

/*
 * Complete the 'findMinimumPlansForBandwidth' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY planSizes
 *  2. INTEGER targetBandwidth
 */
import "math"

func findMinimumPlansForBandwidth(planSizes []int32, targetBandwidth int32) int32 {
	// Write your code here
	if targetBandwidth == 0 {
		return 0
	}

	dp := make([]int32, targetBandwidth+1)
	dp[0] = 0
	for t := int32(1); t <= targetBandwidth; t++ {
		var minNum int32 = math.MaxInt32
		for _, plan := range planSizes {
			if t < plan || dp[t-plan] == -1 {
				continue
			}
			minNum = min(minNum, 1+dp[t-plan])
		}
		if minNum != math.MaxInt32 {
			dp[t] = minNum
		} else {
			dp[t] = -1
		}
	}
	//fmt.Println(dp)
	if dp[targetBandwidth] == 0 {
		return -1
	}
	return dp[targetBandwidth]
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     planSizesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var planSizes []int32

//     for i := 0; i < int(planSizesCount); i++ {
//         planSizesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         planSizesItem := int32(planSizesItemTemp)
//         planSizes = append(planSizes, planSizesItem)
//     }

//     targetBandwidthTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     targetBandwidth := int32(targetBandwidthTemp)

//     result := findMinimumPlansForBandwidth(planSizes, targetBandwidth)

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

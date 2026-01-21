package mergeandsortintervals

/*
 * Complete the 'mergeHighDefinitionIntervals' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts 2D_INTEGER_ARRAY intervals as parameter.
 */

import "slices"

func mergeHighDefinitionIntervals(intervals [][]int32) [][]int32 {
	// Write your code here
	if len(intervals) == 0 {
		return [][]int32{}
	}
	slices.SortFunc(intervals, func(a, b []int32) int {
		if a[0] < b[0] {
			return -1
		}
		if a[0] > b[0] {
			return 1
		}
		return 0
	})
	// fmt.Println(intervals)
	results := make([][]int32, 0)
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] > intervals[i][1] {
			results = append(results, intervals[i])
			continue
		}
		intervals[i+1][0] = min(intervals[i][0], intervals[i+1][0])
		intervals[i+1][1] = max(intervals[i][1], intervals[i+1][1])
	}
	results = append(results, intervals[len(intervals)-1])
	return results
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     intervalsRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     intervalsColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var intervals [][]int32
//     for i := 0; i < int(intervalsRows); i++ {
//         intervalsRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

//         var intervalsRow []int32
//         for _, intervalsRowItem := range intervalsRowTemp {
//             intervalsItemTemp, err := strconv.ParseInt(intervalsRowItem, 10, 64)
//             checkError(err)
//             intervalsItem := int32(intervalsItemTemp)
//             intervalsRow = append(intervalsRow, intervalsItem)
//         }

//         if len(intervalsRow) != int(intervalsColumns) {
//             panic("Bad input")
//         }

//         intervals = append(intervals, intervalsRow)
//     }

//     result := mergeHighDefinitionIntervals(intervals)

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

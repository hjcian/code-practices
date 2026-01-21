package maximumnumberofnonoverlappingintervals

/*
 * Complete the 'maximizeNonOverlappingMeetings' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY meetings as parameter.
 */
import "sort"

func maximizeNonOverlappingMeetings(meetings [][]int32) int32 {
	// Write your code here
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][1] < meetings[j][1] // Use > for descending order
	})
	count := 1
	first := 0
	for i := first + 1; i < len(meetings); i++ {
		if meetings[i][0] >= meetings[first][1] {
			count++
			first = i
		}
	}
	return int32(count)
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     meetingsRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     meetingsColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var meetings [][]int32
//     for i := 0; i < int(meetingsRows); i++ {
//         meetingsRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

//         var meetingsRow []int32
//         for _, meetingsRowItem := range meetingsRowTemp {
//             meetingsItemTemp, err := strconv.ParseInt(meetingsRowItem, 10, 64)
//             checkError(err)
//             meetingsItem := int32(meetingsItemTemp)
//             meetingsRow = append(meetingsRow, meetingsItem)
//         }

//         if len(meetingsRow) != int(meetingsColumns) {
//             panic("Bad input")
//         }

//         meetings = append(meetings, meetingsRow)
//     }

//     result := maximizeNonOverlappingMeetings(meetings)

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

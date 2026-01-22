package findpeakelementinbitonicarray

/*
 * Complete the 'findPeakIndex' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY counts as parameter.
 */

func findPeakIndex(counts []int32) int32 {
	// Write your code here
	left := 0
	right := len(counts) - 1
	for left <= right {
		// fmt.Println("left:", left, "right:", right)
		if left == right {
			return int32(left)
		}
		mid := (right + left) / 2
		if counts[mid] < counts[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     countsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var counts []int32

//     for i := 0; i < int(countsCount); i++ {
//         countsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         countsItem := int32(countsItemTemp)
//         counts = append(counts, countsItem)
//     }

//     result := findPeakIndex(counts)

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

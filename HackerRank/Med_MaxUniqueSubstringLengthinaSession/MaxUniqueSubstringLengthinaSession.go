package maxuniquesubstringlengthinasession

/*
 * Complete the 'maxDistinctSubstringLengthInSessions' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING sessionString as parameter.
 */

func maxDistinctSubstringLengthInSessions(sessionString string) int32 {
	// Write your code here
	if len(sessionString) == 0 || sessionString == "*" {
		return 0
	}
	lastSeen := make(map[byte]int, 0)
	left := 0
	right := 0
	length := 0
	for left < len(sessionString) && right < len(sessionString) {
		//fmt.Println("check:", sessionString[left:right+1])
		char := sessionString[right]
		if char == '*' {
			// reset all
			right++
			left = right
			lastSeen = make(map[byte]int, 0)
			continue
		}
		// Soft Delete Approach //
		if prevIdx, exists := lastSeen[char]; exists && prevIdx >= left {
			left = prevIdx + 1
		}
		lastSeen[char] = right
		if right-left+1 > length {
			length = right - left + 1
		}
		right++

		// Hard Delete Approach //
		// _, exists := counts[sessionString[right]]
		// if !exists {
		//     counts[sessionString[right]]++
		//     if right - left + 1 > length {
		//         length = right - left + 1
		//         //fmt.Println(".  found:", length)
		//     }
		//     right++
		//     continue
		// }
		// for left < right{
		//     delete(counts, sessionString[left])
		//     left++
		//     if sessionString[left-1] == sessionString[right] {
		//         break
		//     }
		// }
		//fmt.Println(". shrink:", sessionString[left:right+1])
	}
	return int32(length)
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     sessionString := readLine(reader)

//     result := maxDistinctSubstringLengthInSessions(sessionString)

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

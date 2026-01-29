package longestalternatingbinarysubstringwithlimitedflipslongestincreasingsubsequencelength

/*
 * Complete the 'longestAlternatingSubstring' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func isAlter(leading int, i int, s string) bool {
	// 1010101010
	remain := 0
	if leading == 0 {
		remain = 1
	}
	if i%2 == remain {
		// i = 0, 2, 4
		return s[i] == '1'
	} else {
		// i = 1, 3, 4
		return s[i] == '0'
	}
}

func pattern1(leading int, s string, k int32) int32 {
	i := 0
	longest := 0
	costs := 0
	//fmt.Println(s)
	for j := 0; j < len(s); j++ {
		//fmt.Println("[i:j]", i, j, "s:", s[i:j+1])
		if !isAlter(leading, j, s) {
			costs++
		}
		for costs > int(k) {
			if !isAlter(leading, i, s) {
				costs--
			}
			i++
		}
		//fmt.Println(".  [i:j]", i, j, "s:", s[i:j+1], "costs:", costs)
		if j-i+1 > longest {
			longest = j - i + 1
		}
	}
	return int32(longest)
}

func longestAlternatingSubstring(s string, k int32) int32 {
	// Write your code here
	var longest int32 = 0
	longest = max(longest, pattern1(1, s, k))
	longest = max(longest, pattern1(0, s, k))
	return int32(longest)
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     s := readLine(reader)

//     kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     k := int32(kTemp)

//     result := longestAlternatingSubstring(s, k)

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

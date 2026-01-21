package checkfornonidenticalstringrotation

/*
 * Complete the 'isNonTrivialRotation' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func isNonTrivialRotation(s1 string, s2 string) bool {
	// Write your code here
	if s1 == s2 {
		return false
	}
	counts := len(s1) - 1
	for counts > 0 {
		s2 = rotate1(s2)
		if s1 == s2 {
			return true
		}
		counts--
	}
	return false
}

func rotate1(str string) string {
	return str[1:] + string(str[0])
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     s1 := readLine(reader)

//     s2 := readLine(reader)

//     result := isNonTrivialRotation(s1, s2)

//     fmt.Printf("%s\n", (map[bool]string{true: "1", false: "0"})[result])
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

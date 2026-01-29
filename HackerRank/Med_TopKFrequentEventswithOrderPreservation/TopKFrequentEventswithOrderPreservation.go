package topkfrequenteventswithorderpreservation

/*
 * Complete the 'getTopKFrequentEvents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY events
 *  2. INTEGER k
 */

import "sort"

type Data struct {
	Num   int32
	Idx   int
	Count int
}

func getTopKFrequentEvents(events []int32, k int32) []int32 {
	// Write your code here
	freq := make(map[int32]*Data)
	datas := make([]*Data, 0)
	for i, event := range events {
		data, ok := freq[event]
		if !ok {
			data = &Data{
				Num:   event,
				Idx:   i,
				Count: 0,
			}
			freq[event] = data
			datas = append(datas, data)
		}
		data.Count++
	}
	sort.Slice(datas, func(i, j int) bool {
		if datas[i].Count > datas[j].Count {
			return true
		}
		if datas[i].Count == datas[j].Count {
			if datas[i].Idx < datas[j].Idx {
				return true
			}
		}
		return false
	})
	ret := make([]int32, k)
	for i := int32(0); i < k; i++ {
		ret[i] = datas[i].Num
	}
	// for _, data := range datas {
	//     fmt.Println(data)
	// }
	return ret
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     eventsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var events []int32

//     for i := 0; i < int(eventsCount); i++ {
//         eventsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         eventsItem := int32(eventsItemTemp)
//         events = append(events, eventsItem)
//     }

//     kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     k := int32(kTemp)

//     result := getTopKFrequentEvents(events, k)

//     for i, resultItem := range result {
//         fmt.Printf("%d", resultItem)

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

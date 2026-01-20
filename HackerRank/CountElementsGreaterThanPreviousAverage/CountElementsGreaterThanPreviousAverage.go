package countelementsgreaterthanpreviousaverage

func countResponseTimeRegressions(responseTimes []int32) int32 {
	// Write your code here
	var total, count int64 = 0, 0
	for i, num := range responseTimes {
		if i == 0 {
			total += int64(num)
			continue
		}
		//fmt.Println(num, total/int32(i))
		if int64(num)*int64(i) > total {
			count++
		}
		total += int64(num)
	}
	return int32(count)
}

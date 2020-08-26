package christmastree

import (
	"fmt"
	"strings"
)

func body(levelNum int, levelHeight int) []string {
	level := make([]string, 0, levelHeight*levelHeight)
	for i := 0; i < levelNum; i++ {
		for j := 0; j < levelHeight; j++ {
			spaces := strings.Repeat(" ", levelNum-i-1+levelHeight-j-1)
			stars := strings.Repeat("*", i*2+5+j*2)
			level = append(level, fmt.Sprintf("%s%s", spaces, stars))
		}
	}
	return level
}

func head(maxLen int) []string {
	return []string{
		fmt.Sprintf("%s%s", strings.Repeat(" ", maxLen/2), "*"),
		fmt.Sprintf("%s%s", strings.Repeat(" ", maxLen/2), "*"),
		fmt.Sprintf("%s%s", strings.Repeat(" ", maxLen/2-1), "***"),
	}
}

func foot(levelNum, levelHeight, maxLen int) []string {
	line := ""
	if levelHeight%2 == 1 {
		// odd
		line = strings.Repeat("*", levelHeight)
	} else {
		// even
		line = strings.Repeat("*", levelHeight+1)

	}
	space := strings.Repeat(" ", maxLen/2-len(line)/2)
	ret := make([]string, 0, levelNum)
	for i := 0; i < levelNum; i++ {
		ret = append(ret, fmt.Sprintf("%s%s", space, line))
	}
	return ret
}

func christmasTree(levelNum int, levelHeight int) []string {
	bodyPart := body(levelNum, levelHeight)
	maxLen := len(bodyPart[len(bodyPart)-1])
	headPart := head(maxLen)
	footPart := foot(levelNum, levelHeight, maxLen)
	return append(append(headPart, bodyPart...), footPart...)
}

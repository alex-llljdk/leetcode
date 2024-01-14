package main

import (
	"sort"
	"strconv"
	"strings"
)

//排序法，首先时分的节点最多有24*60个，因此超出该长度即有重复，返回0
//将时间列表转换为分钟制，方便算法计算操作，并进行升序排序
//特殊情况 计算最大值最小值的差值，在最小值中加入24小时的时间即可
//时间复杂度 遍历拆分时间 nlogn  空间复杂度 n

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 24*60 {
		return 0
	}
	var mins []int
	for _, t := range timePoints {
		time := strings.Split(t, ":")
		h, _ := strconv.Atoi(time[0])
		m, _ := strconv.Atoi(time[1])
		mins = append(mins, h*60+m)
	}

	sort.Ints(mins)
	mins = append(mins, mins[0]+24*60)
	ans := 1 << 30
	for i, x := range mins[1:] {
		ans = min(ans, x-mins[i])
	}
	return ans
}

func main() {

}

package main

//时间复杂度 n，空间复杂度 n
//利用栈的思想，每次都与最高位进行比较
func asteroidCollision(asteroids []int) (stk []int) {
	for _, x := range asteroids {
		if x > 0 {
			stk = append(stk, x)
		} else {
			//摧毁体积小的小行星
			for len(stk) > 0 && stk[len(stk)-1] > 0 && stk[len(stk)-1] < -x {
				stk = stk[:len(stk)-1]
			}
			//摧毁相等的小行星
			if len(stk) > 0 && stk[len(stk)-1] == -x {
				stk = stk[:len(stk)-1]
				//加入到行星堆，如果前一个是-的话则加入，如果都不满足，则是对面的行星比较大
			} else if len(stk) == 0 || stk[len(stk)-1] < 0 {
				stk = append(stk, x)
			}
		}
	}
	return
}

func main() {

}

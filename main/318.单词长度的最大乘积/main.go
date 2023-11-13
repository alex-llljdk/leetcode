package main

/*
位运算+枚举法

①、位运算，已知题目word只包含不重复小写字母，那么每个字符串都可以用32位的int表示，每个字符在一个位置表示是否存在
②、枚举，当记录完毕所有字符串的字符bit位，通过与计算，当相与为0时，则不包含相同的字母，通过遍历返回最大值
时间复杂度：O(N*(N+S)),S表示字符串的最大长度, 空间复杂度O(N)
*/
func maxProduct(words []string) (ans int) {
	mask := make([]int, len(words))

	for idx, word := range words {
		for _, v := range word {
			mask[idx] |= 1 << (v - 'a')
		}
	}

	for i, x := range mask {
		for j, y := range mask[:i] {
			if x&y == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}

/*
位运算+ 哈希表

由于计算每个word的mask会造成不必要的计算，如meet和met包含字母相同，只是出现次数和长度不同，且掩码表示相同。在检查是否相同时需要两次循环解决。
由于判断两个单词是否有公共字母是通过判断两个单词的位掩码的按位与运算实现，因此在位掩码相同的情况下，
单词的长度不会影响是否有公共字母，当两个位掩码的按位与运算等于 000 时，为了得到最大单词长度乘积，这两个位掩码对应的单词长度应该尽可能大。 时间复杂度：O(N*(N+S)),S表示字符串的最大长度, 空间复杂度O(N)
*/
// func maxProduct(words []string) (ans int) {
// 	masks := map[int]int{}
// 	for _, word := range words {
// 		mask := 0
// 		for _, ch := range word {
// 			mask |= 1 << (ch - 'a')
// 		}
// 		//记录当前掩码最大长度，避免重复遍历
// 		if len(word) > masks[mask] {
// 			masks[mask] = len(word)
// 		}
// 	}
// 	//缩小遍历空间，避免重复字段多次操作
// 	for x, lenX := range masks {
// 		for y, lenY := range masks {
// 			if x&y == 0 && lenX*lenY > ans {
// 				ans = lenX * lenY
// 			}
// 		}
// 	}

//		return
//	}
func main() {
	maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"})
}

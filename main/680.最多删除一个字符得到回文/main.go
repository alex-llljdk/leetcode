package main

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1

	check := func(i, j int) bool {
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return check(i, j-1) || check(i+1, j)
		}
	}
	return true
}

func main() {

}

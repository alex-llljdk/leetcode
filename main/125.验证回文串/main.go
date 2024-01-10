package main

func isPalindrome(s string) bool {
	m := len(s)

	i, j := 0, m-1

	for i < j {
		for i < j && !isalnum(s[i]) {
			i++
		}
		for i < j && !isalnum(s[j]) {
			j--
		}
		if tolower(s[i]) != tolower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isalnum(b byte) bool {
	return b >= '0' && b <= '9' ||
		b >= 'a' && b <= 'z' ||
		b >= 'A' && b <= 'Z'
}

func tolower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}

func main() {

}

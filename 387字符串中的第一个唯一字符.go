package main

func main() {

}

func firstUniqChar(s string) int {
	index := -1
	ba := make([]int, 26)
	ia := make([]int, 26)

	for i := 0; i < len(s); i++ {
		if ba[s[i]-'a'] == 0 {
			ba[s[i]-'a'] = 1
		} else if ba[s[i]-'a'] == 1 {
			ba[s[i]-'a'] = 2
		}
		ia[s[i]-'a'] = i
	}

	for i := 0; i < len(ba); i++ {
		if ba[i] == 1 {
			if ia[i] < index || index == -1 {
				index = ia[i]
			}
		}
	}

	return index
}

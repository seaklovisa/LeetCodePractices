package main

func main() {

	isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz")
}

func isAlienSorted(words []string, order string) bool {

	orderMap := make(map[byte]int)
	for i, r := range order {
		orderMap[r] = i
	}

	return true
}

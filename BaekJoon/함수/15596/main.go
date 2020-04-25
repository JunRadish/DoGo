package main

func main() {

}

func sum(a []int) int {

	value := 0
	for i := range a {
		value += a[i]
	}

	return value
}

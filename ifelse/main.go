package main

func main() {
	if i := 9; i < 0 {
		println(i, "is negative ")
	} else if i < 10 {
		println(i, "has  1 digit")
	} else {
		println(i, "has multiple digits")
	}

}

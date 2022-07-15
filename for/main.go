package main

func main() {
	i := 1
	for i <= 3 {
		println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		println(j)
	}

	for {
		println("loop")
		break
	}

	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			continue
		}
		println(k)
	}

}

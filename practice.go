package main

func practiceForLoops() {

	//i operates as the index, v operates as the value : Go automaticaly establishes this
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	m := map[int]string{1: "hello", 2: "world"}
	for _, v := range m {
		println(v)
	}

	var i int
	for i < 4 {
		println(i)
		i++
		if i == 3 {
			continue
		}
		println("continuing...")
	}

	for x := 0; x < 4; x++ {

	}

	for {
		if i < 1 {
			break
		}
		println(i)
		i--
	}
}

package main 

func main() {
	compLit()
	makeWay()
}

func compLit() {
	ii := []int{}
	l := len(ii)
	c := cap(ii)
	fmt.Println(l)
	fmt.Println(c)
	for i := 0; i < 10000; i++ {
		ii = append(ii, i)
		tempC := c
		if tempC != cap(ii) {
			c = cap(ii)
			fmt.Println("new cap:", c)
		}
	}
//func makeaWay()

}

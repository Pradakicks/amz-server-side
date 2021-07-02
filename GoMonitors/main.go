package main

import (
	Amz "amz-monitors/amz"
	"fmt"
)

func main() {
	fmt.Println("Testingz")
	go Amz.NewAmzMonitor("B08G9NKWR9", `Wr0RFj%2BisP6XDoYX7Pa9cykvPmqXPMBz2xsQfKuEDYFmtOC5nhLwdsyiTNEO21AnYke1zdAAU3Vw211fnIBH%2FjR5Dx%2FNsPiNq9zeu6NRq1aFc0dqh5Tsuaa3w7GDdk0Fq57g3QXpbBmx2Ry%2FzqIsTA%3D%3D`, "GPU")
	Amz.NewAmzMonitor("B08R81J62G", `dPIvIYRHvMgp7%2BKcLoVDXcE8qpp13AKCYwruZDDqR39vETkAKxHMkCUFsl3kyw%2F04zX2f4RprUI7BVuqL%2F4jozYXu6g5wHuk5JIw%2FG%2FRDF7zUXZsZREGGBtKeZC4Fh0llYVk1nywlWsUd%2FkIRglNAw%3D%3D`, "GPU")
	fmt.Scanln()
}

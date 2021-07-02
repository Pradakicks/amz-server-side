package main

import (
	"fmt"
	Amz "amz-monitors/amz"
)

func main() {
	fmt.Println("Testingz")
	go Amz.NewAmzMonitor("B08G9NKWR9", `Wr0RFj%2BisP6XDoYX7Pa9cykvPmqXPMBz2xsQfKuEDYFmtOC5nhLwdsyiTNEO21AnYke1zdAAU3Vw211fnIBH%2FjR5Dx%2FNsPiNq9zeu6NRq1aFc0dqh5Tsuaa3w7GDdk0Fq57g3QXpbBmx2Ry%2FzqIsTA%3D%3D`, "GPU")
	fmt.Scanln()
}

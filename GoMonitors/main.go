package main

import (
	Amz "amz-monitors/amz"
	"fmt"
)

func main() {
	fmt.Println("Testingz")
	// go Amz.NewAmzMonitor("B08G9NKWR9", `Wr0RFj%2BisP6XDoYX7Pa9cykvPmqXPMBz2xsQfKuEDYFmtOC5nhLwdsyiTNEO21AnYke1zdAAU3Vw211fnIBH%2FjR5Dx%2FNsPiNq9zeu6NRq1aFc0dqh5Tsuaa3w7GDdk0Fq57g3QXpbBmx2Ry%2FzqIsTA%3D%3D`, "GPU")
	go Amz.NewAmzMonitor("B08R81J62G", `dPIvIYRHvMgp7%2BKcLoVDXcE8qpp13AKCYwruZDDqR39vETkAKxHMkCUFsl3kyw%2F04zX2f4RprUI7BVuqL%2F4jozYXu6g5wHuk5JIw%2FG%2FRDF7zUXZsZREGGBtKeZC4Fh0llYVk1nywlWsUd%2FkIRglNAw%3D%3D`, "GPU")
	go Amz.NewAmzMonitor("B0881YZJ45", `eVgvKwk2KNfVckcnVqFmTn034BEaxYMxk4Z0iXluG1mRQqMfZJi1lW32aVYe2ld7rlRx3a55j2uGU69OPN%2FCE9Naz4iUQzkE8rCA65PsZTh0pATGa2rlKwPMoQvRsZsbq%2FX1AdoCg%2BjtNScjsOEP%2FQ%3D%3D`, "GPU")
	go Amz.NewAmzMonitor("B08TJ2BHCQ", `%2BbBuj%2B5lzqsQ14GCnhpcZA0jmUCWxGLgt1JP6No%2BUFUsGO%2FPcFoxOftbK6vtIb5uedA6wNlBwThPf22jtXENYp6S%2BdDGNsw0nXaI77aLMnrwABzZrOPxt7YsVkpYW%2BekOZ63W8%2Bl8Nta0XjeXilv4g%3D%3D`, "GPU")
	Amz.NewAmzMonitor("B08R81J62G", `dPIvIYRHvMgp7%2BKcLoVDXcE8qpp13AKCYwruZDDqR39vETkAKxHMkCUFsl3kyw%2F04zX2f4RprUI7BVuqL%2F4jozYXu6g5wHuk5JIw%2FG%2FRDF7zUXZsZREGGBtKeZC4Fh0llYVk1nywlWsUd%2FkIRglNAw%3D%3D`, "GPU")
	fmt.Scanln()
}

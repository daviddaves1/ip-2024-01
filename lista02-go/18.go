package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Printf("Digite 3 números inteiros diferentes de zero: ")
	fmt.Scan(&n1, &n2, &n3)

	for n1 == 0 || n2 == 0 || n3 == 0 {
		fmt.Printf("Nenhum número pode ser zero. Tente novamente: ")
		fmt.Scan(&n1, &n2, &n3)
	}
	mmc := n1
	for mmc%n2 != 0 || mmc%n3 != 0 {
		mmc += n1
	}

	fmt.Printf("%5d %5d %5d :%d\n", n1, n2, n3, len(fmt.Sprintf("%d", mmc)))
	fmt.Printf("%5d %5d %5d :%d\n", n2, n3, n1, len(fmt.Sprintf("%d", mmc)))
	fmt.Printf("%5d %5d %5d :%d\n", n3, n1, n2, len(fmt.Sprintf("%d", mmc)))
	fmt.Printf("MMC: %d\n", mmc)
}

package main

import "fmt"

func main(){
	var n int

	fmt.Print("Digite um valor inteiro N: ")
	fmt.Scanf("%d", &n)

	for i:=0; i<=n; i++{
		if(i%2==0){
			fmt.Printf("%d^%d = %d    ", i, i, i*i)
		}
	}
}
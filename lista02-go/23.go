package main

import "fmt"

func main() {
    var n int
    fmt.Println("Digite um número inteiro positivo:")
    fmt.Scan(&n)

    found := 0
    current := 1

    for found < n {
        sum1 := sumDivisors(current)
        sum2 := sumDivisors(sum1)

        if sum2 == current && current < sum1 && sum1 <= n {
            fmt.Printf("(%d,%d)\n", current, sum1)
            found++
        }
        current++
    }
}

func sumDivisors(num int) int {
    sum := 1 // 1 é sempre divisor
    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            sum += i
        }
    }
    return sum
}

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        n, _ := strconv.Atoi(scanner.Text())

        if n == 0 {
            break
        }

        scanner.Scan()
        numbers := make([]int, n)
        for i := 0; i < n; i++ {
            scanner.Scan()
            numbers[i], _ = strconv.Atoi(scanner.Text())
        }

        sortedNumbers := countingSort(numbers)
        for i := 0; i < n; i++ {
            fmt.Print(sortedNumbers[i], " ")
        }
        fmt.Println()
    }
}

func countingSort(numbers []int) []int {
    max := findMax(numbers)
    counts := make([]int, max+1)

    for _, number := range numbers {
        counts[number]++
    }

    z := 0
    for i := 0; i <= max; i++ {
        for counts[i] > 0 {
            numbers[z] = i
            z++
            counts[i]--
        }
    }

    return numbers
}

func findMax(numbers []int) int {
    max := numbers[0]
    for _, number := range numbers {
        if number > max {
            max = number
        }
    }
    return max
}

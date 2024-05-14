package main

import "fmt"

func main() {

	var q1, q2 int = 0, 0

  v1 := []int{}
  v2 := []int{} 

  fmt.Scanln(&q1, &q2)

  for i := 0; i < q1; i++ {
    var num int
    fmt.Scanln(&num)
    v1 = append(v1, num)
  }

  for i := 0; i < q2; i++ {
    var num int
    fmt.Scanln(&num)
    v2 = append(v2, num)
  }

  vr := intercalarOrdenado(v1, v2)

  for _, num := range vr {
    fmt.Println(num)
  }

}

func intercalarOrdenado(v1, v2 []int) []int {
	q1 := len(v1)
	q2 := len(v2)
	vr := make([]int, q1+q2)
  
	i, j, k := 0, 0, 0
	for i < q1 && j < q2 {
	  if v1[i] < v2[j] {
		vr[k] = v1[i]
		i++
	  } else {
		vr[k] = v2[j]
		j++
	  }
	  k++
	}
  
	copy(vr[k:], v1[i:])
	copy(vr[k:], v2[j:])
  
	return vr
  }

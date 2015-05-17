package main

import (
    "fmt"
)

func main() {
  var n, m , a, ndiva, mdiva int64

  fmt.Scanf("%d %d %d", &n, &m, &a)

  if n % a == 0 {
    ndiva = n / a
  } else {
    ndiva = n / a + 1
  }

  if m % a == 0 {
    mdiva = m / a
  } else {
    mdiva = m / a + 1
  }

  fmt.Println(ndiva * mdiva)
}

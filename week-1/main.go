package main

import (
  "fmt"
)

func cetak_kotak(angka int){
  if angka < 0 {
    fmt.Println("Angka harus Positif")
  } else if angka % 2 == 0 {
    fmt.Println("Angka harus Ganjil")
  } else {
    if (angka == 1) {
      fmt.Println("*")
    } else {
      mean := angka / 2
      for  x := 0 ; x < angka ; x++ {
        for  y := 0 ; y < angka ; y++ {
          if (x != mean){
            if y > 0 && y < angka-1 {
              fmt.Print("= ")
            } else {
              fmt.Print("* ")
            }
          } else {
            fmt.Print("* ")
          }
        }
        fmt.Println()
      }
    }
  }
  fmt.Println()
}


func main (){
  cetak_kotak(1)
  cetak_kotak(2)
  cetak_kotak(3)
  cetak_kotak(4)
  cetak_kotak(5)
  cetak_kotak(11)

}

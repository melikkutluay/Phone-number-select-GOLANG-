package main

import (
    "fmt"
    "log"
    "net/http"
    "reflect"
)
type Message struct {
    Status string
    Price string
}

func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/check" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        number := r.FormValue("name")
        fmt.Fprintf(w, "Phone Number is = %s\n", number)
        fmt.Println(reflect.TypeOf(number))
        fmt.Fprintf(w, fiyatlandir(number))
   }

func fiyatlandir(number1 string) string {
x := number1
var sonuc string=""
var y [10]string

if x >= "8508850000" && x <="8508853999"{
  sonuc="uygundur"
  for i, r := range x {
      y[i]=string(r)
  }
    for m := 0; m < 10; m++ {
      if y[6] == y[7]&&y[7] == y[8]&&y[8] == y[9] {
        /*fiyat=1000
        KDV = (fiyat*(18/100))
        toplam=fiyat + KDV*/
        sonuc="1180 TL"
        fmt.Println("1180")
      }else if y[3] == y[7]&&y[4] == y[8]&&y[5] == y[9] {
        sonuc="1180 TL"
        fmt.Println("1180")
      }else if y[7] == y[8]&&y[8] == y[9]{
        sonuc="590 TL"
        fmt.Println("590")
      }else if y[6] == y[8]&&y[7] == y[9]{
        sonuc="590 TL"
        fmt.Println("590")
      }else if ((y[6] == y[7] && y[8] == y[9])){
        sonuc="295 TL"
        fmt.Println("295")
      }else if y[7] == "0"  && y[9] == "0"  && y[8] != "0" {
        sonuc="118 TL"
        fmt.Println("118")
      }else if y[6] == "1" && y[7] == "9" && y[8] == "0" || y[8] == "2" && y[9] == "3" || y[9] == "5"{
        sonuc="59 TL"
        fmt.Println("59")
      }else{/*else if  y[7] == y[6]+"1"  && y[9] == y[8]+"1"  && y[8] == y[7]+"1"{
      fmt.Println("100")
    }*/
        sonuc="0 TL"
        fmt.Println("0")
      }
    }

}else{
  sonuc="istenilen sonuc uygun değildir"
}
return  sonuc
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./template")))
    http.HandleFunc("/check", hello)
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
/*  for k := 6; k < 9; k++ {
    for j := 7; j < 10; j++ {
      if k!=j {
      if y[k] == y[j] {
        fmt.Print(k)
        fmt.Print("->")
        fmt.Print(j)
        fmt.Print("=>")
        fmt.Print(y[k]+"-")
        fmt.Print(y[j]+"-")
        fmt.Println("oluyor gibi")
      }else {
        fmt.Print(k)
        fmt.Print("->")
        fmt.Print(j)
        fmt.Print("=>")
        fmt.Print(y[j]+"-")
        fmt.Print(y[k]+"-")
        fmt.Println("olmadı be")
      sayac=sayac+1
        }
      }
    }
  }*/

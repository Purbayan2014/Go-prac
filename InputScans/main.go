package InputScans 

import (
  "fmt"
)

func InputScans() {
  var ans1, ans2, ans3 string 

  fmt.Println("Name : ")
  _, err := fmt.Scan(&ans1)
  if err != nil {
    panic(err)
  }

  fmt.Println("Fav Food : ")
  _, err = fmt.Scan(&ans2)
  if err != nil {
    panic(err)
  }

  fmt.Println("Sport : ")
  _, err = fmt.Scan(&ans3)
  if err != nil {
    panic(err)
  }

  fmt.Println(ans1, ans2, ans3)

}

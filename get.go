package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "math/rand"
    "sync"
    "time"
)
var baseURL = "addme"

func random(min, max int) int {
  return rand.Intn(max - min) + min
}

func url(baseURL) string {
  id := random(1, 999999)
  return fmt.Sprintf("%s%07d", baseURL, id)
}

func HTTPGet(ch chan string, wg *sync.WaitGroup) {

  httpRequest, err:= http.NewRequest("GET", url(), nil)
  httpRequest.Header.Set("Accept", "addme")
  httpRequest.Header.Set("addme", "demo")
  httpResponse, err := http.DefaultClient.Do(httpRequest)
  if err != nil {
    ch <- fmt.Sprintf("%s", err)
  }
  body, _ := ioutil.ReadAll(httpResponse.Body)
  ch <- fmt.Sprintf("%s", string(body))
  httpResponse.Body.Close()
  wg.Done()

}

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  var wg sync.WaitGroup
  ch := make(chan string)
  for j := 1; j <= 10000; j++ {
    for j := 1; j <= 50; j++  {
      wg.Add(1)
      go HTTPGet(ch, &wg)
    }

    // for j := 1; j <= 50; j++  {
    //   wg.Wait()
    //   fmt.Println(<-ch)
    // }
  }
}

package main

import (
  "fmt"
  "regexp"
  "net/http"
  "io/ioutil"
  "os"
)

func main() {
  persist(find())
  // upload new pictures on flickr
}

/*********************************************************/
/* FIND functions                                        */
/*********************************************************/
func find() (pictures []string) {
  re := regexp.MustCompile("http://(.*).staticflickr.com/[\\d]+/[\\d]+_[\\d\\w]+.(?:jpg|gif|png)")
  return re.FindAllString(get("http://www.flickr.com/search/?l=commderiv&q=pug"), -1)
}

func get(url string) (body string) {
  response, err := http.Get(url)

  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  } else {
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)

    if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
    }

    body = string(contents)
  }

  return body
}

/*********************************************************/
/* PERSIST functions                                     */
/*********************************************************/
func persist(pictures []string) {
  for i := 0; i < len(pictures); i++ {
    fmt.Println(pictures[i])
  }
}

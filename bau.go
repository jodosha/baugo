package main

import (
  "fmt"
  "regexp"
  "strings"
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

/*********************************************************/
/* PERSIST functions                                     */
/*********************************************************/
func persist(pictures []string) {
  for i := 0; i < len(pictures); i++ {
    writeFile(pictures[i])
  }
}

func writeFile(picture string) {
  dirname := "tmp"
  os.Mkdir(dirname, 0700)

  file, err := os.Create(filePath(picture, dirname))
  handleError(err)

  defer file.Close()

  file.WriteString(get(picture))
}

func filePath(picture string, dirname string) (filePath string) {
  return dirname + "/" + filename(picture)
}

func filename(picture string) (filename string) {
  tokens := strings.Split(picture, "/")
  return tokens[len(tokens) - 1]
}

/*********************************************************/
/* GENERAL functions                                     */
/*********************************************************/

func get(url string) (body string) {
  response, err := http.Get(url)
  handleError(err)

  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)

  handleError(err)

  return string(contents)
}

func handleError(err error) {
  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  }
}

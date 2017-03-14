package main
 
import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "path"
)
 
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome, %!", r.URL.Path[1:])
}
 
func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/about/", about)
  http.ListenAndServe(":8080", nil)
}
 
type Message struct {
  Message string
  Version string
  Source string
}
 
func about (w http.ResponseWriter, r *http.Request) {
  ex, err := os.Executable()
  if err != nil {
     panic(err)
  }
  exPath := path.Dir(ex)

  m := Message{"This is based on the blog post http://nordicapis.com/writing-microservices-in-go/", "0.0.1", exPath}
  b, err := json.Marshal(m)

  if err != nil {
      panic(err)
  }
 
  w.Write(b)
}
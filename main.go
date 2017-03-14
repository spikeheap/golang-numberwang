package main
 
import (
  "encoding/json"
  "net/http"
  "os"
  "path"
  "math/rand"
)
 
func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/about/", about)
  http.ListenAndServe(":8080", nil)
}

type NumberWangMessage struct {
  Value int
  NumberWang bool
  WangerNum bool
}

func handler(w http.ResponseWriter, r *http.Request) {
  //rand.Intn(len(answers))
  number := rand.Int()
  isNumberWang := false
  isWangerNum := false
  m := NumberWangMessage{number, isNumberWang, isWangerNum}
  b, err := json.Marshal(m)

  if err != nil {
      panic(err)
  }
  
  w.Write(b)
}
 
type AboutMessage struct {
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

  m := AboutMessage{"This is based on the blog post http://nordicapis.com/writing-microservices-in-go/", "0.0.1", exPath}
  b, err := json.Marshal(m)

  if err != nil {
      panic(err)
  }
 
  w.Write(b)
}
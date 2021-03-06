package main
 
import (
  "encoding/json"
  "log"
  "math/rand"
  "net/http"
  "os"
  "strconv"
)
 
func main() {
  port := getenv("PORT", "8080")
  http.HandleFunc("/", handler)
  http.HandleFunc("/about/", about)
  http.HandleFunc("/favicon.ico", favicon)

  log.Println("Listening on "+port)
  http.ListenAndServe(":"+port, nil)
  log.Println("Exiting...")
}

func favicon(w http.ResponseWriter, r *http.Request) {
  // NOOP
}

type NumberWangMessage struct {
  Value string
  NumberWang bool
  WangerNum bool
}

func handler(w http.ResponseWriter, r *http.Request) {
  betterNumbers:= [4]string{ "twentington", "shinty-six", "filth hundred and knee", "noot" }

  var number string
  if(rand.Intn(100) > 85) {
    number = betterNumbers[rand.Intn(len(betterNumbers))]
  }else{
    number = strconv.Itoa(rand.Intn(1000))
  }

  isNumberWang := rand.Intn(100) > 10
  isWangerNum := rand.Intn(100) > 90

  m := NumberWangMessage{number, isNumberWang, isWangerNum}
  b, err := json.Marshal(m)

  if err != nil {
    log.Println(err)
  }else{
    log.Println(string(b))
    w.Write(b)
  }
}
 
type AboutMessage struct {
  Message string
  Version string
  License string
}
 
func about (w http.ResponseWriter, r *http.Request) {
  m := AboutMessage{"Based on this post: http://nordicapis.com/writing-microservices-in-go/", "0.0.1", "MIT (https://opensource.org/licenses/MIT)"}
  b, err := json.Marshal(m)

  if err != nil {
    log.Println(err)
  }else{
    log.Println(string(b))
    w.Write(b)
  }
}

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
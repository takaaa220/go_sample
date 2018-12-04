package main

import(
  "log"
  "fmt"
  "net/http"
)

type templateHandler struct {
  once sync.Once
  filename string
  temp1 *template.Template
}
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  t.once.Do(func(){
    t.temp1 =
      template.Must(template.ParseFiles(filepath.Join("tmplates", t.filename)))
  })
  t.temp1.Execute(w, nil)
}

func main() {
  http.HandleFunc("/", handler)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndserve:", err)
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello Hello!")
}

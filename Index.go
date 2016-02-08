package main

import (
  "log"
  "fmt"
  "net/http"
  "strings"
  _ "html/template"
  _ "io/ioutil"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Bienvenido %s!</h1>", r.URL.Path[1:])
}

func formularioCarros(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Insertar</h1>"+
    "<form action=\"/save_carro/\" method=\"GET\">"+
    "<label>Marca</label>"+
    "<input type=\"text\" name=\"marca\"><br>"+
    "<label>Modelo</label>"+
    "<input type=\"text\" name=\"modelo\"><br>"+
    "<input type=\"submit\" value=\"Enviar\">"+
    "</form>")
}

func carros(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  marcaCarro:=strings.Join(r.Form["marca"], "")
  modeloCarro:=strings.Join(r.Form["modelo"], "")
  
  database, err := sql.Open("mysql", "fco:fco@/carros_go?charset=utf8")
  if err!=nil{
    log.Println("Bade de datos... ",err)
  }
  statment, err := database.Prepare("INSERT clasicos SET marca=?,modelo=?")
  res, err := statment.Exec(marcaCarro, modeloCarro)
  if err!=nil{
    log.Println("Insertar...",err)
  }
  id, err := res.LastInsertId()
  log.Println("Insertado con exito... ",id," ",marcaCarro," ",modeloCarro)
  fmt.Fprintf(w,"<a href=\"/\">Inicio</a>")
}



func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/carros/", formularioCarros)
  http.HandleFunc("/save_carro/", carros)
  log.Println("ServeMux of Golang run in :9876")
  http.ListenAndServe(":9876", nil)
  log.Println("Server stop")
}
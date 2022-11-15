package main

import (
	_ "io"
	"github.com/Unleash/unleash-client-go/v3"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type metricsInterface struct {
}

func init() {
    unleash.Initialize(
        unleash.WithUrl("https://gitlab.com/api/v4/feature_flags/unleash/40951967"),
        unleash.WithInstanceId("8bZJ99faxtsW3anLf2ak"),
        unleash.WithAppName("production"), // Set to the running environment of your application
        unleash.WithListener(&metricsInterface{}),
    )
}


func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", CatHandler)
	log.Printf("Listening on port 5000!")
	http.ListenAndServe(":5000", nil)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func CatHandler(w http.ResponseWriter, r *http.Request) {
	//Fetch hostname of container

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	// Choose random catpicture
	
	catpic := random(1, 1)
	if unleash.IsEnabled("catpic-random"){
		catpic := random(2, 14)
		fmt.Println(catpic)
	} else {
        catpic := 1
		fmt.Println(catpic)
	}


	//Parse index.html template
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}

	items := struct {
		Url      int
		Hostname string
	}{
		Url:      catpic,
		Hostname: name,
	}

	t.Execute(w, items)
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

}
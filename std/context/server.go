package main

import (
	"context"
	"go-study/std/context/google"
	"go-study/std/context/userip"
	"html/template"
	"log"

	"net/http"
	"time"
)
var resultsTemplate  = template.Must(template.New("results").Parse(`
<html>
	<head/>
	<body>
		<ol>
		{{range .Results}}
			<li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a> </li>
		{{end}}
		</ol>
		<p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}} </p>
	</body>
</htnl>
`))

func main(){
	// https://go.dev/blog/context

	// test:
	// http://localhost:8080/search?q=headache
	// http://localhost:8080/search?q=headache&timeout=100ms
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSearch(w http.ResponseWriter,
	req *http.Request){
		var (
			ctx context.Context
			cancel context.CancelFunc
		)

		timeout, err := time.ParseDuration(req.FormValue("timeout"))
		if err == nil{
			ctx, cancel = context.WithTimeout(context.Background(), timeout)
		}else{
			ctx, cancel = context.WithCancel(context.Background())
		}
		defer cancel()
		query := req.FormValue("q")
		if query  == ""{
			http.Error(w, "no query", http.StatusBadRequest)
			return
		}

		userIp, err := userip.FromRequest(req)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		//store the user ip in ctx for use by other pkg
		ctx = userip.NewContext(ctx, userIp)

		start := time.Now()
		// Search will get ip from ctx
		results, err := google.Search(ctx, query)
		elapsed := time.Since(start)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := resultsTemplate.Execute(w, struct{
			Results google.Results 
			Timeout, Elapsed time.Duration 
		}{
			Results: results,
			Timeout: timeout,
			Elapsed: elapsed,
		}); err !=nil{
			log.Print(err)
			return
		}
}


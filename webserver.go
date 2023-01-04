// 参考
// https://zenn.dev/skonb/articles/0bad1d59371d09
// https://docs.kilvn.com/build-web-application-with-golang/ja/03.2.html

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	customsearch "google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

func hoge(w http.ResponseWriter, r *http.Request) {
	search_id := os.Getenv("SEARCH_ID")
	fmt.Println("search_engine_id: ", search_id)
	token := os.Getenv("CUSTOMSEARCH_API_TOKEN")
	fmt.Println("API_token", token)

	// using customsearch golang api
	r.ParseForm()
	word := strings.Join(r.Form["word"], "")
	fmt.Println("search_term: ", word)
	result_html := "<ul>"
	if word != "" {
		ctx := context.Background()
		cseService, _ := customsearch.NewService(ctx, option.WithAPIKey(token))
		search := cseService.Cse.List().Q(word)
		search.Cx(search_id)
		search.Start(1)
		call, err := search.Do()
		if err != nil {
			log.Fatal(err)
		}
		for _, r := range call.Items {
			result_html += "<li><a href=" + r.Link + ">" + r.HtmlTitle + "</a></li>"
		}
	}
	result_html += "</ul>"

	// using custom search engine
	res, err := http.Get("https://customsearch.googleapis.com/customsearch/v1?cx=" + search_id + "&key=" + token + "&q=" + word)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	search_body, _ := io.ReadAll(res.Body)

	// using customsearch through googleAPI to get result with json
	res_json, err := http.Get("https://www.googleapis.com/customsearch/v1?key=" + token + "&q=" + word + "&cx=" + search_id)
	res_json_body, _ := io.ReadAll(res_json.Body)
	defer res_json.Body.Close()
	template := `
	<html>
	<head><title>hoge</title></head>
	<body>
	<div name="custom search">
	<h3>1. Programmable Search Element Control API</h3>
	Implemented by JavaScript in this html.
	You can customize this code on the control pannel of the programmable search engine.
		<script async src="https://cse.google.com/cse.js?cx=%s">
		</script>
		<div class="gcse-search"></div>
	</div>
	<div>
	<h3>Search from form</h3>
	<div>
	<form>
	<input id="word" name="word" />
	<input type="submit" />
	</form>
	</div>
	<h4>2-1. result from API</h4>
	<a href="https://www.googleapis.com/customsearch/v1">endpoint</a>
	<div>
	%s
	</div>
	<h4>2-2. result from JSON API</h4>
	<a href="https://customsearch.googleapis.com/customsearch/v1">endpoint</a>
	<div>
	%s
	</div>
	<h4>3. result from golang API (list of link)</h4>
	<div>
	%s
	</div>
	</div>
	</body>
	`
	fmt.Fprintf(w, template, search_id, search_body, res_json_body, result_html)
}

func main() {
	http.HandleFunc("/", hoge)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal("ListenAndserve: ", err)
	}
}

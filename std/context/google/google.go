package google

import (
	"context"
	"encoding/json"
	"go-study/std/context/userip"
	"net/http"
	"os"
)

type Result struct{
	Title, URL string 
}
type Results []Result

func Search(ctx context.Context,  query string)(Results, error){
	// https://developers.google.com/custom-search/v1/overview?hl=zh-cn
	google_custom_search_cx := os.Getenv("GOOGLE_CUSTOM_SEARCH_CX")
	if google_custom_search_cx == ""{
		panic("google custom search cx not exist")
	}
	google_custom_search_key := os.Getenv("GOOGLE_CUSTOM_SEARCH_KEY")
	if google_custom_search_key == ""{
		panic("google custom search key not exist")
	}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/customsearch/v1",nil)
	if err != nil{
		return nil, err 
	}
	//set query string
	q := req.URL.Query()
	q.Set("q", query)
	q.Set("key", google_custom_search_key)
	q.Set("cx", google_custom_search_cx)

	if userIP, ok := userip.FromContext(ctx); ok{
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()
	var results Results 
	err = httpDo(ctx, req, func(resp *http.Response, err error) error{
		if err != nil{
			return err 
		}
		defer resp.Body.Close()
		var data struct{
				Items[]  struct{
					Title string `json:"title"`
					Link string `json:"link"`
				} `json:"items"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err !=nil{
			return err 
		}
		for _, res := range data.Items{
			results = append(results, Result{Title: res.Title,
			URL: res.Link })
		}
		return nil 
	})
	return results, err 
}

// make req in a goroutine and pass response to to func f
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error)error) error{
	c := make(chan error, 1)
	req = req.WithContext(ctx)
	go func(){
		c <- f(http.DefaultClient.Do(req))
	}()
	select{
	case <- ctx.Done(): //maybe request is cancelled
		<-c  //wait for func f to return
		return ctx.Err()
	case err := <-c:
		return err 
	}



}
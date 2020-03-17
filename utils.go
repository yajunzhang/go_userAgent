package go_userAgent

import (
	"fmt"
	"github.com/gocolly/colly"
	"encoding/json"
	"os"
	"io/ioutil"
)

func Get(name string){
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.2.3 Safari/537.02"),
	)

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML("[id=liste]", func(e *colly.HTMLElement) {		
		var agent []string
		e.ForEach("li", func(_ int, e1 *colly.HTMLElement){
			agent = append(agent, e1.Text)
		})

		jsons, errs := json.Marshal(agent)
		if errs != nil {
			fmt.Println(errs.Error())
		}

		write(name, jsons)
	})

	c.Visit("http://useragentstring.com/pages/useragentstring.php?name="+name)
}

func write(name string, content []byte){
	f, err := os.OpenFile(cachePath+name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
    defer f.Close()
    if err != nil {
        fmt.Println(err.Error())
	} 
	f.Write(content)
}

func Read(name string) []string {
	bytes, err := ioutil.ReadFile(cachePath+name)
	if err != nil {
		fmt.Println(err)
	}

	m := make([]string, 0)
    err = json.Unmarshal(bytes, &m)
    if err != nil {
        fmt.Println("Umarshal failed:", err)
        return nil
	}
	
	return m
}


// func main(){
// 	utils := NewUtils()
// 	utils.get("Safari")
// }
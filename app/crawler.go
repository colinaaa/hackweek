package crawler

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"regexp"
// 	"strings"

// 	"github.com/gocolly/colly"
// )

// var count int

// func oldvisit(s string) chan struct{} {
// 	ch := make(chan struct{})
// 	go func() {
// 		c := colly.NewCollector(
// 			colly.URLFilters(
// 				regexp.MustCompile("http://www\\.douguo\\.com/caipu/.*?/[1-9+]"),
// 				regexp.MustCompile("https://www\\.douguo\\.com/caipu"),
// 			),
// 			colly.Async(false),
// 			colly.CacheDir("./cache"+s),
// 		)
// 		c.OnRequest(func(r *colly.Request) {
// 			r.Ctx.Put("category", getCategory(r.URL.String()))
// 			// fmt.Println(r.URL)
// 			count++
// 		})
// 		c.OnHTML("ul[class='sortlist clearfix']", func(e *colly.HTMLElement) {
// 			e.ForEach("li > a", func(i int, e *colly.HTMLElement) {
// 				link := e.Request.AbsoluteURL(e.Attr("href"))
// 				// log.Println("visiting" + getCategory(link))
// 				e.Request.Visit(link)
// 			})
// 		})
// 		c.OnHTML("a[class='cover']", func(e *colly.HTMLElement) {
// 			src := e.ChildAttr("img", "src")
// 			name := e.ChildAttr("img", "alt")
// 			name = strings.Replace(name, "/", "_", -1)
// 			category := e.Response.Ctx.Get("category")
// 			getImg(src, category+"-----"+name+".jpg", category)
// 		})
// 		c.OnHTML("div[class='pages']", func(e *colly.HTMLElement) {
// 			e.ForEach("div > a[class='anext']", func(i int, e *colly.HTMLElement) {
// 				link := e.Request.AbsoluteURL(e.Attr("href"))
// 				// fmt.Println(link)
// 				e.Request.Visit(link)
// 			})
// 		})
// 		c.Visit(s)
// 		c.Wait()
// 		close(ch)
// 	}()
// 	return ch
// }
// func getCategory(u string) string {
// 	s := strings.Split(u, "/")
// 	if len(s) > 4 {
// 		text, err := url.QueryUnescape(s[4])
// 		if err != nil {
// 			return ""
// 		}
// 		return text
// 	}
// 	return ""
// }
// func main() {
// 	for range oldvisit(`https://www.douguo.com/caipu/fenlei`) {
// 		log.Println("hello")
// 	}
// 	fmt.Println(count)
// }

// func getImg(url, name, category string) (n int64, err error) {
// 	// path := strings.Split(url, "/")
// 	// fmt.Println(name)
// 	_ = os.Mkdir(category, 0666)
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Println(111, err)
// 		return -1, err
// 	}
// 	defer response.Body.Close()

// 	//open a file for writing
// 	file, err := os.Create(name)
// 	if err != nil {
// 		log.Println(222, err)
// 		return -1, err
// 	}
// 	defer file.Close()

// 	// Use io.Copy to just dump the response body to the file. This supports huge files
// 	_, err = io.Copy(file, response.Body)
// 	if err != nil {
// 		log.Println(333, err)
// 		return -1, err
// 	}
// 	return -1, nil
// }

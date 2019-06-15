package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/xerrors"
)

func GetPage(url string, atr string) ([]string, error) {
	var src []string
	res, err := http.Get(url)
	fmt.Println(res)
	if err != nil {
		return src, xerrors.Errorf("err request[URL:%s]: %w", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr(atr)
		src = append(src, url)
	})
	return src, nil
}

func main() {
	url := "google"
	s, _ := GetPage(url, "")
	fmt.Println(s)

}

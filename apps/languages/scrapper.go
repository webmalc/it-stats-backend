package languages

import (
	"fmt"

	"github.com/pkg/errors"
)

// const redmonkLink = "https://redmonk.com/sogrady/"

// type RedmonkScrapper struct {
// 	Link      string
// 	collector *colly.Collector
// }

// func (s *RedmonkScrapper) Scrap() ([]*Language, error) {
// 	result := make([]*Language, 0)
// 	articleSelector := "a[title*='The RedMonk Programming Language Rankings']"
// 	ratingSelector := "p"
// 	pageSelector := "a.page-numbers"
// 	articleFound := false
// 	detailCollector := s.collector.Clone()

// 	detailCollector.OnHTML(ratingSelector, func(e *colly.HTMLElement) {
// 		if len(e.Text) > 10 && e.Text[0:2] == "1 " {
// 			scanner := bufio.NewScanner(strings.NewReader(e.Text))
// 			for scanner.Scan() {
// 				fmt.Println(strings.Split(scanner.Text(), " ")[1])
// 			}
// 		}
// 	})

// 	s.collector.OnHTML(articleSelector, func(e *colly.HTMLElement) {
// 		if !articleFound {
// 			fmt.Println(e.Attr("href"))
// 			detailCollector.Visit(e.Attr("href"))
// 			articleFound = true
// 		}
// 	})

// 	s.collector.OnHTML(pageSelector, func(e *colly.HTMLElement) {
// 		if !articleFound {
// 			e.Request.Visit(e.Attr("href"))
// 		}
// 	})

// 	err := s.collector.Visit(s.Link)
// 	return result, err
// }

// const tiobeLink = "https://tiobe.com/tiobe-index/"

// type TiobeScrapper struct {
// 	Link      string
// 	collector *colly.Collector
// }

// func (s *TiobeScrapper) Scrap() ([]*Language, error) {
// 	result := make([]*Language, 0)
// 	s.collector.OnHTML("table#top20 tr", func(e *colly.HTMLElement) {
// 		pos, _ := strconv.Atoi(e.ChildText("td:nth-child(1)"))
// 		title := e.ChildText("td:nth-child(4)")
// 		if pos > 0 && title != "" {
// 			result = append(result, &Language{
// 				Title:    title,
// 				Position: uint(pos),
// 				Source:   "tiobe",
// 			})
// 		}
// 	})
// 	err := s.collector.Visit(s.Link)
// 	return result, err
// }

// const pyplLink = "https://pypl.github.io/PYPL.html"

// type PyplScrapper struct {
// 	Link string
// }

// func (s *PyplScrapper) Scrap() ([]*Language, error) {
// 	result := make([]*Language, 0)
// 	resp, _ := http.Get(s.Link)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	defer resp.Body.Close()

// 	re := regexp.MustCompile(`(?s)begin section All-->\\(.*?)<!-- end section All`)

// 	sdt := string(re.FindSubmatch(body)[1])
// 	sdt = strings.ReplaceAll(sdt, "\\\"", "\"")
// 	sdt = strings.ReplaceAll(sdt, ">\\", ">")
// 	sdt = "<html><body><table>" + sdt + "</table></body></html>"

// 	doc := soup.HTMLParse(sdt)
// 	for _, el := range doc.FindAll("tr") {
// 		tds := el.FindAll("td")
// 		pos, _ := strconv.Atoi(tds[0].Text())

// 		result = append(result, &Language{
// 			Title:    tds[2].Text(),
// 			Position: uint(pos),
// 			Source:   "pypl",
// 		})
// 	}
// 	return result, nil
// }

type ScrappersRunner struct {
	scrappers []Scrapper
	logger    Logger
}

func (s *ScrappersRunner) Run() {
	for _, scrapper := range s.scrappers {
		s.logger.Infof("Started a language scrapper: %#v", scrapper)
		languages, err := scrapper.Scrap()
		if err != nil {
			s.logger.Error(errors.Wrap(err, "languages scrapper"))
		}
		for _, l := range languages {
			fmt.Println(l)
		}
	}
}

func NewScrapperRunner(logger Logger) *ScrappersRunner {
	// c := colly.NewCollector()
	return &ScrappersRunner{logger: logger, scrappers: []Scrapper{
		// &TiobeScrapper{Link: tiobeLink, collector: c},
		// &PyplScrapper{Link: pyplLink},
		// &RedmonkScrapper{Link: redmonkLink, collector: c},
	}}
	// spectrum.ieee.org
	// githut.info
}

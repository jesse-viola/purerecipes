package scrape

type ScrapeManager struct {
}

func NewScrapeManager() *ScrapeManager {
	return &ScrapeManager{}
}

func (sm *ScrapeManager) ScrapeWebUrl(url string) string {
	return url + "backend is fake scraping"
}

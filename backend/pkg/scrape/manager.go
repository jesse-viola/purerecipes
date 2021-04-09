package scrape

// ScrapeManager is the implementation of ScrapeCruder interface
type ScrapeManager struct {
}

// NewScrapeManager returns pointer to ScrapeManager.
func NewScrapeManager() *ScrapeManager {
	return &ScrapeManager{}
}

// ScrapeWebUrl webcrawls and searches for the ingredients.
func (sm *ScrapeManager) ScrapeWebUrl(url string) string {
	return url + " " + "backend is fake scraping"
}

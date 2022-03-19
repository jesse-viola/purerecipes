from fastapi import APIRouter, Depends

from server.scrape.services import WebsiteCrawlerInterface as WebsiteCrawler
from server.url.schemas import URL

router = APIRouter()


@router.post("clean")
async def clean_url(url_in: URL, website_crawler: WebsiteCrawler = Depends()):
    website_crawler.crawl(url_in.link)

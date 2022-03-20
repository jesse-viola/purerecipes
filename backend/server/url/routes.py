from fastapi import APIRouter, Depends

from server.scrape.services import WebsiteCrawlerInterface as WebsiteCrawler
from server.url import schemas as url_schemas

router = APIRouter()


@router.post("/clean")
async def clean_url(url_in: url_schemas.URL, website_crawler: WebsiteCrawler = Depends()):
    website_crawler.crawl(url_in.link)

import requests

from fastapi import HTTPException
from bs4 import BeautifulSoup

from .services import WebsiteCrawlerInterface


class BeautifulSoupManager(WebsiteCrawlerInterface):
    def _get_website_html(self, website_link: str) -> requests.Response:
        response = requests.get(website_link)
        if response.status_code != 200:
            raise HTTPException(
                status_code=500, detail=f"Failed to retrieve data from {website_link}"
            )
        return response

    def crawl(self, website_link: str) -> None:
        html_page = self._get_website_html(website_link)
        soup = BeautifulSoup(html_page.content, "lxml")

        print(soup)
        
website_crawler = BeautifulSoupManager()

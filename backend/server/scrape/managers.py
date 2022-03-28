from fastapi import HTTPException

from internal import recipe_scrapers

from .services import IScrape


class ScrapeManager(IScrape):
    def get_recipe(self, link: str) -> None:
        """
        soup.findAll() maybe check all the headers? then regex string search
        the results for keywords.

        Good example https://github.com/hhursev/recipe-scrapers/tree/main/recipe_scrapers
        powered by https://github.com/scrapinghub/extruct
        structured by https://schema.org/docs/gs.html#schemaorg

        """
        recipe_scrapers.scrape_me()
        
        return soup
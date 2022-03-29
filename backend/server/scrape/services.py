from abc import ABC, abstractmethod

from .schemas import Recipe


class IScrape(ABC):
    @abstractmethod
    def get_recipe(self, website_link: str) -> Recipe:
        """Takes a given website link parses with recipe_scrapers module

        Args:
            website_link (str): link of website to parse
        """
        raise NotImplementedError

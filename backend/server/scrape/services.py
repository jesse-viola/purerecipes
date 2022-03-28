from abc import ABC, abstractmethod


class IScrape(ABC):
    @abstractmethod
    def get_recipe(self, website_link: str) -> None:
        """Takes a given website link parses with recipe_scrapers module

        Args:
            website_link (str): link of website to parse
        """
        raise NotImplementedError

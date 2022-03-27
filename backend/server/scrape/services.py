from abc import ABC, abstractmethod


class WebsiteCrawlerInterface(ABC):
    @abstractmethod
    def crawl(self, website_link: str) -> None:
        """Takes a given website link and parses into readable data

        Args:
            website_link (str): link of website to parse
        """
        raise NotImplementedError

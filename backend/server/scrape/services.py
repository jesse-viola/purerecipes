from abc import ABCMeta, abstractmethod


class WebsiteCrawlerInterface(ABCMeta):
    @abstractmethod
    def crawl(self, website_link: str) -> None:
        """Takes a given website link and parses into readable data

        Args:
            website_link (str): link of website to parse
        """
        raise NotImplementedError

from typing import Dict, List, Optional
from pydantic import BaseModel


class URL(BaseModel):
    """Post body for scraping route"""

    link: str


class Recipe(BaseModel):
    """Recipe schema based on :py:class:`~recipe-scrapers.SchemaScraper`"""

    title: str = ""
    description: str = ""
    ingredients: List[str] = []
    instructions: str = ""
    total_time: int = 0
    prep_time: int = 0
    cook_time: int = 0
    image: str = ""
    category: str = ""
    yields: str = ""
    author: Optional[str]
    language: Optional[str]
    nutrients: Optional[Dict[str, str]]
    cuisine: Optional[str]

from typing import List, Optional, Dict, Union
from pydantic import BaseModel

class URL(BaseModel):
    link: str


class Recipe(BaseModel):
    """Recipe schema based on :py:class:`~recipe-scrapers.SchemaScraper`"""

    title: str
    description: str
    ingredients: List[str]
    instructions: str
    total_time: int
    prep_time: int
    cook_time: int
    image: str
    category: str
    yields: str
    author: Optional[str]
    language: Optional[str]
    nutrients: Optional[Dict[str, str]]
    cuisine: Optional[str]

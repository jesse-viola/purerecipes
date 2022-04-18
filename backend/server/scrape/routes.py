from fastapi import APIRouter

from server.managers import scrape_manager
from .schemas import URL, Recipe

router = APIRouter()


@router.post("/recipe", response_model=Recipe)
async def get_recipe(url_in: URL) -> Recipe:
    return scrape_manager.get_recipe(url_in.link)

from fastapi import APIRouter

from .schemas import URL, Recipe
from server.managers import scrape_manager

router = APIRouter()


@router.post("/recipe", response_model=Recipe)
async def get_recipe(url_in: URL) -> Recipe:
    return scrape_manager.get_recipe(url_in.link)

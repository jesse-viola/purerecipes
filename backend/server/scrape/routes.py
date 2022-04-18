from typing import Optional
from fastapi import APIRouter

from server.managers import scrape_manager
from .schemas import URL, Recipe

router = APIRouter()


@router.post("/recipe", response_model=Recipe)
async def get_recipe(url_in: URL) -> Optional[Recipe]:
    """Returns a recipe from the given URL

    Args:
        url_in (`schemas.URL`): model URL

    Returns:
        Recipe: Represents a recipe from the given URL in string format.
    """
    return scrape_manager.get_recipe(url_in.link)

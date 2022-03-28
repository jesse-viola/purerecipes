from fastapi import APIRouter

from .schemas import URL
from server.managers import scrape_manager

router = APIRouter()


@router.post("/recipe")
async def get_recipe(url_in: URL):
    scrape_manager.get_recipe(url_in.link)

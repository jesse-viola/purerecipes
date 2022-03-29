from fastapi import APIRouter

from server.scrape.routes import router as scrape_router

api_router = APIRouter(prefix="/api")
api_router.include_router(scrape_router, prefix="/scrape")

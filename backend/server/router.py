from fastapi import APIRouter

from server.url.routes import router as url_router

api_router = APIRouter()
api_router.include_router(url_router, prefix="/url")

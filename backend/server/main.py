from fastapi import FastAPI

from .router import api_router as v1_router

app = FastAPI()

app.include_router(v1_router)

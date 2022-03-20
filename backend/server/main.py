from fastapi import FastAPI
import uvicorn

from server.router import api_router as v1_router

app = FastAPI()

app.include_router(v1_router)

# For debugging
if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)

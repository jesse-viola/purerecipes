from pydantic import BaseModel


class URL(BaseModel):
    link: str

class Recipe(BaseModel):
    pass

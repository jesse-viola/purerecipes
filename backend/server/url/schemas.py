from pydantic import BaseModel


class URL(BaseModel):
    link: str

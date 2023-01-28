from fastapi import FastAPI
from pydantic import BaseSettings


class Settings(BaseSettings):
    steam_key: str = "55F10D9E9AC8E75027FDE67170A39A78"
    steam_public_api: str = "http://store.steampowered.com/api"
    steam_private_api: str = "https://api.steampowered.com"
    my_user: int = 76561198139366501


settings = Settings()

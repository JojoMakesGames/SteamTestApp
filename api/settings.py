import configparser
from fastapi import FastAPI
from pydantic import BaseSettings


class Settings(BaseSettings):
    config = configparser.ConfigParser()  
    config.read('api/local.config.ini')
    steam_key: str = config.get('API', 'STEAM_API_TOKEN')
    steam_public_api: str = "http://store.steampowered.com/api"
    steam_private_api: str = "https://api.steampowered.com"
    neo4j_uri: str = config.get('DB', 'NEO4J_URI')
    my_user: int = 76561198045509954


settings = Settings()

import requests
from typing import List
from api.settings import settings

from api.types.steam_types import Game

def get_owned_games(steam_id: int) -> List[Game]:
        """
        Get a list of games owned by a user
        """
        url = f"{settings.steam_private_api}/IPlayerService/GetOwnedGames/v1/"
        params = {
            "key": settings.steam_key,
            "steamid": steam_id,
            "include_appinfo": 1,
            "format": "json",
        }
        response = requests.get(url, params=params)
        response.raise_for_status()
        response = response.json()
        return sorted([Game(play_time=game["playtime_forever"], 
            name=game["name"]) for game in response["response"]["games"]],
            key=lambda game: game.name)
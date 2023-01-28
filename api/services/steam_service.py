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
        return response['response']['games']
        return sorted([Game(_play_time=game["playtime_forever"], 
            name=game["name"]) for game in response["response"]["games"]],
            key=lambda game: game.name)



def get_game_details(app_id: int) -> List[Game]:
        """
        Get a list of games owned by a user
        """
        url = f"{settings.steam_public_api}/appdetails/"
        params = { 
            "appids": app_id,
            "format": "json",
        }
        response = requests.get(url, params=params)
        response.raise_for_status()
        response = response.json()
        return response

def get_steam_username(steam_id: int) -> str:
        """
        Get a list of games owned by a user
        """
        url = f"{settings.steam_private_api}/ISteamUser/GetPlayerSummaries/v2/"
        params = {
            "key": settings.steam_key,
            "steamids": steam_id,
            "format": "json",
        }
        response = requests.get(url, params=params)
        response.raise_for_status()
        response = response.json()
        return response['response']['players'][0]['personaname']
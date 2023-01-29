
from typing import List
from api.database.games_dao import get_all_owned_games
from api.services.steam_service import get_owned_games, get_steam_username
from api.types.steam_types import Game


def get_user_games(steam_id: float) -> List[Game]:
    """
    Get a list of games owned by a user
    """
    username = get_steam_username(steam_id)
    games = get_all_owned_games(username)
    # if not games:
    #     return get_owned_games(steam_id)
    return games

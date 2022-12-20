import strawberry
from fastapi import Depends
from strawberry.fastapi import GraphQLRouter
from strawberry.types import Info
from typing import List

from api.database.session import get_session
from api.types.steam_types import Game
from api.services.steam_service import get_owned_games
  

@strawberry.type
class Query:
    @strawberry.field
    async def Games(self, info: Info) -> List[Game]:
        return get_owned_games(steam_id=76561198045509954)



async def get_context(
    session=Depends(get_session),
):
    return {
        "session": session,
    }

schema = strawberry.Schema(query=Query)

graphql_app = GraphQLRouter(schema, context_getter=get_context)
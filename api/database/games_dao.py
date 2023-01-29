from typing import List
from api.database.session import get_graph_session
from api.types.steam_types import Game

from neo4j import GraphDatabase


def get_all_owned_games(username: str) -> List[Game]:
        """
        Get a list of games owned by a user
        """
        uri = "bolt://localhost:7687"
        driver = GraphDatabase.driver(uri)
        print(username)
        with driver.session() as session:
            query = """
            MATCH (u:User)-[r:OWNS]->(g:Game)
            WHERE u.name = $username
            RETURN g
            """
            result = session.run(query, username=username)
            games = result.data()
        return [Game(name=game['g']['name']) for game in games]

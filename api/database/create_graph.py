import json
from typing import Any, Dict, List, Set
from api.services.steam_service import get_game_details, get_owned_games, get_steam_username
from api.settings import settings
from neo4j import GraphDatabase
import time
import re

uri             = "bolt://localhost:7687"
userName        = "neo4j"
password        = "password"
#region methods
def clean_string(string: str):
    return re.sub('[^a-zA-Z1-9]+', '', string)

def add_spaces(string: str):
    return re.sub('([a-z])([A-Z])', r'\1 \2', string)

def write_file(file_name: str, game_details: List[Dict[str, Any]]):
    with open(file_name, "w") as file:
        for dete in game_details:
            file.write(f"{dete['name']};, {dete['release_date']};, {dete['publisher']};, {dete['developer']};, {dete['genre']}\n")

def read_file(file_name: str) -> List[Dict[str, Any]]:
    game_details = list()
    hasAdded = set()
    try:
        with open(file_name, "r") as file:
            while True:
                line = file.readline().split(";, ")
                if not line or line == ['']:
                    break
                if line[0] in hasAdded:
                    continue
                detail = {
                    'name': line[0],
                    'release_date': line[1],
                    'publisher': eval(line[2]),
                    'developer': eval(line[3]),
                    'genre': eval(line[4])
                }
                hasAdded.add(line[0])
                game_details.append(detail)
    except Exception as e:
        print(e)
           
    return game_details

def fetch_games():
    game_details = []
    games = get_owned_games(settings.my_user)
    for i, game in enumerate(games):
        print(i, game['appid'])
        details = get_game_details(game['appid'])
        game_details.append((game['appid'],details))
        time.sleep(1.5)
    return game_details

def fetch_game_details(games: List[Dict[str, Any]])-> List[Dict[str, Any]]:
    detes = []
    publishers = set()
    developers = set()
    genres = set()
    for id, detail in games:
        if not detail[str(id)]['success']:
            continue
        developer = detail[str(id)]['data'].get('developers', None)
        publisher = detail[str(id)]['data'].get('publishers', None)
        genre = detail[str(id)]['data'].get('genres', None)
        if developer:
            for dev in developer:
                developers.add(dev)
        if publisher:
            for pub in publisher:
                publishers.add(pub)
        if genre:
            for gen in genre:
                genres.add(gen['description'])
        dete = {
            'name': detail[str(id)]['data']['name'],
            'publisher': publisher,
            'developer': developer,
            'release_date': detail[str(id)]['data']['release_date']['date'],
            'genre': genre
        }
        detes.append(dete)
    return detes

def build_create_query(game_details: List[Dict[str, Any]], genres: Set, publishers: Set, developers: Set, steam_id: int) -> str:
    username = get_steam_username(steam_id)
    create_query = """CREATE"""
    create_query += f"""
    (u{username}user:User {{name: "{username}"}}),
    """
    for pub in publishers.union(developers):
        create_query += f"""
        (c{clean_string(pub)}company:Company {{name: "{add_spaces(pub)}"}}),
        """
    for i, gen in enumerate(genres):
        create_query += f"""
        ({clean_string(gen)}genre:Genre {{name: "{add_spaces(gen)}"}}),
        """
    detail_dict = {dete['name']: f"game{i}" for i, dete in enumerate(game_details)}
    for i, dete in enumerate(game_details):
        create_query += f"""
        ({detail_dict[dete['name']]}:Game {{name: "{dete['name'].replace('"', '')}", release_date: "{dete['release_date']}"}}),
        """
        if username:
            create_query += f"""
            ({detail_dict[dete['name']]})-[:OWNED_BY]->(u{username}user),
            """
        if dete['publisher']:
            for pub in dete['publisher']:
                create_query += f"""
                ({detail_dict[dete['name']]})-[:PUBLISHED_BY]->(c{clean_string(pub)}company),
                """
        if dete['developer']:
            for dev in dete['developer']:
                create_query += f"""
                ({detail_dict[dete['name']]})-[:DEVELOPED_BY]->(c{clean_string(dev)}company),
                """
        if dete['genre']:
            for gen in dete['genre']:
                create_query += f"""
                ({detail_dict[dete['name']]})-[:GAME_TYPE]->({clean_string(gen['description'])}genre),
                """
    create_query = create_query.strip()[:-1]
    return create_query
#endregion

publishers = set()
developers = set()
genres = set()
details = read_file("games.txt")
if not details:
    games = fetch_games()
    details = fetch_game_details(games)
    write_file("games.txt", details)

for dete in details:
    if dete['publisher']:
        for pub in dete['publisher']:
            publishers.add(clean_string(pub))
    if dete['developer']:
        for dev in dete['developer']:
            developers.add(clean_string(dev))
    if dete['genre']:
        for gen in dete['genre']:
            genres.add(clean_string(gen['description']))
create_query = build_create_query(details, genres, publishers, developers, settings.my_user)

driver = GraphDatabase.driver(uri, auth=(userName, password))
with driver.session(database="neo4j") as session:
    session.run(create_query)


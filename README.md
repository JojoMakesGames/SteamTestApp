# SteamTestApp
Test app using steam integration

Run docker command:
docker run \
    --publish=7474:7474 --publish=7687:7687 \
    --volume=$HOME/neo4j/data:/data \
    --env=NEO4J_AUTH=none neo4j

set 
run python ./api/database/create_graph.py

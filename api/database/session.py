from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession

from contextlib import asynccontextmanager, contextmanager

from configparser import ConfigParser
from api.settings import settings
from neo4j import GraphDatabase


@asynccontextmanager
async def get_session() -> AsyncSession:
    config = ConfigParser()
    config.read('local.ini')
    db_uri = config.get('connection', 'db_uri')

    engine = create_async_engine(db_uri, future=True, echo=True)
    session_maker = sessionmaker(bind=engine, expire_on_commit=False, class_=AsyncSession)
    async with session_maker() as session:
        try:
            yield session
        except Exception as e:
            await session.rollback()
            raise e

def get_graph_session():
    uri = settings.neo4j_uri
    driver = GraphDatabase.driver(uri)
    return driver.session()
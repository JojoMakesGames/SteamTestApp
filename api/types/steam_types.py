import strawberry

@strawberry.type
class Game:
    play_time: int
    name: str
    
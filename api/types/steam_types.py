import strawberry

@strawberry.type
class Game:
    _play_time: int
    name: str

    @strawberry.field
    def play_time(self, info) -> str:
        return f"{self._play_time} minutes"
    
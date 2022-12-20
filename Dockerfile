FROM python:3.10

WORKDIR /app

COPY poetry.lock pyproject.toml ./

RUN pip install poetry
RUN poetry install

COPY ./api /app/api


CMD ["poetry", "run", "app"]
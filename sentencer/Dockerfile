FROM python:3-alpine

WORKDIR /app

ENV POETRY_HOME=/opt/poetry
ENV PATH=${POETRY_HOME}/bin:${PATH}

RUN apk add curl gcc linux-headers \
    && curl -sSL https://install.python-poetry.org | python3 - \
    && poetry config virtualenvs.create false

COPY pyproject.toml poetry.lock /app/

RUN poetry install -qn --only main --no-root

COPY . /app/

RUN poetry install --only-root

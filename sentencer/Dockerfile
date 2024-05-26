FROM python:3-alpine

WORKDIR /app

ENV POETRY_HOME=/opt/poetry
ENV PATH=${POETRY_HOME}/bin:${PATH}

RUN apk add curl gcc python3-dev musl-dev linux-headers make g++ \
    && curl -sSL https://install.python-poetry.org | python3 - \
    && poetry config virtualenvs.create false

COPY pyproject.toml poetry.lock Makefile /app/

ARG DEBUG=""

RUN if [ -z "${DEBUG}" ]; then \
    make install_deps_only \
    ; else \
    make install_deps_full \
    ; fi

COPY sentencer ./sentencer
COPY README.md .

RUN make install_project && make post_install

RUN apk del curl gcc linux-headers make

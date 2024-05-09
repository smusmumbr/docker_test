FROM python:3-alpine

WORKDIR /app

ENV POETRY_HOME=/opt/poetry
ENV PATH=${POETRY_HOME}/bin:${PATH}

RUN apk add curl gcc linux-headers make \
    && curl -sSL https://install.python-poetry.org | python3 - \
    && poetry config virtualenvs.create false

COPY pyproject.toml poetry.lock Makefile /app/

RUN make install_deps_only

COPY . /app/

RUN make install_project && make post_install

RUN apk del curl gcc linux-headers make

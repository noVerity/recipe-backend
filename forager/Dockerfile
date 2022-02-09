FROM python:3.9-slim-buster

WORKDIR /app

RUN pip install 'poetry==1.1.0'

COPY pyproject.toml pyproject.toml
COPY poetry.lock poetry.lock 

RUN poetry install

COPY ./ ./

EXPOSE 8080

CMD [ "poetry", "run", "gunicorn", "forager:create_app()", "--bind=0.0.0.0:8080" ]
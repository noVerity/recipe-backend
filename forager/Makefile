start:
	FLASK_ENV=development poetry run gunicorn forager:app --bind=0.0.0.0:8000

test:
	poetry run pytest

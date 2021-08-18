start:
	APP_TOKEN="TEST" \
	APP_OUT_QUEUE="ingredients_results" \
	APP_IN_QUEUE="ingredients_lookup" \
	poetry run gunicorn "forager:create_app()" --bind=0.0.0.0:8000

develop:
	FLASK_ENV=development \
	FLASK_APP="forager:create_app()" \
	APP_TOKEN="TEST" \
	APP_OUT_QUEUE="ingredients_results" \
	APP_IN_QUEUE="ingredients_lookup" \
	poetry run flask run --host=0.0.0.0 --port=8000

test:
	poetry run pytest

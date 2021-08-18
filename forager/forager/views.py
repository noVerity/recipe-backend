from flask import Blueprint, request
index = Blueprint('index', __name__)
from forager.fooddata.lookup import lookup
import os


@index.route("/lookup")
def Echoer():
    food_token = os.environ.get("APP_FOODDATA_TOKEN")
    if not food_token:
        return { "error": "no food token available" }

    api_token = os.environ.get("APP_TOKEN")
    if not api_token:
        return { "error": "no connection token available" }
    
    if request.headers.get("Authorization") != "Bearer " + api_token:
        return { "error": "invalid token" }

    item = request.args.get('item')
    if not isinstance(item, str):
        return { "error": "item needs to be a string" }
    
    try:
        return lookup(item)
    except IOError:
        return { "error": "could not retrieve url" }
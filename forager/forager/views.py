from flask import Blueprint, request, url_for, current_app
index = Blueprint('index', __name__)
from forager.fooddata.lookup import lookup
import os


def has_no_empty_params(rule):
    defaults = rule.defaults if rule.defaults is not None else ()
    arguments = rule.arguments if rule.arguments is not None else ()
    return len(defaults) >= len(arguments)

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


@index.route("/site-map")
def site_map():
    links = []
    for rule in current_app.url_map.iter_rules():
        # Filter out rules we can't navigate to in a browser
        # and rules that require parameters
        if "GET" in rule.methods and has_no_empty_params(rule):
            url = url_for(rule.endpoint, **(rule.defaults or {}))
            links.append((url, rule.endpoint))
    # links is now a list of url, endpoint tuples
    return { "links": links }
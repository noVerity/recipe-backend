import requests
import json
from urllib.parse import urlencode
from forager.fooddata.convert import convert
import os

def lookup(item: str):
    with requests.get(
        "https://api.nal.usda.gov/fdc/v1/foods/search?" + urlencode({"query": item}),
        headers={"X-Api-Key": os.environ.get("APP_FOODDATA_TOKEN")}
        ) as response:
        return convert(json.loads(response.text), item)
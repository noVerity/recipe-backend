def convert(food_data, search_term: str):
    if not food_data:
        raise Exception('nothing passed')
    
    if not 'foods' in food_data.keys():
        print(food_data)
        raise Exception('no result was found for the search term')

    if len(food_data.get("foods")) < 1:
        print(food_data)
        raise Exception('no food was found for the search term')

    found_food_item = food_data.get("foods")[0]
    nutrient_list = found_food_item.get("foodNutrients")

    return {
        "name": found_food_item.get("description"),
        "searchTerm": search_term,
        "calories": fetch_nutrient(nutrient_list, "energy"),
        "protein": fetch_nutrient(nutrient_list, "protein"),
        "fat": fetch_nutrient(nutrient_list, "Fatty acids, total polyunsaturated"),
        "carbohydrates": fetch_nutrient(nutrient_list, "Carbohydrate, by difference"),
    }

def fetch_nutrient(nutrient_list, name):
    result = list(filter(lambda x: x.get("nutrientName").lower() == name.lower(), nutrient_list))
    if len(result) <= 0:
        return 0
    return result[0].get("value")
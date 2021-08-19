import unittest

from forager.fooddata.convert import convert, fetch_nutrient

class ConvertTestCase(unittest.TestCase):
    def test_convert_invalid_inputs(self):
        self.assertRaises(Exception, convert, None, "Banana") # Nothing to convert
        self.assertRaises(Exception, convert, {}, "Banana") # The result did not contain a food item
        self.assertRaises(Exception, convert, { "foods": [] }, "Banana") # The result did not contain a food item

    def test_convert_valid_input(self):
        result = convert({
            "foods": [
                {
                    "description": "Banana, raw",
                    "foodNutrients": [
                        { "nutrientName": "Energy", "value": 1 },
                        { "nutrientName": "Protein", "value": 2 },
                        { "nutrientName": "Fatty acids, total polyunsaturated", "value": 3 },
                        { "nutrientName": "Carbohydrate, by difference", "value": 4 },
                    ]
                }
            ]
        }, "Banana")

        self.assertDictEqual(result, {
            "name": "Banana, raw",
            "searchTerm": "Banana",
            "calories": 1,
            "protein": 2,
            "fat": 3,
            "carbohydrates": 4,
        })

    def test_fetch_nutrient_invalid_inputs(self):
        assert fetch_nutrient([], "Banana") == 0

    def test_fetch_nutrient_valid(self):
        assert fetch_nutrient([
            { "nutrientName": "Banana", "value": 3.124 }
        ], "Banana") ==  3.124
import unittest
from unittest.mock import patch
from json import loads
from forager import __version__, create_app

class ForagerTestCase(unittest.TestCase):

    def test_version(self):
        assert __version__ == '0.1.0'
    
    @patch('forager.is_running_from_reloader', lambda: True)
    def test_available_routes(self):
        with create_app().test_client() as c:
            response = c.get('/site-map')
            self.assertEqual(response.status_code, 200)
            result = loads(response.data)
            assert [
            "/lookup", 
            "index.Echoer"
            ] in result.get("links")
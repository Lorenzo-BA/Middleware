import json
import requests
from marshmallow import EXCLUDE

from src.schemas.user import UserSchema
from src.models.user import User as UserModel


ratings_url = "https://ratings-juliette.edu.forestier.re"  # URL de l'API ratings (golang) fournis.


def get_ratings_by_song_id(song_id):
    response = requests.request(method="GET", url=ratings_url+"/songs/"+song_id+"/ratings")
    return response.json(), response.status_code


def get_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    response = requests.request(method="GET", url=ratings_url+"/songs/"+song_id+"/ratings/"+rating_id)
    return response.json(), response.status_code


def add_ratings_with_song_id(song_id, rating):
    rating_model = UserModel.from_dict_with_clear_password(rating)
    rating_schema = UserSchema().loads(json.dumps(rating_model), unknown=EXCLUDE)
    response = requests.request(method="POST", url=ratings_url+"/songs/"+song_id+"/ratings", json=rating_schema)
    return response.json(), response.status_code


def delete_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    response = requests.request(method="DELETE", url=ratings_url+"/songs/"+song_id+"/ratings/"+rating_id)
    return "", response.status_code


def update_ratings_by_song_id_and_ratings_id(song_id, rating_id, rating):
    rating_model = UserModel.from_dict_with_clear_password(rating)
    rating_schema = UserSchema().loads(json.dumps(rating_model), unknown=EXCLUDE)
    response = requests.request(method="PUT", url=ratings_url+"/songs/"+song_id+"/ratings/"+rating_id, json=rating_schema)
    return response.json(), response.status_code
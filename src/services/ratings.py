import json
import requests
from marshmallow import EXCLUDE

from src.schemas.rating import RatingSchema
from src.models.user import User as UserModel


ratings_url = "http://localhost:8081"  # URL de l'API ratings (golang).


def get_ratings_by_song_id(song_id):
    response = requests.request(method="GET", url=ratings_url+"/songs/"+song_id+"/ratings")
    return response.json(), response.status_code


def get_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    response = requests.request(method="GET", url=ratings_url+"/songs/"+song_id+"/ratings/"+rating_id)
    return response.json(), response.status_code


def add_ratings_with_song_id(song_id, rating, user_id):
    rating_schema = RatingSchema().loads(json.dumps(rating), unknown=EXCLUDE)
    rating_schema["user_id"] = user_id
    response = requests.request(method="POST", url=ratings_url+"/songs/"+song_id+"/ratings", json=rating_schema)
    return response.json(), response.status_code


def delete_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    response = requests.request(method="DELETE", url=ratings_url+"/songs/"+song_id+"/ratings/"+rating_id)
    return "", response.status_code


def update_ratings_by_song_id_and_ratings_id(song_id, rating_id, rating):
    rating_schema = RatingSchema().loads(json.dumps(rating), unknown=EXCLUDE)
    response = requests.request(method="PUT", url=ratings_url+"/songs/"+song_id+"/ratings/"+rating_id, json=rating_schema)
    return response.json(), response.status_code
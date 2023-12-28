import json
import requests
from marshmallow import EXCLUDE

from src.schemas.rating import RatingSchema
from src.models.http_exceptions import *


ratings_url = "http://localhost:8081/songs/"  # URL de l'API ratings (golang).


def get_ratings_by_song_id(song_id):
    response = requests.get(ratings_url+song_id+"/ratings")
    return response.json(), response.status_code


def get_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    response = requests.get(ratings_url+song_id+"/ratings/"+rating_id)
    return response.json(), response.status_code


def add_ratings_with_song_id(song_id, rating, user_id):
    rating["user_id"] = user_id
    response = requests.post(ratings_url+song_id+"/ratings", json=rating)
    return response.json(), response.status_code


def delete_ratings_by_song_id_and_ratings_id(song_id, rating_id, current_user_id):
    song_data, song_status = get_ratings_by_song_id_and_ratings_id(song_id, rating_id)
    if song_status != 200:
        return song_data, song_status
    if song_data["user_id"] != current_user_id:
        raise Forbidden
    response = requests.request(method="DELETE", url=ratings_url+song_id+"/ratings/"+rating_id)
    return "", response.status_code


def update_ratings_by_song_id_and_ratings_id(song_id, rating_id, rating, current_user_id):
    song_data, song_status = get_ratings_by_song_id_and_ratings_id(song_id, rating_id)
    if song_status != 200:
        return song_data, song_status
    if song_data["user_id"] != current_user_id:
        raise Forbidden
    response = requests.request(method="PUT", url=ratings_url+song_id+"/ratings/"+rating_id, json=rating)
    return response.json(), response.status_code

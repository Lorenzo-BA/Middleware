import requests
from flask_login import current_user

from src.models.http_exceptions import *


RATINGS_URL = "http://localhost:8081/songs/"  # URL de l'API ratings (golang).


def get_ratings_by_song_id(song_id):
    response = requests.get(RATINGS_URL + song_id + "/ratings")
    return response.json(), response.status_code


def get_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    response = requests.get(RATINGS_URL + song_id + "/ratings/" + rating_id)
    return response.json(), response.status_code


def add_ratings_with_song_id(song_id, rating):
    rating["user_id"] = current_user.id
    response = requests.post(RATINGS_URL + song_id + "/ratings", json=rating)
    return response.json(), response.status_code


def delete_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    song_data, song_status = get_ratings_by_song_id_and_ratings_id(song_id, rating_id)
    if song_status != 200:
        return song_data, song_status
    if song_data["user_id"] != current_user.id:
        raise Forbidden
    response = requests.delete(RATINGS_URL + song_id + "/ratings/" + rating_id)
    return "", response.status_code


def update_ratings_by_song_id_and_ratings_id(song_id, rating_id, rating):
    song_data, song_status = get_ratings_by_song_id_and_ratings_id(song_id, rating_id)
    if song_status != 200:
        return song_data, song_status
    if song_data["user_id"] != current_user.id:
        raise Forbidden
    response = requests.put(RATINGS_URL + song_id + "/ratings/" + rating_id, json=rating)
    return response.json(), response.status_code

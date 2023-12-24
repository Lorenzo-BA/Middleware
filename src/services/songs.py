import json
import requests
from marshmallow import EXCLUDE

from src.schemas.song import *


songs_url = "http://localhost:8079/songs/"  # URL de l'API songs (golang).


def get_songs():
    songs_response = requests.request(method="GET", url=songs_url)
    songs_data = songs_response.json()

    songs_with_ratings = []

    for song_data in songs_data:
        song_id = song_data["id"]

        ratings_response = requests.get("https://ratings-juliette.edu.forestier.re/songs/" + song_id + "/ratings")
        ratings_data = ratings_response.json()

        song_schema = SongWithRatingSchema().load(song_data, unknown=EXCLUDE)
        ratings_schema = RatingSchema(many=True).load(ratings_data, unknown=EXCLUDE)

        song_schema["ratings"] = ratings_schema

        songs_with_ratings.append(song_schema)

    return songs_with_ratings, songs_response.status_code


def get_song(song_id):
    response = requests.request(method="GET", url=songs_url+song_id)
    return response.json(), response.status_code


def add_songs(song):
    song_schema = SongSchema().loads(json.dumps(song), unknown=EXCLUDE)
    response = requests.request(method="POST", url=songs_url, json=song_schema)
    return response.json(), response.status_code


def delete_song(song_id):
    response = requests.request(method="DELETE", url=songs_url+song_id)
    return "", response.status_code


def update_song(song_id, song):
    song_schema = SongSchema().loads(json.dumps(song), unknown=EXCLUDE)
    response = requests.request(method="PUT", url=songs_url+song_id, json=song_schema)
    return response.json(), response.status_code
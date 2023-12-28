import json
import requests
from marshmallow import EXCLUDE

from src.schemas.song import *


songs_url = "http://localhost:8079/songs/"  # URL de l'API songs (golang).
ratings_url = "http://localhost:8081"  # URL de l'API ratings (golang).


def get_songs():
    songs_response = requests.get(songs_url)
    songs_data = songs_response.json()

    songs_with_ratings = []

    for song_data in songs_data:
        song_id = song_data["id"]

        ratings_response = requests.get(ratings_url + "/songs/" + song_id + "/ratings")
        ratings_data = ratings_response.json()

        song_schema = SongWithRatingSchema().load(song_data, unknown=EXCLUDE)
        ratings_schema = RatingSchema(many=True).load(ratings_data, unknown=EXCLUDE)

        song_schema["ratings"] = ratings_schema

        songs_with_ratings.append(song_schema)

    return songs_with_ratings, songs_response.status_code


def get_song(song_id):
    song_response = requests.get(songs_url + song_id)
    if song_response.status_code != 200:
        return song_response.json(), song_response.status_code
    song_data = song_response.json()
    song_id = song_data["id"]
    ratings_response = requests.get(ratings_url + "/songs/" + song_id + "/ratings")
    ratings_data = ratings_response.json()
    song_schema = SongWithRatingSchema().load(song_data, unknown=EXCLUDE)
    ratings_schema = RatingSchema(many=True).load(ratings_data, unknown=EXCLUDE)
    song_schema["ratings"] = ratings_schema
    return song_schema, song_response.status_code


def create_song(song):
    song_schema = SongWithRatingSchema().loads(json.dumps(song), unknown=EXCLUDE)
    response = requests.post(songs_url, json=song_schema)

    if response.status_code == 201:
        created_song_schema = SongWithRatingSchema().load(response.json(), unknown=EXCLUDE)
        created_song_schema["ratings"] = []
        return created_song_schema, response.status_code
    else:
        return response.json(), response.status_code


def delete_song(song_id):
    response = requests.delete(songs_url + song_id)
    if response.status_code != 204:
        return response.json(), response.status_code
    return "", response.status_code


def update_song(song_id, song):
    song_schema = SongWithRatingSchema().loads(json.dumps(song), unknown=EXCLUDE)
    response = requests.put(songs_url + song_id, json=song_schema)

    if response.status_code == 200:
        created_song_schema = SongWithRatingSchema().load(response.json(), unknown=EXCLUDE)
        created_song_schema["ratings"] = []
        return created_song_schema, response.status_code
    else:
        return response.json(), response.status_code

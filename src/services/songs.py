import requests
from marshmallow import EXCLUDE

from src.schemas.song import *
from src.services.ratings import get_ratings_by_song_id


songs_url = "http://localhost:8079/songs/"  # URL de l'API songs (golang).


def get_songs():
    songs_response = requests.get(songs_url)
    if songs_response.status_code != 200:
        return songs_response.json(), songs_response.status_code

    songs_data = songs_response.json()
    songs_with_ratings = []
    for song_data in songs_data:
        song_id = song_data["id"]
        ratings_data, ratings_status = get_ratings_by_song_id(song_id)
        if ratings_status != 200:
            return ratings_data, ratings_status

        song_with_ratings = {
            **SongSchema().load(song_data, unknown=EXCLUDE),
            "ratings": ratings_data,
        }
        songs_with_ratings.append(song_with_ratings)

    return songs_with_ratings, songs_response.status_code


def get_song(song_id):
    song_response = requests.get(songs_url + song_id)
    if song_response.status_code != 200:
        return song_response.json(), song_response.status_code

    song_data = song_response.json()
    song_id = song_data["id"]
    ratings_data, ratings_status = get_ratings_by_song_id(song_id)
    if ratings_status != 200:
        return ratings_data, ratings_status

    return {
        **SongSchema().load(song_data, unknown=EXCLUDE),
        "ratings": ratings_data,
    }, song_response.status_code


def create_song(song):
    response = requests.post(songs_url, json=song)
    if response.status_code != 201:
        return response.json(), response.status_code
    created_song = response.json()
    created_song["ratings"] = []
    return created_song, response.status_code


def delete_song(song_id):
    response = requests.delete(songs_url + song_id)
    if response.status_code != 204:
        return response.json(), response.status_code
    return "", response.status_code


def update_song(song_id, song):
    response = requests.put(songs_url + song_id, json=song)
    if response.status_code != 200:
        return response.json(), response.status_code

    song_response, song_status = get_song(song_id)
    if song_status != 200:
        return song_response, song_status
    return song_response, response.status_code


def song_exists(song_id):
    response = requests.get(songs_url + song_id)
    return response.status_code == 200

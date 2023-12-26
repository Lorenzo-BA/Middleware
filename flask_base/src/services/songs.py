import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.user import UserSchema
from src.schemas.song import *
from src.models.user import User as UserModel
from src.models.http_exceptions import *
import src.repositories.users as users_repository



SONGS_API_URL = "http://localhost:8089/songs/"  # URL de l'API songs (golang)

def get_all_songs():
    response = requests.get(SONGS_API_URL)
    response.raise_for_status()
    return response.json(), response.status_code

def get_song(song_id):
    response = requests.get(SONGS_API_URL + song_id)
    response.raise_for_status()
    return response.json(), response.status_code


def create_song(song):

    # on récupère le schéma utilisateur pour la requête vers l'API users
    # on crée l'utilisateur côté API users

    song_schema = SongSchema().loads(json.dumps(song, default=str), unknown=EXCLUDE)
    print(song_schema)

    response = requests.post(SONGS_API_URL, json=song)


    return response.json(), response.status_code


def modify_song(song_id, song_update):
    # on vérifie que l'utilisateur se modifie lui-même
    if song_id != current_song.id:
        raise Forbidden

    # s'il y a quelque chose à changer côté API (username, name)
    song_schema = SongSchema().loads(json.dumps(song_update), unknown=EXCLUDE)
    response = None
    if not SongSchema.is_empty(song_schema):
        # on lance la requête de modification
        response = requests.put(SONGS_API_URL + song_id, json=song_schema)
        response.raise_for_status()

        if response.status_code != 200:
            return response.json(), response.status_code

    # s'il y a quelque chose à changer côté BDD
    song_model = SongModel.from_dict_with_clear_password(song_update)
    if not song_model.is_empty():
        song_model.id = song_id
        found_song = songs_repository.get_song_from_id(song_id)

        if not song_model.title:
            song_model.title = found_song.title


        try:
            songs_repository.update_song(song_model)
        except exc.IntegrityError as e:
            if "NOT NULL" in e.orig.args[0]:
                raise UnprocessableEntity
            raise Conflict

    return (response.json(), response.status_code) if response else get_song(song_id)


def delete_song(song_id):
    response = requests.delete(SONGS_API_URL + song_id)
    response.raise_for_status()

    return "", response.status_code

def get_song_from_db(songname):
    return songs_repository.get_song(songname)


def song_exists(songTitle):
    return get_song_from_db(songTitle) is not None

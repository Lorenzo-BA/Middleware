import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.user import UserSchema
from src.models.user import User as UserModel
from src.models.http_exceptions import *
import src.repositories.users as users_repository


USERS_URL = "http://localhost:8080/users/"  # URL de l'API users (golang)


def get_users():
    response = requests.get(USERS_URL)
    return response.json(), response.status_code


def get_user(user_id):
    response = requests.get(USERS_URL + user_id)
    return response.json(), response.status_code


def create_user(user_register):
    # on récupère le modèle utilisateur pour la BDD
    user_model = UserModel.from_dict_with_clear_password(user_register)
    # on récupère le schéma utilisateur pour la requête vers l'API users
    user_schema = UserSchema().loads(json.dumps(user_register), unknown=EXCLUDE)

    # on crée l'utilisateur côté API users
    response = requests.request(method="POST", url=USERS_URL, json=user_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    # on ajoute l'utilisateur dans la base de données
    # pour que les données entre API et BDD correspondent
    try:
        user_model.id = response.json()["id"]
        users_repository.add_user(user_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code


def update_user(user_id, user):
    if user_id != current_user.id:
        raise Forbidden

    user_model = UserModel.from_dict_with_clear_password(user)
    user_model.id = user_id
    found_user = users_repository.get_user_from_id(user_id)
    if not user_model.username:
        user_model.username = found_user.username
        user["username"] = found_user.username
    if not user_model.encrypted_password:
        user_model.encrypted_password = found_user.encrypted_password

    if not ("name" in user and user["name"] is not None):
        user_data, status_code = get_user(user_id)
        if status_code != 200:
            return user_data, status_code
        user["name"] = user_data["name"]

    response = requests.put(USERS_URL + user_id, json=user)
    if response.status_code != 200:
        return response.json(), response.status_code

    try:
        users_repository.update_user(user_model)
    except exc.IntegrityError as e:
        if "NOT NULL" in e.orig.args[0]:
            raise UnprocessableEntity
        raise Conflict

    return response.json(), response.status_code


def delete_user(user_id):
    if user_id != current_user.id:
        raise Forbidden

    response = requests.delete(USERS_URL + user_id)
    if response.status_code != 204:
        return response.json(), response.status_code

    try:
        users_repository.delete_user(user_id)
    except Exception:
        raise SomethingWentWrong

    return "", response.status_code


def get_user_from_db(username):
    return users_repository.get_user(username)


def user_exists(username):
    return get_user_from_db(username) is not None



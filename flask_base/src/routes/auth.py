import json
from flask import Blueprint, request, json
from marshmallow import ValidationError
from flask_login import login_user, logout_user, login_required, current_user

from src.models.http_exceptions import *
from src.schemas.errors import *
from src.schemas.song import *
from src.schemas.user_auth import *
import src.services.users as users_service
import src.services.songs as songs_service
import src.services.auth as auth_service


auth = Blueprint(name="login", import_name=__name__)


@auth.route('/login', methods=['POST'])
def login():
    """
    ---
    post:
      description: Login
      requestBody:
        required: true
        content:
            application/json:
                schema: UserLogin
      responses:
        '200':
          description: Ok
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Already logged in
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - auth
          - users
    """
    if current_user.is_authenticated:
        error = ForbiddenSchema().loads(json.dumps({"message": "Already logged in"}))
        return error, error.get("code")

    # parser le body
    try:
        # it is possible to use marshmallow Schemas validation (used also for doc) or custom classes
        user_login = UserLoginSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # logger l'utilisateur
    try:
        user = auth_service.login(user_login)
    except (NotFound, Unauthorized):
        error = UnauthorizedSchema().loads("{}")
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

    login_user(user, remember=True)
    return "", 200


@auth.route('/logout', methods=['POST'])
@login_required
def logout():
    """
    ---
    post:
      description: Logout
      responses:
        '200':
          description: Ok
        '401':
          description: Unauthorized
      tags:
          - auth
          - users
    """
    logout_user()
    return "", 200


@auth.route('/register', methods=['POST'])
def register():
    """
    ---
    post:
      description: Register
      requestBody:
        required: true
        content:
            application/json:
                schema: UserRegister
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Already logged in
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '409':
          description: User already exists
          content:
            application/json:
              schema: Conflict
            application/yaml:
              schema: Conflict
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - auth
          - users
    """
    if current_user.is_authenticated:
        error = ForbiddenSchema().loads(json.dumps({"message": "Already logged in"}))
        return error, error.get("code")

    # parser le body
    try:
        user_register = UserRegisterSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # enregistrer l'utilisateur
    try:
        return auth_service.register(user_register)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "User already exists"}))
        return error, error.get("code")
    except SomethingWentWrong:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")


@auth.route('/introspect', methods=["GET"])
@login_required
def introspect():
    """
    ---
    get:
      description: Getting authenticated user
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
      tags:
          - auth
          - users
    """
    return users_service.get_user(current_user.id)

@auth.route('/users/', methods=["GET"])
def get_users():
    return users_service.get_users()


@auth.route('/users/<user_id>', methods=["GET"])
def get_user(user_id):
    return users_service.get_user(user_id)


@auth.route('/users/<user_id>', methods=["DELETE"])
def delete_user(user_id):
    return users_service.delete_user(user_id)


@auth.route('/users/<user_id>', methods=["PUT"])
def update_user(user_id):
    try:
        user_modified = UserRegisterSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")
    return users_service.modify_user(user_id, user_modified)


@auth.route('/songs/<song_id>/ratings', methods=["GET"])
def get_ratings_by_song_id(song_id):
    return ratings_service.get_ratings_by_song_id(song_id)


@auth.route('/songs/<song_id>/ratings', methods=["POST"])
def add_ratings_with_song_id(song_id):
    try:
        rating = RatingAddSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")
    return ratings_service.add_ratings_with_song_id(song_id, rating)


@auth.route('/songs/<song_id>/ratings/<rating_id>', methods=["DELETE"])
def delete_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    return ratings_service.delete_ratings_by_song_id_and_ratings_id(song_id, rating_id)


@auth.route('/songs/<song_id>/ratings/<rating_id>', methods=["GET"])
def get_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    return ratings_service.get_ratings_by_song_id_and_ratings_id(song_id, rating_id)


@auth.route('/songs/<song_id>/ratings/<rating_id>', methods=["PUT"])
def update_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    try:
        rating = RatingAddSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")
    return ratings_service.update_ratings_by_song_id_and_ratings_id(song_id, rating_id, rating)

####################################################################################
#____________________________________SONGS__________________________________________
####################################################################################

@auth.route('/songs/', methods=["GET"])
def get_all_songs():
    """
    ---
    get:
      description: Getting all song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized

      tags:
          - songs
    """
    return songs_service.get_all_songs()


@auth.route('/songs/<id>', methods=["GET"])
def get_single_song(id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.get_song(id)



@auth.route('/songs/<id>', methods=["DELETE"])
def delete_song(id):
    """
    ---
    get:
      description: Deleting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '204':
          description: No content
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    return songs_service.delete_song(id)


@auth.route('/songs/', methods=["POST"])
def create_song():
    """
    ---
    post:
      description: Creating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - songs
    """
    try:
        json_data = request.data.decode('utf-8')
        print(f"Received JSON data: {json_data}")
        song = BaseSongSchema().loads(json_data)
    except ValidationError as e:
        # Handle validation error
        print(f"Validation Error: {e}")
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")
    return songs_service.create_song(song)


@auth.route('/songs/<id>', methods=["PUT"])
def update_song(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: SongUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
    """
    try:
        song_modified = SongUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")
    return songs_service.modify_song(id, song_modified)
import json

from flask import Blueprint, request
from flask_login import login_user, logout_user, login_required, current_user
from json.decoder import JSONDecodeError

from src.models.http_exceptions import *
from src.schemas.errors import *
from src.schemas.user_auth import UserLoginSchema, UserRegisterSchema
from src.schemas.user import UserUpdateSchema
from src.schemas.rating import *
from src.schemas.song import *
import src.services.users as users_service
import src.services.auth as auth_service
import src.services.songs as songs_service
import src.services.ratings as ratings_service
from src.utils.content_negotiation import negotiate_content

auth = Blueprint(name="login", import_name=__name__)


####################################################################################
# ____________________________________AUTH________________________________________ #
####################################################################################


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
          description: OK
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
            application/yaml:
              schema:
                type: object
                properties:
                  message:
                    type: string
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
        return negotiate_content(error, error.get("code"))

    try:
        user_login = UserLoginSchema().loads(json_data=request.data.decode('utf-8'))
    except JSONDecodeError:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "Invalid JSON format"}))
        return negotiate_content(error, error.get("code"))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return negotiate_content(error, error.get("code"))

    try:
        user = auth_service.login(user_login)
    except (NotFound, Unauthorized):
        error = UnauthorizedSchema().loads("{}")
        return negotiate_content(error, error.get("code"))
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))

    login_user(user, remember=True)
    return negotiate_content({"message": "Authentication successful. Welcome, {0}!".format(user.username)}, 200)


@auth.route('/logout', methods=['POST'])
@login_required
def logout():
    """
    ---
    post:
      description: Logout
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                type: object
                properties:
                  message:
                    type: string
            application/yaml:
              schema:
                type: object
                properties:
                  message:
                    type: string
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
    logout_user()
    return negotiate_content({"message": "Logout successful. Goodbye !"}, 200)


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
            application/yaml:
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
        return negotiate_content(error, error.get("code"))

    try:
        user_register = UserRegisterSchema().loads(json_data=request.data.decode('utf-8'))
    except JSONDecodeError:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "Invalid JSON format"}))
        return negotiate_content(error, error.get("code"))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return negotiate_content(error, error.get("code"))

    try:
        response_data, status_code = auth_service.register(user_register)
        return negotiate_content(response_data, status_code)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "User already exists"}))
        return negotiate_content(error, error.get("code"))
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


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
    try:
        response_data, status_code = users_service.get_user(current_user.id)
        return negotiate_content(response_data, status_code)
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


####################################################################################
# ____________________________________USERS_______________________________________ #
####################################################################################


@auth.route('/users/', methods=["GET"])
@login_required
def get_users():
    """
    ---
    get:
      description: Getting all users
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
               type: array
               items:
                 $ref: "#/components/schemas/User"
            application/yaml:
              schema:
               type: array
               items:
                 $ref: "#/components/schemas/User"
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
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
    try:
        response_data, status_code = users_service.get_users()
        return negotiate_content(response_data, status_code)
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


@auth.route('/users/<user_id>', methods=["GET"])
@login_required
def get_user(user_id):
    """
    ---
    get:
      description: Getting a user
      parameters:
        - in: path
          name: user_id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
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
        '404':
          description: Not Found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable Entity
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
    try:
        response_data, status_code = users_service.get_user(user_id)
        return negotiate_content(response_data, status_code)
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


@auth.route('/users/<user_id>', methods=["DELETE"])
@login_required
def delete_user(user_id):
    """
    ---
    delete:
      description: Deleting oneself
      parameters:
        - in: path
          name: user_id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
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
    try:
        response_data, status_code = users_service.delete_user(user_id)
        return negotiate_content(response_data, status_code)
    except Forbidden as e:
        error = ForbiddenSchema().loads(json.dumps({"message": "Forbidden: Resource is locked."}))
        return negotiate_content(error, error.get("code"))
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


@auth.route('/users/<user_id>', methods=["PUT"])
@login_required
def update_user(user_id):
    """
    ---
    put:
      description: Updating oneself
      parameters:
        - in: path
          name: user_id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      requestBody:
        required: true
        content:
          application/json:
            schema: UserUpdateSchema
      responses:
        '200':
          description: OK
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
          description: Forbidden
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
          description: Unprocessable Entity
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
    try:
        user_modified = UserUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except JSONDecodeError:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "Invalid JSON format"}))
        return negotiate_content(error, error.get("code"))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return negotiate_content(error, error.get("code"))

    try:
        response_data, status_code = users_service.update_user(user_id, user_modified)
        return negotiate_content(response_data, status_code)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "User already exists"}))
        return negotiate_content(error, error.get("code"))
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Forbidden: Resource is locked."}))
        return negotiate_content(error, error.get("code"))
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


####################################################################################
# ____________________________________SONGS_______________________________________ #
####################################################################################


@auth.route('/songs/', methods=["GET"])
@login_required
def get_songs():
    """
    ---
    get:
      description: Getting all song
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
               type: array
               items:
                 $ref: "#/components/schemas/Song"
            application/yaml:
              schema:
               type: array
               items:
                 $ref: "#/components/schemas/Song"
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - songs
    """
    try:
        response_data, status_code = songs_service.get_songs()
        return negotiate_content(response_data, status_code)
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


@auth.route('/songs/<song_id>', methods=["GET"])
@login_required
def get_song(song_id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: song_id
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
        '422':
          description: Unprocessable Entity
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
          - songs
    """
    try:
        response_data, status_code = songs_service.get_song(song_id)
        return negotiate_content(response_data, status_code)
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


@auth.route('/songs/<song_id>', methods=["DELETE"])
@login_required
def delete_song(song_id):
    """
    ---
    delete:
      description: Deleting a song
      parameters:
        - in: path
          name: song_id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '204':
          description: No content
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - songs
    """
    try:
        response_data, status_code = songs_service.delete_song(song_id)
        return negotiate_content(response_data, status_code)
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


@auth.route('/songs/', methods=["POST"])
@login_required
def create_song():
    """
    ---
    post:
      description: Adding a song
      requestBody:
        required: true
        content:
            application/json:
                schema: SongAdding
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
        song = SongAddingSchema().loads(json_data=request.data.decode('utf-8'))
    except JSONDecodeError:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "Invalid JSON format"}))
        return negotiate_content(error, error.get("code"))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return negotiate_content(error, error.get("code"))

    try:
        response_data, status_code = songs_service.create_song(song)
        return negotiate_content(response_data, status_code)
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


@auth.route('/songs/<song_id>', methods=["PUT"])
@login_required
def update_song(song_id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: song_id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: SongAdding
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - songs
    """
    try:
        song = SongAddingSchema().loads(json_data=request.data.decode('utf-8'))
    except JSONDecodeError:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "Invalid JSON format"}))
        return negotiate_content(error, error.get("code"))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return negotiate_content(error, error.get("code"))

    try:
        response_data, status_code = songs_service.update_song(song_id, song)
        return negotiate_content(response_data, status_code)
    except Exception as e:
        error = SomethingWentWrongSchema().loads(json.dumps({"message": str(e)}))
        return negotiate_content(error, error.get("code"))


####################################################################################
# ____________________________________RATINGS_____________________________________ #
####################################################################################


@auth.route('/songs/<song_id>/ratings', methods=["GET"])
@login_required
def get_ratings_by_song_id(song_id):
    """
    ---
    get:
      description: Getting ratings of a song
      parameters:
        - in: path
          name: song_id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
               type: array
               items:
                 $ref: "#/components/schemas/Rating"
            application/yaml:
              schema:
               type: array
               items:
                 $ref: "#/components/schemas/Rating"
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - auth
          - songs
          - ratings
    """
    if not songs_service.song_exists(song_id):
        error = NotFoundSchema().loads(json.dumps({"message": "Not Found"}))
        return error, error.get("code")
    return ratings_service.get_ratings_by_song_id(song_id)


@auth.route('/songs/<song_id>/ratings', methods=["POST"])
@login_required
def add_ratings_with_song_id(song_id):
    """
    ---
    post:
      description: Adding a rating
      parameters:
        - in: path
          name: song_id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
          application/json:
            schema: RatingAdding
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - auth
          - songs
          - ratings
    """
    if not songs_service.song_exists(song_id):
        error = NotFoundSchema().loads(json.dumps({"message": "Not Found"}))
        return error, error.get("code")
    try:
        rating = RatingAddingSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")
    return ratings_service.add_ratings_with_song_id(song_id, rating, current_user.id)


@auth.route('/songs/<song_id>/ratings/<rating_id>', methods=["DELETE"])
@login_required
def delete_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    """
    ---
    delete:
      description: Deleting one's rating
      parameters:
        - in: path
          name: song_id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - auth
          - songs
          - ratings
    """
    if not songs_service.song_exists(song_id):
        error = NotFoundSchema().loads(json.dumps({"message": "Not Found"}))
        return error, error.get("code")
    try:
        return ratings_service.delete_ratings_by_song_id_and_ratings_id(song_id, rating_id, current_user.id)
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Forbidden: Resource is locked."}))
        return error, error.get("code")


@auth.route('/songs/<song_id>/ratings/<rating_id>', methods=["GET"])
@login_required
def get_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    """
    ---
    get:
      description: Getting a specific rating
      parameters:
        - in: path
          name: song_id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - auth
          - songs
          - ratings
    """
    if not songs_service.song_exists(song_id):
        error = NotFoundSchema().loads(json.dumps({"message": "Not Found"}))
        return error, error.get("code")
    return ratings_service.get_ratings_by_song_id_and_ratings_id(song_id, rating_id)


@auth.route('/songs/<song_id>/ratings/<rating_id>', methods=["PUT"])
@login_required
def update_ratings_by_song_id_and_ratings_id(song_id, rating_id):
    """
    ---
    put:
      description: Updating one's rating
      parameters:
        - in: path
          name: song_id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - auth
          - songs
          - ratings
    """
    if not songs_service.song_exists(song_id):
        error = NotFoundSchema().loads(json.dumps({"message": "Not Found"}))
        return error, error.get("code")
    try:
        rating = RatingAddingSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")
    try:
        return ratings_service.update_ratings_by_song_id_and_ratings_id(song_id, rating_id, rating, current_user.id)
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Forbidden: Resource is locked."}))
        return error, error.get("code")

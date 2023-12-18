from flask_login import SongMixin
from werkzeug.security import generate_password_hash
from src.helpers import db


# modèle de donnée pour la base de donnée utilisateur
# vous pouvez lier les utilisateurs de cette API à ceux de la vôtre (Golang) avec leur ID ou leur username
class Song(SongMixin, db.Model):
    __tablename__ = 'songs'

    id = db.Column(db.String(255), primary_key=True)
    songTitle = db.Column(db.String(255), unique=True, nullable=False)
    published_date = db.Column(db.String(255), nullable=False)

    def __init__(self, uuid, songTitle, published_date):
        self.id = uuid
        self.Title = songTitle
        self.published_date = published_date

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.songTitle or self.songTitle == "") and \
               (not self.published_date or self.published_date == "")

    @staticmethod
    def from_dict_with_clear_password(obj):
        songTitle = obj.get("songTitle") if obj.get("songTitle") != "" else None
        return Song(None, songTitle)

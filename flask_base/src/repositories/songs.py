from src.helpers import db
from src.models.song import Song


def get_song(songTitle):
    return db.session.query(Song).filter(Song.songTitle == songTitle).first()


def get_song_from_id(id):
    return Song.query.get(id)


def add_song(song):
    db.session.add(song)
    db.session.commit()


def update_song(song):
    existing_song = get_song_from_id(song.id)
    existing_song.songTitle = song.songTitle
    existing_song.published_date = song.published_date
    db.session.commit()


def delete_song(id):
    db.session.delete(get_song_from_id(id))
    db.session.commit()

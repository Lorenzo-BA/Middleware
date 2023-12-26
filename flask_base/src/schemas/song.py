from marshmallow import Schema, fields, validates_schema, ValidationError

# Schéma utilisateur de sortie (renvoyé au front)
class SongSchema(Schema):
    id = fields.String(description="UUID")
    content = fields.String(description="content")
    title = fields.String(description="title")
    file_name = fields.String(description="file_name")
    artist = fields.String(description="artist")


    @staticmethod
    def is_empty(obj):
        return (
                not obj.get("title") or obj["title"] == ""
        ) and not (
                obj.get("file_name") or obj["file_name"] == ""
        ) and not (
                obj.get("artist") or obj["artist"] == ""
        )

class BaseSongSchema(Schema):
    title = fields.String(description="title")
    file_name = fields.String(description="file_name")
    artist = fields.String(description="artist")

# Schéma utilisateur de modification (name, username, password)
class SongUpdateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not any(key in data and data[key] != "" for key in ["title", "file_name", "artist"]):
            raise ValidationError("At least one of ['title', 'file_name', 'artist'] must be specified")

from marshmallow import Schema, fields, validates_schema, ValidationError

# Schéma utilisateur de sortie (renvoyé au front)
class SongSchema(Schema):
    music_title = fields.String(description="Music_title")
    artist = fields.String(description="Artist_name")
    file_name = fields.String(description="File_name")

    @staticmethod
    def is_empty(obj):
        return (
                not obj.get("Music_title") or obj["Music_title"] == ""
        ) and not (
                obj.get("Artist_name") or obj["Artist_name"] == ""
        ) and not (
                obj.get("File_name") or obj["File_name"] == ""
        )

class BaseSongSchema(Schema):
    artist = fields.String(description="Music_title")
    file_name = fields.String(description="Artist_name")
    music_title = fields.String(description="File_name")

# Schéma utilisateur de modification (name, username, password)
class SongUpdateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not any(key in data and data[key] != "" for key in ["Music_title", "Artist_name", "File_name"]):
            raise ValidationError("At least one of ['Music_title', 'Artist_name', 'File_name'] must be specified")

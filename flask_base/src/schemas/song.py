from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class SongSchema(Schema):
    id = fields.String(description="UUID")
    Title = fields.String(description="Title")
    published_date = fields.DateTime(published="published date")

    @staticmethod
    def is_empty(obj):
        return (
                not obj.get("id") or obj["id"] == ""
        ) and not (
                obj.get("title") or obj["title"] == ""
        ) and not (
                obj.get("published_date") or obj["published_date"] == ""
        )

class BaseSongSchema(Schema):
    id = fields.String(description="id")
    Title = fields.String(description="Title")
    published_date = fields.String(description="published date")


# Schéma utilisateur de modification (name, username, password)
class SongUpdateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not any(key in data and data[key] != "" for key in ["id", "title", "published_date"]):
            raise ValidationError("at least one of ['id','Title','published_date'] must be specified")

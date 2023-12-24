from marshmallow import Schema, fields, validates_schema, ValidationError


class SongSchema(Schema):
    id = fields.String(description="UUID")
    title = fields.String(description='Title')
    file_name = fields.String(description='File name')

    @staticmethod
    def is_empty(obj):
        return not obj.get("id") or obj.get("id") == ""



class SongAddingSchema(Schema):
    title = fields.String(description="Music title", required=True)
    artist = fields.String(description="Artist name", required=True)
    file_name = fields.String(description="File name", required=True)


class SongAddSchema(SongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("id" in data and data["id"] != "") or
                ("title" in data and data["title"] != "")):
            raise ValidationError("at least one argument must be specified")


class RatingSchema(Schema):
    id = fields.String(description="UUID")
    comment = fields.String(description="Comment")
    rating = fields.Float(description="Rating")
    rating_date = fields.DateTime(description="Rating date")
    song_id = fields.String(description="Song ID")
    user_id = fields.String(description="User ID")


class SongWithRatingSchema(Schema):
    id = fields.String(description="UUID")
    title = fields.String(description="Music title", required=True)
    artist = fields.String(description="Artist name", required=True)
    file_name = fields.String(description="File name", required=True)
    published_date = fields.DateTime(description="Published date")
    ratings = fields.Nested(RatingSchema, many=True)

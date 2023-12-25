from marshmallow import Schema, fields, validates_schema, ValidationError


class RatingSchema(Schema):
    comment = fields.String(description="Comment")
    id = fields.String(description="UUID")
    rating = fields.Int(description='Rating')
    rating_date = fields.String(description='Rating date')
    song_id = fields.String(description='Song id')
    user_id = fields.String(description='User id')

    @staticmethod
    def is_empty(obj):
        return not obj.get("id") or obj.get("id") == ""


class RatingAddSchema(RatingSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("comment" in data and data["comment"] != "") or
                ("id" in data and data["id"] != "") or
                ("rating" in data and data["rating"] != "") or
                ("rating_date" in data and data["rating_date"] != "") or
                ("song_id" in data and data["song_id"] != "") or
                ("user_id" in data and data["user_id"] != "")):
            raise ValidationError("at least one argument must be specified")
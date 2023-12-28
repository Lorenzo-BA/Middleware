from marshmallow import Schema, fields, ValidationError, validates_schema


class RatingSchema(Schema):
    comment = fields.String(description="Comment")
    id = fields.String(description="UUID")
    rating = fields.Int(description='Rating')
    rating_date = fields.String(description='Rating date')
    song_id = fields.String(description='Song id')
    user_id = fields.String(description='User id')


class RatingAddingSchema(Schema):
    comment = fields.String(description="Comment", required=True)
    rating = fields.Int(description="Rating", required=True)
    user_id = fields.String(description='User id')

    @validates_schema
    def validate_rating_range(self, data, **kwargs):
        if "rating" in data:
            rating_value = data["rating"]
            if not 1 <= rating_value <= 5:
                raise ValidationError("Rating must be in the range of 0 to 5.")

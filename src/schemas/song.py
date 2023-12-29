from marshmallow import Schema, fields, ValidationError, validates_schema
from src.schemas.rating import RatingSchema


class SongSchema(Schema):
    id = fields.String(description="UUID")
    title = fields.String(description="Music title")
    artist = fields.String(description="Artist name")
    file_name = fields.String(description="File name")
    Published_date = fields.DateTime(description="Published date")
    ratings = fields.Nested(RatingSchema, many=True)


class SongAddingSchema(Schema):
    title = fields.String(description="Music title", required=True)
    artist = fields.String(description="Artist name", required=True)
    file_name = fields.String(description="File name", required=True)
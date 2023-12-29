from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class UserSchema(Schema):
    id = fields.UUID(description="UUID")
    name = fields.String(description="Name")
    username = fields.String(description="Username")
    inscription_date = fields.DateTime(description="Inscription Date")


class BaseUserSchema(Schema):
    name = fields.String(description="Name")
    password = fields.String(description="Password")
    username = fields.String(description="Username")


class UserUpdateSchema(BaseUserSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("name" in data and data["name"] != "") or
                ("username" in data and data["username"] != "") or
                ("password" in data and data["password"] != "")):
            raise ValidationError("at least one of ['name','username','password'] must be specified")

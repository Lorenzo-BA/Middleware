from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class UserSchema(Schema):
    id = fields.String(description="UUID")
    name = fields.String(description="Name")



class BaseUserSchema(Schema):
    name = fields.String(description="Name")
    password = fields.String(description="Password")
    username = fields.String(description="Username")


# Schéma utilisateur de modification (name, username, password)
class UserUpdateSchema(BaseUserSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("name" in data and data["name"] != "") or
                ("username" in data and data["username"] != "") or
                ("password" in data and data["password"] != "")):
            raise ValidationError("at least one of ['name','username','password'] must be specified")

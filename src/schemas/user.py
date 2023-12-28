from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class UserSchema(Schema):
    id = fields.UUID(description="UUID")
    name = fields.String(description="Name")
    username = fields.String(description="Username")
    inscription_date = fields.DateTime(description="Inscription Date")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("name") or obj.get("name") == "") and \
               (not obj.get("username") or obj.get("username") == "")


class BaseUserSchema(Schema):
    name = fields.String(description="Name")
    password = fields.String(description="Password")
    username = fields.String(description="Username")


# Schéma utilisateur de modification (name, username, password)
class UserUpdateSchema(BaseUserSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if "name" not in data or data["name"] == "" or \
                "username" not in data or data["username"] == "" or \
                "password" not in data or data["password"] == "":
            raise ValidationError("['name','username','password'] must all be specified")

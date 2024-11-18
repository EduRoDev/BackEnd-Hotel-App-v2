package esquemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Administrador struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id_administrador"`
	Nombre   string             `bson:"nombre" json:"nombre"`
	Apellido string             `bson:"apellido" json:"apellido"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

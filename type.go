package ryaas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Lapangan struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama  string             `bson:"nama" json:"nama"`
	Harga string             `bson:"harga" json:"harga"`
}

type Kategori struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama" json:"nama"`
	Turnamen string             `bson:"turnamen" json:"turnamen"`
}

type Kontak struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama" json:"nama"`
	Phone_number string             `bson:"phone_number" json:"phone_number"`
}

type Bank struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Bank string             `bson:"nama_bank" json:"nama_bank"`
	Atas_Nama string             `bson:"atas_nama" json:"atas_nama"`
}

type Diskon struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Dikurang string             `bson:"dikurang" json:"dikurang"`
}

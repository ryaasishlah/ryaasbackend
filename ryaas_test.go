package ryaas

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertLapangan(t *testing.T) {
	dbname := "proyek1"
	lapangan := Lapangan{
		ID:    primitive.NewObjectID(),
		Nama:  "Farhan Rizki Maulana",
		Harga: "150.000",
	}
	insertedID := InsertLapangan(dbname, lapangan)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertKategori(t *testing.T) {
	dbname := "proyek1"
	kategori := Kategori{
		ID:       primitive.NewObjectID(),
		Nama:     "tournament",
		Turnamen: "tournaments",
	}
	insertedID := InsertKategori(dbname, kategori)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertKontak(t *testing.T) {
	dbname := "proyek1"
	kontak := Kontak{
		ID:           primitive.NewObjectID(),
		Nama:         "WAWAN",
		Phone_number: "0821247838192",
	}
	insertedID := InsertKontak(dbname, kontak)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertBank(t *testing.T) {
	dbname := "proyek1"
	bank := Bank{
		ID:        primitive.NewObjectID(),
		Nama_Bank: "Microtfon",
		Atas_Nama: "MUSALODON",
	}
	insertedID := InsertBank(dbname, bank)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertDiskon(t *testing.T) {
	dbname := "proyek1"
	diskon := Diskon{
		ID:       primitive.NewObjectID(),
		Dikurang: "35000",
	}
	insertedID := InsertDiskon(dbname, diskon)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

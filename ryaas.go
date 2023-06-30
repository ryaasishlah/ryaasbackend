package ryaasbackend

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertLapangan(db string, lapangan Lapangan) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("lapangan").InsertOne(context.TODO(), lapangan)
	if err != nil {
		fmt.Printf("InsertLapangan: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKategori(db string, kategori Kategori) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("kategori").InsertOne(context.TODO(), kategori)
	if err != nil {
		fmt.Printf("InsertKategori: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertKontak(db string, kontak Kontak) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("kontak").InsertOne(context.TODO(), kontak)
	if err != nil {
		fmt.Printf("InsertKontak: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertBank(db string, bank Bank) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("bank").InsertOne(context.TODO(), bank)
	if err != nil {
		fmt.Printf("InsertBank: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDiskon(db string, diskon Diskon) (InsertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("diskon").InsertOne(context.TODO(), diskon)
	if err != nil {
		fmt.Printf("InsertDiskon: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetDataLapangan(stats string) (data []Lapangan) {
	user := MongoConnect("proyek1").Collection("lapangan")
	filter := bson.M{"nama": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataLapangan :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataKategori(stats string) (data []Kategori) {
	user := MongoConnect("proyek1").Collection("kategori")
	filter := bson.M{"turnamen": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataKategori :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataKontak(stats string) (data []Kontak) {
	user := MongoConnect("proyek1").Collection("kontak")
	filter := bson.M{"phone_number": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataKontak :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataBank(stats string) (data []Bank) {
	user := MongoConnect("proyek1").Collection("bank")
	filter := bson.M{"nama_bank": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataBank :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetDataDiskon(stats string) (data []Diskon) {
	user := MongoConnect("proyek1").Collection("diskon")
	filter := bson.M{"dikurang": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataDiskon :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caslanqa/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").Find(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"%s\n",uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user data: %v\n", err)
		return
	}

	// Yeni bir ObjectId oluştur ve kullanıcıya ata
	u.Id = bson.NewObjectId()

	// Kullanıcıyı veritabanına ekle
	if err := uc.session.DB("mongo-golang").C("users").Insert(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error inserting user: %v\n", err)
		return
	}

	// Kullanıcıyı JSON formatına dönüştür
	uj, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshaling user: %v\n", err)
		return
	}

	// Yanıt başlığını ayarla ve kullanıcıyı yanıt olarak gönder
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// URL'den kullanıcı kimliğini al
	id := p.ByName("id")

	// Kimliğin geçerli bir BSON ObjectId olup olmadığını kontrol et
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// ObjectId'ye dönüştür
	oid := bson.ObjectIdHex(id)

	// Veritabanından kullanıcıyı sil
	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error deleting user: %v\n", err)
		return
	}

	// Başarılı bir silme işlemi için yanıt gönder
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User deleted successfully\n")
}
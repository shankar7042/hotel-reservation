package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/shankar7042/hotel-reservation-golang/db"
	"github.com/shankar7042/hotel-reservation-golang/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	db.UserStore
}

const (
	testdburi = "mongodb://localhost:27017"
	dbname    = "hotel-reservation-testing"
)

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		t.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)

	app.Post("/", userHandler.HandlePostUser)
	body := types.CreateUserParams{
		FirstName: "testing",
		LastName:  "Again testing",
		Email:     "test@test.com",
		Password:  "12121212121",
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(res.Body).Decode(&user)
	if user.FirstName != body.FirstName {
		t.Errorf("expected firstName %s but got %s \n", body.FirstName, user.FirstName)
	}
	if user.LastName != body.LastName {
		t.Errorf("expected lastName %s but got %s \n", body.LastName, user.LastName)
	}
	if user.FirstName != body.FirstName {
		t.Errorf("expected email %s but got %s \n", body.Email, user.Email)
	}
}

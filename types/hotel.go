package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int                  `bson:"rating" json:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// small, medium, king
	Size    string             `bson:"size" json:"size"`
	SeaSide bool               `bson:"seaside" json:"seaside"`
	Price   float64            `bson:"price" json:"price"`
	HotelID primitive.ObjectID `bson:"hotelID" json:"hotelID"`
}

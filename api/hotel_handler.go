package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shankar7042/hotel-reservation-golang/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	hoelStore db.HotelStore
	roomStore db.RoomStore
}

func NewHotelHandler(hs db.HotelStore, rs db.RoomStore) *HotelHandler {
	return &HotelHandler{
		hoelStore: hs,
		roomStore: rs,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.hoelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.roomStore.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}

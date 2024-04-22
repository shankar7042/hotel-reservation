package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shankar7042/hotel-reservation-golang/db"
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

type GetHotelQueryParams struct {
	Rooms  bool
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qparams GetHotelQueryParams
	if err := c.QueryParser(&qparams); err != nil {
		return err
	}
	hotels, err := h.hoelStore.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

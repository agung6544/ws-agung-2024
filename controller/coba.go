package controller

import (
	"github.com/agung6544/ws-agung-2024/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	cek "github.com/indrariksa/cobapakcage/module"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

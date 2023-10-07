package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"main/ec2"
	"main/eks"
	"main/pricing"
	"main/vpc"
)

func main() {
	app := fiber.New()
	app.Get("/vpcs", func(ctx *fiber.Ctx) error {
		return ctx.JSON(vpc.GetVpc())
	})
	app.Get("/vpcSubnets", func(ctx *fiber.Ctx) error {
		return ctx.JSON(vpc.GetVpcSubnets())
	})
	app.Get("/instanceTypes", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ec2.GetEC2InstanceTypes())
	})
	app.Get("/ami", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ec2.GetEC2AMI())
	})
	app.Get("/eksClusterVersion", func(ctx *fiber.Ctx) error {
		return ctx.JSON(eks.GetEKSClusterVersion())
	})
	app.Get("/costUsage", func(ctx *fiber.Ctx) error {
		return ctx.JSON(pricing.GetCostUsage())
	})
	log.Fatal(app.Listen(":3000"))
}

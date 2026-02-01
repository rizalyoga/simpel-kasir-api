package repositories

import "kasir-api-bootcamp/models"

var Categories = []models.Category{
	{
		ID:          1,
		Name:        "Bahan Pokok",
		Description: "Bahan pokok sehari - hari.",
	},
	{
		ID:          2,
		Name:        "Snack",
		Description: "Snack atau jajanan.",
	},
	{
		ID:          3,
		Name:        "Soda",
		Description: "Minuman bersoda.",
	},
}

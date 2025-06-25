package service

import "newProject/cafe/manager"

func DrinkLatte() {
	coffee := manager.NewLatte()
	coffee.Drink()
}

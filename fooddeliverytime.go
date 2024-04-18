package piscine

type food struct {
	preptime int
	name     string
}

func FoodDeliveryTime(order string) int {
	// Declaring the menu items and the preperation time of each item
	burger := food{}
	chips := food{}
	nuggets := food{}

	burger.preptime = 15
	burger.name = "burger"
	chips.preptime = 10
	chips.name = "chips"
	nuggets.preptime = 12
	nuggets.name = "nuggets"

	menu := []food{burger, chips, nuggets}

	for _, item := range menu {
		if order == item.name {
			return item.preptime
		}
	}
	return 404
}

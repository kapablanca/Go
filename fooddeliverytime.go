package piscine

type food struct {
	name     string
	preptime int
}

// Function to create menu item
func MenuItem(name string, time int) *food {
	i := food{name: name, preptime: time}
	return &i
}

// Given the following menu with the coresponding times that each item takes to cook
// (burger takes 15 min, chips takes 10 min and nuggets takes 12 min), return the time
//
//	that each order item takes to be prepared and the toteal amount of time for the order
//
// to be ready assuming the items are cooked one after the other. If any of the order items
//
//	don't exist in the menu above return the error message 404
func FoodDeliveryTime(order string) int {
	// Creating given menu items
	burger := MenuItem("burger", 15)
	chips := MenuItem("chips", 10)
	nuggets := MenuItem("nuggets", 12)

	menu := []*food{burger, chips, nuggets}

	for _, item := range menu {
		if item.name == order {
			return item.preptime
		}
	}
	return 404
}

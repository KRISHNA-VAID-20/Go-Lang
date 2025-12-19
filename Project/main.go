package main

import (
	"fmt"
)

// Structure for each item
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var inventory []Item // slice to hold items

func main() {
	fmt.Println("--- SHOP INVENTORY SYSTEM ---")
	for {
		fmt.Println("1. Add Item")
		fmt.Println("2. Display Items")
		fmt.Println("3. Search Item")
		fmt.Println("4. Update Quantity")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addItem()
		case 2:
			displayItems()
		case 3:
			searchItem()
		case 4:
			updateQuantity()
		case 5:
			fmt.Println("Exiting...\n Thank you!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

// 1. Add item
func addItem() {
	var id, qty int
	var name string
	var price float64

	fmt.Print("Enter Item ID: ")
	fmt.Scanln(&id)

	fmt.Print("Enter Item Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&qty)

	fmt.Print("Enter Price: ")
	fmt.Scanln(&price)

	newItem := Item{id, name, qty, price}
	inventory = append(inventory, newItem)

	fmt.Println("Item added successfully!")
}

// 2. Display all items
func displayItems() {
	fmt.Println("\n--- INVENTORY LIST ---")
	for _, item := range inventory {
		fmt.Printf("ID: %d | Name: %s | Qty: %d | Price: %.2f\n",
			item.ID, item.Name, item.Quantity, item.Price)
	}
}

// 3. Search item by ID
func searchItem() {
	var id int
	fmt.Print("Enter Item ID to search: ")
	fmt.Scanln(&id)

	for _, item := range inventory {
		if item.ID == id {
			fmt.Println("\nItem Found!")
			fmt.Printf("ID: %d | Name: %s | Qty: %d | Price: %.2f\n",
				item.ID, item.Name, item.Quantity, item.Price)
			return
		}
	}
	fmt.Println("Item not found.")
}

// 4. Update quantity of item
func updateQuantity() {
	var id, newQty int
	fmt.Print("Enter Item ID to update quantity: ")
	fmt.Scanln(&id)

	for i := range inventory {
		if inventory[i].ID == id {
			fmt.Print("Enter new quantity: ")
			fmt.Scanln(&newQty)
			inventory[i].Quantity = newQty
			fmt.Println("Quantity updated successfully!")
			return
		}
	}
	fmt.Println("Item not found.")
}

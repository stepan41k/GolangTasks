package main

import "fmt"

type Observer interface {
	Update(itemName string)
	GetID() string
}

type Subject interface {
	Subscribe(Observer)
	Unsubscribe(Observer)
	NotifyAll()
}

type Item struct {
	observers []Observer
	name      string
	inStock   bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Товар '%s' теперь в наличии!\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

func (i *Item) Subscribe(o Observer) {
	i.observers = append(i.observers, o)
	fmt.Printf("Подписчик %s подписался на %s\n", o.GetID(), i.name)
}

func (i *Item) Unsubscribe(o Observer) {
	for idx, observer := range i.observers {
		if observer.GetID() == o.GetID() {
			i.observers = append(i.observers[:idx], i.observers[idx+1:]...)
			fmt.Printf("Подписчик %s отписался от %s\n", o.GetID(), i.name)
			break
		}
	}
}

func (i *Item) NotifyAll() {
	for _, observer := range i.observers {
		observer.Update(i.name)
	}
}

type Customer struct {
	id string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf(" [Уведомление для %s]: Товар '%s' появился на складе, успейте купить!\n", c.id, itemName)
}

func (c *Customer) GetID() string {
	return c.id
}

func main() {
	shirt := NewItem("Футболка Go")

	customer1 := &Customer{id: "user_101"}
	customer2 := &Customer{id: "user_202"}

	shirt.Subscribe(customer1)
	shirt.Subscribe(customer2)

	fmt.Println("---")

	shirt.UpdateAvailability()

	fmt.Println("---")

	shirt.Unsubscribe(customer1)

	fmt.Println("---")

	shirt.UpdateAvailability()
}

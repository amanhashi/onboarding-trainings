package main

import (
	"fmt"
	"strings"
)

// Contact represents an individual contact
type Contact struct {
	Name  string
	Email string
	Phone string
}

// ContactBook holds a slice of contacts
type ContactBook struct {
	contacts map[string]*Contact
}

// AddContact adds a new contact to the book
func (cb *ContactBook) AddContact(c *Contact) {
	fmt.Println("====>", c)
	cb.contacts[strings.ToLower(c.Name)] = c
}

// GetAllContacts returns all contacts
func (cb *ContactBook) GetAllContacts() []*Contact {
	var list []*Contact
	for _, c := range cb.contacts {
		list = append(list, c)
	}
	return list
}

// SearchContact searches for a contact by name
func (cb *ContactBook) SearchContact(name string) *Contact {
	return cb.contacts[strings.ToLower(name)]
}

// Formatter interface for printing
type Formatter interface {
	Format() string
}

// Format implements Formatter for Contact
func (c *Contact) Format() string {
	return fmt.Sprintf("Name: %s, Email: %s, Phone: %s", c.Name, c.Email, c.Phone)
}

func main() {
	cb := &ContactBook{contacts: make(map[string]*Contact)}

	// Using pointer to modify data
	contact1 := &Contact{Name: "Alice", Email: "AlicE@example.com", Phone: "1234565467890"}
	contact2 := &Contact{Name: "Bob", Email: "bob@example.com", Phone: "9874566543210"}

	cb.AddContact(contact1)
	cb.AddContact(contact2)

	fmt.Println("All Contacts:")
	for _, c := range cb.GetAllContacts() {
		fmt.Println(c.Format())
	}

	fmt.Println("\nSearch Result:")
	name := "alice"
	result := cb.SearchContact(name)
	if result != nil {
		fmt.Println("Found:", result.Format())
	} else {
		fmt.Println("Contact not found")
	}

	// Arrays (demo use)
	arr := [3]string{"alpha", "beta", "gamma"}
	fmt.Println("\nStatic Array Example:")
	for i, v := range arr {
		fmt.Printf("%d: %s\n", i, v)
	}
}

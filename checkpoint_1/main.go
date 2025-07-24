package main

import (
	"fmt"
	"strings"
	"sync"
)

// Contact represents an individual contact
type Contact struct {
	Name  string
	Email string
	Phone string
}

// ContactBook holds a map of contacts and a mutex for safe access
type ContactBook struct {
	contacts map[string]*Contact
	mu       sync.Mutex
}

// AddContact adds a new contact safely using mutex
func (cb *ContactBook) AddContact(c *Contact) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.contacts[strings.ToLower(c.Name)] = c
}

// GetAllContacts returns all contacts safely
func (cb *ContactBook) GetAllContacts() []*Contact {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	var list []*Contact
	for _, c := range cb.contacts {
		list = append(list, c)
	}
	return list
}

// SearchContact safely looks up a contact by name
func (cb *ContactBook) SearchContact(name string) *Contact {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.contacts[strings.ToLower(name)]
}

// Formatter interface for output
type Formatter interface {
	Format() string
}

// Format implements Formatter for Contact
func (c *Contact) Format() string {
	return fmt.Sprintf("Name: %s, Email: %s, Phone: %s", c.Name, c.Email, c.Phone)
}

func main() {
	cb := &ContactBook{contacts: make(map[string]*Contact)}

	// Simulate adding contacts
	contact1 := &Contact{Name: "Alice", Email: "alice@example.com", Phone: "1234567890"}
	contact2 := &Contact{Name: "Bob", Email: "bob@example.com", Phone: "9876543210"}

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

	// Array example
	arr := [3]string{"alpha", "beta", "gamma"}
	fmt.Println("\nStatic Array Example:")
	for i, v := range arr {
		fmt.Printf("%d: %s\n", i, v)
	}
}

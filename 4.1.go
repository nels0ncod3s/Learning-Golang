package main

import "fmt"

var contacts = map[string]string{}

func addContact() {
 var name, phone string

 fmt.Print("Enter name: ")
 fmt.Scanln(&name)

 fmt.Print("Enter phone: ")
 fmt.Scanln(&phone)

 contacts[name] = phone

 fmt.Println("✅ Contact added!")
}

func viewContact() {
 var name string

 fmt.Print("Enter name: ")
 fmt.Scanln(&name)

 phone, exists := contacts[name]

 if exists {
  fmt.Println("Phone:", phone)
 } else {
  fmt.Println("❌ Contact not found.")
 }
}

func updateContact() {
 var name string

 fmt.Print("Enter name: ")
 fmt.Scanln(&name)

 _, exists := contacts[name]

 if !exists {
  fmt.Println("❌ Contact not found.")
  return
 }

 var newPhone string

 fmt.Print("Enter new phone: ")
 fmt.Scanln(&newPhone)

 contacts[name] = newPhone

 fmt.Println("✅ Updated successfully!")
}

func deleteContact() {
 var name string

 fmt.Print("Enter name: ")
 fmt.Scanln(&name)

 delete(contacts, name)

 fmt.Println("🗑️ Contact deleted.")
}

func listContacts() {

 if len(contacts) == 0 {
  fmt.Println("No contacts saved.")
  return
 }

 fmt.Println("\n--- Contacts ---")

 for name, phone := range contacts {
  fmt.Println(name, "-", phone)
 }
}

func main() {

 for {

  fmt.Println("\n===== CONTACT BOOK =====")
  fmt.Println("1. Add Contact")
  fmt.Println("2. View Contact")
  fmt.Println("3. Update Contact")
  fmt.Println("4. Delete Contact")
  fmt.Println("5. List Contacts")
  fmt.Println("6. Exit")

  var choice int

  fmt.Print("Choose: ")
  fmt.Scanln(&choice)

  switch choice {

  case 1:
   addContact()

  case 2:
   viewContact()

  case 3:
   updateContact()

  case 4:
   deleteContact()

  case 5:
   listContacts()

  case 6:
   fmt.Println("Goodbye 👋")
   return

  default:
   fmt.Println("Invalid choice.")
  }
 }
}
package main

// in go lang functions in struct is named as methods
import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in go lang; no super or no parent exists in go

	vishal := User{"Vishal", "vishal@gmail.com", true, 24}
	fmt.Println(vishal)

	// this will print the details with the fields created in the struct
	fmt.Printf("Vishal details are: %+v", vishal)

	fmt.Printf("Name is: %v and email is: %v\n", vishal.Name, vishal.Email)

	vishal.GetStatus()
	vishal.NewMail()
	fmt.Printf("Name is: %v and email is: %v\n", vishal.Name, vishal.Email)


}

// defining a structure
// things which are having first letter as capital means they are public
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// now it is method as we have passed the struct User to it and we can either pass a property or the whole struct to this method and u is the object of the User struct
func (u User) GetStatus(){
	fmt.Println("Is user active: ",u.Status)
}

// in this function it is not changing the value of the actual object as the copy is passed to the function
// hence we need to use pointer inorder to change the original object
func (u User) NewMail(){
	u.Email = "v@test.dev"
	fmt.Println("Email of the user is: ",u.Email)
}
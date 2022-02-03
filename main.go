package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var bitcoin_conference string = "Bitcoin Conference"

const bitcoin_conference_ticket int = 50

var remaining_tickets uint = 50
var bookings = make([]UserData, 0)

var wait_group = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	emailAddress    string
	numberOfTickets uint
}

func main() {

	greet_users(bitcoin_conference, bitcoin_conference_ticket, remaining_tickets)

	// fmt.Printf("Bitcoin Conference Ticket Type is: %T\n", bitcoin_conference_ticket)
	// fmt.Printf("Remaining Ticket Type is: %T\n", remaining_tickets)
	// fmt.Printf("Bitcoin Conference Type is: %T\n", bitcoin_conference)

	firstName, lastName, emailAddress, userTickets := validateAll()

	// isValid_city := city == "Manilla" || city == "Abuja"
	isValid_name, isValid_email, isValid_ticket := get_user_input(firstName, lastName, emailAddress, userTickets, remaining_tickets)

	if isValid_name && isValid_email && isValid_ticket {

		wait_group.Add(1)

		go sendTickets(userTickets, firstName, lastName, emailAddress)

		var usersData = UserData{
			firstName:       firstName,
			lastName:        lastName,
			emailAddress:    emailAddress,
			numberOfTickets: userTickets,
		}
		// usersData["firstName"] = firstName
		// usersData["lastName"] = lastName
		// usersData["emailAddress"] = emailAddress
		// usersData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		remaining_tickets = remaining_tickets - userTickets
		bookings = append(bookings, usersData)

		// fmt.Printf("The whole arrays: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice size: %v\n", len(bookings))

		fmt.Printf("Mr/Mrs %v %v booked %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("You will recieve a confirmation email at %v\n", emailAddress)
		fmt.Printf("Current users in our booking software are => %v\n", bookings)
		fmt.Printf("%v tickets remaining for our %v\n", remaining_tickets, bitcoin_conference)

		// firstNames := []string{}
		// for _, booking := range bookings {
		// 	var names = strings.Fields(booking)
		// 	// var firstName = names[0]
		// 	firstNames = append(firstNames, names[0])
		// }
		// fmt.Printf("The first names in our booking app are: %v\n", firstNames)

		// firstNames := get_firstNames(bookings)
		// fmt.Printf("The first names in our bookings software are: %v\n", firstNames)

		if remaining_tickets == 0 {
			fmt.Println("Our Conference Is Booked Out. Come back next year")
			// break
		}
	} else {
		if !isValid_name {
			fmt.Println("The first name or last name you enetered is too short")
		}
		if !isValid_email {
			fmt.Println("The email address must contain your first name")
		}
		if !isValid_ticket {
			fmt.Println("number of tickets is invalid")
		}
		// fmt.Println("There was an error fetching your input that's all we know")
		// continue
	}

	// var bitcoin_city string

	// fmt.Println("What's your city? ")
	// fmt.Scan(&bitcoin_city)

	// switch bitcoin_city {
	// case "New York":
	// 	fmt.Println("Welcome to new york")
	// case "Singapore":
	// 	fmt.Println("Welcome to singapore")
	// case "Berlin":
	// 	fmt.Println("Welcome to berlin")
	// case "Austria":
	// 	fmt.Println("Welcome to austria")
	// case "Perth":
	// 	fmt.Println("Welcome to perth")
	// case "Lagos":
	// 	fmt.Println("Welcome to lagos")
	// default:
	// 	fmt.Println("No valid city selected")
	// }

	wait_group.Wait()
}

func greet_users(bitconf string, confTickets int, remainingTickets uint) {
	// fmt.Println("Welcome to our conference")
	fmt.Printf("Welcome to our %v booking software\n", bitconf)
	fmt.Printf("we have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here and attend")
}

func get_user_input(firstName string, lastName string, emailAddress string, userTickets uint, remaining_tickets uint) (bool, bool, bool) {
	isValid_name := len(firstName) >= 2 && len(lastName) >= 2
	isValid_email := strings.Contains(emailAddress, firstName)
	isValid_ticket := userTickets > 0 && userTickets <= remaining_tickets

	return isValid_name, isValid_email, isValid_ticket
}

func get_firstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func validateAll() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint
	// var city string

	// fmt.Scanf(&userName)
	fmt.Println("What's Your First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("What's Your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("What's Your Email Address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("Please Enter Number of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

func sendTickets(userTickets uint, firstName string, lastName string, emailAddress string) {
	time.Sleep(50 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", tickets, emailAddress)
	fmt.Println("###############")
	wait_group.Done()
}

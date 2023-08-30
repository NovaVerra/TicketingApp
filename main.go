package main

import (
	"fmt"
	"strings"
)

func main() {

	// Conference Information
	var		conference_name string = "Defense Con"
	var		conference_year uint = 2023;
	var		conference_location string = "San Diego"
	const	conference_tickets uint = 50
	var		remaining_conference_tickets uint = conference_tickets;

	// List of Attendees
	var attendees []string;

	// User Input
	var user_first_name string;
	var user_last_name string;
	var user_email string;
	var user_ticket_count uint;

	fmt.Printf("Welcome to Nova ticketing!\n")
	fmt.Printf("Would you like to purchase your ticket for the %v %v in %v\n", conference_year, conference_name, conference_location)
	fmt.Printf("There are a total of %v conference_tickets available and we currently have %v left\n", conference_tickets, remaining_conference_tickets)

	for {
		if remaining_conference_tickets > 0 {
			fmt.Printf("Please enter your first name: ")
			fmt.Scanf("%s", &user_first_name)
		
			fmt.Printf("Please enter your last name: ")
			fmt.Scanf("%s", &user_last_name)
		
			attendees = append(attendees, user_first_name + " " + user_last_name)
		
			fmt.Printf("Please enter your email: ")
			fmt.Scanf("%s", &user_email)
		
			fmt.Printf("Please enter the number of tickets you'd like to purchase: ")
			fmt.Scanf("%d", &user_ticket_count)

			if user_ticket_count > remaining_conference_tickets {
				fmt.Printf("We only have %d tickets remaining", remaining_conference_tickets)
				} else {
					remaining_conference_tickets -= user_ticket_count
					fmt.Printf("User: %s %s purchased %d tickets. You'll receive your confirmation at %s\n", user_first_name, user_last_name, user_ticket_count, user_email)
					fmt.Printf("There's a total of %d tickets left\n", remaining_conference_tickets)
				}
		} else {
			fmt.Printf("We are out of tickets now, please come back next year")
			break;
		}
		var list_user []string;
		for _, attendee := range attendees {
			var name = strings.Fields(attendee)
			list_user = append(list_user, name[0])
		}
		fmt.Printf("%s\n", list_user)
	}
}

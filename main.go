package main

import (
	"fmt"
	"log"
	"time"

	"github.com/civera17/astana-extractor/client"
)

func main() {
	astanaC := client.NewAstanaClient("https://app.asana.com/api/1.0")

	// Return each project/user in its own JSON
	projects, err := astanaC.GetAllProjects("1")
	if err != nil {
		log.Fatalf("Req to astana not successful: %s", err.Error())
	}

	fmt.Printf("%+v\n", projects)
	
	users, err := astanaC.GetAllUsers("10")
	if err != nil {
		log.Fatalf("Req to astana not successful: %s", err.Error())
	}
	fmt.Printf("%+v\n", users)


	// Show periodic extraction for users with 5sec interval
	usersPeriodicExtraction(5)
	projectsPeriodicExtraction(5)
	// Sleep for 15 sec for testing purposes
	time.Sleep(15 * time.Second)

	// Show periodic extraction for users with 20sec interval
	// usersPeriodicExtraction(20)
	// projectsPeriodicExtraction(20)
	// // Sleep for 15 sec for testing purposes
	// time.Sleep(60 * time.Second)
}

// usersPeriodicExtraction implements periodic extraction for users, takes interval as arg
func usersPeriodicExtraction(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	astanaC := client.NewAstanaClient("https://app.asana.com/api/1.0")
	go func() {
		for {
			select {
			case <- ticker.C:
				users, err := astanaC.GetAllUsers("10")
					if err != nil {
						log.Fatalf("Req to astana not successful: %s", err.Error())
					}
					fmt.Printf("%+v\n", users)
					fmt.Println("tick users")
			default:

			}
		}
	}()
}

// projectsPeriodicExtraction implements periodic extraction for projects, takes interval as arg
func projectsPeriodicExtraction(interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	astanaC := client.NewAstanaClient("https://app.asana.com/api/1.0")
	go func() {
		for {
			select {
			case <- ticker.C:
				users, err := astanaC.GetAllUsers("10")
					if err != nil {
						log.Fatalf("Req to astana not successful: %s", err.Error())
					}
					fmt.Printf("%+v\n", users)
					fmt.Println("tick users")
			default:

			}
		}
	}()
}
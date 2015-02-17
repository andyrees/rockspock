package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func generate_computer_choice() (choice string, err error) {
	choices := []string{"rock", "paper", "scissors", "lizard", "spock"}
	rand.Seed(time.Now().Unix())
	myindex := rand.Intn(4)
	return choices[myindex], nil
}

// returns int 1=user, 2=computer, 3=draw, and win string
func winner(userchoice, computerchoice string) (int, string) {
	winner := "Draw"
	switch {
	case userchoice == computerchoice:
		return 3, "draw"
	case userchoice == "scissors" && computerchoice == "paper":
		return 1, "User Wins: Scissors cut Paper"
	case computerchoice == "scissors" && userchoice == "paper":
		return 2, "Computer Wins: Scissors cut Paper"
	case userchoice == "paper" && computerchoice == "rock":
		return 1, "User Wins: Paper covers Rock"
	case computerchoice == "paper" && userchoice == "rock":
		return 2, "Computer Wins: Paper covers Rock"
	case userchoice == "rock" && computerchoice == "lizard":
		return 1, "User Wins: Rock crushes Lizard"
	case computerchoice == "rock" && userchoice == "lizard":
		return 2, "Computer Wins: Rock crushes Lizard"
	case userchoice == "lizard" && computerchoice == "spock":
		return 1, "User Wins: Lizard poisons Spock"
	case computerchoice == "lizard" && userchoice == "spock":
		return 2, "Computer Wins: Lizard poisons Spock"
	case userchoice == "spock" && computerchoice == "scissors":
		return 1, "User Wins: Spock smashes Scissors"
	case computerchoice == "spock" && userchoice == "scissors":
		return 2, "Computer Wins: Spock smashes Scissors"
	case userchoice == "scissors" && computerchoice == "lizard":
		return 1, "User Wins: Scissors decapitate Lizard"
	case computerchoice == "scissors" && userchoice == "lizard":
		return 2, "Computer Wins: Scissors decapitate Lizard"
	case userchoice == "lizard" && computerchoice == "paper":
		return 1, "User Wins: Lizard eats Paper"
	case computerchoice == "lizard" && userchoice == "paper":
		return 2, "Computer Wins: Lizard eats Paper"
	case userchoice == "paper" && computerchoice == "spock":
		return 1, "User Wins: Paper disproves Spock"
	case computerchoice == "paper" && userchoice == "spock":
		return 2, "Computer Wins: Paper disproves Spock"
	case userchoice == "spock" && computerchoice == "rock":
		return 1, "User Wins: Spock vaporizes Rock"
	case computerchoice == "spock" && userchoice == "rock":
		return 2, "Computer Wins: Spock vaporizes Rock"
	case userchoice == "rock" && computerchoice == "scissors":
		return 1, "User Wins: Rock crushes Scissors"
	case computerchoice == "rock" && userchoice == "scissors":
		return 2, "Computer Wins: Rock crushes Scissors"
	default:
		return 3, winner
	}
	return 3, winner
}

func who_won_then(userChoice string) int {
	fmt.Println("You chose     : ", userChoice)
	ch, err := generate_computer_choice()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Computer chose: ", ch)
	who_won, win_string := winner(strings.ToLower(userChoice), ch)
	fmt.Printf("-----\n%s\n", win_string)
	return who_won
}

func stringInSlice(a string) bool {
	choices := []string{"rock", "paper", "scissors", "lizard", "spock"}
	for _, b := range choices {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	user_score := 0
	computer_score := 0
	draw_score := 0
	fmt.Println("Choose: Rock | Paper | Scissors | Lizard | Spock | Quit")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		user_choice := strings.ToLower(scanner.Text())
		if user_choice == "quit" {
			fmt.Println("Good Game")
			fmt.Printf("User: %d, Computer: %d, Draw: %d\n", user_score, computer_score, draw_score)
			return
		} else if stringInSlice(user_choice) {
			score := who_won_then(user_choice)
			switch score {
			case 1:
				user_score = user_score + 1
			case 2:
				computer_score = computer_score + 1
			case 3:
				draw_score = draw_score + 1
			}
		} else {
			fmt.Println("Choose: Rock | Paper | Scissors | Lizard | Spock | Quit")
		}
		fmt.Printf("User: %d, Computer: %d, Draw: %d\n", user_score, computer_score, draw_score)
		fmt.Println("")
	}

}

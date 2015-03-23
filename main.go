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

func generateComputerChoice() (choice string, err error) {
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
}

func whoWonThen(userChoice string) int {
	fmt.Println("You chose     : ", userChoice)
	ch, err := generateComputerChoice()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Computer chose: ", ch)
	whoWon, winString := winner(strings.ToLower(userChoice), ch)
	fmt.Printf("-----\n%s\n", winString)
	return whoWon
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
	userScore := 0
	computerScore := 0
	drawScore := 0
	fmt.Println("Choose: Rock | Paper | Scissors | Lizard | Spock | Quit")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userChoice := strings.ToLower(scanner.Text())
		if userChoice == "quit" {
			fmt.Println("Good Game")
			fmt.Printf("User: %d, Computer: %d, Draw: %d\n", userScore, computerScore, drawScore)
			return
		} else if stringInSlice(userChoice) {
			score := whoWonThen(userChoice)
			switch score {
			case 1:
				userScore = userScore + 1
			case 2:
				computerScore = computerScore + 1
			case 3:
				drawScore = drawScore + 1
			}
		} else {
			fmt.Println("Choose: Rock | Paper | Scissors | Lizard | Spock | Quit")
		}
		fmt.Printf("User: %d, Computer: %d, Draw: %d\n", userScore, computerScore, drawScore)
		fmt.Println("")
	}

}

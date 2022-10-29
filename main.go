package main

import (
	"fmt"
	"os"
	"strconv"
)

// global variable declaration
var (
	defaultPin = 9876
	oldPin     int
	newPin     int
	USD        = 10000
	choice     string
	i          = 0
)

// main function
func main() {
	welcome()
	for {
		fmt.Println("Welcome to your account! Please enter your 4 digits default pin. Press 5 to exit program.")
		var inputPin int

		_, err := fmt.Scan(&inputPin)
		// fmt.Println(inputPin)
		numberOfdigit := len(strconv.Itoa(inputPin))
		// fmt.Println(numberOfdigit)

		if err != nil {
			fmt.Println("Error: Please use numbers 0-9")
		} else if inputPin == 5 {
			exit()
		} else if numberOfdigit != 4 {
			fmt.Println("Error: The pin entered is not 4 digits")
		} else if inputPin != defaultPin {
			fmt.Println("Error: Wrong default pin")
		} else {
			newline(1)
			displayMenu()
		}
	}
}

// welcome function: displays a banner
func welcome() {
	newline(1)
	fmt.Println("\t█████╗ ████████╗███╗   ███╗\tBy @cyberspades")
	fmt.Println("\t██╔══██╗╚══██╔══╝████╗ ████║")
	fmt.Println("\t███████║   ██║   ██╔████╔██║")
	fmt.Println("\t██╔══██║   ██║   ██║╚██╔╝██║")
	fmt.Println("\t██║  ██║   ██║   ██║ ╚═╝ ██║")
	fmt.Println("\t╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚═╝")
	newline(1)

}

// newline function: prints a new line
func newline(numberOfLines int) {
	// i := 0
	for i < numberOfLines {
		fmt.Println("\n")
		i++
	}
}

// displayMenu function: displays the atm menu
func displayMenu() {
	fmt.Println("Select an operation")
	fmt.Println("1. Change pin.\t\t\t2. Account balance.")
	fmt.Println("3. Deposit money\t\t4. Withdraw money")
	fmt.Println("5. Exit ")

	var menuNumber int
	_, err := fmt.Scan(&menuNumber)

	if err != nil {
		fmt.Println("Error: Please enter a number between 1 and 5")
	}

	//switch for menu options
	switch menuNumber {
	case 1:
		checkPin()
	case 2:
		accountBalance()
	case 3:
		deposit()
	case 4:
		withdraw()
	case 5:
		exit()
	default:
		fmt.Println("Error: Please enter a number between 1 and 5")
		displayMenu()
	}

}

// checkPin function: verifies that the old is not
func checkPin() {
	fmt.Println("Enter your pin")
	// var oldPin int
	_, err := fmt.Scan(&oldPin)

	if err != nil {
		fmt.Println("Error: Please use numbers 0-9")
		checkPin()
	} else if oldPin != defaultPin {
		fmt.Println("Error: Enter the default or correct pin")
		checkPin()
	} else {
		changePin()
	}
}

// changePin function: sets the new pin
func changePin() {
	fmt.Println("Enter your new pin")
	_, err := fmt.Scan(&newPin)
	numberOfdigit2 := len(strconv.Itoa(newPin))

	if err != nil {
		fmt.Println("Error: Please use numbers 0-9")
		changePin()
	} else if numberOfdigit2 != 4 {
		fmt.Println("Error: The pin entered is not 4 digits")
		fmt.Println(numberOfdigit2)
		changePin()
	} else if newPin == oldPin {
		fmt.Println("Error: New pin can not be the same with the old pin")
		changePin()
	}
	defaultPin = newPin
	fmt.Println(defaultPin)
	displayMenu()
}

// accountBalance function: prints the exisiting USD Balance to the console
func accountBalance() {
	fmt.Println("Your account balance is $", USD)
	anotherOption()
}

// menuAuth function: It authenticates the user to the display menu once the correct pin has been inputed
func menuAuth() {
	fmt.Println("To perform another operation please enter your pin")
	var inputPin int
	_, err := fmt.Scan(&inputPin)
	fmt.Println(inputPin)

	if err != nil {
		fmt.Println("Error: Please use numbers 0-9")
		menuAuth()
	} else if inputPin != defaultPin {
		fmt.Println("Error: Wrong pin")
		menuAuth()
	} else {
		displayMenu()
	}
}

// anotherOption function: It allows the user to choose between performing another fucntion or exitting the program
func anotherOption() {
	fmt.Println("Do you want to perform another operation? Enter yes or no (in lowercase)")
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println("Error: Enter yes or no")
		anotherOption()
	}

	if choice == "yes" {
		menuAuth()
	}

	if choice == "no" {
		exit()
	} else {
		fmt.Println("Error: Please enter yes or no in lowercase!")
		anotherOption()
	}
}

// deposit function: adds the input to the exisiting USD Balance
func deposit() {
	fmt.Println("Enter the amount(USD) of money to be deposited")
	var amount int
	_, err := fmt.Scan(&amount)

	if err != nil {
		if i < 2 {
			fmt.Println("Error: Please enter the amount of money(USD) in numbers")
			deposit()
			i++
		}
	}
	USD = amount + USD
	fmt.Println("Your account balance is $", USD)
	anotherOption()
}

// withdraw function: subtracts the input from the exisiting USD Balance
func withdraw() {
	fmt.Println("Enter the amount of money(USD) to be withdrawn from your account")
	var withdrawnAmount int
	_, err := fmt.Scan(&withdrawnAmount)

	if err != nil {
		if i < 2 {
			fmt.Println("Error: Please enter an amount in words")
			deposit()
		}
		anotherOption()
	}
	USD = USD - withdrawnAmount
	fmt.Println("Your account balance is $", USD)
	anotherOption()
}

// exit function: exits the program
func exit() {
	fmt.Println("Thanks for using our service.")
	newline(2)
	fmt.Println("\t██████╗  ██████╗  ██████╗ ██████╗     ██████╗ ██╗   ██╗███████╗")
	fmt.Println("\t██╔════╝ ██╔═══██╗██╔═══██╗██╔══██╗    ██╔══██╗╚██╗ ██╔╝██╔════╝")
	fmt.Println("\t██║  ███╗██║   ██║██║   ██║██║  ██║    ██████╔╝ ╚████╔╝ █████╗")
	fmt.Println("\t██║   ██║██║   ██║██║   ██║██║  ██║    ██╔══██╗  ╚██╔╝  ██╔══╝")
	fmt.Println("\t╚██████╔╝╚██████╔╝╚██████╔╝██████╔╝    ██████╔╝   ██║   ███████╗")
	fmt.Println("\t ╚═════╝  ╚═════╝  ╚═════╝ ╚═════╝     ╚═════╝    ╚═╝   ╚══════╝")
	os.Exit(0)
}

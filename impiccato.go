package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

var mappa map[int]bool

func main() {

	tentativi := 10
	isComplete := false 
	parola, err := getParola();

	if err != nil {
		fmt.Println("C'è stato un errore uagliu")
		os.Exit(0)	
	} 

	parolaLen := len(parola)
	mappa = initialize(parolaLen)

	fmt.Println("Perfetto, Giocatore 2, hai 10 vite a disposizione, Buona Fortuna!")
	printParola(parola, mappa)

	for (tentativi > 0) && (!isComplete) {

		userInput := getUserInput()
		userInputLen := len(userInput)
		
		if userInputLen == 1 {

			letterFound := findLetter(userInput, parola)

			if letterFound {
				fmt.Println("Hai indovinato la lettera!")
			} else {
				tentativi -= 1
				fmt.Println("Non hai trovato la lettera, ritenta.")
			}

			isComplete = checkParola(parola)

			fmt.Println("Hai ancora", tentativi, "vite.")
			printParola(parola, mappa)

		} else if (userInputLen != 1) && (userInput != parola) {
			tentativi -= 1
			fmt.Println("La parola è errata, hai ancora", tentativi, "tentativi, ritenta.")
			printParola(parola, mappa)
		} else if userInput == parola {
			fmt.Println("La parola è corretta, complimenti!")
			isComplete = true
		}
	}

	if isComplete {
		fmt.Println("Complimenti hai completato la parola!")
	}
}

func printParola(parola string, mappa map[int]bool) {

	parolaLen := len(parola)

	for i := 0; i < parolaLen; i++ {
		if mappa[i] == true {
			fmt.Print(string(parola[i]))
		} else {
			fmt.Print(" _ ")
		}
	} 
	fmt.Println()
}

func getParola() (string, error) {
		
	fmt.Println("Giocatore 1, inserisci una parola: ")
	byteArr, err := term.ReadPassword(0)
		
	if err != nil {
		return "", err 
	} 

	str := string(byteArr)
	inputLen := len(str)

	if inputLen > 2 {
		return str, nil
	} else {
		fmt.Println("Inserisci una parola con più di due lettere")
		return getParola()
	}
}

func initialize(parolaLen int) map[int]bool {
	
	mappa := make(map[int]bool)
	
	for i := 0; i < parolaLen; i++ {
		if i == 0 || i == (parolaLen - 1) {
			mappa[i] = true
		} else {
			mappa[i] = false
		}
	} 
	
	return mappa
}

func getUserInput() string {
	
	var userInput string

	fmt.Println("Inserisci una lettera da azzeccare o la parola intera: ")
	fmt.Scan(&userInput)
	
	return userInput
} 

func findLetter(input string, parola string) (bool) {

	letterFound := false
	parolaLen := len(parola)

	for i := 0; i < parolaLen; i++ {
		if string(parola[i]) == input {
			mappa[i] = true
			letterFound = true
		} 
	} 
	
	return letterFound
}

func checkParola(parola string) bool {
	
	parolaLen := len(parola)

	for i := 0; i < parolaLen; i++ {
		if !mappa[i] {
			return false
		}
	}
	
	return true
}
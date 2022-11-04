package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"time"
	"math/rand"
	"strings"
)

var mappa map[int]bool

func main() {

	var parola string
	var mod int

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Ciao! Vuoi giocare in modalità singola o multigiocatore?")
	fmt.Println("1) Modalità giocatore singolo.\n2) Modalità multigiocatore.\nInserisci 1 o 2")
	fmt.Scan(&mod)
	if mod == 1 {
		wordBuffer, err := os.ReadFile("./listaparola.txt")
		if err != nil {
			fmt.Println("C'è stato un errore uagliu")
			os.Exit(0) 
		}
		wordList := strings.Split(string(wordBuffer), "\n")
		randomIndex := rand.Intn(len(wordList))
		parola = wordList[randomIndex]
		fmt.Println("Perfetto, abbiamo generato una parola casuale, hai 10 vite a disposizione, Buona Fortuna!")
	} else if mod == 2 {
		var err error
		parola, err = getParola();
		if err != nil {
			fmt.Println("C'è stato un errore uagliu")
			os.Exit(0)	
		} 
		fmt.Println("Perfetto, Giocatore 2, hai 10 vite a disposizione, Buona Fortuna!")
	} else {
		fmt.Println("Sei un coglione, ti avevo detto di mettere o 1 o 2.")
		os.Exit(0)
	}
	
	tentativi := 10
	isComplete := false 
	parolaLen := len(parola)
	mappa = initialize(parolaLen)
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
	} else {
		fmt.Println("La parola era: ", parola)
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

	fmt.Print("Inserisci una lettera da azzeccare o la parola intera: ")
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

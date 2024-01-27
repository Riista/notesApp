package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func loadNotes(collectionName string) []string {
	notesFile := fmt.Sprintf("%s.txt", collectionName)
	var notes []string

	file, err := os.Open(notesFile)
	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			notes = append(notes, scanner.Text())
		}
	}

	return notes
}

func saveNotes(collectionName string, notes []string) {
	notesFile := fmt.Sprintf("%s.txt", collectionName)
	file, err := os.Create(notesFile)
	if err != nil {
		fmt.Println("Error creating notes file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, note := range notes {
		fmt.Fprintln(writer, note)
	}
	writer.Flush()
}

func displayNotes(notes []string) {
	if len(notes) == 0 {
		fmt.Println("Collection is empty, no notes available.")
	} else {
		for i, note := range notes {
			fmt.Printf("%03d - %s\n", i+1, note)
		}
	}
}

func addNote(notes []string, newNote string) []string {
	return append(notes, newNote)
}

func removeNote(notes []string, noteNumber int) []string {
	if noteNumber > 0 && noteNumber <= len(notes) {
		return append(notes[:noteNumber-1], notes[noteNumber:]...)
	}
	fmt.Println("Invalid note number.")
	return notes
}

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" {
		fmt.Println("Usage: ./notestool [COLLECTION_NAME]")
		return
	}

	collectionName := os.Args[1]
	notes := loadNotes(collectionName)

	fmt.Printf("Welcome to the Notes Tool!\n\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Select operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("\nNotes:")
			displayNotes(notes)
			fmt.Println("----------------------------")
		case "2":
			fmt.Print("\nEnter the note text:\n")
			scanner.Scan()
			newNote := scanner.Text()
			notes = addNote(notes, newNote)
			fmt.Println("\nNote added successfully.")
			fmt.Println("----------------------------")
		case "3":
			fmt.Print("\nEnter the number of note to remove or 0 to cancel:\n")
			scanner.Scan()
			noteNumber, _ := strconv.Atoi(scanner.Text())
			notes = removeNote(notes, noteNumber)
			fmt.Println("\nNote removed successfully.")
			fmt.Println("----------------------------")

		case "4":
			saveNotes(collectionName, notes)
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}
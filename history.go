package main

import (
	"bufio"
	"github.com/sbinet/liner"
	"log"
	"os"
	"os/user"
)

func LoadHistory(liner *liner.State, currentUser *user.User) {
	if currentUser != nil && currentUser.HomeDir != "" {
		ReadHistoryFromFile(liner, currentUser.HomeDir+"/.unql_history")
	}
}

func UpdateHistory(liner *liner.State, currentUser *user.User, line string) {
	liner.AppendHistory(line)
	if currentUser != nil && currentUser.HomeDir != "" {
		WriteHistoryToFile(liner, currentUser.HomeDir+"/.unql_history")
	}
}

func WriteHistoryToFile(liner *liner.State, path string) {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return
	}

	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = liner.WriteHistory(writer)
	if err != nil {
		log.Printf("Error updating .unql_history file: %v", err)
	} else {
		writer.Flush()
	}

}

func ReadHistoryFromFile(liner *liner.State, path string) {

	f, err := os.Open(path)
	if err != nil {
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	liner.ReadHistory(reader)
}

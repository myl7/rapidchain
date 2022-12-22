package main

import (
	"encoding/csv"
	"os"
	"strings"
)

// Loaded as global vars for easier tx gen interception
var txGenDatasetUsers []string
var txGenDatasetUserIndex map[string]int
var txGenDatasetTxs [][]string

const txGenDatasetUserFile = "users.txt"
const txGenDatasetTxFile = "txs.csv"

func loadTxGenDataset(nUser int) error {
	usersB, err := os.ReadFile(txGenDatasetUserFile)
	if err != nil {
		return err
	}

	usersS := strings.TrimSpace(string(usersB))
	users := strings.Split(usersS, "\n")

	txGenDatasetUsers = users
	for i := 0; i < nUser; i++ {
		txGenDatasetUserIndex[users[i]] = i
	}

	txsF, err := os.Open(txGenDatasetTxFile)
	if err != nil {
		return err
	}

	txsR := csv.NewReader(txsF)
	txs, err := txsR.ReadAll()
	if err != nil {
		return err
	}

	txGenDatasetTxs = txs
	return nil
}

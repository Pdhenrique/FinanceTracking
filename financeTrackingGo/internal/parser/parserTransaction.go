package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Pdhenrique/FinanceTracking/domain"
)

func Parse(reader io.Reader) ([]*domain.Transaction, error) {
	transactions := make([]*domain.Transaction, 0, 1024)
	scanner := bufio.NewScanner(reader)

	var currentTransaction domain.Transaction
	fmt.Println("PARSE INICIADO")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.Contains(line, "Agência:") {
			parts := strings.Split(line, "Conta:")
			currentTransaction.AGENCY = strings.TrimSpace(strings.Split(parts[0], ":")[1])
			currentTransaction.ACCOUNT_ID = strings.TrimSpace(parts[1])
			continue
		}

		if strings.HasPrefix(line, "ID,") {
			continue
		}

		columns := strings.Split(line, ",")
		if len(columns) < 8 {
			continue
		}

		// monta Transaction
		currentTransaction := &domain.Transaction{
			ID:              strings.TrimSpace(columns[0]),
			RELEASE_DATE:    strings.TrimSpace(columns[1]),
			ACCOUNTING_DATE: strings.TrimSpace(columns[2]),
			TITLE:           strings.TrimSpace(columns[3]),
			DESCRIPTION:     strings.TrimSpace(columns[4]),
		}

		// parse dos números (income, expense, balance)
		income, _ := strconv.ParseFloat(strings.ReplaceAll(columns[5], ",", "."), 64)
		expense, _ := strconv.ParseFloat(strings.ReplaceAll(columns[6], ",", "."), 64)
		balance, _ := strconv.ParseFloat(strings.ReplaceAll(columns[7], ",", "."), 64)

		currentTransaction.INCOME = income
		currentTransaction.EXPENSE = expense
		currentTransaction.DAILY_BALANCE = balance

		transactions = append(transactions, currentTransaction)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

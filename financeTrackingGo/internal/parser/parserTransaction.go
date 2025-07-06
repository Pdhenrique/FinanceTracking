package parser

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Pdhenrique/FinanceTracking/domain"
)

func Parse(reader io.Reader) ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s", line)
	}

	return transactions, nil
}

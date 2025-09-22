package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/Pdhenrique/FinanceTracking/domain"
	"github.com/google/uuid"
)

func Parse(r io.Reader) ([]*domain.Transaction, error) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var agency, accountID string

	txs := make([]*domain.Transaction, 0, 1024)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		// "Agência: X / Conta: Y"
		if strings.Contains(line, "Agência:") && strings.Contains(line, "Conta:") {
			parts := strings.SplitN(line, "Conta:", 2)
			if len(parts) == 2 {
				left := strings.TrimSpace(parts[0])
				right := strings.TrimSpace(parts[1])
				kv := strings.SplitN(left, ":", 2)
				if len(kv) == 2 {
					agency = strings.TrimSpace(strings.Replace(kv[1], "/", "", -1))
				}
				accountID = strings.TrimSpace(right)
			}
			continue
		}

		// pula header da grade
		up := strings.ToUpper(line)
		if strings.HasPrefix(up, "DATA LAN") {
			continue
		}

		cols := strings.Split(line, ",")
		if len(cols) < 7 {
			continue
		}

		income, _ := parseMoney(cols[4])
		expense, _ := parseMoney(cols[5])
		balance, _ := parseMoney(cols[6])
		releaseDate, err := time.Parse("02/01/2006", strings.TrimSpace(cols[0]))
		if err != nil {
			return nil, fmt.Errorf("erro ao converter release_date %s: %w", cols[0], err)
		}
		accountingDate, err := time.Parse("02/01/2006", strings.TrimSpace(cols[1]))
		if err != nil {
			return nil, fmt.Errorf("erro ao converter accounting_date %s: %w", cols[1], err)
		}

		txs = append(txs, &domain.Transaction{
			ID:              uuid.NewString(),
			AGENCY:          agency,
			ACCOUNT_ID:      accountID,
			RELEASE_DATE:    releaseDate.Format("2006-01-02"),
			ACCOUNTING_DATE: accountingDate.Format("2006-01-02"),
			TITLE:           strings.TrimSpace(cols[2]),
			DESCRIPTION:     strings.TrimSpace(cols[3]),
			INCOME:          income,
			EXPENSE:         expense,
			DAILY_BALANCE:   balance,
		})
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return txs, nil
}

func parseMoney(s string) (float64, error) {
	s = strings.TrimSpace(s)
	if strings.Contains(s, ",") {
		s = strings.ReplaceAll(s, ".", "")
		s = strings.ReplaceAll(s, ",", ".")
	}
	return strconv.ParseFloat(s, 64)
}

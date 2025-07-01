package loader

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func (t *LoaderTicketCSV) LoadTicketsToDB(ctx context.Context, db *sql.DB, csvPath string) (err error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("erro ao ler CSV: %w", err)
	}

	for i, record := range records[0:] {
		price, _ := strconv.ParseFloat(record[5], 64)
		_, err := db.ExecContext(ctx,
			`INSERT INTO tickets (name, email, country, hour, price) VALUES (?, ?, ?, ?, ?)`,
			record[1], record[2], record[3], record[4], price)
		if err != nil {
			return fmt.Errorf("erro ao inserir linha %d: %w", i+2, err)
		}
	}
	return nil
}

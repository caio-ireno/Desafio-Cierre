package loader

import (
	"app/internal"
	"context"
	"database/sql"
	"fmt"
)

func LoadTicketsToDB(ctx context.Context, db *sql.DB, records map[int]internal.Ticket) error {
	for id, record := range records {
		_, err := db.ExecContext(ctx,
			`INSERT INTO tickets (name, email, country, hour, price) VALUES (?, ?, ?, ?, ?)`,
			record.Attributes.Name,
			record.Attributes.Email,
			record.Attributes.Country,
			record.Attributes.Hour,
			record.Attributes.Price,
		)
		if err != nil {
			return fmt.Errorf("erro ao inserir ticket id %d: %w", id, err)
		}
	}
	return nil
}

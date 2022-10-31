package postgres

import (
	"evo_fintech/config"
	"evo_fintech/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(cfg *config.Config) (*Postgres, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s ",
		cfg.Postgres.DBName,
		cfg.Postgres.User,
		cfg.Postgres.Pass,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.SSLmode))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Postgres{db: db}, nil
}

func (p *Postgres) Insert(data *entity.Data) error {
	_, err := p.db.Exec("INSERT INTO data VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21);",
		data.TransactionId,
		data.RequestId,
		data.TerminalId,
		data.PartnerObjectId,
		data.AmountTotal,
		data.AmountOriginal,
		data.CommissionPS,
		data.CommissionClient,
		data.CommissionProvider,
		data.DateInput,
		data.DatePost,
		data.Status,
		data.PaymentType,
		data.PaymentNumber,
		data.ServiceId,
		data.Service,
		data.PayeeId,
		data.PayeeName,
		data.PayeeBankMfo,
		data.PayeeBankAccount,
		data.PaymentNarrative)
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) Select(filter string) ([]*entity.Data, error) {
	var dataSlice []*entity.Data
	err := p.db.Select(&dataSlice, fmt.Sprintf("SELECT * FROM data %s;", filter))
	if err != nil {
		return nil, err
	}
	return dataSlice, nil
}

package gorm

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"

	"github.com/rs/zerolog"

	env "sp-consumer/internal/env"

	models "sp-consumer/internal/gorm/models"
	"sp-consumer/internal/gorm/repositories"
)

type DBHandler struct {
	Db  *gorm.DB
	cfg *env.Config
	l   *zerolog.Logger

	OrderRepository repositories.OrderRepository
}

func NewDbHandler(cfg *env.Config, l *zerolog.Logger) *DBHandler {
	return &DBHandler{
		cfg: cfg,
		l:   l,
	}
}

// Init tries to connect to mysql and then creates gorm.DB
func (m *DBHandler) Init() error {
	// define connection string using config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?"+"parseTime=true&loc=Local",
		m.cfg.MySqlUsername,
		m.cfg.MySqlPassword,
		m.cfg.MySqlHost,
		m.cfg.MySqlPort,
		m.cfg.MySqlDbName,
	)

	m.l.Info().Msgf("dsn: ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		m.l.Error().Msgf("Error connecting to database: error=%v\n", err)
		return err
	}

	m.Db = db

	return nil
}

// InitRepositories fill repository variables in structure
func (m *DBHandler) InitRepositories() {
	m.OrderRepository = repositories.NewOrderRepository(m.Db)
}

// Migrate database and add new tables
func (m *DBHandler) Migrate() {
	m.Db.AutoMigrate(&models.Order{})
}

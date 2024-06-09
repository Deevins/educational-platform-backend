package pgmock

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"

	"github.com/spf13/viper"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	database "gitlab.ozon.dev/daker255/homework-8/pkg/database/clients"
)

type TDB struct {
	sync.Mutex
	DB *database.PostgresqlDatabase
}

func NewMokeDB() *TDB {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := initConfig(); err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	cfg := models.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Dbname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	fmt.Println(cfg)

	DB, err := database.NewDB(ctx, cfg)
	if err != nil {
		fmt.Println("err")
	}

	return &TDB{DB: DB}
}

func (b *TDB) SetUp(t *testing.T) {
	t.Helper()
	ctx := context.Background()

	b.Lock()
	b.Truncate(ctx)
}

func (b *TDB) TearDown() {
	defer b.Unlock()
	b.Truncate(context.Background())
}

func (b *TDB) Truncate(ctx context.Context) {
	tables := make([]string, 0)
	err := b.DB.Select(ctx, &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name!= 'goose_db_version'")
	if err != nil {
		log.Println(err)
	}
	if len(tables) == 0 {
		fmt.Println("run migrations")
	}

	q := fmt.Sprintf("Truncate table %s RESTART IDENTITY", strings.Join(tables, ","))
	if _, err := b.DB.Exec(ctx, q); err != nil {
		log.Fatalf("qwe")
	}
}

func initConfig() error {
	viper.SetConfigType("yml")
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

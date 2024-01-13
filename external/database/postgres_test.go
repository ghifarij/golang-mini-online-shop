package database

import (
	"testing"

	"github.com/ghifarij/golang-mini-online-shop/internal/config"
	"github.com/stretchr/testify/require"
)

func Init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, err := ConnectPostgres(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})
}

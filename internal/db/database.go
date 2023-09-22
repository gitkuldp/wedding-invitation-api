package db

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDsn(dialect string, user string, password string, host string, port string, db string) (dsn string) {
	return fmt.Sprintf(`%s://%s:%s@%s:%s/%s`, dialect, user, password, host, port, db)
}
func InitDB(env *Env) *gorm.DB {
	// create dsn
	dsn := NewDsn(env.DBDialect, env.PostgresUser, env.PostgresPassword, env.DBHost, env.DBPort, env.PostgresDB)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Kathmandu")
			return time.Now().In(ti)
		},
	})

	if err != nil {
		logrus.Error(err)
		return nil
	}
	return db
}

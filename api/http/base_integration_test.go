//+build integration

package http

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/volodimyr/Symphony/app"
	"github.com/volodimyr/Symphony/models"

	_ "github.com/lib/pq"
)

var (
	postgresDBTest *sql.DB
	//redisClientTest *redis.Client
)

func TestMain(m *testing.M) {
	var err error

	postgresDBTest, err = app.ConnectPostgresDB(app.PostgresDBConfig{
		Host: "localhost",
		Port: "5432",
		Name: "symphony-rent",
		Usr:  "postgres",
		Pwd:  "postgres",
	})

	if err != nil {
		panic("failed to connect to test postgres db")
	}

	//redisClientTest, err = app.NewRedisClient(app.RedisStoreConfig{
	//	Addr: "localhost",
	//	Port: "6379",
	//})
	//if err != nil {
	//	panic("failed to connect to test redis storage")
	//}

	os.Exit(m.Run())
}

func migrateTestData(t *testing.T, subDir string) {
	migrateEmptyTestData(t)

	fixtr, err := testfixtures.New(
		testfixtures.Database(postgresDBTest),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("testdata/"+subDir),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		t.Fatalf("failed to create test fixtures due to %v", err)
	}

	if err := fixtr.Load(); err != nil {
		t.Fatalf("failed to load fixtures due to %v", err)
	}
}

func migrateEmptyTestData(t *testing.T) {
	for _, table := range []string{models.TableNames.CarOrders, models.TableNames.Cars} {

		if _, err := postgresDBTest.Exec(`TRUNCATE TABLE ` + table + ` CASCADE`); err != nil {
			t.Fatalf("failed to truncate test table=%s due ot %v", table, err)
		}
	}
}

//func migrateTestData(t *testing.T, subDir string) {
//	//truncate all table
//	migrateEmptyTestData(t)
//
//	fixtr, err := testfixtures.New(
//		testfixtures.Database(migrConn),
//		testfixtures.Dialect("postgres"),
//		testfixtures.Directory("testdata/"+subDir),
//		testfixtures.UseAlterConstraint(),
//	)
//	if err != nil {
//		t.Fatalf("failed to create test fixtures due to %v", err)
//	}
//
//	if err := fixtr.Load(); err != nil {
//		t.Fatalf("failed to load fixtures due to %v", err)
//	}
//}
//
//func migrateEmptyTestData(t *testing.T) {
//	for _, table := range domain.Tables {
//		_, err := migrConn.Exec(`TRUNCATE TABLE ` + table + ` CASCADE`)
//		if err != nil {
//			t.Fatalf("failed to truncate test table=%s due ot %v", table, err)
//		}
//	}
//}
//
//func getEnvVariable(key, fallback string) string {
//	if value, ok := os.LookupEnv(key); ok {
//		return value
//	}
//	return fallback
//}

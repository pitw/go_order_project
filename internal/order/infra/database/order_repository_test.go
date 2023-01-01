package database

import (
	"database/sql"
	"testing"

	"github.com/axnd/goaxnd/internal/order/entity"
	"github.com/stretchr/testify/suite"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	// adds several properties
	suite.Suite
	Db *sql.DB
}

// executes before running a test
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders(id varchar(255) NOT NULL,price float NOT NULL,tax foat NOT NULL,final_price NOT NULL,PRIMARY KEY(id))")
	suite.Db = db
}

// executes after running a test
func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("select id, price, tax, final_price from orders where id = ?", order.ID).Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

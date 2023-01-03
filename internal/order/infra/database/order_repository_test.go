package database

import (
	"database/sql"
	"testing"

	"github.com/axnd/goaxnd/internal/order/entity"
	"github.com/axnd/goaxnd/shared"
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
/* func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders(id varchar(255) NOT NULL,price float NOT NULL,tax foat NOT NULL,final_price NOT NULL,PRIMARY KEY(id))")
	suite.Db = db
} */

// executes after running a test
/* func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
} */

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	shared.Connect()

	//order, err := entity.NewOrder("123", 10.0, 2.0)
	newOrder := entity.Order{
		Price: 23,
		Tax:   6,
	}
	db := shared.SQL.Model(&entity.Order{}).Create(&newOrder)

	suite.NoError(db.Error)

	suite.NoError(newOrder.CheckFinalPrice())
	//repo := NewOrderRepository(suite.Db)
	//err = repo.Save(order)
	//suite.NoError(err)

	var orderResult entity.Order
	dbQuery := shared.SQL.Model(&entity.Order{}).Where("id = ?", newOrder.ID).First(&orderResult)
	dbQueryAlternative := shared.SQL.Model(&entity.Order{}).First(&orderResult, newOrder.ID)

	//err = suite.Db.QueryRow("select id, price, tax, final_price from orders where id = ?", order.ID).Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(dbQuery.Error)
	suite.NoError(dbQueryAlternative.Error)

	suite.Equal(newOrder.ID, orderResult.ID)
	suite.Equal(newOrder.Tax, orderResult.Tax)
	suite.Equal(newOrder.Price, orderResult.Price)
	suite.Equal(newOrder.FinalPrice, orderResult.FinalPrice)
}

package repos

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"math"
	"nft_auction/pkg/models"
	"runtime/debug"
	"time"
)

const (
	generalQueryTimeout = 60 * time.Second
	defaultPageSize     = 30
	maxPageSize         = 1000
)

func NewPGRepo(db *gorm.DB) PGInterface {
	return &RepoPG{
		db: db,
	}
}

type PGInterface interface {
	// Database
	Transaction(ctx context.Context, f func(rp PGInterface) error) (err error)

	// User
	LoginUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserProfile(ctx context.Context, id string) (*models.User, error)

	// Collections
	CreateCollection(ctx context.Context, req *models.Collection) (*models.Collection, error)
	GetCollection(ctx context.Context, cid *uuid.UUID) (*models.Collection, error)
	QueryCollections(ctx context.Context, req *models.QueryCollectionReq) (*models.QueryCollectionRes, error)
	UpdateCollection(ctx context.Context, req *models.Collection) (*models.Collection, error)
	DeleteCollection(ctx context.Context, cid *uuid.UUID) error

	// Item
	CreateItem(ctx context.Context, req *models.Item) (*models.Item, error)
	GetItem(ctx context.Context, id *uuid.UUID) (*models.Item, error)
	Like(ctx context.Context, req *models.ItemLike) error
	QueryItems(ctx context.Context, req *models.QueryItemReq) (*models.QueryItemRes, error)
	UpdateItem(ctx context.Context, req *models.Item) (*models.Item, error)
	DeleteItem(ctx context.Context, cid *uuid.UUID) error

	// Order
	CreateOrder(ctx context.Context, req *models.Order) (*models.Order, error)
	GetOrder(ctx context.Context, id *uuid.UUID) (*models.Order, error)
	QueryOrders(ctx context.Context, req *models.QueryOrderReq) (*models.QueryOrderRes, error)
	UpdateOrder(ctx context.Context, req *models.Order) (*models.Order, error)
	DeleteOrder(ctx context.Context, cid *uuid.UUID) error

	CreateSale(ctx context.Context, req *models.Sale) (*models.Sale, error)
	GetSale(ctx context.Context, id *uuid.UUID) (*models.Sale, error)
	QuerySales(ctx context.Context, req *models.QuerySaleReq) (*models.QuerySaleRes, error)
	UpdateSale(ctx context.Context, req *models.Sale) (*models.Sale, error)
	DeleteSale(ctx context.Context, cid *uuid.UUID) error

	DB() *gorm.DB
}

type RepoPG struct {
	db *gorm.DB
}

func (r *RepoPG) Transaction(ctx context.Context, f func(rp PGInterface) error) (err error) {
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	// create new instance to run the transaction
	repo := *r
	tx = tx.Begin()
	repo.db = tx
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = errors.New(fmt.Sprint(r))
			log.Println("error_500: Panic when run Transaction err=" + err.Error())
			debug.PrintStack()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	err = f(&repo)
	if err != nil {
		log.Println("error_500: Error when run Transaction err=" + err.Error())
		return err
	}
	return nil
}

func (r *RepoPG) DB() *gorm.DB {
	return r.db
}

func (r *RepoPG) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return r.db.WithContext(ctx), cancel
}

func (r *RepoPG) GetPage(page int) int {
	if page <= 0 {
		return 1
	}
	return page
}

func (r *RepoPG) GetOffset(page int, pageSize int) int {
	return (page - 1) * pageSize
}

func (r *RepoPG) GetPageSize(pageSize int) int {
	if pageSize <= 0 {
		return defaultPageSize
	}
	if pageSize > maxPageSize {
		return maxPageSize
	}
	return pageSize
}

func (r *RepoPG) GetTotalPages(totalRows int64, pageSize int) int {
	return int(math.Ceil(float64(totalRows) / float64(pageSize)))
}

func (r *RepoPG) GetPaginationInfo(totalRow int64, page, pageSize int) map[string]interface{} {
	return map[string]interface{}{
		"page":        page,
		"page_size":   pageSize,
		"total_pages": r.GetTotalPages(totalRow, pageSize),
		"total_rows":  totalRow,
	}
}

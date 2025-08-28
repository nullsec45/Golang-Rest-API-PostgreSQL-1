package service

import (
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/config"
	"context"
	"github.com/google/uuid"
	"database/sql"
	"time"
	"path"
    // "fmt"
)

type BookService struct {
	config *config.Config
	bookRepository domain.BookRepository
	bookStockRepository domain.BookStockRepository
	mediaRepository domain.MediaRepository
}

func NewBook(
	config *config.Config,
	bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository,
	mediaRepository domain.MediaRepository,
) domain.BookService {
	return &BookService{
		config:config,
		bookRepository: bookRepository,
		bookStockRepository: bookStockRepository,
		mediaRepository: mediaRepository,
	}
}

func (bs BookService) Index(ctx context.Context)  ([]dto.BookData, error) {
	books, err := bs.bookRepository.FindAll(ctx)
	if err != nil {
		// Handle error (e.g., log it, return an error response)
		return nil, err
	}

	coverId := make([]string, 0)
	for _, v := range books {
		if v.CoverId.Valid {
			coverId = append(coverId, v.CoverId.String)
		}
	}

	covers := make(map[string]string)
	if len(coverId) > 0 {
		media, _ := bs.mediaRepository.FindByIds(ctx, coverId)
		
		for _, v := range media {
			covers[v.Id] = path.Join(bs.config.Server.Asset, v.Path)	
		}
	}

	var bookData []dto.BookData

	for _, v := range books {
		var coverUrl string

		if v2, e := covers[v.CoverId.String]; e {
			coverUrl = v2
		}

		bookData = append(bookData, dto.BookData{
			Id:   v.Id,
			Isbn: v.Isbn,
			Title: v.Title,
			CoverUrl: coverUrl,
			Description: v.Description,
		})
	}
	
	return bookData, nil
}

func (bs BookService) Show (ctx context.Context, id string) (dto.BookShowData, error) {
    data, err := bs.bookRepository.FindById(ctx,id)

    if err != nil && data.Id == "" {
        return dto.BookShowData{}, domain.BookNotFound
    }
    
    if err != nil {
        return dto.BookShowData{}, err
    }

	stocks, err := bs.bookStockRepository.FindByBookId(ctx, data.Id)

	if err != nil {
		return dto.BookShowData{}, err
	}
	
	stocksData := make([]dto.BookStockData, 0)
	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code:v.Code,
			Status:v.Status,
		})
	}

	var coverUrl string

	if data.CoverId.Valid {
		cover, _ := bs.mediaRepository.FindById(ctx, data.CoverId.String)

		if cover.Path != "" {
			coverUrl=path.Join(bs.config.Server.Asset, cover.Path)	
		}
	}


    return dto.BookShowData{
		BookData:dto.BookData{
			Id:   data.Id,
			Isbn: data.Isbn,
			Title: data.Title,
			CoverUrl: coverUrl,
			Description: data.Description,
		},
		Stocks:stocksData,
    }, nil
}

func (bs BookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	coverId := sql.NullString{Valid:false, String:req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

    book := domain.Book{
        Id: uuid.New().String(),
       	Isbn: req.Isbn,
		Title: req.Title,
		Description: req.Description,
		CreatedAt:sql.NullTime{Valid:true, Time:time.Now()},	
    }

	return bs.bookRepository.Save(ctx, &book)
}

func (bs BookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
    // Cari data book
    exist, err := bs.bookRepository.FindById(ctx, req.Id)

    // Jika book tidak ditemukan
    if err != nil && exist.Id == "" {
        return  domain.BookNotFound
    }
    
    if err != nil {
        return  err
    }

	coverId := sql.NullString{Valid:false, String:req.CoverId}
	if req.CoverId != "" {
		coverId.Valid = true
	}

	exist.Isbn=req.Isbn
	exist.Title=req.Title
	exist.Description=req.Description
	exist.UpdatedAt=sql.NullTime{Valid:true, Time:time.Now()}
	exist.CoverId=coverId

    return bs.bookRepository.Update(ctx, &exist)
}

func (bs BookService) Delete (ctx context.Context, id string) error {
    exist, err := bs.bookRepository.FindById(ctx,id)

    if err != nil && exist.Id == "" {
        return  domain.BookNotFound
    }
    
    if err != nil {
        return err
    }

	err = bs.bookRepository.Delete(ctx,exist.Id)

	if err != nil {
		return err
	}

    return bs.bookStockRepository.DeleteByBookId(ctx, exist.Id)
}
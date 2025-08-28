package service


import (
	"context"
	"time"
    "github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/domain"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/dto"
	"github.com/nullsec45/Golang-Rest-API-PostgreSQL-1/internal/config"
	"github.com/google/uuid"
	"database/sql"
	"path"
)

type MediaService struct {
	config *config.Config
	mediaRepository domain.MediaRepository
}

func NewMedia(config *config.Config, mediaRepository domain.MediaRepository) domain.MediaService {
	return &MediaService{
		config:config,
		mediaRepository:mediaRepository,
	}
}

func (m MediaService) Create(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error) {
	media := domain.Media{
		Id: uuid.NewString(),	
		Path:req.Path,
		CreatedAt:sql.NullTime{Time:time.Now(), Valid:true},
	}

	err := m.mediaRepository.Save(ctx, &media)
	if err != nil {
		return dto.MediaData{}, err
	}

	url := path.Join(m.config.Server.Asset, media.Path)	
	return dto.MediaData{
		Id: media.Id,	
		Path:media.Path,
		Url:url,	
	}, nil
}

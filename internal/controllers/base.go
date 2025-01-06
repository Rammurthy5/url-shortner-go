package controllers

import (
	"github.com/Rammurthy5/url-shortner-go/config"
	urls_mapping "github.com/Rammurthy5/url-shortner-go/internal/db/sqlc"
	"go.uber.org/zap"
	"net/http"
)

type Controller interface {
	ServeHandle(w http.ResponseWriter, r *http.Request)
}
type BaseController struct {
	Cfg config.Config
	Log *zap.Logger
	Db  *urls_mapping.Queries
}

func NewBaseController(cfg *config.Config, log *zap.Logger, db *urls_mapping.Queries) *BaseController {
	return &BaseController{}
}

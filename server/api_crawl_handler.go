package server

import "github.com/seesawlabs/crawlfish/shared/database"

type ApiV1CrawlHandler struct {
	Db database.IDatabase
}

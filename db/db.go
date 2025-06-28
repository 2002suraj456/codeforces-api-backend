package db

import (
	"sync"

	"cf.practice.com/models"
)

var (
	instance models.ProblemDBInterface
	once     sync.Once
)

// GetInstance returns a singleton instance of the problem database
func GetInstance() models.ProblemDBInterface {
	once.Do(func() {
		db := &models.ProblemDB{}
		db.Init()
		instance = db
	})
	return instance
}

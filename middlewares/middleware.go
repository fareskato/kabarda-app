package middlewares

import (
	"github.com/fareskato/kabarda"
	"myapp/data"
)

type Middleware struct {
	App    *kabarda.Kabarda
	Models data.Models
}

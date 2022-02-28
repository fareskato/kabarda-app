package middlewares

import (
	"myapp/data"
)

type Middleware struct {
	App    *kabarda.Kabarda
	Models data.Models
}

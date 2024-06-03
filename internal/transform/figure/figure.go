package figure

import (
	"transformMDLink/internal/transform/config"
)

type figure interface {
	upload(path *string)
}

var Figures figure

func Init(config *config.Config) {
	token := config.Token()
	ih := config.Ih()
	if token == "" {
		// TODO 添加日志
	}
	if ih == "" {
		ih = "0"
	}
	switch ih {
	case "0":
		Figures = &imageTP{token: token}
	}
}

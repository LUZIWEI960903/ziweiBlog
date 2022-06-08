package comment

import (
	"ziweiBlog/config"
	"ziweiBlog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
}

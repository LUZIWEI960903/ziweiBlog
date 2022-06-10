package comment

import (
	"sync"
	"ziweiBlog/config"
	"ziweiBlog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}

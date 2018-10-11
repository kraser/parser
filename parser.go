// parser
package parser

import (
	"webreader"
	//"fmt"
	log "logger"
	"math/rand"
	"os"
	"os/signal"
	"priceloader"
	"time"
)

const (
	ENDMESSAGE = "PARSE_DONE"
)

type InterfaceCustomParser interface {
	ParseCategories(param string)
	ParseItems()
	ParserInit(*ParserObject)
	ParserRun()
}

type ParserOptions struct {
	Name           string
	URL            string
	Loaders        int
	LoaderCapacity int
}

type ParserObject struct {
	Options             *webreader.RequestOptions
	CustomParserOptions *ParserOptions
	CustomParserActions InterfaceCustomParser
}

func (pParser *ParserObject) init() {
	log.Info("PARSER_INIT_START")
	priceloader.PriceList.PriceList(pParser.CustomParserOptions.Name)
	rand.Seed(time.Now().UnixNano())
	pParser.Options = webreader.GetOptions()
	pParser.SetUserAgent(useragents[rand.Intn(len(useragents))])

	pParser.Options.Url = pParser.CustomParserOptions.URL
	pParser.CustomParserActions.ParserInit(pParser)
	log.Info("PARSER_INIT_DONE")
}

func (pParser *ParserObject) SetUserAgent(ua string) {
	log.Debug("UA:", ua)
	pParser.Options.UserAgent = ua
}

func (pParser *ParserObject) Run() {
	pParser.init()
	/*
		//Подготовим каналы и регулятор
		taskChan := make(chan priceloader.LoadTask)
		quitChan := make(chan bool)
		pController := &LoadController{Loaders: pParser.CustomParserOptions.Loaders, LoaderCapacity: pParser.CustomParserOptions.LoaderCapacity}
		pController.init(taskChan)

		//Приготовимся перехватывать сигнал останова в канал keys
		keys := make(chan os.Signal, 1)
		signal.Notify(keys, os.Interrupt)

		//go pController.balance(quitChan)
		//go pParser.taskGenerator(taskChan)

		log.Info("MAIN_CYCLE_START")
		//Основной цикл программы:
		for {
			select {
			case <-keys: //пришла информация от нотификатора сигналов:
				log.Info("CTRL-C: Ожидаю завершения активных загрузок")
				quitChan <- true //посылаем сигнал останова балансировщику

			case <-quitChan: //пришло подтверждение о завершении от балансировщика
				log.Info("MAIN_CYCLE_END")
				return
			}
		}
	*/
	//	result := webreader.DoRequest(pParser.CustomParserOptions.URL, pParser.Options)
	//	pParser.CustomParserActions.ParseCategories(result)
	//	pParser.CustomParserActions.ParseItems()
}

/*
func (pParser *ParserObject) taskGenerator(out chan priceloader.LoadTask) {
	pPriceList := priceloader.PriceList
	for _, value := range pPriceList.Categories {
		log.Info(value.Name)
		for _, subCat := range value.Categories {
			subCat.URL = URL + subCat.URL
			log.Info("  ", subCat.Name, subCat.URL)
			task := priceloader.LoadTask{Pointer: subCat, Message: "TASK"}
			out <- task
		}
	}
	endTask := priceloader.LoadTask{Pointer: nil, Message: ENDMESSAGE}
	out <- endTask
}
*/
func LoadAndParse(itemLoadTask priceloader.LoadTask) {

}

package main

import (
	"log"
	"sync"

	"printraduga_parser/excel"
	parsers "printraduga_parser/parsers"
	"printraduga_parser/shared"
)

func main() {
	parsers := []shared.Parser{
		parsers.DigitalTranslusentParser{},
		parsers.CoralTranslusentParser{},
		parsers.GcTranslusentParser{},
	}
	var wg sync.WaitGroup
	var resultMutex sync.Mutex
	results := make(map[string][]shared.CostData)
	for _, parser := range parsers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Printf("Parser in process: %T\n", parser)
			result, err := parser.Parse()
			if err != nil {
				log.Printf("Error while parsing: %T; Error: %v", parser, err)
				return
			}
			resultMutex.Lock()
			if results[result.ParserType] == nil {
				results[result.ParserType] = make([]shared.CostData, 0)
			}
			results[result.ParserType] = append(results[result.ParserType], result.Data)
			resultMutex.Unlock()
		}()
	}
	wg.Wait()
	var writer shared.ExcelWriter = excel.DefaultExcelWriter{}
	err := writer.Write("parsing_result.xlsx", results)
	if err != nil {
		log.Printf("Writing error: %v", err)
	}
}

//  NOTE: Test code

// func main() {
// 	parser := parsers.GcTranslusentParser{}
// 	data, err := parser.Parse()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Print(data)
//
// }

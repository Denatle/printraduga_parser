package main

import (
	"log"
	"printraduga_parser/excel"
	parsers "printraduga_parser/parsers"
	"printraduga_parser/shared"
	"sync"
)

func main() {
	parsers := []shared.Parser{
		parsers.DigitalTranslusentParser{},
		parsers.CoralTranslusentParser{},
	}
	var wg sync.WaitGroup
	var resultMutex sync.Mutex
	var results []shared.CostInfo
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
			results = append(results, result)
			resultMutex.Unlock()
		}()
	}
	wg.Wait()
	var writer shared.ExcelWriter = excel.DefaultExcelWriter{}
	err := writer.Write("balls.xlsx", results)
	if err != nil {
		log.Printf("Writing error: %v", err)
	}
}

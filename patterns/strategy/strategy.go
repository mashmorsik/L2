package main

import (
	"fmt"
	"time"
)

type Report struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"dateCreated"`
	Type        int       `json:"type"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Status      int       `json:"status"`
}

type ReportStrategy interface {
	GenerateReport(report *Report) error
}

// CSVReport - стратегия генерации отчета в формате CSV
type CSVReport struct{}

func (c *CSVReport) GenerateReport(report *Report) error {
	fmt.Printf("Generating CSV file from report: %v", report)
	return nil
}

// PDFReport - стратегия генерации отчета в формате PDF
type PDFReport struct{}

func (c *PDFReport) GenerateReport(report *Report) error {
	fmt.Printf("Generating PDF file from report: %v", report)
	return nil
}

// HTMLReport - стратегия генерации отчета в формате HTML
type HTMLReport struct{}

func (c *HTMLReport) GenerateReport(report *Report) error {
	fmt.Printf("Generating HTML file from report: %v", report)
	return nil
}

// ReportContext - объект, который использует стратегии
type ReportContext struct {
	Strategy ReportStrategy
}

// SetStrategy - метод для установки стратегии
func (s *ReportContext) SetStrategy(strategy ReportStrategy) {
	s.Strategy = strategy
}

// Execute - метод для выполнения стратегии
func (s *ReportContext) Execute(report *Report) error {
	return s.Strategy.GenerateReport(report)
}

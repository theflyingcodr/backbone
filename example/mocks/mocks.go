// Package mocks requires the github.com/matryer/moq lib to be installed.
package mocks

//go:generate moq -pkg mocks -out thing_reader_writer.go ../ ThingReaderWriter
//go:generate moq -pkg mocks -out thing_service.go ../ ThingService

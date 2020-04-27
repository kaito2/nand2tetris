package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Aggregator struct {
	parsers []Parser
	// outputFilepath string
}

func NewAggregator(path string) (Aggregator, error) {
	filePaths, err := getVMFilenames(path)
	if err != nil {
		return Aggregator{}, fmt.Errorf("failed to getVMFilenames: %w", err)
	}
	log.Printf(".vm files: %v\n", filePaths)

	var parsers []Parser
	for _, filePath := range filePaths {
		parser, err := NewParser(filePath)
		if err != nil {
			return Aggregator{}, fmt.Errorf("failed to get NewParser: %w", err)
		}
		parsers = append(parsers, parser)
	}

	return Aggregator{
		parsers: parsers,
	}, nil
}

func (a Aggregator) ParseAll(outputFilePath string) error {
	_, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to os.Create: %w", err)
	}

	for _, parser := range a.parsers {
		err := parser.Parse(outputFilePath)
		if err != nil {
			return fmt.Errorf("failed to parse: %w", err)
		}
	}
	return nil
}

func getVMFilenames(path string) (filenames []string, err error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to os.Stat: %w", err)
	}
	switch mode := stat.Mode(); {
	case mode.IsDir():
		// ディレクトリの場合はディレクトリ内の .vm ファイルを返す
		filenames, err = getVMFilenamesFromDir(path)
		if err != nil {
			return nil, fmt.Errorf("failed to getFilenamesFromDir: %w", err)
		}
		return
	default: // case mode.IsRegular():
		// ファイルの場合はそのパスを直接 対象ファイルとして返す
		if isVMFile(path) {
			filenames = append(filenames, path)
		}
		return
	}
}

func getVMFilenamesFromDir(dirPath string) (filenames []string, err error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to ioutil.ReadDir: %w", err)
	}
	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		if !isVMFile(filePath) {
			continue
		}
		filenames = append(filenames, filePath)
	}
	return
}

func isVMFile(filepath string) bool {
	// See https://qiita.com/takashi/items/25900c31c5ad73276b76
	ext := filepath[strings.LastIndex(filepath, "."):]
	return ext == ".vm"
}

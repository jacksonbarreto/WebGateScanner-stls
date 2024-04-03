package scanner

import (
	"github.com/jacksonbarreto/WebGateScanner-DNSSECAnalyzer/pkg/logservice"
	"github.com/jacksonbarreto/WebGateScanner-stls/config"
	"os"
	"os/exec"
)

type Scanner struct {
	pathToResults        string
	readyToProcessSuffix string
	processFileExtension string
	log                  logservice.Logger
}

func NewScanner(pathToResults string, logService logservice.Logger, readyToProcessSuffix string,
	processFileExtension string) *Scanner {
	return &Scanner{
		pathToResults:        pathToResults,
		readyToProcessSuffix: readyToProcessSuffix,
		processFileExtension: processFileExtension,
		log:                  logService,
	}
}

func NewScannerDefault(pathToResults string) *Scanner {
	return NewScanner(pathToResults, logservice.NewLogService(config.App().Id),
		config.App().ReadyToProcessSuffix, config.App().ProcessFileExtension)
}

func (s *Scanner) Scan(host string) error {
	fileName := s.pathToResults + host + s.processFileExtension
	fileNameDone := fileName + s.readyToProcessSuffix
	s.log.Info("Scanning host '%s'...", host)
	cmd := exec.Command("/home/stls/testssl.sh/testssl.sh", "--jsonfile-pretty", fileName, host)

	err := cmd.Run()
	if err != nil {
		return err
	}

	file, err := os.Create(fileNameDone)
	if err != nil {
		s.log.Error("Error creating file '%s': %v", fileNameDone, err)
	}
	defer file.Close()
	return nil
}

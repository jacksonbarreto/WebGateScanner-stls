package scanner

import (
	"github.com/jacksonbarreto/WebGateScanner-DNSSECAnalyzer/pkg/logservice"
	"os"
	"os/exec"
)

type Scanner struct {
	PathToResults string
	Log           logservice.Logger
}

func NewScanner(pathToResults string, logService logservice.Logger) *Scanner {
	return &Scanner{
		PathToResults: pathToResults,
		Log:           logService,
	}
}

func NewScannerDefault(pathToResults string) *Scanner {
	return NewScanner(pathToResults, logservice.NewLogService("scanner"))

}

func (s *Scanner) Scan(host string) error {
	fileName := s.PathToResults + host + ".json"
	fileNameDone := fileName + ".done"
	s.Log.Info("Scanning host '%s'...", host)
	cmd := exec.Command("/home/stls/testssl.sh/testssl.sh", "--jsonfile-pretty", fileName, host)

	err := cmd.Run()
	if err != nil {
		return err
	}

	file, err := os.Create(fileNameDone)
	if err != nil {
		s.Log.Error("Error creating file '%s': %v", fileNameDone, err)
	}
	defer file.Close()
	return nil
}

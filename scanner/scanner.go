package scanner

import "os/exec"

type Scanner struct {
	PathToResults string
}

func NewScanner(pathToResults string) *Scanner {
	return &Scanner{
		PathToResults: pathToResults,
	}
}

func (s *Scanner) Scan(host string) error {

	cmd := exec.Command("/home/stls/testssl.sh/testssl.sh", "--jsonfile-pretty", s.PathToResults, host)

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

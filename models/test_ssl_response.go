package models

type TestSSLResponse struct {
	ClientProblem1 []ClientProblem1 `json:"clientProblem1"`
	Invocation     string           `json:"Invocation"`
	At             string           `json:"at"`
	Version        string           `json:"version"`
	Openssl        string           `json:"openssl"`
	StartTime      string           `json:"startTime"`
	ScanResult     []ScanResult     `json:"scanResult"`
	ScanTime       int              `json:"scanTime"`
}
type ClientProblem1 struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
}
type Pretest struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
}
type Protocols struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
}
type Ciphers struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Cwe      string `json:"cwe,omitempty"`
	Finding  string `json:"finding"`
}
type ServerPreferences struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
}
type Fs struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
}
type ServerDefaults struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
}
type HeaderResponse struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
	Cwe      string `json:"cwe,omitempty"`
}
type Vulnerabilities struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Cve      string `json:"cve,omitempty"`
	Cwe      string `json:"cwe,omitempty"`
	Finding  string `json:"finding"`
}
type BrowserSimulations struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
}
type Rating struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Finding  string `json:"finding"`
}
type ScanResult struct {
	TargetHost         string               `json:"targetHost"`
	IP                 string               `json:"ip"`
	Port               string               `json:"port"`
	RDNS               string               `json:"rDNS"`
	Service            string               `json:"service"`
	Pretest            []Pretest            `json:"pretest"`
	Protocols          []Protocols          `json:"protocols"`
	Grease             []any                `json:"grease"`
	Ciphers            []Ciphers            `json:"ciphers"`
	ServerPreferences  []ServerPreferences  `json:"serverPreferences"`
	Fs                 []Fs                 `json:"fs"`
	ServerDefaults     []ServerDefaults     `json:"serverDefaults"`
	HeaderResponse     []HeaderResponse     `json:"headerResponse"`
	Vulnerabilities    []Vulnerabilities    `json:"vulnerabilities"`
	CipherTests        []any                `json:"cipherTests"`
	BrowserSimulations []BrowserSimulations `json:"browserSimulations"`
	Rating             []Rating             `json:"rating"`
}

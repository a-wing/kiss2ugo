package config

import (
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	opts *Options
}

func NewParser() *Parser {
	return &Parser{
		opts: NewOptions(),
	}
}

func (p *Parser) ParseEnvironmentVariables() (*Options, error) {
	err := p.parseLines(os.Environ())
	if err != nil {
		return nil, err
	}
	return p.opts, nil
}

func (p *Parser) parseLines(lines []string) (err error) {
	var port string
	for _, line := range lines {
		fields := strings.SplitN(line, "=", 2)
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		switch key {
		case "PORT":
			port = value
		case "LISTEN_ADDR":
			p.opts.listenAddr = parseString(value, defaultListenAddr)
		case "DATABASE_URL":
			p.opts.databaseURL = parseString(value, defaultDatabaseURL)
		case "LILAC_LOG":
			p.opts.lilacLog = parseString(value, defaultLilacLog)
		case "LILAC_REPO":
			p.opts.lilacRepo = parseString(value, defaultLilacRepo)
		}
	}

	if port != "" {
		p.opts.listenAddr = ":" + port
	}
	return nil
}

func parseBool(value string, fallback bool) bool {
	if value == "" {
		return fallback
	}

	value = strings.ToLower(value)
	if value == "1" || value == "yes" || value == "true" || value == "on" {
		return true
	}

	return false
}

func parseInt(value string, fallback int) int {
	if value == "" {
		return fallback
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return v
}

func parseString(value string, fallback string) string {
	if value == "" {
		return fallback
	}

	return value
}

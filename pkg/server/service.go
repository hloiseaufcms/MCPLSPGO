package server

// This file provides the expected API for main.go

// NewService creates a new MCP service for gopls integration
func NewService() (*Service, error) {
	// We're reusing the existing Service struct and its creation logic
	// from server.go
	logFile, err := setupLogger()
	if err != nil {
		return nil, err
	}

	svc := &Service{
		logFile: logFile,
	}

	if err := svc.initLSPClient(); err != nil {
		if logFile != nil {
			logFile.Close()
		}
		return nil, err
	}

	svc.server = setupServer()
	return svc, nil
}
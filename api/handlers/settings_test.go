package handlers

import "github.com/stretchr/testify/mock"

// func setupSettingHandler(t *testing.T) (Auth, *authmocks.MockService) {
// 	mockSvc := authmocks.NewMockService(t)
// 	jwtCfg := setupJWTConfig()
// 	handler := NewAuth(mockSvc, jwtCfg)
// 	return handler, mockSvc
// }

type mockSettingsService struct {
	mock.Mock
}

//TODO: complete this

package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/FabricioCosati/onfly-test/internal/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type LogTest struct {
	Time      string `json:"time"`
	Level     string `json:"level"`
	Msg       string `json:"msg"`
	Timestamp string `json:"Timestamp"`
	Method    string `json:"Method"`
	Request   string `json:"Request"`
	Response  string `json:"Response"`
	Path      string `json:"Path"`
	Status    int    `json:"Status"`
}

func TestOperationLogsSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	os.Setenv("LOGS_DIR", "logs_test")
	defer os.RemoveAll("logs_test")

	now := time.Now()
	date := now.Format("2006_01_02")
	name := fmt.Sprintf("logs_test/operation_log_%s.log", date)

	engine := gin.New()
	engine.Use(middlewares.OperationLogs("operation"))
	engine.GET("/operation-log", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "success")
	})

	request, err := http.NewRequest("GET", "/operation-log", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, request)

	log, err := os.ReadFile(name)
	assert.NoError(t, err)

	var expectedLog LogTest

	err = json.Unmarshal(log, &expectedLog)
	assert.NoError(t, err)

	assert.Equal(t, "INFO", expectedLog.Level)
	assert.Equal(t, "[OPERATION]", expectedLog.Msg)
	assert.Equal(t, "GET", expectedLog.Method)
	assert.Equal(t, "success", expectedLog.Response)
	assert.Equal(t, "/operation-log", expectedLog.Path)
	assert.Equal(t, 200, expectedLog.Status)
}

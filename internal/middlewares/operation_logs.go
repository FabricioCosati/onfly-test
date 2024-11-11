package middlewares

import (
	"bytes"
	"io"
	"log"
	"log/slog"
	"time"

	"github.com/FabricioCosati/onfly-test/internal/utils"
	"github.com/gin-gonic/gin"
)

type CustomWriter struct {
	gin.ResponseWriter
	Buff *bytes.Buffer
}

func (cw CustomWriter) Write(b []byte) (int, error) {
	cw.Buff.Write(b)
	return cw.ResponseWriter.Write(b)
}

func OperationLogs(fileName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := getRequestBody(ctx)
		writer := &CustomWriter{
			ResponseWriter: ctx.Writer,
			Buff:           &bytes.Buffer{},
		}
		ctx.Writer = writer
		ctx.Next()

		utils.CreateFolderIfNotExists()
		file := utils.GetFileToSave(fileName)
		defer file.Close()

		logger := slog.New(slog.NewJSONHandler(io.MultiWriter(file), nil))

		level := utils.GetLogLevel(ctx.Writer.Status())
		logger.Log(ctx, level, "[OPERATION]",
			"Timestamp", time.Now().Format("2006-01-02 15:04:05"),
			"Method", ctx.Request.Method,
			"Request", body,
			"Response", writer.Buff.String(),
			"Path", ctx.FullPath(),
			"Status", ctx.Writer.Status())
	}
}

func getRequestBody(ctx *gin.Context) string {
	if ctx.Request.Method == "GET" {
		return ""
	}

	body, err := ctx.GetRawData()
	if err != nil {
		log.Fatalf("error on get request data: %s", err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
	return string(body)
}

package handler

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSendEmailHandler_SendEmail(t *testing.T) {
	r := gin.New()

	sendEmailHandler := &Handler{}
	r.POST("/api/mail/file", sendEmailHandler.SendEmail)

	t.Run("ValidRequest", func(t *testing.T) {

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fileWriter, _ := writer.CreateFormFile("file", "document.pdf")
		fileWriter.Write([]byte("binary_data_of_file"))
		writer.WriteField("emails", "elonmusk@x.com,jeffbezos@amazon.com,zuckerberg@meta.com")
		writer.Close()

		req := httptest.NewRequest("POST", "/api/mail/file", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, but got %d", http.StatusOK, w.Code)
		}
		expectedResponse := `{"message":"Email sent successfully"}`
		if w.Body.String() != expectedResponse {
			t.Errorf("Expected response %s, but got %s", expectedResponse, w.Body.String())
		}
	})

	t.Run("InvalidMimeType", func(t *testing.T) {

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fileWriter, _ := writer.CreateFormFile("file", "invalid.jpg")
		fileWriter.Write([]byte("binary_data_of_file"))
		writer.WriteField("emails", "elonmusk@x.com,jeffbezos@amazon.com,zuckerberg@meta.com")
		writer.Close()

		req := httptest.NewRequest("POST", "/api/mail/file", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, but got %d", http.StatusBadRequest, w.Code)
		}
	})

}

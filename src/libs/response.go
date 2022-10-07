package libs

import (
	"encoding/json"
	"net/http"

	"github.com/rs/cors"
)

type Resp struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Message     string      `json:"message"`
	Error       interface{} `json:"error,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func (res *Resp) Send(w http.ResponseWriter) {
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// w.Header().Set("Access-Control-Allow-Methods", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	c.HandlerFunc(w,&http.Request{})

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(res.Status)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("error encode"))
	}
}

func Response(data interface{}, status int, message string, err error) *Resp {
	if err != nil {
		return &Resp{
			Status:      status,
			Description: statusDescription(status),
			Message:     message,
			Error:       err.Error(),
			Data:        data,
		}
	}

	return &Resp{
		Status:      status,
		Description: statusDescription(status),
		Message:     message,
		Error:       err,
		Data:        data,
	}
}

func statusDescription(status int) string {
	switch status {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 304:
		return "Not Modified"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	case 501:
		return "Bad Gateway"
	default:
		return ""
	}
}

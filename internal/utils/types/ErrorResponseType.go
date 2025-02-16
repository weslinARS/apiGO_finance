package types

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors" jsonapi:"attr,errors"`
}
type ErrorDetail struct {
	Status string      `json:"status"`
	Title  string      `json:"title"`
	Detail interface{} `json:"detail"`
}

func NewErrorResponse(status int, detail interface{}, title string) *fiber.Map {
	return &fiber.Map{
		"errors": []ErrorDetail{
			{
				Status: strconv.Itoa(status),
				Title: func() string {
					if title != "" {
						return title
					}
					return ErrorCodeToText(status)
				}(),
				Detail: detail,
			},
		},
	}
}

func ErrorCodeToText(code int) string {
	switch code {
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 405:
		return "Method Not Allowed"
	case 406:
		return "Not Acceptable"
	case 408:
		return "Request Timeout"
	case 409:
		return "Conflict"
	case 410:
		return "Gone"
	case 411:
		return "Length Required"
	case 412:
		return "Precondition Failed"
	case 413:
		return "Payload Too Large"
	case 414:
		return "URI Too Long"
	case 415:
		return "Unsupported Media Type"
	case 416:
		return "Range Not Satisfiable"
	case 417:
		return "Expectation Failed"
	case 418:
		return "I'm a teapot"
	case 422:
		return "Unprocessable Entity"
	case 425:
		return "Too Early"
	case 426:
		return "Upgrade Required"
	case 428:
		return "Precondition Required"
	case 429:
		return "Too Many Requests"
	case 431:
		return "Request Header Fields Too Large"
	case 451:
		return "Unavailable For Legal Reasons"
	case 500:
		return "Internal Server Error"
	case 501:
		return "Not Implemented"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	case 504:
		return "Gateway Timeout"
	case 505:
		return "HTTP Version Not Supported"
	case 506:
		return "Variant Also Negotiates"
	case 507:
		return "Insufficient Storage"
	case 508:
		return "Loop Detected"
	case 510:
		return "Not Extended"
	case 511:
		return "Network Authentication Required"
	default:
		return "Unknown Error"
	}
}

package core

import (
	"strconv"
)

type Response struct {
	Version string
	Code    int
	Reason  string
	Headers map[string]string
	Body    []byte
	Size    int
}

func CreateResponse(code int, version string, reason string, headers map[string]string, body []byte) *Response {
	resp := Response{
		Version: version,
		Code:    code,
		Reason:  reason,
		Headers: headers,
		Body:    body,
	}

	if _, ok := headers["Content-Length"]; !ok && len(body) != 0 {
		headers["Content-Length"] = strconv.Itoa(len(body))
	}

	return &resp
}

func (r *Response) ToBytes() []byte {
	resp_str := ""

	resp_str += r.Version + " " + strconv.Itoa(r.Code) + " " + r.Reason + "\r\n"

	for k, v := range r.Headers {
		resp_str += k + ": " + v + "\r\n"
	}

	if len(r.Body) != 0 {
		resp_str += "\r\n"
	}

	resp_str += string(r.Body)

	return []byte(resp_str)

}

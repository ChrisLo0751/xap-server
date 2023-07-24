package errors

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrorResponse struct {
	HttpStatusCode int
	Err            *Error
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpStatusCode: 400,
		Err: &Error{
			Message: "Request body parse failed",
			Code:    1,
		},
	}
	ErrorNotAuthUser = ErrorResponse{
		HttpStatusCode: 401,
		Err: &Error{
			Message: "User authentication failed",
			Code:    2,
		},
	}
	ErrorDBError = ErrorResponse{
		HttpStatusCode: 500,
		Err: &Error{
			Message: "Database ops failed",
			Code:    3,
		},
	}
	ErrorInternalFaults = ErrorResponse{
		HttpStatusCode: 500,
		Err: &Error{
			Message: "Internal service error",
			Code:    4,
		},
	}
)

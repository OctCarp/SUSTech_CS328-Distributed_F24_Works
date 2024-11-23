package utils

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleDbErr(c *gin.Context, err error) {
	// convert gRPC error to HTTP status code and API error
	st, ok := status.FromError(err)
	if !ok {
		// if not a gRPC error, send internal server error
		SendInternalErr(c, "Database internal Error")
	}

	message := "From db: " + st.Message()

	switch st.Code() {
	case codes.NotFound:
		SendNotFoundErr(c)
	case codes.InvalidArgument:
		SendBadRequestErr(c, message)
	case codes.PermissionDenied:
		SendBadRequestErr(c, message)
	case codes.Unauthenticated:
		SendUnauthorizedErr(c)
	default:
		SendInternalErr(c, message)
	}
}

package middlewares

// HttpMiddleware represent the data-struct for middleware
type HttpMiddleware struct {
	// another stuff , may be needed by middleware
}

// InitHttpMiddleware InitMiddleware initialize the middleware
func InitHttpMiddleware() *HttpMiddleware {
	return &HttpMiddleware{}
}

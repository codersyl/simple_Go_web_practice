type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// A Handler responds to an HTTP request.
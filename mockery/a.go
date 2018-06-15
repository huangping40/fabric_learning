package mockery

// ServerRequestContext defines the functionality of a server request context object
type ServerRequestContext interface {
	BasicAuthentication() (string, error)
	Get(string, int64) (string, error)
}

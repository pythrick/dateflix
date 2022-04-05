package repository

const (
	// ErrReadTimeout should be used when we received a timeout from kafka consumer
	ErrReadTimeout = PkgError("timeout while connecting database")

	// ErrEmptyConfig should be used when a required config is empty or nil
	ErrEmptyConfig = PkgError("required config is empty")

	// ErrConsumerConnection should be used when we received a expected consumer connection error
	ErrDatabaseConnection = PkgError("error while connecting database")
)

type PkgError string

func (e PkgError) Error() string {
	return string(e)
}

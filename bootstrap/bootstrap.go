package bootstrap

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	db := newDatabase()

	srv := New(host, port, db)
	return srv.Run()
}

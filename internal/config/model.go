package config

type Config struct {
	App   App
	HTTP  HTTP
	Redis Redis
	NATS  NATS
	DSN   DSN
}

type App struct {
	Name string
	Env  string
}

type HTTP struct {
	Listen string
}

type Redis struct {
	Addr string
	DB   int
}

type NATS struct {
	URL string
}

type DSN struct {
	Postgres string
}

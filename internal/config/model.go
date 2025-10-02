package config

type Config struct {
	App       App
	HTTP      HTTP
	Redis     Redis
	NATS      NATS
	DSN       DSN
	Auth      Auth
	SnowFlake SnowFlake
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

type Auth struct {
	Keys string
	TTL  string
}

type SnowFlake struct {
	WorkerIdBitLength string
	WorkerID          string
	SeqBitLength      string
	MinSeqNumber      string
	MaxSeqNumber      string
	BaseTime          string
}

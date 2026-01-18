package config

type Config struct {
	App       App       `yaml:"app"`
	HTTP      HTTP      `yaml:"http"`
	Redis     Redis     `yaml:"redis"`
	NATS      NATS      `yaml:"nats"`
	Database  Database  `yaml:"database"`
	Auth      Auth      `yaml:"auth"`
	Snowflake Snowflake `yaml:"snowflake"`
	Hash      Hash      `yaml:"hash"`
}

type App struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
}

type HTTP struct {
	Listen string `yaml:"listen"`
}

type Redis struct {
	Addr string `yaml:"addr"`
	DB   int    `yaml:"db"`
}

type NATS struct {
	URL string `yaml:"url"`
}

type Database struct {
	DSN string `yaml:"dsn"`
}

type Auth struct {
	Keys string `yaml:"keys"`
	TTL  uint32 `yaml:"ttl"`
}

type Snowflake struct {
	WorkerID          int64  `yaml:"worker_id"`
	WorkerIDBitLength int    `yaml:"worker_id_bit_length"`
	SeqBitLength      int    `yaml:"seq_bit_length"`
	BaseTime          string `yaml:"base_time"`
}

type Hash struct {
	Memory     uint32 `yaml:"memory"`
	Inerations uint32 `yaml:"iterations"`
	SaltLength uint32 `yaml:"salt_length"`
	KeyLength  uint32 `yaml:"key_length"`
}

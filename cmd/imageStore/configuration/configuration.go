package configurations

import (
	"flag"
)

const (
	ServerAddress = "localhost:8080"
	GrpcPort      = 50051
	PathToFile    = "/Users/pl_chkn/Dev/Go/ImageStore/downloaded/"
	//DataBaseURI  = "postgresql://postgres:1234@localhost:5432?sslmode=disable"
	//AccrualSystemAdress        = "http://localhost:8080/"
	//AccessTokenLiveTimeMinutes = 15
	//RefreshTokenLiveTimeDays   = 7
	//AccessTokenSecret          = "jdnfksdmfksd"
	//RefreshTokenSecret         = "mcmvmkmsdnfsdmfdsjf"
	//NumOfWorkers               = 10
	//PoolBuffer                 = 1000
	//MaxJobRetryCount           = 5
)

type Config struct {
	ServerAddress string `env:"RUN_ADDRESS"`
	GrpcPort      int    `env:"GRPC_PORT"`
	PathToFile    string `env:"PATH_TO_FILE"`
	//AccrualSystemAdress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	//Token               ConfigToken
	//DataBase            ConfigDatabase
	//WorkerPool          ConfigWorkerPool
}

//type ConfigToken struct {
//	AccessTokenLiveTimeMinutes int    `env:"ACCESS_TOKEN_LIVE_TIME_MINUTES"`
//	RefreshTokenLiveTimeDays   int    `env:"REFRESH_TOKEN_LIVE_TIME_DAYS"`
//	AccessTokenSecret          string `env:"ACCESS_TOKEN_SECRET"`
//	RefreshTokenSecret         string `env:"REFRESH_TOKEN_SECRET"`
//}

//type ConfigDatabase struct {
//	DataBaseURI string `env:"DATABASE_URI"`
//}
//
//type ConfigWorkerPool struct {
//	NumOfWorkers     int `env:"num_of_workers"`
//	PoolBuffer       int `env:"pool_buffer"`
//	MaxJobRetryCount int `env:"max_job_retry_count"`
//}

func New() *Config {
	//dbCfg := ConfigDatabase{
	//	DataBaseURI: DataBaseURI,
	//}

	//tokenCfg := NewTokenConfig()

	flagServerAddress := flag.String("a", ServerAddress, "server address")
	flagGrpcPort := flag.Int("gp", GrpcPort, "the grpc server port")
	flagPathToFile := flag.String("ptf", PathToFile, "path to files")
	//flagDataBaseURI := flag.String("d", DataBaseURI, "URI for database")
	//flagAccrualSystemAdress := flag.String("r", AccrualSystemAdress, "URL for accrual system")
	flag.Parse()

	//if *flagDataBaseURI != DataBaseURI {
	//	dbCfg.DataBaseURI = *flagDataBaseURI
	//}
	//
	//wpConf := ConfigWorkerPool{
	//	NumOfWorkers:     NumOfWorkers,
	//	PoolBuffer:       PoolBuffer,
	//	MaxJobRetryCount: MaxJobRetryCount,
	//}

	cfg := Config{
		ServerAddress: ServerAddress,
		GrpcPort:      GrpcPort,
		PathToFile:    PathToFile,
		//AccrualSystemAdress: AccrualSystemAdress,
		//DataBase:            dbCfg,
		//Token:               tokenCfg,
		//WorkerPool:          wpConf,
	}

	if *flagServerAddress != ServerAddress {
		cfg.ServerAddress = *flagServerAddress
	}
	if *flagGrpcPort != GrpcPort {
		cfg.GrpcPort = *flagGrpcPort
	}
	if *flagPathToFile != PathToFile {
		cfg.PathToFile = *flagPathToFile
	}
	//if *flagAccrualSystemAdress != AccrualSystemAdress {
	//	cfg.AccrualSystemAdress = *flagAccrualSystemAdress
	//}
	//
	//err := env.Parse(&cfg.DataBase)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = env.Parse(&cfg)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//cfg.ServerAddress += "/api/image/"

	return &cfg
}

//func NewTokenConfig() ConfigToken {
//	tokenCfg := ConfigToken{
//		AccessTokenLiveTimeMinutes: AccessTokenLiveTimeMinutes,
//		RefreshTokenLiveTimeDays:   RefreshTokenLiveTimeDays,
//		AccessTokenSecret:          AccessTokenSecret,
//		RefreshTokenSecret:         RefreshTokenSecret,
//	}
//	err := env.Parse(&tokenCfg)
//	if err != nil {
//		log.Fatal()
//	}
//	return tokenCfg
//}

package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// Application
	AppName    string
	AppEnv     string
	AppDebug   bool
	AppVersion string
	AppPort    string
	LogLevel   string

	// Security
	JWTSecret            string
	JWTExpiration        time.Duration
	JWTRefreshExpiration time.Duration
	EncryptionKey        string
	HSMEnabled           bool

	// Database - PostgreSQL
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBSSLMode        string
	DBMaxConnections int
	DBIdleTimeout    time.Duration

	// Redis
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	// MongoDB
	MongoURI string
	MongoDB  string

	// Elasticsearch
	ElasticHost     string
	ElasticPort     string
	ElasticUsername string
	ElasticPassword string

	// Message Queues
	KafkaBrokers   []string
	RabbitMQHost   string
	RabbitMQPort   string
	RabbitMQUser   string
	RabbitMQPass   string
	RabbitMQVHost  string

	// Blockchain
	BTCNodeURL   string
	BTCNetwork   string
	ETHNodeURL   string
	ETHChainID   int
	ETHNetwork   string
	TRONNodeURL  string
	BNBNodeURL   string
	SOLRPCURL    string
	PolygonURL   string

	// Compliance
	KYCLevelMin       int
	AMLScreeningEnabled bool
	AMLThresholdDaily int64

	// Trading
	MakerFeePercent float64
	TakerFeePercent float64
	MinOrderSize    float64
	MaxOrderSize    float64

	// Features
	FeatureP2PTrading      bool
	FeatureAgentNetwork    bool
	FeatureRemittance      bool
	FeatureNFTMarketplace  bool
	FeatureStaking         bool

	// Development
	DevelopmentMode bool
	DebugLogging    bool
}

var AppConfig *Config

func LoadConfig() (*Config, error) {
	// Load from .env file
	godotenv.Load()

	c := &Config{
		// Application
		AppName:    getEnv("APP_NAME", "tm-money-exchange"),
		AppEnv:     getEnv("APP_ENV", "development"),
		AppDebug:   getEnvBool("APP_DEBUG", true),
		AppVersion: getEnv("APP_VERSION", "1.0.0"),
		AppPort:    getEnv("APP_PORT", "8080"),
		LogLevel:   getEnv("APP_LOG_LEVEL", "info"),

		// Security
		JWTSecret:            getEnv("JWT_SECRET", "super-secret-key-change-in-production"),
		JWTExpiration:        parseDuration(getEnv("JWT_EXPIRATION", "24h")),
		JWTRefreshExpiration: parseDuration(getEnv("JWT_REFRESH_EXPIRATION", "7d")),
		EncryptionKey:        getEnv("ENCRYPTION_KEY", "your-encryption-key-32-chars-long"),
		HSMEnabled:           getEnvBool("HSM_ENABLED", false),

		// Database
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "tm_money"),
		DBPassword:       getEnv("DB_PASSWORD", "secure_password"),
		DBName:           getEnv("DB_NAME", "tm_money_production"),
		DBSSLMode:        getEnv("DB_SSL_MODE", "disable"),
		DBMaxConnections: getEnvInt("DB_MAX_CONNECTIONS", 100),
		DBIdleTimeout:    parseDuration(getEnv("DB_IDLE_TIMEOUT", "10m")),

		// Redis
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

		// MongoDB
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:  getEnv("MONGO_DB", "tm_money_logs"),

		// Elasticsearch
		ElasticHost:     getEnv("ELASTIC_HOST", "localhost"),
		ElasticPort:     getEnv("ELASTIC_PORT", "9200"),
		ElasticUsername: getEnv("ELASTIC_USERNAME", "elastic"),
		ElasticPassword: getEnv("ELASTIC_PASSWORD", "elastic_password"),

		// Message Queues
		KafkaBrokers:  []string{getEnv("KAFKA_BROKERS", "localhost:9092")},
		RabbitMQHost:  getEnv("RABBITMQ_HOST", "localhost"),
		RabbitMQPort:  getEnv("RABBITMQ_PORT", "5672"),
		RabbitMQUser:  getEnv("RABBITMQ_USER", "guest"),
		RabbitMQPass:  getEnv("RABBITMQ_PASSWORD", "guest"),
		RabbitMQVHost: getEnv("RABBITMQ_VHOST", "/"),

		// Blockchain
		BTCNodeURL:   getEnv("BTC_NODE_URL", "http://localhost:8332"),
		BTCNetwork:   getEnv("BTC_NETWORK", "testnet"),
		ETHNodeURL:   getEnv("ETH_NODE_URL", "http://localhost:8545"),
		ETHChainID:   getEnvInt("ETH_CHAIN_ID", 5),
		ETHNetwork:   getEnv("ETH_NETWORK", "goerli"),
		TRONNodeURL:  getEnv("TRON_NODE_URL", "http://localhost:8090"),
		BNBNodeURL:   getEnv("BNB_NODE_URL", "https://data-seed-prebsc-1-b.binance.org:8545"),
		SOLRPCURL:    getEnv("SOL_RPC_URL", "https://api.devnet.solana.com"),
		PolygonURL:   getEnv("POLYGON_NODE_URL", "https://rpc-mumbai.maticvigil.com"),

		// Compliance
		KYCLevelMin:         getEnvInt("KYC_LEVEL_MIN", 1),
		AMLScreeningEnabled: getEnvBool("AML_SCREENING_ENABLED", true),
		AMLThresholdDaily:   getEnvInt64("AML_THRESHOLD_DAILY", 50000),

		// Trading
		MakerFeePercent: getEnvFloat("DEFAULT_MAKER_FEE_PERCENT", 0.10),
		TakerFeePercent: getEnvFloat("DEFAULT_TAKER_FEE_PERCENT", 0.15),
		MinOrderSize:    getEnvFloat("MIN_ORDER_SIZE", 0.001),
		MaxOrderSize:    getEnvFloat("MAX_ORDER_SIZE", 1000000),

		// Features
		FeatureP2PTrading:     getEnvBool("FEATURE_P2P_TRADING", true),
		FeatureAgentNetwork:   getEnvBool("FEATURE_AGENT_NETWORK", false),
		FeatureRemittance:     getEnvBool("FEATURE_INTERNATIONAL_REMITTANCE", true),
		FeatureNFTMarketplace: getEnvBool("FEATURE_NFT_MARKETPLACE", false),
		FeatureStaking:        getEnvBool("FEATURE_STAKING", true),

		// Development
		DevelopmentMode: getEnvBool("DEVELOPMENT_MODE", true),
		DebugLogging:    getEnvBool("DEBUG_LOGGING", true),
	}

	AppConfig = c
	return c, nil
}

// Helper functions
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvBool(key string, defaultVal bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true" || value == "1" || value == "yes"
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultVal
}

func getEnvInt64(key string, defaultVal int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intVal
		}
	}
	return defaultVal
}

func getEnvFloat(key string, defaultVal float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
			return floatVal
		}
	}
	return defaultVal
}

func parseDuration(value string) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return time.Hour * 24
	}
	return duration
}

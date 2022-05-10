package config

import (
	"flag"
)

var (
	blockDivisorDefault           = 3
	consumerPendingTimeoutDefault = 60
	logFolderDefault              = "./logs/"
)

// RedisConfig all redis related config
type RedisConfig struct {
	RedisURL               string
	Password               string
	BlockDivisor           int // can be set to any divisor (decrease specimen production throughput)
	ConsumerPendingTimeout int // defaults to 1 min
}

// CodecConfig all codec related config
type CodecConfig struct {
	AvroCodecPath string
}

// StorageConfig composes of configs needed by differtn stores
type StorageConfig struct {
	// local
	BinaryFilePath string

	// gcp
	GcpSvcAccountAuthFile string
	ReplicaBucketLoc      string

	// ipfs
	IpfsServiceType  string
	IpfsServiceToken string
}

// ProofchainConfig all proof chain related configurations
type ProofchainConfig struct {
	ProofChainAddr string
}

// ChainConfig config around the supported blockchains
type ChainConfig struct {
	RPCURL       string
	PrivateKey   string
	KeystorePath string
	KeyStorePwd  string

	// for elrond
	WebsocketURLs string
}

// MetricsConfig config around collecting metrics from agent
type MetricsConfig struct {
	Enabled        bool
	HTTPServerAddr string
	HTTPServerPort string
}

// AgentConfig composes all the different configs into a single config for agent node
type AgentConfig struct {
	RedisConfig      RedisConfig
	CodecConfig      CodecConfig
	StorageConfig    StorageConfig
	ProofchainConfig ProofchainConfig
	ChainConfig      ChainConfig
	MetricsConfig    MetricsConfig

	LogFolder string
}

// NewAgentConfig create a new empty config
func NewAgentConfig() *AgentConfig {
	return &AgentConfig{}
}

// LoadConfig populates the config from env flags and cli arguments
func (ac *AgentConfig) LoadConfig() {
	envConfig, err := loadEnvConfig()
	if err != nil {
		panic(err)
	}

	ac.populateFromEnvConfig(envConfig)
	ac.populateFromCliFlags()
}

// SegmentLength #block-specimen within a single proofchain tx
func (ac *AgentConfig) SegmentLength() int {
	// number of block specimen/results within a single uploaded avro encoded object
	// defaults to 1 block per segment in bsp-agent live mode
	return 1
}

func (ac *AgentConfig) populateFromEnvConfig(config *EnvConfig) {
	ac.StorageConfig.IpfsServiceToken = config.IpfsConfig.ServiceToken

	ac.RedisConfig.Password = config.RedisConfig.Password

	ac.ChainConfig.RPCURL = config.EthConfig.RPCURL
	ac.ChainConfig.PrivateKey = config.EthConfig.PrivateKey
	ac.ChainConfig.KeystorePath = config.EthConfig.KeystorePath
	ac.ChainConfig.KeyStorePwd = config.EthConfig.KeyStorePwd
}

func (ac *AgentConfig) populateFromCliFlags() {
	// redis
	flag.StringVar(&ac.RedisConfig.RedisURL, "redis-url", LookupEnvOrString("RedisURL", ""), "redis consumer stream url")
	flag.IntVar(&ac.RedisConfig.BlockDivisor, "block-divisor", LookupEnvOrInt("BlockDivisor", blockDivisorDefault), "integer divisor that allows for selecting only block numbers divisible by this number")
	flag.IntVar(&ac.RedisConfig.ConsumerPendingTimeout, "consumer-timeout", LookupEnvOrInt("ConsumerPendingTimeout", consumerPendingTimeoutDefault), "number of seconds to wait before pending messages consumer timeout")

	// codec
	flag.StringVar(&ac.CodecConfig.AvroCodecPath, "avro-codec-path", LookupEnvOrString("CodecPath", ""), "local path to AVRO .avsc files housing the specimen/result schemas")

	// chain
	flag.StringVar(&ac.ProofchainConfig.ProofChainAddr, "proof-chain-address", LookupEnvOrString("ProofChain", ""), "hex string address for deployed proof-chain contract")
	flag.StringVar(&ac.ChainConfig.WebsocketURLs, "websocket-urls", LookupEnvOrString("WebsocketURLs", ""), "url to websockets clients separated by space")

	// log
	flag.StringVar(&ac.LogFolder, "log-folder", LookupEnvOrString("LogFolder", logFolderDefault), "Location where the log files should be placed")

	// storage
	flag.StringVar(&ac.StorageConfig.IpfsServiceType, "ipfs-service", LookupEnvOrString("IpfsService", ""), "Allowed values are 'web3.storage', 'pinata' and 'others'")
	flag.StringVar(&ac.StorageConfig.BinaryFilePath, "binary-file-path", LookupEnvOrString("BinaryFilePath", ""), "local path to AVRO encoded binary files that contain block-replicas")
	flag.StringVar(&ac.StorageConfig.GcpSvcAccountAuthFile, "gcp-svc-account", LookupEnvOrString("GcpSvcAccount", ""), "local path to google cloud platform service account auth file")
	flag.StringVar(&ac.StorageConfig.ReplicaBucketLoc, "replica-bucket", LookupEnvOrString("ReplicaBucket", ""), "google cloud platform object store target for specimen")

	// metrics
	flag.BoolVar(&ac.MetricsConfig.Enabled, "metrics", false, "enable metrics reporting and collection")
	flag.StringVar(&ac.MetricsConfig.HTTPServerAddr, "metrics.addr", LookupEnvOrString("MetricsHttpServerAddr", "127.0.0.1"), "Enable stand-alone metrics HTTP server listening interface (default: \"127.0.0.1\")")
	flag.StringVar(&ac.MetricsConfig.HTTPServerPort, "metrics.port", LookupEnvOrString("MetricsHttpServerPort", "6061"), "Metrics HTTP server listening port (default: 6061)")

	flag.Parse()
}

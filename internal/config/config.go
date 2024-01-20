package config

type KDFConfiguration struct {
	CostFactor            int
	KeySize               int
	SaltSize              int
	BlockSizeFactor       int
	ParallelizationFactor int
}

// Configuration to generate AES encryption key
var AESConfig = KDFConfiguration{
	CostFactor:            32768,
	KeySize:               32, // 32 bytes = 256 bits key. For AES key
	SaltSize:              8,  // 8 bytes = 64 bits of salt
	BlockSizeFactor:       8,
	ParallelizationFactor: 1,
}

package types

const (
	// ModuleName defines the module name
	ModuleName = "hellochain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hellochain"
)

var (
	ParamsKey = []byte("p_hellochain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

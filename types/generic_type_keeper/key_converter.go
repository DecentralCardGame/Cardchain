package generic_type_keeper

type KeyConver[T any] func([]byte, T) []byte

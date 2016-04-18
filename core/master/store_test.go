// Author hoenig

package master

import "github.com/shoenig/subspace/core/common/stream"

type MockStore struct {
}

func (s *MockStore) CreateStream(stream.Metadata) error {
	return nil
}

func (s *MockStore) ContainsStream(string) bool {
	return false
}

func (s *MockStore) GetStreams() []stream.Metadata {
	return nil
}

func (s *MockStore) AddPack(stream.Bundle) (uint64, error) {
	return 0, nil
}

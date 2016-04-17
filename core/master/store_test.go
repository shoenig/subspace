// Author hoenig

package master

import "github.com/shoenig/subspace/core/common/stream"

type MockStore struct {
}

func (s *MockStore) CreateStream(stream.Stream) error {
	return nil
}

func (s *MockStore) ContainsStream(string) bool {
	return false
}

func (s *MockStore) GetStreams() []stream.Stream {
	return nil
}

func (s *MockStore) AddPack(stream.Pack) (uint64, error) {
	return 0, nil
}

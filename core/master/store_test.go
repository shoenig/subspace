// Author hoenig

package master

import "github.com/shoenig/subspace/core/common/stream"

type MockStore struct {
}

func (s *MockStore) NewStream(stream.Metadata) error {
	return nil
}

func (s *MockStore) ContainsStream(string) bool {
	return false
}

func (s *MockStore) AllStreams() []stream.Metadata {
	return nil
}

func (s *MockStore) NewGeneration(stream.Generation) (uint64, error) {
	return 0, nil
}

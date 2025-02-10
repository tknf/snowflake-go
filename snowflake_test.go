package snowflake

import "testing"

func TestSnowflake(t *testing.T) {
	s := NewSnowflake(1)
	id := s.Generate()
	if id <= 0 {
		t.Errorf("invalid id: %d", id)
	}

	id2 := s.Generate()
	if id2 <= 0 {
		t.Errorf("invalid id2: %d", id2)
	}

	if id == id2 {
		t.Errorf("duplicated id: %d", id)
	}
}

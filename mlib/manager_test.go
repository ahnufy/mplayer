package mlib

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("Test Build")
	}
	if mm.Len() != 0 {
		t.Error("Test Len")
	}
	m := &MusicEntry{"1", "SONG", "FY", "WY", "GD"}
	mm.Add(m)
	if mm.Len() != 1 {
		t.Error("Test Add")
	}
}

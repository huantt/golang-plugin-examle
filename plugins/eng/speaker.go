package main

type speaker struct {
}

func (s *speaker) Speak() string {
	return "hello"
}

// Exported
var Speaker speaker
var SpeakerName = "Alice"

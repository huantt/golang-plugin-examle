package main

type speaker struct {
}

func (s *speaker) Speak() string {
	return "xin chào"
}

// Exported
var Speaker speaker
var SpeakerName = "Anh Thư"

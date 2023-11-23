package main

import "fmt"

type AudioPlayer struct{}

type VideoPlayer struct{}

type SubtitleManager struct{}

func (a *AudioPlayer) PlayAudio() {
	fmt.Println("Audio start")
}

func (a *AudioPlayer) StopAudio() {
	fmt.Println("Audio stop")
}

func (v *VideoPlayer) PlayVideo() {
	fmt.Println("Video start")
}

func (a *VideoPlayer) StopVideo() {
	fmt.Println("Video stop")
}

func (s *SubtitleManager) LoadSubtitle() {
	fmt.Println("Subtitle Load")
}

func (s *SubtitleManager) DisplaySubtitle() {
	fmt.Println("Subtitle Display")
}

type MediaPlayerFacade struct {
	AudioPlayer
	VideoPlayer
	SubtitleManager
}

type MediaPlayer interface {
	PlayAudio()
	PlayVideo()
	LoadandDisplaySubtitle()
}

func (m *MediaPlayerFacade) PlayAudio() {
	m.AudioPlayer.PlayAudio()
}

func (m *MediaPlayerFacade) PlayVideo() {
	m.VideoPlayer.PlayVideo()
}

func (m *MediaPlayerFacade) LoadandDisplaySubtitle() {
	m.SubtitleManager.LoadSubtitle()
	m.SubtitleManager.DisplaySubtitle()
}

func main() {
	// Create an instance of MediaPlayerFacade
	mediaPlayer := &MediaPlayerFacade{}

	// Play audio
	mediaPlayer.PlayAudio()

	// Play video
	mediaPlayer.PlayVideo()

	// Load and display subtitles
	mediaPlayer.LoadandDisplaySubtitle()
}

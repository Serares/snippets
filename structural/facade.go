package structural

import "fmt"

type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player is on")
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player is off")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Println("Playing movie:", movie)
}

// Subsystem 2: Sound System
type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("Sound System is on")
}

func (s *SoundSystem) Off() {
	fmt.Println("Sound System is off")
}

func (s *SoundSystem) SetVolume(volume int) {
	fmt.Printf("Setting volume to %d\n", volume)
}

// Subsystem 3: Projector
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is on")
}

func (p *Projector) Off() {
	fmt.Println("Projector is off")
}

func (p *Projector) SetInput(source string) {
	fmt.Println("Setting projector input to", source)
}

type HomeTheaterFacade struct {
	dvdPlayer   *DVDPlayer
	soundSystem *SoundSystem
	projector   *Projector
}

func NewHomeTheaterFacade(d *DVDPlayer, s *SoundSystem, p *Projector) *HomeTheaterFacade {
	return &HomeTheaterFacade{dvdPlayer: d, soundSystem: s, projector: p}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.projector.On()
	h.projector.SetInput("DVD")
	h.soundSystem.On()
	h.soundSystem.SetVolume(10)
	h.dvdPlayer.On()
	h.dvdPlayer.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	h.dvdPlayer.Off()
	h.soundSystem.Off()
	h.projector.Off()
}

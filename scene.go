package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
	. "fmt"

	"time"
	"log"
)

type scene struct {
	time int
	bg *sdl.Texture
	healers  *healer

}

func newScene(r *sdl.Renderer) (*scene,error){
// Background//
	bg, err := img.LoadTexture(r,"res/img/dungeon1.png")
	if err != nil {
		 Errorf("no regreso background: %v",nil,err)
	}
////fin del background///

/// inicio de healers
	h, err := newhealer(r)
	if err != nil {
		Errorf("cosas",err)
	}


	return &scene{bg:bg,healers:h},nil
	//asigna los valores a la estructura
}

func (s *scene) run(events <-chan sdl.Event,r *sdl.Renderer)<-chan error{
	errc := make(chan error)

	go func() {
		defer close(errc)
		for range time.Tick(10*time.Millisecond){
			select {

			case e := <-events:
				log.Printf("evento : %T",e)
				return
			default:

				if err := s.paint(r); err != nil {
					errc <- err
				}
			}
		}
	}()

	return errc
}


func (s *scene) paint(r *sdl.Renderer)error{
	//s.time++
	r.Clear()
	if err:= r.Copy(s.bg,nil,nil);err !=nil{
		return Errorf(" no se puedo copiar la textura %v",err)
	}

	if err := s.healers.paint2(r);err != nil{
		return err
	}

	r.Present()
	return nil
}


func (s *scene)destroy(){
	s.bg.Destroy()
	s.healers.desroy()
}
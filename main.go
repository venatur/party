package main

import (
	. "fmt"
	. "os"
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"github.com/veandco/go-sdl2/ttf"

	"runtime"
)

func main() {

	if err := run(); err != nil {
		Fprint(Stderr, "no se inicio la aplicacion %v", err)
		Exit(2)
	}
}
func run() error {
	ttf.Init()

	if err := ttf.Init(); err != nil {
		return Errorf(" no se inicio fuente %v",err)
	}
	defer ttf.Quit()
	
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return Errorf("no se inicio la pantalla %v", err)
		Exit(2)
	}

	defer sdl.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return Errorf( "no se mostro la ventana %v", err)
	}

	defer w.Destroy()
	if err := drawTitle(r);err!=nil{
		Errorf("no se pinto la fuente %v",err)
	}

	time.Sleep(time.Second*2)

	s,err := newScene(r)

	if err != nil {
		 Errorf("no se realizo el fondo %v",err)
	}

	defer s.destroy()
	events := make(chan sdl.Event)
	errc:=s.run(events,r)
		runtime.LockOSThread()
		for{
			select {
			case events <-sdl.WaitEvent():
			case err:= <-errc:
				return Errorf("error de evento %v",err)
			}

		}


	/*
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

 	time.AfterFunc(5*time.Second,cancel)*/




	if err := s.paint(r);err !=nil{
		return Errorf("no se puede cargar la scene")
	}
	time.Sleep(time.Second*5)
	return nil

}

func drawTitle(r *sdl.Renderer) error  {
	r.Clear()
	font, err := ttf.OpenFont("res/fonts/mariachi.ttf",20)

	if err != nil {

		return Errorf("no se coloco la fuente %v",err)

	}
	defer font.Close()
	c := sdl.Color{R: 255, G: 158, B: 0, A: 255}
	s, err :=font.RenderUTF8_Solid("mariachi fiesta",c)
	if err != nil {
		return Errorf("cannot render surface %v",err)
	}
	defer s.Free()
	t,err := r.CreateTextureFromSurface(s)
	if err != nil {
		Errorf("canno create texture %v",err)
	}
	defer t.Destroy()

	if err = r.Copy(t,nil,nil);err != nil{
		Errorf("could not copy %v",err)
	}

	r.Present()

	return  nil
}


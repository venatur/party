package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
	."fmt"

)

type healer struct {
	time int
	textures []*sdl.Texture
	y float64

}


func newhealer(r *sdl.Renderer)(*healer,error){
	println("entre a metodo de nuevo healer")
	var textures[] *sdl.Texture
	for i:=1;i<3;i++ {
		path := Sprintf("res/img/healerw%d.png",i)
		texture, err := img.LoadTexture(r, path)

		if err != nil {
			Errorf("no cargó imagen de healer: %v", nil, err)
		}
		textures = append(textures,texture)

	}

	return &healer{textures:textures,y:300}, nil
}

func (h *healer)paint2(r *sdl.Renderer)error {
	h.time++
	h.y ++
	rect:= &sdl.Rect{200,int32(h.y),50,70}
	i := h.time /10 % len(h.textures)
	if err:= r.Copy(h.textures[i],nil,rect);err !=nil{
		return Errorf(" no se puedo copiar la textura %v",err)
	}
return nil
}

func (h *healer)desroy(){
	println("entre al destructor")
	for _, t:=range h.textures{
		t.Destroy()
	}
}
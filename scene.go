package neutron

import (
	"errors"
)

var (
	scenes map[string]*Scene
)

type Scene struct {
	Name    string
	OnStart func(scene *Scene)
}

func NewScene(name string, onStart func(scene *Scene)) *Scene {
	//make new scene
	scene := Scene{
		Name:    name,
		OnStart: onStart,
	}

	//add scene to game
	if scenes == nil {
		scenes = make(map[string]*Scene)
	}
	scenes[name] = &scene

	return &scene

}

func (scene *Scene) load() {

	//unload previous scene
	depthmap = make([]neutronRef, 0)

	//start scene
	scene.OnStart(scene)

}

func LoadScene(name string) error {
	scene, exists := scenes[name]
	if !exists {
		return errors.New("attempting to load scene but, scene with the name \"" + name + "\" couldnt be found. ")
	}
	nxscene = scene
	return nil
}

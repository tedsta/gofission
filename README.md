Theodore DeRego

2013

fission
=======

A flexible game engine written in Go and based on the entity component system (ECS) model.

The Concept
===========

In an ECS, the scene contains any number of entities, which, by themselves, are just abstract objects with no functionality. Functionality is added to entities by adding components to them. In that way, entities are composed of and defined by their components. Alone, however, components do nothing. Components only contain data. You should implement as few methods for a component as possible. So, where does all of the meat of your game happen? In systems. A system processes entities that satisfy the system's component requirements. For instance, a PlayerControlSystem would process all entities with a PlayerControlComponent (probably only one entity for this example).

Why Go?
=======

Something that really got me going was Go's philosophy - composition over inheritance. As you can see, I really like entity component systems (this is my third attempt at writing fission, the first two were in C++). Entities are composed of components. Composition. The entity component system design fits perfectly with the underlying ideas of Go. And that's why I chose to rewrite fission in Go.

Features
========

Fission is designed to be a very extensible game engine. The framework is the core module, which defines entities, components, the scene, and the event manager. Functionality, such as rendering and input, are added as separate packages. You can extend fission with your own packages. Want to make a physics package? Great! Because I probably won't, as my current game doesn't call for a realistic physics simulation.

Note: For now, most of these are unimplemented, but I implement things pretty fast :)

 - It's an entity component system, is that a feature?
 - Rendering (sprites, spritesheets, shapes)
 - All sorts of input (keyboard, mouse, joystick, etc.?)
 - Easy networking for all of your multiplayer games

Dependencies
============

go get github.com/go-gl/gl

go get github.com/go-gl/glfw3

Core
 - Just the Go standard lib :)

Input
 - GLFW3

Rendering
 - GLFW3
 - GL

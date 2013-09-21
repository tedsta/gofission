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

A few people have asked me - why Go? Go is such a young language, I only just began learning it, and I already have 9 years of C++ experience under my belt. I was a C++ purist. C++ or the highway. But as I was doing a rewrite of the old fission in C++, I became dissatisfied with C++. C++ was bloated, and C++11 was even more bloated. Heavy parallelism in C++ is a nightmare. And, apparently, if I didn't jump on the smart pointer bandwagon, I was a bad programmer. I was finally beginning to realize why some people didn't like C++. I had grown up with C++, so everything about the language was natural to me - writing C++ code was perhaps easier for me than writing English sentences.

Then I remembered Go. A coworker introduced it to me a some months ago, but I originally brushed it off as a crappy remake of C that took away manual memory management. Of course, I hadn't made it past the front page of the website. Anyway, while I was on my search for a new language, I decided to give it another "go";). I quickly fell in love. Built-in concurrency, interfaces, and dynamic arrays, a juicy standard library, even more strongly typed than C++, and no more inheritance diagrams. It was a dream.

On that note, something that really got me going was Go's philosophy - composition over inheritance. As you can see, I really like entity component systems (why else would I write the same engine 3 times). Entities are composed of components. Composition. The entity component system design fits perfectly with the underlying ideas of Go. And that's why I chose to rewrite fission in Go.

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

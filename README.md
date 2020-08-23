
# GOLF Engine Examples

This repo contains several demo examples for the [GoLF engine](https://github.com/bjatkin/golf-engine). These examples
were designed to showcase the types of games that can be made with GoLF as well as providing developers with a reference
for how to use the API. Currently, this repo contains 4 different projects. Each of these mini project is explained below.
The different features used in the games are also listed to make it easy to find code examples to help you with the
GoLF engine.

# Running GoLF Examples

Running these game demos is easy. First, install [go](https://golang.org/) on your computer if it is not already.
Next clone this repo as well as the [GoLF engine repo](https://github.com/bjatkin/golf-engine). Then, open terminal
and navigate into the golf-examples/demos repository. From here run the golf_toolkit binary located 
in the golf-engine/util/[your os]/ directory. This will start the golf toolkit which you can use to 
start a simple web server. Simply run play [demo name] and a localhost server will be started. 
You're defalut browser will also be opened to http://localhost:8080 where you will be able to play the demo.

# One Page API Demo
![Demo Screen Shot](https://github.com/bjatkin/golf-examples/blob/master/images/APIDemo.png)

This is a simple demo of several of the most important GoLF API calls. It demonstrates drawing a tile man, drawing, scaling and flipping sprites, and moving the camera. This demo also enables the mouse. you can change
the screen pallet by clicking the up and down buttons
on the sides of the screen. You can also pan the screen
by using the arrow keys and clip / un-clip the screen
with the Z key. Also, be careful, Slimy is watching you.

### Controls
 * arrow keys - pan the screen.
 * X key - recenter the screen.
 * Z key - enable / diable screen cliping.

### Key Features
 * Using the map function
 * Using the Spr and SSpr function
 * Using the Camera function
 * Getting mouse input
 * Getting keyboard input
 * Swapping color pallets
 * Screen clipping
 * Drawing rectangles and circles
 * A simple particle system

# BiBiDuck Demo

![BiBiDuck Screeshot 1](https://github.com/bjatkin/golf-examples/blob/master/images/bibiDuck1.png)
![BiBiDuck Screeshot 2](https://github.com/bjatkin/golf-examples/blob/master/images/bibiDuck2.png)
![BiBiDuck Screeshot 3](https://github.com/bjatkin/golf-examples/blob/master/images/bibiDuck3.png)
![BiBiDuck Screeshot 4](https://github.com/bjatkin/golf-examples/blob/master/images/bibiDuck4.png)

This is a basic platformer game that plays similarly to 
2D mario and other platforming games. It has multiple
collectibles to gather and an NPC to interact with.

### Controls
 * Arrow Keys - move the player
 * Z Key - Jump
 * X Key - Talk to an NPC

### Key Features
 * Plaformer physics (friction, gravity etc.)
 * Simple HUD
 * 2 types of collectables
 * NPC interaction / dialouge system
 * Tile based collision detection

# Blood Demo

![Blood Screeshot 1](https://github.com/bjatkin/golf-examples/blob/master/images/blood1.png)
![Blood Screeshot 2](https://github.com/bjatkin/golf-examples/blob/master/images/blood2.png)
![Blood Screeshot 3](https://github.com/bjatkin/golf-examples/blob/master/images/blood3.png)
![Blood Screeshot 4](https://github.com/bjatkin/golf-examples/blob/master/images/blood4.png)

This game is a small demo of an ARPG game. It is inspired
visually by the diablo series and plays similarly to
those games. You can move you player arround with the
arrow keys and bolth Z and X can be used to fire
projectiles at your enemies. When enemies die they
drow 'blood' which you must pick up to fill your mana
gauge.

### Controls
 * Arrow Keys - move the player
 * Z Key - light attack
 * X Key - heavy attack

### Key Features
 * Particle system (blood)
 * AABB collision detection
 * Simple Entity Component System
 * Fading colors using pallet cycling
 * A main menue, death screen and victory screen
 * Simple enemy AI
 * An animation system
 * A dablo inspired HUD

# Wander Demo

![Wander Screeshot 1](https://github.com/bjatkin/golf-examples/blob/master/images/wander1.png)
![Wander Screeshot 2](https://github.com/bjatkin/golf-examples/blob/master/images/wander2.png)
![Wander Screeshot 3](https://github.com/bjatkin/golf-examples/blob/master/images/wander3.png)
![Wander Screeshot 4](https://github.com/bjatkin/golf-examples/blob/master/images/wander4.png)

This is a small game that plays somewhat like a Pokemon
game. You can move around and interact with various people and objects. You can also go into houses and 
talk to the people inside. Be sure to revisit old
conversations every once in a while as dialog can
change over time. Anything with an exclamation mark
above it is interactable.

### Controls
 * Arrow Keys - move the player
 * Z Key - intract with the object you are facing

### Key Features
 * Simple dialoge system
 * Simple event system for tracking story events
 * Tile based collision detection
 * Simple system with scean swaping
 * Drawing subsets of the tile map
 * Following the player with a camera
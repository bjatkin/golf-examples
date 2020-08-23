
# GOLF Engine Examples

This repo contains several demo examples for the [GoLF engine](https://github.com/bjatkin/golf-engine). These examples
were designed to showcase the types of games that can be made with GoLF as well as providing developers with a reference
for how to use the GoLF APIs. Currently, this repo contains 4 different projects. Each of these mini projects is explained below.
The different features used in the games are also listed to make it easy to find code examples to help you build your games.

# Running GoLF Examples

Running these game demos is easy. First, install [go](https://golang.org/) on your computer if it is not already.
Next clone this repo as well as the [GoLF engine repo](https://github.com/bjatkin/golf-engine). Then, open terminal
and navigate into the golf-examples/demos directory. From here run the golf_toolkit binary located 
in the golf-engine/util/[your os]/ directory. This will start the golf toolkit which you can use to 
start a simple web server. From here simply run play [demo name] and a localhost server will be started. 
You're default browser will also be opened to http://localhost:8080 where you will be able to play the demo.

# One Page API Demo
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/APIDemo.png" alt="Demo Screenshot" width="350" >

This is a simple demo of several of the most important GoLF API calls. It demonstrates drawing a tile map, drawing, scaling and flipping sprites, and moving the camera. This demo also enables the mouse. you can change
the engine pallets by clicking the up and down buttons
on the sides of the screen. You can also pan the screen
by using the arrow keys and clip / unclip the screen
with the Z key. Also, be careful, Slimy is watching you.

### Controls
 * arrow keys - pan the screen.
 * X key - recenter the screen.
 * Z key - enable / disable screen clipping.

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


<img src="https://github.com/bjatkin/golf-examples/blob/master/images/bibiDuck1.png" alt="BiBiDuck Screenshot 1" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/bibiDuck2.png" alt="BiBiDuck Screenshot 2" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/bibiDuck3.png" alt="BiBiDuck Screenshot 3" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/bibiDuck4.png" alt="BiBiDuck Screenshot 4" width="350" >

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
 * NPC interaction / dialogue system
 * Tile based collision detection

# Blood Demo

<img src="https://github.com/bjatkin/golf-examples/blob/master/images/blood1.png" alt="Blood Screenshot 1" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/blood2.png" alt="Blood Screenshot 2" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/blood3.png" alt="Blood Screenshot 3" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/blood4.png" alt="Blood Screenshot 4" width="350" >

This game is a small demo of an ARPG game. It is inspired
visually by the diablo series and plays similarly to
those games. You can move the player around with the
arrow keys and both Z and X can be used to fire
projectiles at your enemies. When enemies die they
drop blood pools which you must pick up to fill up your mana.

### Controls
 * Arrow Keys - move the player
 * Z Key + Arrow Keys - light attack
 * X Key - heavy attack

### Key Features
 * Particle system (blood)
 * AABB collision detection
 * Simple Entity Component System
 * Fading colors using pallet cycling
 * A main menu, death screen and victory screen
 * Simple enemy AI
 * An animation system
 * A dablo inspired HUD

# Wander Demo

<img src="https://github.com/bjatkin/golf-examples/blob/master/images/wander1.png" alt="Wander Screenshot 1" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/wander2.png" alt="Wander Screenshot 2" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/wander3.png" alt="Wander Screenshot 3" width="350" >
<img src="https://github.com/bjatkin/golf-examples/blob/master/images/wander4.png" alt="Wander Screenshot 4" width="350" >

This is a small game that plays somewhat like Pokemon.
You can move around and interact with various people and objects. You can also go into houses and 
talk to the people inside. Be sure to revisit old
conversations every once in a while as dialog can
change over time. Anything with an exclamation mark
above it is interactable.

### Controls
 * Arrow Keys - move the player
 * Z Key - interact with the object you are facing

### Key Features
 * Simple dialogue system
 * Simple event system for tracking story events
 * Tile based collision detection
 * Simple system with scene swapping
 * Drawing subsets of the tile map
 * Following the player with a camera
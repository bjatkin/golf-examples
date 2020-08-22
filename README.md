
# GOLF Engine examples
this repo contains several example games desinged to show how the golf fantasy engine works. it contains one main project right now located in the blood directory.
This is a small ARPG game where you play a vampire killing monsters and collecting there blood. It demonstrates a simple ECS system as well as many of the features
of the GOLF engine including things like sprite sheets, map files and drawing text on screen. It also demonstrates common RPG game elements like character movement
collision detection, simple enemy AI, dialouge boxes and more. the code is currently in development so to see the final version you'll need to check back later.

### TODO
[] Rewrite golf demos intro
[] Swap the imports on the demo to use the github version
[] update the github version of golf (go get -u)
[] Add a section on running these demos
[] Add some screen shots of the different games to the main read me
[] Add a link to the golf readme
[] Add a description to the bibi duck project
[] Add a features list to the bibi duck project
[] Add some screen shots of the bibi duck project
[] Add a description to the wander project
[] Add a features list to the wander project
[] Add some screen shots of the wander project
[] Add a description to the blood project
[] Add a features list to the blood project
[] Add some screen shots of the wander project
[] Finish the blood demo

### WANDER DEMO
---
DONE

### BIBI DUCK DEMO
---
DONE

### API DEMO
---
DONE

### BLOOD DEMO
---
MVP todo  list
[x] spawn projectiles pointing in the direction youre faceing (attack 1)
[x] spawn diagonal projectiles as well (attack 1)
[x] spawn a ring of projectiles (attack 2)
[x] when a projectile hits an enemy, do damage + knockback
[] when an enemies hp hits zero, kill them and create lots of blood
[] increase mob cap over time (but set a max so the game dosen't crash)
[] split the code into more files (e.g. player file, enemy file etc.)
[] split initalization code into several smaller functions.
[] make a better UI for the blood bank
[] remove the coord UI at the top of the screen
[] update the golf engine import
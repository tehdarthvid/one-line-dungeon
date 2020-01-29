# ピザを焼く

> baking pizza

A project where we start with a "hello world" function and go on a journey to making it an API and getting it deployed into a decent environment.

The project name is from a storyline in the 銀の匙 (Silver Spoon) manga by 荒川 弘 (Hiromu Arakawa).

## The Journey

1. [[done]](https://github.com/tehdarthvid/pizaoyaku/commit/1ef972f179e43c8d0e55550e4aefb2cbed848837) ["hello world" in Golang](#hello-world-in-golang)
1. [[todo]]() [The App](#the-app-one-line-dungeon)

## The Journal

[newest entry first]

### The App: One Line Dungeon

The app is a one line dungeon generator.

The app will generate a JSON object containing details of a procedurally generated rouguelike dungeon representing an arena encounter. The specs are:

1. The name of the world where this arena exists will come from an environment variable.
1. The arena is one dimentional, 16 characters long.
1. The arena generation is deterministic.
1. There is some form of identifier for each generated arena.
1. The following are procedurally generated/placed:
   1. exit
   1. Amulet of Yendor
   1. monster/s
   1. item/s
   1. trap/s

This app was inspired by [Jim Shepard](https://twitter.com/madjackmcmad)'s [Storytelling, World Building, and You](https://www.youtube.com/watch?v=jd7K0EmkvPk) talk at [Roguelike Celebration](https://roguelike.club) 2018.

### "hello world" in Golang

Why Golang?

1.  Using JS may lessen the need for a build step. I want to make sure builds and binary deployments can be handled.
1.  I want to learn Golang. :P

Speedbumps in this chapter:

1.  issues with 'go install' due to GOBIN env
1.  issues with 'go build' default output and '-o'

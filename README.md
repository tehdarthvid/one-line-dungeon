# One Line Dungeon

This project is a one line dungeon generator.

The app will generate a JSON object containing details of a procedurally generated "dungeon" with only one-dimension. This one-dimensional dungeon represents a scenario, a small snippet of what one would encounter in a traditional roguelike. The scenario should be visualizable by a short string of characters.

This app was inspired by [Jim Shepard](https://twitter.com/madjackmcmad)'s [Storytelling, World Building, and You](https://www.youtube.com/watch?v=jd7K0EmkvPk) talk at [Roguelike Celebration](https://roguelike.club) 2018.

### Specs

1. The name of the world where this scenario exists will come from an environment variable.
1. The scenario is one dimentional, not more than 16 characters long.
1. Scenario generation is deterministic.
1. There is some form of unique identifier for each generated scenario.
1. The following are procedurally generated/placed:
   1. exit/s
   1. monster/s
   1. item/s
   1. trap/s
   1. Amulet of Yendor

### Purpose

> ピザを焼く

The intent is to start with the most basic "hello world" and then evolve it into a sample deployed service, applying/demonstrating workflow concepts along the way. The journey is recorded so people can follow the evolution, and hopefully learn from it.

『ピザを焼く』 ("baking pizza") is a story arc from 銀の匙 (Silver Spoon), a manga by 荒川 弘 (Hiromu Arakawa). Hachiken, the protagonist, enrolled to an agricultural college knowing essentially nothing about agriculture. In the "Baking Pizza" arc, Hachiken stumbles into an old, broken, stone oven and is thrust into leading an effort to create pizza wiht it.

## The Journey

1. [[done]](https://github.com/tehdarthvid/pizaoyaku/releases/tag/v0.1.0) "hello world" in Golang. [[journal]](#hello-world-in-golang)
1. [todo] Define the initial [app scope](#the-app-one-line-dungeon) and [specs](#specs).

## The Journal

[newest entry first]

### "hello world" in Golang

Why Golang?

1.  Using JS may lessen the need for a build step. I want to make sure builds and binary deployments can be handled.
1.  I want to learn Golang. :P

Speedbumps in this chapter:

1.  issues with 'go install' due to GOBIN env
1.  issues with 'go build' default output and '-o'

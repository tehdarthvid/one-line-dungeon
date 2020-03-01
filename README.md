# One Line Dungeon

The intent is to start with the most basic "hello world" and then evolve it into a sample deployed service, applying/demonstrating workflow concepts along the way. The journey is recorded so people can follow the evolution, and hopefully learn from it. It's gonna be dirty and messy, but you'll be part of the ride, bumps and all! :D

In terms of functionality, it's currently aiming to be a "one line dungeon" generator.

The app will generate a JSON object containing details of a procedurally generated "dungeon" with only one-dimension. This one-dimensional dungeon represents a scenario, a small snippet of what one would encounter in a traditional roguelike. The scenario should be visualizable by a short string of characters, similar to ASCII or char rendered roguelikes.

The functionality was inspired by [Jim Shepard](https://twitter.com/madjackmcmad)'s [Storytelling, World Building, and You](https://www.youtube.com/watch?v=jd7K0EmkvPk) talk at [Roguelike Celebration](https://roguelike.club) 2018.

## Specs

1. The name of the world where this scenario exists will come from an environment variable.
1. The scenario is one dimentional, not more than 16 characters long.
1. Scenario generation is deterministic.
1. There is some form of unique identifier for each generated scenario.
1. The following are procedurally generated/placed:
   1. exit/s
   1. monster/s
   1. item/s
   1. trap/s
   1. "Amulet of Yendor"

## Status: pre-alpha

Nothing really usable yet. More of you can see it evolve slowly.

### Building

![Go Unit and Build Tests](https://github.com/tehdarthvid/one-line-dungeon/workflows/Go%20Unit%20and%20Build%20Tests/badge.svg)

1. clone the repo: `git clone https://github.com/tehdarthvid/one-line-dungeon`
1. run all unit tests: `go test ./src/... -v`
1. build: `go build -v -o ./bin/old ./src`
1. run : `./bin/old`

## The Journey

1. [[done]](https://github.com/tehdarthvid/pizaoyaku/releases/tag/v0.1.0)[[journal]](#hello-world-in-golang) "hello world" in Golang.
1. [[done]](https://github.com/tehdarthvid/pizaoyaku/commit/52cfd10d9781eecd72a197057bb15a4b34b4a62a) Define the initial [app scope](#one-line-dungeon) and [specs](#specs).
1. [[done]](https://github.com/tehdarthvid/pizaoyaku/commit/2e7bacfe4962d727dc8d7bfae8ca8699c65e2c51) Play around on implementing the specs in Go to learn the laguage basics.
   1. functions
   1. splices (lists)
   1. getting time
   1. getting environment variables
1. [[done]](https://github.com/tehdarthvid/pizaoyaku/commit/15035a79c90060bd035cedef2ed80248905f28e3) Play around on Unit Testing features of Go.
1. [[done]](https://github.com/tehdarthvid/one-line-dungeon/commit/24aa5827bd3d9edbbdc0073dfcdae275bf38b86b)[[journal]](#first-github-actions) Investigate GitHub CI options.
1. [todo] Redesign output into meta-data.
1. [todo] Fill dummy output from Encounter package and generate a JSON from its caller.

## The Journal

[newest entry first]

### "hello world" in Golang

Why Golang?

1. Using JS may lessen the need for a build step. I want to make sure builds and binary deployments can be handled.
1. I want to learn Golang. :P

Speedbumps in this chapter:

1. issues with 'go install' due to GOBIN env undefined
1. issues with 'go build'
   1. avoid params after the '-o' param
   1. relative folders need to start with "." (`./src`)
   1. no need to `*` (no need for `./src/*`)
   1. default output name is folder name

### First GitHub Actions

GitHub Actions helpfully provided a [starter workflow for Go](https://github.com/actions/starter-workflows/blob/master/ci/go.yml). This starter workflow just runs a Go build assuming files on root. We customized this to the build params based on this project's directory structure. Also added running the unit tests.

There are two unit tests to check for the ability to set and retrive environment variables:

1. [custom environment variables](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/using-environment-variables)
   These are in your workflow YAML, so they're visible.
1. [custom environment variables for credtentials/tokens/secrets](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-and-using-encrypted-secrets)
   They SHOULD NOT be in your source files. They are stored in the system either via the GitHub or APIs. Once set, they cannot be retrieved, only removed. They are "not passed to workflows that are triggered by a pull request from a fork." They are also masked if you try to display them in the action logs, which is good.

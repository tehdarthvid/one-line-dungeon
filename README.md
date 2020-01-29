# ピザを焼く

> baking pizza

A project where we start with a "hello world" function and go on a journey to making it an API and getting it deployed into a decent environment.

The project name is from a storyline in the 銀の匙 (Silver Spoon) manga by 荒川 弘 (Hiromu Arakawa).

## The Journey

1. [[1ef972f]](https://github.com/tehdarthvid/pizaoyaku/commit/1ef972f179e43c8d0e55550e4aefb2cbed848837) "hello world" in Golang.

   Why:

   1. Using JS may lessen the need for a build step. I want to make sure builds and binary deployments can be handled.
   1. I want to learn Golang. :P

   Speedbumps:

   1. issues with 'go install' due to GOBIN env
   1. issues with 'go build' default output and '-o'

1. [todo] Function that returns time and some form of instance identifier.

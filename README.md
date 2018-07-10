# Multistage Scratch Non-Privileged

This demonstrates a [multi-stage docker build](https://docs.docker.com/engine/userguide/eng-image/multistage-build/) to [create a scratch container](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/), [running as a non-privileged user](https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341).

The aim is to create a tiny image, with no exta files, and only one non-privileged, `nobody` user, with the ability to use SSL.

## Usage

To build the image:

    docker build --tag multistage .

To run the image:

    docker run -it multistage

To make a request:

    curl localhost:7000/google

## Demo

If you'd like to walk through the steps that go into creating this image, [there is a demonstration repository here](https://github.com/davidcarboni/ddd). The [built image can be found on Docker Hub](https://hub.docker.com/r/davidcarboni/go-scratch-nonprivileged-multistage/) and there's [an accompanying blog post](https://medium.com/@davidcarboni/simplicity-wins-4fc54efae1f0) too.

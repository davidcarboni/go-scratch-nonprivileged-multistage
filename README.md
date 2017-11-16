# Multistage Scratch Non-Privileged

This demonstrates a [multi-stage docker build](https://docs.docker.com/engine/userguide/eng-image/multistage-build/) to [create a scratch container](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/), [running as a non-privileged user](https://medium.com/@lizrice/non-privileged-containers-based-on-the-scratch-image-a80105d6d341).

The aim is to create a tiny image, with no exta files, and only one non-privileged, `nobody` user, with the ability to use SSL.

#!/bin/sh

docker kill $(docker ps | grep soup_bot | cut -d " " -f1)

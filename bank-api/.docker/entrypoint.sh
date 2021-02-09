#!/bin/bash

yarn
yarn typeorm migration:run
yarn console fake
yarn start:dev

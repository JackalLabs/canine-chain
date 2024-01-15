#!/bin/bash

check_chain() {
    ret=$(command -v canined > /dev/null 2>&1)
    if [ ret ]; then
        canined version
    fi
}

check_provider() {
    ret=$(command -v jprovd > /dev/null 2>&1)
    if [ ret ]; then
        jprovd version
    fi
}



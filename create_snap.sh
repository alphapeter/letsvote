#!/bin/bash
docker run --rm -v $(pwd):/go/src/github.com/alphapeter/letsvote -it -w /go/src/github.com/alphapeter/letsvote snapcore/snapcraft snapcraft 


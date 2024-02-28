#!/bin/bash

# generate CV
pyenv local 3.11.5
python data/cv/cv_generator.py

# compile GO binary
CGO_ENABLED=0 
GOOS="linux" 
GOARCH="amd64"
go build -ldflags="-s -w" -o dac.ac .

# generate the HTML files for tailwind
./dac.ac &
# Save the PID of the server process
SERVER_PID=$!

# Wait for 10 seconds
sleep 3

# Kill the server process
kill $SERVER_PID

# generate CSS using generated files
NODE_ENV=production npx tailwindcss -i ./static/assets/tailwind/src/styles.css -o ./static/assets/tailwind/styles.css

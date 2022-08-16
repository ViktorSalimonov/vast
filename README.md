# VAST generator
The service creates VAST tag based on the input parameters such as video-file and landing page.

# Preview
![Screenshot 2022-08-05 at 14 22 39](https://user-images.githubusercontent.com/16116989/183068095-e7d8e886-cfa4-4758-b43d-3b8650248187.png)
![Screenshot 2022-08-05 at 14 23 22](https://user-images.githubusercontent.com/16116989/183068107-5d577397-bc9c-491a-ad8f-4efdcfdccaad.png)

# Installation

## Dockerfile

1. ```docker build --tag vast .```
2. ```docker run --publish 8080:8080 vast```
3. Open http://localhost:8080/

## Source code

### Pre-requirements
1. PostgreSQL (14.4)
2. ffprobe/ffmpeg (https://ffmpeg.org/download.html)

### Installation
1. Clone the repo ```git clone https://github.com/ViktorSalimonov/vast.git```
2. Install dependrencies ```go mod download```
3. Run the service ```go run .```
4. Open http://localhost:8080/

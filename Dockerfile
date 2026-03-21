FROM golang:1.26.1
LABEL authors="litiano"

ENTRYPOINT ["top", "-b"]
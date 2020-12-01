FROM faddat/sos-base AS builder

MAINTAINER jacobgadikian@gmail.com

RUN pacman -Syyu go git

COPY . /spn

RUN go mod vendor && \
		go build /spn/spnd

FROM faddat/sos



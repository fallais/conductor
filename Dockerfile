FROM        frolvlad/alpine-glibc:latest
MAINTAINER  Francois ALLAIS <francois.allais@sogeti.com>

ADD orchestrator /usr/bin
ADD scenarii.yml /usr/bin

EXPOSE     5001
CMD        [ "/usr/bin/orchestrator", "--c", "/usr/bin/scenarii.yml" ]

FROM alpine:latest

MAINTAINER Sergio Vallejo <vallejosergio@gmail.com>

# Rethinkdb
ADD rethinkdb /rethinkdb/
ADD entrypoint.sh /entrypoint.sh

# Backend
ADD backend /backend/
ENV GOPATH=$HOME/go 
ENV PATH=$GOPATH/bin:$PATH

WORKDIR /backend

# Frontend
ADD frontend /frontend/

# Install packages, compile and clean
RUN apk --no-cache add rethinkdb ruby py-pip git \
&& apk --no-cache add --virtual .go-dependencies go libc-dev \
&& pip install rethinkdb \
&& chmod +x /entrypoint.sh \
&& go get gopkg.in/gorethink/gorethink.v3 \
&& go build \
&& apk del .go-dependencies  && rm -r $GOPATH && rm -rf /var/cache/apk/*

# Entrypoint
ENTRYPOINT /entrypoint.sh
CMD ["rethinkdb", "--bind", "all"]

# Expose RethinkDB
EXPOSE 28015 29015 8080
# Expose Frontend
EXPOSE 8000
# Expose Backend
EXPOSE 8001

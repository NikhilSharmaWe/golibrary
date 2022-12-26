# FROM golang:1.19

# ENV MYSQL_USER=admin \
#     MYSQL_PASSWORD=password

# RUN mkdir /build
# WORKDIR /build
# RUN go mod init github.com/NikhilSharmaWe/golibrary

# RUN export GO111MODULE=on
# RUN go get github.com/NikhilSharmaWe/golibrary/...
# RUN cd /build && git clone https://github.com/NikhilSharmaWe/golibrary.git

# RUN cd /build/golibrary && go build

# EXPOSE 8080
# ENTRYPOINT [ "/build/URLShortener/golibrary" ]
FROM golang:1.19

WORKDIR /home
COPY ./ /home

RUN cd /home && go build -o golibrary

CMD ["/home/golibrary"]
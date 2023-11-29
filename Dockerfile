# syntax=docker/dockerfile:1
FROM golang:1.21

# Update package list and install jq
RUN apt-get update && apt-get install -y jq

WORKDIR /canine-chain
COPY . /canine-chain
RUN ls
RUN go install ./cmd/canined
CMD ["sh", "./scripts/test-node-dev.sh"]
EXPOSE 26657
EXPOSE 26656
EXPOSE 26658
EXPOSE 1317
EXPOSE 6060
EXPOSE 9090


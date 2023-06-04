FROM golang:1.20

# Create app directory 
WORKDIR /app

# Install app dependencies 
COPY go.mod go.sum ./
RUN go mod download

# Bundle app source 
COPY . .

# Build the Go app
RUN go build -o ./main ./app

# Expose the port that the express are running 
EXPOSE 13002

# Run migrations and start program (with shell to expand environment variables)
ENTRYPOINT ["./main"]

# Step 1: Use the official Go image as the base image
FROM golang:1.20

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the rest of the application code
COPY . .

# Step 5: Build the Go application
RUN go build -o e-wallet

# Step 6: Specify the command to run the application
CMD ["./e-wallet"]

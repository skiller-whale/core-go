FROM golang:1.16.5-alpine3.13

# Copy files for the sync script
COPY _sync /app/sync
WORKDIR /app/sync
RUN go mod download
RUN go build sync.go

# Set the working directory to be the exercises dir (when sh is run)
WORKDIR /app/exercises
RUN ln -s /root/.ash_history /app/.command_history

# Clear the history on startup, and run the sync
CMD > /root/.ash_history && /app/sync/sync

FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Build Args
ARG LOG_DIR=/app/logs

# Create Log Directory
RUN mkdir -p ${LOG_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log

# Copy the source from the current directory to the Working Directory inside the container
# Copy binary file
COPY main  .
# Copy config file
COPY app/config/config.yaml .

# Declare volumes to mount
VOLUME [${LOG_DIR}]

# Command to run the executable
CMD ["./main", "--config", "./config.yaml"]
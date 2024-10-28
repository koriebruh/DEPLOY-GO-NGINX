# Gunakan base image Go
FROM golang:1.22.6

# Set working directory
WORKDIR /app

# Copy seluruh isi proyek ke dalam container
COPY . .

# Install dependencies
RUN go mod download

# Compile aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["./main"]

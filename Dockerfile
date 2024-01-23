# Sử dụng hình ảnh golang:1.21 làm hình ảnh cơ sở
FROM golang:1.21.3-alpine AS builder

# Đặt thư mục làm thư mục làm việc hiện tại trong container
WORKDIR /app

# Sao chép go.mod và go.sum vào container
COPY go.mod go.sum /app/

# Tải dependencies
RUN go mod download
COPY . .
# Biên dịch ứng dụng Golang
RUN go build -o wan-api-news .

# Tạo hình ảnh cuối cùng
FROM alpine:3.18.4 AS final

# Đặt thư mục làm thư mục làm việc hiện tại trong container
WORKDIR /app

# Sao chép tệp thực thi từ stage trước
COPY --from=builder /app/wan-api-news .

# Cài đặt các gói cần thiết
RUN apk --no-cache add curl iputils net-tools 

# Khởi chạy ứng dụng Golang khi container được chạy
CMD ["/app/wan-api-news"]

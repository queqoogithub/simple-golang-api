# 🐹 SIMPLE CRUD API
การพัฒนา Golang API อย่างง่าย โดยเน้นไปที่การปรับใช้หลักการ SOLID และ Testing

## ก่อนอื่นให้ติดตั้ง Redis จาก [Docker Hub](https://hub.docker.com/_/redis) ก่อน
1. ติดตั้ง redis image
    ```sh
    go mod download
    ```
2. รัน redis ด้วย docker
    ```sh
    docker run --name some-redis -p 6379:6379 -d redis
    ```

3. ตรวจสอบว่า redis ทำงานอยู่หรือไม่
    ```sh
    docker ps
    ```

4. การเข้าถึง redis cli (command line interface)
    ```sh
    docker exec -it some-redis redis-cli
    ```

## หลังจาก Clone โปรเจคมาแล้ว
1. ติดตั้ง dependencies:
    ```sh
    docker pull redis
    ```

2. สร้าง binary:
    ```sh
    go build -o simple-crud-api
    ```

3. รันโปรเจค:
    ```sh
    go run main.go
    ```

4. ทดสอบ API:
    ```sh
    curl -X GET http://localhost:8000/items
    curl -X GET http://localhost:8000/items/1
    curl -X POST http://localhost:8000/items -H "Content-Type: application/json" -d '{"id":"1","name":"Item One","price":"$30"}'
    curl -X PUT http://localhost:8000/items/1 -H "Content-Type: application/json" -d '{"name":"Updated Item One","price":"$15"}'
    curl -X DELETE http://localhost:8000/items/1
    ```

5. รัน test
    ```sh
    go test ./handlers
    ```

## 🐳 หากต้องการใช้ Docker ในการรันโปรเจค
1. สร้าง Docker image:
    ```sh
    docker build -t simple-crud-api .
    ```

2. รัน Docker container:
    ```sh
    docker run -p 8000:8000 simple-crud-api
    ```
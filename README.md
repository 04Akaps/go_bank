# go_bank

Golang, Docker, PostgreSQL을 활용한 은행

# Data Schme

- https://dbdiagram.io/d/63ad40e57d39e42284e80351

# Docker Setting

- docker를 설치 후 다음과 같은 명령어를 실행

1. docker pull postgress:12-alpine

   - 도커 이미지를 가져 온다.

2. docker run --name <name> -p <port번호> -e POSTGRES_USER=<id = root> -e POSTGRES_PASSWORD=<password> -d postgres:12-alpine

   - docker를 실행 시킨다.
   - imgName을 설정하고, 원하는 port번호를 설정한다.
   - 이후 패스워드 및 백그라운드 실행과 img를 선택 해 준다.

3. docker ps

   - 컨테이너에 등록된 이미지를 확인
   - 만약 잘못 올라가 있다면, docker stop, docker rm을 활용

4. docker exec -it <imgName> psql -U root

   - 이후 컨테이너를 실행 시킨다

5. docker logs <img>

   - 해당 이미지의 로그를 확인 가능하다.

# go library

1. brew install golang-migrate

- DB schema migrate
- `migrate create -ext sql -dir db/migration -seq init_schema`

2. brew install kyleconroy/sqlc/sqlc

- DB Connect library

3. go get github.com/lib/pq

- db connect driver

4. go get github.com/stretchr/testify

- To Check the test result

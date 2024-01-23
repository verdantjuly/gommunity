# GOmuunity

### go를 사용하면서 발생하는 에러를 교환하는 커뮤니티

# Coding Convention

참고 자료 : https://google.github.io/styleguide/go/best-practices

- 함수 이름에서 패키지 이름을 반복하지 않는다.
- 리시버함수에 리시버 구조체 이름을 사용하지 않는다.
- 파라미터, 변수 등의 이름을 함수 이름에 사용하지 않는다.
- 리턴 값의 이름이나 타입을 함수 이름에 사용하지 않는다.
- Get을 함수 이름에 쓰는 것을 지양하나 필요한 경우에는 허용한다.

# Tech Spec

- Backend Language : GO 1.21.5
- Frontend Language : HTML, javascript
- DB : MySQL(ver 8.0.33)
- Cache : Redis(ver 6.2.6)

# Wireframe

[figma](https://www.figma.com/file/Ypk12AKtXHeJqkC4PHklvn/Untitled?type=design&node-id=0%3A1&mode=design&t=wFf0jDXzPFOQ2qVM-1)
![스크린샷 2024-01-16 오전 10 32 19](https://github.com/verdantjuly/gommunity/assets/131671804/bd55744c-de91-4895-a0d4-1e2885a13f24)

![스크린샷 2024-01-16 오전 10 32 35](https://github.com/verdantjuly/gommunity/assets/131671804/89e4da49-7667-4820-9a3a-5e6e5afc46ec)


# API

# ERD

![gommunity_ERD](https://github.com/verdantjuly/gommunity/assets/131671804/cf9394c2-32d3-4cf1-a784-3bd66a894142)


# Directory Structure

```
gommunity
├─ .gitignore
├─ README.md
├─ go.mod
├─ server
│  ├─ api
│  │  ├─ handler.go
│  │  └─ http.go
│  ├─ cache
│  │  └─ resource.go
│  ├─ main.go
│  └─ util
│     └─ file.go
├─ static
└─ templates
   └─ index.html

```

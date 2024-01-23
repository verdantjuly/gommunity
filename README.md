# GOmuunity

### go를 사용하면서 발생하는 에러를 교환하는 커뮤니티

# Coding Convention

참고 자료 : https://google.github.io/styleguide/go/best-practices

- 함수 이름에서 패키지 이름을 반복하지 않는다.

  ```go
  // Bad:
  package yamlconfig

  func ParseYAMLConfig(input string) (*Config, error)

  // Good:
  package yamlconfig

  func Parse(input string) (*Config, error)
  ```

- 리시버함수에 리시버 구조체 이름을 사용하지 않는다.

  ```go
  // Bad:
  func (c *Config) WriteConfigTo(w io.Writer) (int64, error)

  // Good:
  func (c *Config) WriteTo(w io.Writer) (int64, error)
  ```

- 파라미터, 변수 등의 이름을 함수 이름에 사용하지 않는다.

  ```go
  // Bad:
  func OverrideFirstWithSecond(dest, source *Config) error

  // Good:
  func Override(dest, source *Config) error
  ```

- 리턴 값의 이름이나 타입을 함수 이름에 사용하지 않는다.

  ```go
  // Bad:
  func TransformYAMLToJSON(input *Config) *jsonconfig.Config

  // Good:
  func Transform(input *Config) *jsonconfig.Config
  ```

- Get을 함수 이름에 쓰는 것을 지양하나 필요한 경우에는 허용한다.

  ```go
  // Good:
  func (c *Config) JobName(key string) (value string, ok bool)

  ```

# Tech Spec

- Backend Language : GO 1.21.5
- Frontend Language : HTML, javascript
- DB : MySQL(ver 8.0.33)
- Cache : Redis(ver 6.2.6)

# Wireframe

[figma]https://www.figma.com/file/Ypk12AKtXHeJqkC4PHklvn/Untitled?type=design&node-id=0%3A1&mode=design&t=wFf0jDXzPFOQ2qVM-1

![스크린샷 2024-01-16 오전 10.32.19.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/41e7b361-b2f9-488e-a1c9-a6387c941d75/6740a8a5-cd8f-4fa5-b1d5-4b94d86ae0ea/%E1%84%89%E1%85%B3%E1%84%8F%E1%85%B3%E1%84%85%E1%85%B5%E1%86%AB%E1%84%89%E1%85%A3%E1%86%BA_2024-01-16_%E1%84%8B%E1%85%A9%E1%84%8C%E1%85%A5%E1%86%AB_10.32.19.png)

![스크린샷 2024-01-16 오전 10.32.35.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/41e7b361-b2f9-488e-a1c9-a6387c941d75/34332050-31f5-45ed-95d1-84ae5329be44/%E1%84%89%E1%85%B3%E1%84%8F%E1%85%B3%E1%84%85%E1%85%B5%E1%86%AB%E1%84%89%E1%85%A3%E1%86%BA_2024-01-16_%E1%84%8B%E1%85%A9%E1%84%8C%E1%85%A5%E1%86%AB_10.32.35.png)

# API

# ERD

![gommunity_ERD.png](https://prod-files-secure.s3.us-west-2.amazonaws.com/41e7b361-b2f9-488e-a1c9-a6387c941d75/601884f7-f569-4da7-9f7f-dff4592cd199/gommunity_ERD.png)

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

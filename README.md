# 📝 Go-Todolist

<p align="center">
Go 언어와 Fyne 라이브러리를 활용하여 개발된 크로스 플랫폼 GUI ToDoList
애플리케이션입니다.
</p>

## ✨ Features

이 애플리케이션은 직관적인 GUI를 통해 일정을 관리할 수 있도록 설계되었으며, 두
개의 주요 탭으로 구성되어 있습니다.

- ➕ Add 탭: 일정 생성

    - 새로운 할 일 항목을 입력하고 저장할 수 있습니다.

- 📋 Display 탭: 일정 확인 및 관리

    - 저장된 전체 할 일 목록을 표시합니다.

    - 목록에서 완료된 항목을 개별적으로 확인하거나 삭제할 수 있습니다.

- 💾 데이터 지속성

    - 모든 일정 데이터는 프로젝트 루트 디렉터리에 위치한 todolist.json 파일에 JSON 형식으로 자동 저장됩니다.


## 🚀 Installation & Execution

이 프로젝트는 Fyne 라이브러리를 사용하므로, 시스템에 기본적인 GUI 라이브러리
의존성 패키지들이 설치되어 있어야 합니다.

#### 1. 시스템 의존성 설치

Fyne 애플리케이션 빌드를 위해 필요한 핵심 라이브러리들을 설치합니다.
(데비안/우분투 계열 기준)

```bash
# 필요한 기본 라이브러리 설치
sudo apt update
sudo apt install libx11-dev libgl1-mesa-dev xorg-dev
```

#### 2. Go 의존성 설정

Go 모듈을 초기화하고 Fyne 라이브러리를 프로젝트에 추가합니다.

```bash
# Go 모듈 관리 활성화 및 Fyne 라이브러리 가져오기
go mod tidy
go get fyne.io/fyne/v2
```

#### 3. PKG_CONFIG_PATH 설정 (Optional)

시스템에 따라 gl.pc 파일 경로를 환경 변수에 추가해야 할 수도 있습니다.

```bash
# gl.pc 파일의 위치를 찾습니다.
find /usr -name gl.pc

# **주의:** 찾은 실제 경로로 대체하여 PKG_CONFIG_PATH에 추가합니다.
# 예시:
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/lib/x86_64-linux-gnu/pkgconfig

# 설정이 올바른지 확인합니다.
pkg-config --modversion gl
```

#### 4. 빌드 및 실행

프로젝트의 루트 디렉터리에서 애플리케이션을 빌드하고 실행합니다.

```bash
# 빌드 (실행 파일 이름은 todolist)
go build -o todolist ./src

# 실행
./todolist
```

## 🔗 Related Technologies
- 언어: [Go(Golang)](https://go.dev/)
- GUI 라이브러리: [Fyne](https://fyne.io/) (크로스 플랫폼 GUI 툴킷)

## 📄 License

이 프로젝트는 MIT 라이선스를 따릅니다.

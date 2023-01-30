# Visual studio

## Extension
Extension 은 설치하고나면 화면 왼쪽에 탭으로 생기는 것과 바로 코드에서
동작하는 것들이 있다.

### 탭 형태로 생성
- docker
- kubernetes
- cloud code
- git lens

### 바로 코드에서 적용
- go
- ctags
- vim
- kubernetes-snippets (yaml 에서 k- 시작)
- markdown short cut (md 파일에서 마우스 우클릭)
- tabnine


## Short cut 
```
F12: Go to Definition
Ctrl + -: Back to last cursor position
Shift + F12: Go to reference
CMD + P: file 찾기

ctrl + ]: tags 에서 정의 찾아서 이동 (F12)
ctrl + t: 이전 위치로 돌아오기

shift + cmd + f: 파일들에서 내용 찾기
```

## setting
- vim 모드에서 key repeat 안될때
```bash
# https://vimforvscode.com/enable-key-repeat-vim
defaults write com.microsoft.VSCode ApplePressAndHoldEnabled -bool false
```

- Code > Preference > Key binding
```
# 예제
F12 --> ctrl + O
Ctrl + - --> ctrl + t
```

## get remote url of current file
```
GitLens Extension 받은뒤에
메뉴 바 - Code - 기본 설정 - 설정

혹시 "gitlens.integrations.enabled"를 검색했을때 체크 안되어있으면 체크하기
"gitlens.remote"를 검색하여 맨 아래 Remote 설정 변경을 위해 settings.json 열기

아래와 같이 입력
"gitlens.remotes": [
    {
        "domain": "github.daumkakao.com",
            "type": "GitHub"
    }
]
"Open File on Remote", "Copy Remote File URL" 기능이 쓸만하다.
```

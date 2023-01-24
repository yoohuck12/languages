# leetcode

## leetcode cli 실행
- https://skygragon.github.io/leetcode-cli/commands#submission
```bash
# https://skygragon.github.io/leetcode-cli/install
$ npm install -g leetcode-cli
$ leetcode version
$ leetcode -h
$ leetcode list
$ leetcode show <num> -g -x
$ leetcode submission -l <lang> <num>
$ leetcode submit <filename>
```

- 그러나 보니까 python3 는 기본적으로 지원이 잘 안되는 것 같음.
submission 에서 python3 를 지원하지 않는 문제들이 많음.


### 설치 및 로그인 trouble shooting
- 다음과 같이 login 을 하려고하는데 warning 때문에 로그인이 안됨.
```bash
❯ leetcode user -l
login: (node:39920) Warning: Accessing non-existent property 'padLevels' of module exports inside circular dependency
(Use `node --trace-warnings ...` to show where the warning was created)

❯ node --trace-warnings ...
node:internal/modules/cjs/loader:1042
  throw err;
  ^

Error: Cannot find module '/Users/kakao/test'
    at Module._resolveFilename (node:internal/modules/cjs/loader:1039:15)
    at Module._load (node:internal/modules/cjs/loader:885:27)
    at Function.executeUserEntryPoint [as runMain] (node:internal/modules/run_main:82:12)
    at node:internal/main/run_main_module:23:47 {
  code: 'MODULE_NOT_FOUND',
  requireStack: []
}
```
- 그래서 아애 소스코드 레벨에서 설치
```bash
$ git clone http://github.com/skygragon/leetcode-cli
$ npm install underscore
$ cd leetcode-cli && node ./bin/pkg
```
- 그러나 같은 warning. 아래를 해도 같음.
```bash
$ npm audict fix --force
```
- 그냥 redirection 으로 해결하고자 했으나.. 그게 문제가 아닌듯.
```
$ leetcode user -l 2> /tmp/err.txt
```
- 검색해보니 [다음](https://github.com/skygragon/leetcode-cli/issues/153#issuecomment-590818198) 과 같이 해결한 경우가 있음.
```
just found a solution and it worked for me.
Steps:
	Login into leetcode on chrome
	install cookie.chrome pluging with command => leetcode plugin -i cookie.chrome
	install libsecret tools =>
	linux (sudo apt-get install libsecret-tools)
	mac ( brew install libsecret )
	now try to log in
Ref link: https://github.com/skygragon/leetcode-cli-plugins/blob/master/docs/cookie.chrome.md
```
- 추가.
```
$ leetcode plugin -i cooke.chrome
$ leetcode plugin -e cookie.chrome
$ brew install libsecret
$ brew upgrade npm
$ 다시 설치
$ leetcode user -l -vv
$ leetcode plugin -d lintcode
$ leetcode plugin -d leetcode.cn
$ leetcode seesion get
leetcode 홈페이지 들어가서 로그인 한 후 F12 눌러서 Network 탭으로 간 후
Graphql 과 같은 것을 누른후 Header 를 확인하면, cookie 에 LEETCODE_SESSION
이 적혀있고, 해당 값을 환경 변수로 세팅하면됨.
# Setting leetcode session
$ LEETCODE_SESSION=""
$ leetcode user -l -vv
```
- warning 을 없애지 않고는 로그인 안되는 것 같음.
```

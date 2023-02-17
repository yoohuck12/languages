# vim

##  vim 명령어 알아두면 좋을 것들
- [line break with N chars](https://stackoverflow.com/questions/1272173/in-vim-how-do-i-break-one-really-long-line-into-multiple-lines)
```
First set your vim so that it understands that you want 80 characters:

:set tw=80
then, hilight the line:

V
and make vim reformat it:

gq
```

## vim 을 강력하게 쓸 수 있게 만드는 marks
vim marks 는 vim 으로 연 파일에 대한 기억(mark)을 하게 만들어주는
단순한 도구다. `m+[a-zA-Z]` 로 등록할 수 있으며 접근시에는 `'+[a-zA-Z]`
로 할 수 있다. `:marks` 를 통해서 등록된 마크를 사용할 수도 있다.

참고로 소문자는 현재 파일에서만 작동하므로, 다른 폴더의 파일을 열기 위해서
대문자로 등록해야 한다.

엄청나게 단순한 도구로 보이지만, 이 기능을 사용하면 git repo 별, 폴더별로
파일을 열고 해당 파일부터 작업을 시작할 수 있다는 장점이 있다. 특히
NerdTree 의 기능과 `:!<command>` 기능까지 함께 쓴다면 어느 위치에서든 원하는
파일을 열고 편집할 수 있다.

.vimrc 에 등록된 다음의 기능이 무척 중요하다.
```
set browsedir = current
```

vim 이 열릴때 현재 폴더를 기준으로 열리게 되는건데, 이 기능을 통해서
열린 파일을 기준으로 NerdTree 가 정렬되고, command 도 해당 파일이 있는
폴더에서 사용할 수 있게 된다.

가장 좋은 사용법은 여러 repo 의 README.md 에 마크를 해두고, vi 를
아무데서나 킨 후 마크로 이동해서 NerdTree 를 켜고 command 를 사용하는 방식이다.

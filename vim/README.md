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

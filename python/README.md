# PYTHON


## ipython
```python3
# 저장은 띄어쓰기로 구분됨
# $save <filename> line line-line
$ %save my_useful_session 10-20 23
```

## pylint
- 기본적으로 사용하고 있는 것.
```python3
pylint -d too-many-arguments,too-many-function-args,logging-fstring-interpolation,missing-function-docstringinvald-name,missing-function-docstring,unused-argument  <file>
```

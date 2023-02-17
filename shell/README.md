# bash


## 용량 가져오기
```bash
du -had1 dir/
(maybe like "do you had 1")

-h: human readable sizes
-a: show files, not just directories
-d1: show totals only at depth 1, i.e. the current directory's contents
```

## 이전 커멘드에서 사용한 변수 가져오기
```bash
$ mkdir a && cd $_
$ echo "asdf" | grep $_ *.py
```

## 결과값을 파이프라인으로 가져오기
```bash
$  kubectl get pods --namespace <namespace> | awk '{ print $1 }' | xargs -I {} kubectl exec --namespace <namespace> {}  -- bash -c "wc -l /tmp/*.csv"
```

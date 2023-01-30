# k8s

## stern
log 를 보는 방법
```bash
$ stern -n <namespace> <pod prefix> --color always -exclude 'GET ' -t --tail 1 > /tmp/XXXX.txt
```

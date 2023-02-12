# k8s

## stern
log 를 보는 방법
```bash
$ stern -n <namespace> <pod prefix> --color always -exclude 'GET ' -t --tail 1 > /tmp/XXXX.txt
```

## access service from pod
- 팟에서 클러스터 내부의 서비스 찾아가기
```
http://<namespace>.<service>.svc.cluster.local:<port>
```

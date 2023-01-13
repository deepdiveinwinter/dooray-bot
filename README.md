# Dooray Messenger Bot Client

Dooray Incoming Hook을 활용해 Dooray Bot이 메세지를 보내도록 할 수 있습니다.<br/>
일정한 주기로 반복적으로 보내야하는 공지사항 발송을 자동화하기 위해 만들어졌습니다.

## 개발 환경 구성
* Go 1.19
* docker-ce 20.10.20

## 로컬 환경에서 실행하는 방법

1. Makefile에 **DOORAY_HOOK_URL** 환경변수의 **{SERVICE_HOOK}** 값을 수정합니다.<br/>
```
DOORAY_HOOK_URL="https://hook.dooray.com/services/{SERVICE_HOOK}"
```

2. 빌드 명령어를 사용해 컨테이너 이미지를 빌드합니다.
```
make build
```

빌드된 이미지를 기반으로 컨테이너를 생성해 메세지를 발송할 수 있습니다.<br/>

3. 실행 명령어를 사용해 dooray-bot 컨테이너를 실행합니다.
```
make run
```
dooray-bot 컨테이너는 메세지 발송 후 정상 종료 시 Exit Code 0을 리턴합니다.<br/>
실행이 완료된 컨테이너는 자동으로 삭제됩니다.<br/>

## Kubernetes 환경에서 실행하는 방법

dooray-bot 컨테이너 이미지를 Job 또는 CronJob으로 생성해 메세지를 보낼 수 있습니다.<br>
**CronJob은 kube-controller-manager의 TimeZone을 기준으로 이벤트가 발생합니다.**

(예제) 매 주 월요일 오전 10시에 특정 메세지를 발송하는 CronJob YAML

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: dooray-bot
spec:
  schedule: "0 10 * * 1"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: dooray-bot
              image: deepdiveinwinter/dooraybot:latest
              env:
                - name: DOORAY_HOOK_URL
                  value: https://hook.dooray.com/services/{SERVICE_HOOK}
          restartPolicy: OnFailure
```

## 오늘 하루

- 하루를 나의 상태에 따라 행동하는 나의 모습을 표현하기 위해 다음의 프로그램을 기획
- 24시간을 분 단위로 하루를 기록
- 특정시간마다 행동 선택이 가능
- 시간, 행동에 따라 내 status가 바뀜

## 행동 목록

- 성장
- 식사
- 화장실
- 숙면
- 놀이
- 씻기
- 취준

## 기능 목록

- [x] status 초기화하기 - init: initStatus()
- [x] 하루 시작하기 - init: startDay()
  - [x] 다음 time으로 가기 - timer: nextTime()
    - [x] 시간에 따른 상태변화 - timer: changeStatusByTime()
    - [x] 시간 계산하기 - timer: getCurrentTime()
    - [x] 시간 출력하기 - outputView: printTime()
  - [x] 현재 상태 진단하기 - diagnoseis: diagnose()
    - [x] 현재 상태 출력하기 - outputView: printStatus()
  - [x] 행동하기 - action: action()
    - [x] 행동 출력하기 - outputView: printAction()
    - [x] 행동 선택하기 - inputView: chooseAction()
- [x] 하루 결과 보여주기 - outputView: printDay()

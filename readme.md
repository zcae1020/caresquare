## 오늘 하루

- 24시간을 분 단위로 하루를 기록
- 특정시간마다 행동 선택이 가능
- 시간, 행동에 따라 내 status가 바뀜

## 기능 목록

- [ ] status 초기화하기 - init: initStatus()
- [ ] 하루 시작하기 - init: startDay()
  - [ ] 다음 time으로 가기 - timer: nextTime()
    - [ ] 시간 계산하기 - timer: getCurrentTime()
    - [ ] 시간 출력하기 - outputView: printTime()
  - [ ] 현재 상태 진단하기 - diagnoseis: diagnose()
    - [ ] 현재 상태 출력하기 - outputView: printStatus()
  - [ ] 행동하기 - action: action()
    - [ ] 행동 출력하기 - outputView: printAction()
    - [ ] 행동 선택하기 - action: chooseAction()
- [ ] 하루 결과 보여주기 - outputView: printDay()
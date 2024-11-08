
# 2390. 문자열에서 별표 제거하기

문자열 `s`가 주어지며, 이 문자열에는 별표 `*`가 포함되어 있습니다.

한 번의 연산에서 다음을 수행할 수 있습니다:
* `s`에서 별표 하나를 선택합니다.
* 해당 별표의 **왼쪽**에서 가장 가까운 **별표가 아닌** 문자와 선택한 별표를 함께 제거합니다.

**모든** 별표를 제거한 후의 문자열을 반환하세요.

참고:
* 입력값은 항상 위의 연산이 가능하도록 주어집니다.
* 결과 문자열은 항상 유일함이 보장됩니다.

예제 1:
```
Input: s = "leet**cod*e"
Output: "lecoe"
설명: 왼쪽에서 오른쪽으로 제거를 수행합니다:
- 첫 번째 별표에 가장 가까운 문자는 "leet**cod*e"의 't'입니다. s는 "lee*cod*e"가 됩니다.
- 두 번째 별표에 가장 가까운 문자는 "lee*cod*e"의 'e'입니다. s는 "lecd*e"가 됩니다.
- 세 번째 별표에 가장 가까운 문자는 "lecd*e"의 'd'입니다. s는 "lecoe"가 됩니다.
더 이상 별표가 없으므로 "lecoe"를 반환합니다.
```

예제 2:
```
Input: s = "erase*****"
Output: ""
설명: 문자열 전체가 제거되므로 빈 문자열을 반환합니다.
```



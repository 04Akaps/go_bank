Paseto가 JWT Token보다 더 뛰어난 이유

기본적으로 더 좋은 보안을 가지고 있기 떄문이다.

JWT는 다음과 같은 문제를 가지고 있다.

1. 알고리즘이 공격에 취약하다.
2. header에 알로기릊믕ㄹ 기록하지 떄문에 실수가 있다면 공격에 취약해진다.
   - 해당 부분은 사실 라이브러리를 사용하면 문제는 없음

Paseto는 다음과 같은 장점이 있다.

1. 알고리즘이 굉장히 튼튼하다. -> AEAD 알고리즘을 사용
2. 헤더에 알고리즘 정보를 저장 하지 않는다.

추가적인 정보는 해당 링크를 참고

# https://github.com/o1egl/paseto#key-differences-between-paseto-and-jwt

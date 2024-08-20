-- https://school.programmers.co.kr/learn/courses/30/lessons/131120, [3월에 태어난 여성 회원 목록 출력하기]

SELECT
    MEMBER_ID,
    MEMBER_NAME,
    GENDER,
    DATE_FORMAT(DATE_OF_BIRTH, '%Y-%m-%d') as DATE_OF_BIRTH
FROM MEMBER_PROFILE
WHERE MONTH(DATE_OF_BIRTH) = 03
    AND TLNO is not NULL
    AND GENDER = 'W'
ORDER BY MEMBER_ID ASC;
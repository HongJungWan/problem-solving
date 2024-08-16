-- https://school.programmers.co.kr/learn/courses/30/lessons/133025 [과일로 만든 아이스크림 고르기]

SELECT F.FLAVOR
FROM FIRST_HALF F
JOIN ICECREAM_INFO I ON F.FLAVOR = I.FLAVOR
WHERE F.FLAVOR IN('strawberry', 'melon')
ORDER BY F.TOTAL_ORDER DESC;
-- https://school.programmers.co.kr/learn/courses/30/lessons/131117 [5월 식품들의 총매출 조회하기]

SELECT E.EMP_NO, EMP_NAME,
        CASE WHEN AVG(SCORE) >= 96 THEN 'S'
            WHEN AVG(SCORE) >= 90 THEN 'A'
            WHEN AVG(SCORE) >= 80 THEN 'B'
            ELSE 'C'
        END AS 'GRADE',
        CASE WHEN AVG(SCORE) >= 96 THEN SAL * 0.2
            WHEN AVG(SCORE) >= 90 THEN SAL * 0.15
            WHEN AVG(SCORE) >= 80 THEN SAL * 0.1
            ELSE 0
        END AS 'BONUS'
FROM HR_EMPLOYEES AS E
LEFT JOIN HR_GRADE
ON E.EMP_NO = HR_GRADE.EMP_NO
GROUP BY E.EMP_NO
ORDER BY EMP_NO ASC;
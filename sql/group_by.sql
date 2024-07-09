[https://school.programmers.co.kr/learn/courses/30/lessons/284528], 연간 평가점수에 해당하는 평가 등급 및 성과금 조회하기

SELECT
    E.EMP_NO,
    EMP_NAME,
    CASE
        WHEN AVG(SCORE) >= 96 THEN 'S'
        WHEN AVG(SCORE) >= 90 THEN 'A'
        WHEN AVG(SCORE) >= 80 THEN 'B'
        ELSE 'C'
    END AS 'GRADE',
    CASE
        WHEN AVG(SCORE) >= 96 THEN SAL * 0.2
        WHEN AVG(SCORE) >= 90 THEN SAL * 0.15
        WHEN AVG(SCORE) >= 80 THEN SAL * 0.1
        ELSE 0
    END AS 'BONUS'
FROM HR_EMPLOYEES AS E
LEFT JOIN HR_GRADE ON E.EMP_NO = HR_GRADE.EMP_NO
GROUP BY E.EMP_NO
ORDER BY EMP_NO ASC;

---

[https://school.programmers.co.kr/learn/courses/30/lessons/276036], 언어별 개발자 분류하기

WITH front_end AS (
    SELECT SUM(CODE)
    FROM SKILLCODES
    WHERE CATEGORY = 'Front End'
)

SELECT CASE
    WHEN SKILL_CODE & (SELECT * FROM front_end)
        AND SKILL_CODE & (SELECT CODE FROM SKILLCODES WHERE NAME = 'python')
        THEN 'A'
    WHEN SKILL_CODE & (SELECT CODE FROM SKILLCODES WHERE NAME = 'C#')
        THEN 'B'
    WHEN SKILL_CODE & (SELECT * FROM front_end)
        THEN 'C'
    END AS GRADE, ID, EMAIL
FROM DEVELOPERS
GROUP BY GRADE, ID, EMAIL
HAVING GRADE IS NOT NULL
ORDER BY GRADE, ID
-- https://school.programmers.co.kr/learn/courses/30/lessons/59409, [중성화 여부 파악하기]

SELECT
    ANIMAL_ID,
    NAME,
    CASE
        WHEN SEX_UPON_INTAKE LIKE 'Intact%'
        THEN 'X'
        ELSE 'O'
    END AS '중성화'
FROM ANIMAL_INS
ORDER BY ANIMAL_ID ASC
-- https://school.programmers.co.kr/learn/courses/30/lessons/59045, [보호소에서 중성화한 동물]

SELECT
    a.ANIMAL_ID,
    a.ANIMAL_TYPE,
    a.NAME
FROM ANIMAL_INS as a
JOIN ANIMAL_OUTS as b
ON a.ANIMAL_ID = b.ANIMAL_ID
WHERE a.SEX_UPON_INTAKE != b.SEX_UPON_OUTCOME
ORDER BY a.ANIMAL_ID ASC;
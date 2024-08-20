-- https://school.programmers.co.kr/learn/courses/30/lessons/59411, [오랜 기간 보호한 동물(2)]

SELECT
    outs.animal_id AS ANIMAL_ID,
    outs.name AS NAME
FROM ANIMAL_INS AS ins
JOIN ANIMAL_OUTS AS outs
ON ins.ANIMAL_ID = outs.ANIMAL_ID
ORDER BY DATEDIFF(outs.DATETIME, ins.DATETIME) DESC limit 2;
-- https://school.programmers.co.kr/learn/courses/30/lessons/164673 [조건에 부합하는 중고거래 댓글 조회하기]

SELECT
    ub.title as TITLE, ub.board_id as BOARD_ID,
    ur.reply_id as REPLY_ID, ur.writer_id as WRITER_ID,ur.contents as CONTENTS,
    DATE_FORMAT(ur.created_date, "%Y-%m-%d") as CREATED_DATE
FROM USED_GOODS_BOARD as ub
JOIN USED_GOODS_REPLY as ur
ON ub.board_id = ur.board_id
WHERE DATE_FORMAT(ub.CREATED_DATE, "%Y-%m") = "2022-10"
ORDER BY ur.created_date asc, ub.title asc
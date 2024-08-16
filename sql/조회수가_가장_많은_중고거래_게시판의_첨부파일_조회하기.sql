-- https://school.programmers.co.kr/learn/courses/30/lessons/164671 [조회수가 가장 많은 중고거래 게시판의 첨부파일 조회하기]

select
    concat('/home/grep/src/', ugf.BOARD_ID, '/', ugf.FILE_ID, ugf.FILE_NAME, ugf.FILE_EXT) as FILE_PATH
from USED_GOODS_BOARD as ugb
join USED_GOODS_FILE as ugf
on ugb.BOARD_ID = ugf.BOARD_ID
where VIEWS = (select MAX(VIEWS) from USED_GOODS_BOARD)
ORDER By ugf.FILE_EXT DESC
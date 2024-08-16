-- https://school.programmers.co.kr/learn/courses/30/lessons/144854 [조건에 맞는 도서와 저자 리스트 출력하기]

SELECT B.BOOK_ID, A.AUTHOR_NAME, DATE_FORMAT(B.PUBLISHED_DATE,"%Y-%m-%d") as PUBLISHED_DATE
FROM BOOk B
JOIN AUTHOR A on B.AUTHOR_ID = A.AUTHOR_ID
WHERE B.CATEGORY = '경제'
ORDER BY B.PUBLISHED_DATE ASC;
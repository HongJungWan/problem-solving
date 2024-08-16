-- https://school.programmers.co.kr/learn/courses/30/lessons/144855 [카테고리 별 도서 판매량 집계하기
]

SELECT B.category as GATEGORY, sum(BS.sales) as TOTAL_SALES
FROM BOOK as B
JOIN BOOK_SALES as BS ON B.BOOK_ID = BS.BOOK_ID
WHERE BS.SALES_DATE like '2022-01%'
GROUP BY B.category
ORDER BY B.category ASC;
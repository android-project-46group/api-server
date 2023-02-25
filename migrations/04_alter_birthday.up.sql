ALTER TABLE member_infos ALTER COLUMN birthday 
TYPE date USING CASE 
    WHEN birthday ~ E'^\\d{1,2}/\\d{1,2}/\\d{4}$' THEN TO_DATE(birthday, 'MM/DD/YYYY')
    WHEN birthday ~ E'^\\d{4}年\\d{1,2}月\\d{1,2}日$' THEN TO_DATE(birthday, 'YYYY年MM月DD日')
END;

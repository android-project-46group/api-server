ALTER TABLE member_infos ALTER COLUMN birthday 
TYPE VARCHAR (11) USING CASE 
    -- locale_id == 1 → en
    WHEN locale_id = '1' THEN TO_CHAR(birthday, 'MM/DD/YYYY')
    -- locale_id == 2 → ja
    WHEN locale_id = '2' THEN TO_CHAR(birthday, 'YYYY年MM月DD日')
END;

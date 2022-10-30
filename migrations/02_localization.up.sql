CREATE TABLE IF NOT EXISTS locales(
    locale_id   serial PRIMARY KEY,
    name        VARCHAR (20) UNIQUE NOT NULL
);

ALTER TABLE member_infos
ADD locale_id INT NOT NULL;

ALTER TABLE member_infos ADD CONSTRAINT fk_locale
FOREIGN KEY(locale_id)
REFERENCES locales(locale_id);

ALTER TABLE member_infos
ADD CONSTRAINT member_info_id_locale_unique
UNIQUE (member_id, locale_id);

ALTER TABLE member_infos
ALTER COLUMN generation TYPE VARCHAR (20),
ALTER COLUMN blood_type TYPE VARCHAR (20);

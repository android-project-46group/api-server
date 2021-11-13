CREATE TABLE IF NOT EXISTS groups(
    group_id    serial PRIMARY KEY,
    group_name  VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS members(
    member_id   serial PRIMARY KEY,
    group_id    INT NOT NULL,
    name_ja     VARCHAR (20) UNIQUE NOT NULL,
    joined_at   TIMESTAMP,
    left_at     TIMESTAMP,
    CONSTRAINT fk_member
        FOREIGN KEY(group_id)
            REFERENCES groups(group_id)
);

CREATE TABLE IF NOT EXISTS member_infos(
    member_info_id  serial PRIMARY KEY,
    member_id       INT NOT NULL,
    birthday        VARCHAR (11) NOT NULL,
    blood_type      VARCHAR (3) NOT NULL,
    height          VARCHAR (10) NOT NULL,
    generation      VARCHAR (10) NOT NULL,
    blog_url        VARCHAR (200),
    img_url         VARCHAR (200),
    CONSTRAINT fk_member_info
        FOREIGN KEY(member_id)
            REFERENCES members(member_id)
);

CREATE TABLE IF NOT EXISTS blogs(
    blog_id         serial PRIMARY KEY,
    member_id       INT NOT NULL,
    blog_url        VARCHAR (200) UNIQUE NOT NULL,
    last_blog_img   VARCHAR (200) NOT NULL,
    last_updated_at VARCHAR (20) NOT NULL,
    CONSTRAINT fk_blog
        FOREIGN KEY(member_id)
            REFERENCES members(member_id)
);

CREATE TABLE IF NOT EXISTS formations(
    formation_id    serial PRIMARY KEY,
    first_row_num   SMALLINT,
    second_row_num  SMALLINT,
    third_row_num   SMALLINT
);

CREATE TABLE IF NOT EXISTS songs(
    song_id         serial PRIMARY KEY,
    group_id        INT NOT NULL,
    formation_id    INT NOT NULL,
    title           VARCHAR (20) UNIQUE NOT NULL,
    single          VARCHAR (10) NOT NULL,
    CONSTRAINT fk_song_g
        FOREIGN KEY(group_id)
            REFERENCES groups(group_id),
    CONSTRAINT fk_song_f
        FOREIGN KEY(formation_id)
            REFERENCES formations(formation_id)
);

CREATE TABLE IF NOT EXISTS positions(
    position_id     serial PRIMARY KEY,
    song_id         INT NOT NULL,
    member_id       INT NOT NULL,
    position        VARCHAR (10) NOT NULL,
    is_center       BOOLEAN,
    CONSTRAINT fk_position_s
        FOREIGN KEY(song_id)
            REFERENCES songs(song_id),
    CONSTRAINT fk_position_m
        FOREIGN KEY(member_id)
            REFERENCES members(member_id)
);

CREATE TABLE IF NOT EXISTS tags(
    tag_id    serial PRIMARY KEY,
    tag_name  VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS member_tags(
    member_tag_id   serial PRIMARY KEY,
    member_id       INT NOT NULL,
    tag_id          INT NOT NULL,
    CONSTRAINT fk_member_tag_m
        FOREIGN KEY(member_id)
            REFERENCES members(member_id),
    CONSTRAINT fk_member_tag_t
        FOREIGN KEY(tag_id)
            REFERENCES tags(tag_id)
);

CREATE TABLE IF NOT EXISTS api_keys(
    key_id   serial PRIMARY KEY,
    key_val VARCHAR(70) UNIQUE NOT NULL
);

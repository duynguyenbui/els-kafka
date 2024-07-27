--- Create the hotels Database
CREATE DATABASE hotels
    WITH OWNER = 'teknix'
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

\c hotels

CREATE TABLE "hotels" (
  "address" text,
  "amenity_groups" json,
  "check_in_time" text,
  "check_out_time" text,
  "description_struct" json,
  "id" text,
  "images" json,
  "kind" text,
  "latitude" double precision,
  "longitude" double precision,
  "name" text,
  "phone" text NULL,
  "policy_struct" json,
  "postal_code" text,
  "room_groups" json,
  "region" json,
  "star_rating" bigint,
  "email" text,
  "serp_filters" json,
  "is_closed" boolean,
  "is_gender_specification_required" boolean,
  "metapolicy_struct" json,
  "metapolicy_extra_info" text NULL,
  "star_certificate" text NULL,
  "facts" json,
  "payment_methods" json,
  "hotel_chain" text,
  "front_desk_time_start" text NULL,
  "front_desk_time_end" text NULL,
  "semantic_version" bigint
);
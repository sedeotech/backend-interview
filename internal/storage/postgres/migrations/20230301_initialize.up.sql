CREATE EXTENSION IF NOT EXISTS "pgcrypto";

/* Table 'address' */
CREATE TABLE IF NOT EXISTS public.address (
   uuid UUID DEFAULT gen_random_uuid() NOT NULL,
   line1 text NOT NULL,
   line2 text,
   line3 text,
   city text NOT NULL,
   province text,
   zipcode text NOT NULL,
   country_code text NOT NULL,
   other text,
   created timestamptz NOT NULL,
   updated timestamptz NOT NULL,
   PRIMARY KEY(uuid)
);

CREATE TABLE IF NOT EXISTS public.item (
	uuid UUID DEFAULT gen_random_uuid() NOT NULL,
	name text NOT NULL,
	description text,
	price double precision NOT NULL,
	created timestamptz NOT NULL,
	updated timestamptz NOT NULL,
	PRIMARY KEY(uuid)
);

/* Table cart and item relation not done 😉 */
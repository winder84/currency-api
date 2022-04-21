CREATE TABLE IF NOT EXISTS public.currencies
(
    currency_from character(3) COLLATE pg_catalog."default" NOT NULL,
    currency_to character(3) COLLATE pg_catalog."default" NOT NULL,
    well double precision NOT NULL DEFAULT 0,
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT currencies_pkey PRIMARY KEY (currency_from, currency_to)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.currencies
    OWNER to postgres;
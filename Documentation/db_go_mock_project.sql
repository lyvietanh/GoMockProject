toc.dat                                                                                             0000600 0004000 0002000 00000035223 14317501275 0014451 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        PGDMP       -                	    z            db_go_mock_project     12.10 (Debian 12.10-1.pgdg110+1)    14.4 +    ?           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false         ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false         ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false         ?           1262    73728    db_go_mock_project    DATABASE     f   CREATE DATABASE db_go_mock_project WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';
 "   DROP DATABASE db_go_mock_project;
                postgres    false         ?            1259    82228    booking_flights    TABLE     ?   CREATE TABLE public.booking_flights (
    id bigint NOT NULL,
    booking_transaction_id bigint,
    flight_id bigint,
    currency_code character varying(3) NOT NULL,
    price double precision NOT NULL
);
 #   DROP TABLE public.booking_flights;
       public         heap    postgres    false         ?            1259    82226    booking_flights_id_seq    SEQUENCE        CREATE SEQUENCE public.booking_flights_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.booking_flights_id_seq;
       public          postgres    false    209         ?           0    0    booking_flights_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.booking_flights_id_seq OWNED BY public.booking_flights.id;
          public          postgres    false    208         ?            1259    82236    booking_passengers    TABLE     f  CREATE TABLE public.booking_passengers (
    id bigint NOT NULL,
    booking_transaction_id bigint,
    title character varying(8),
    last_name character varying(25),
    first_name character varying(25),
    date_of_birth date,
    travel_doc_type character varying(1),
    travel_doc_number character varying(16),
    ffp_number character varying(16)
);
 &   DROP TABLE public.booking_passengers;
       public         heap    postgres    false         ?            1259    82234    booking_passengers_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.booking_passengers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 0   DROP SEQUENCE public.booking_passengers_id_seq;
       public          postgres    false    211         ?           0    0    booking_passengers_id_seq    SEQUENCE OWNED BY     W   ALTER SEQUENCE public.booking_passengers_id_seq OWNED BY public.booking_passengers.id;
          public          postgres    false    210         ?            1259    82219    booking_transactions    TABLE     ?  CREATE TABLE public.booking_transactions (
    id bigint NOT NULL,
    reservation_code character varying(8),
    status character varying(32),
    error_code integer,
    error_message character varying(256),
    contact_email character varying(64) NOT NULL,
    contact_phone character varying(16) NOT NULL,
    currency_code character varying(3),
    total_price double precision,
    create_date timestamp without time zone DEFAULT now(),
    modify_date timestamp without time zone
);
 (   DROP TABLE public.booking_transactions;
       public         heap    postgres    false         ?            1259    82217    booking_transactions_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.booking_transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 2   DROP SEQUENCE public.booking_transactions_id_seq;
       public          postgres    false    207         ?           0    0    booking_transactions_id_seq    SEQUENCE OWNED BY     [   ALTER SEQUENCE public.booking_transactions_id_seq OWNED BY public.booking_transactions.id;
          public          postgres    false    206         ?            1259    73835    flights    TABLE     ?  CREATE TABLE public.flights (
    id bigint NOT NULL,
    origin_airport_code character varying(3) NOT NULL,
    destination_airport_code character varying(3) NOT NULL,
    departure_date_time timestamp without time zone NOT NULL,
    booking_class character varying(16) NOT NULL,
    flight_number character varying(8) NOT NULL,
    flight_duration integer NOT NULL,
    enabled boolean DEFAULT true NOT NULL,
    seat_available integer DEFAULT 10 NOT NULL,
    currency_code character varying(3) NOT NULL,
    price double precision DEFAULT 1000000 NOT NULL,
    create_by character varying(64),
    create_date timestamp without time zone DEFAULT now(),
    modify_by character varying(64),
    modify_date timestamp without time zone
);
    DROP TABLE public.flights;
       public         heap    postgres    false         ?            1259    73833    flights_id_seq    SEQUENCE     w   CREATE SEQUENCE public.flights_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.flights_id_seq;
       public          postgres    false    203         ?           0    0    flights_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.flights_id_seq OWNED BY public.flights.id;
          public          postgres    false    202         ?            1259    82176    roles    TABLE     ?   CREATE TABLE public.roles (
    id character varying(32) NOT NULL,
    name character varying(32),
    enabled boolean DEFAULT true NOT NULL,
    create_date timestamp without time zone DEFAULT now(),
    modify_date timestamp without time zone
);
    DROP TABLE public.roles;
       public         heap    postgres    false         ?            1259    82183    users    TABLE     ?  CREATE TABLE public.users (
    email character varying(64) NOT NULL,
    role_id character varying(32),
    title character varying(8),
    last_name character varying(25),
    first_name character varying(25),
    date_of_birth date,
    phone character varying(16),
    travel_doc_type character varying(1),
    travel_doc_number character varying(16),
    ffp_number character varying(16),
    password_salt character varying(16),
    secure_password character varying(256),
    enabled boolean DEFAULT true NOT NULL,
    create_date timestamp without time zone DEFAULT now(),
    modify_date timestamp without time zone,
    last_login_date timestamp without time zone
);
    DROP TABLE public.users;
       public         heap    postgres    false         0           2604    82231    booking_flights id    DEFAULT     x   ALTER TABLE ONLY public.booking_flights ALTER COLUMN id SET DEFAULT nextval('public.booking_flights_id_seq'::regclass);
 A   ALTER TABLE public.booking_flights ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    208    209    209         1           2604    82239    booking_passengers id    DEFAULT     ~   ALTER TABLE ONLY public.booking_passengers ALTER COLUMN id SET DEFAULT nextval('public.booking_passengers_id_seq'::regclass);
 D   ALTER TABLE public.booking_passengers ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    211    210    211         .           2604    82222    booking_transactions id    DEFAULT     ?   ALTER TABLE ONLY public.booking_transactions ALTER COLUMN id SET DEFAULT nextval('public.booking_transactions_id_seq'::regclass);
 F   ALTER TABLE public.booking_transactions ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    206    207    207         %           2604    73838 
   flights id    DEFAULT     h   ALTER TABLE ONLY public.flights ALTER COLUMN id SET DEFAULT nextval('public.flights_id_seq'::regclass);
 9   ALTER TABLE public.flights ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202    203         ?          0    82228    booking_flights 
   TABLE DATA           f   COPY public.booking_flights (id, booking_transaction_id, flight_id, currency_code, price) FROM stdin;
    public          postgres    false    209       3017.dat ?          0    82236    booking_passengers 
   TABLE DATA           ?   COPY public.booking_passengers (id, booking_transaction_id, title, last_name, first_name, date_of_birth, travel_doc_type, travel_doc_number, ffp_number) FROM stdin;
    public          postgres    false    211       3019.dat ?          0    82219    booking_transactions 
   TABLE DATA           ?   COPY public.booking_transactions (id, reservation_code, status, error_code, error_message, contact_email, contact_phone, currency_code, total_price, create_date, modify_date) FROM stdin;
    public          postgres    false    207       3015.dat ?          0    73835    flights 
   TABLE DATA           ?   COPY public.flights (id, origin_airport_code, destination_airport_code, departure_date_time, booking_class, flight_number, flight_duration, enabled, seat_available, currency_code, price, create_by, create_date, modify_by, modify_date) FROM stdin;
    public          postgres    false    203       3011.dat ?          0    82176    roles 
   TABLE DATA           L   COPY public.roles (id, name, enabled, create_date, modify_date) FROM stdin;
    public          postgres    false    204       3012.dat ?          0    82183    users 
   TABLE DATA           ?   COPY public.users (email, role_id, title, last_name, first_name, date_of_birth, phone, travel_doc_type, travel_doc_number, ffp_number, password_salt, secure_password, enabled, create_date, modify_date, last_login_date) FROM stdin;
    public          postgres    false    205       3013.dat ?           0    0    booking_flights_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('public.booking_flights_id_seq', 3, true);
          public          postgres    false    208         ?           0    0    booking_passengers_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.booking_passengers_id_seq', 6, true);
          public          postgres    false    210         ?           0    0    booking_transactions_id_seq    SEQUENCE SET     I   SELECT pg_catalog.setval('public.booking_transactions_id_seq', 3, true);
          public          postgres    false    206         ?           0    0    flights_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.flights_id_seq', 15, true);
          public          postgres    false    202         =           2606    82233 $   booking_flights booking_flights_pkey 
   CONSTRAINT     b   ALTER TABLE ONLY public.booking_flights
    ADD CONSTRAINT booking_flights_pkey PRIMARY KEY (id);
 N   ALTER TABLE ONLY public.booking_flights DROP CONSTRAINT booking_flights_pkey;
       public            postgres    false    209         ?           2606    82241 *   booking_passengers booking_passengers_pkey 
   CONSTRAINT     h   ALTER TABLE ONLY public.booking_passengers
    ADD CONSTRAINT booking_passengers_pkey PRIMARY KEY (id);
 T   ALTER TABLE ONLY public.booking_passengers DROP CONSTRAINT booking_passengers_pkey;
       public            postgres    false    211         ;           2606    82225 .   booking_transactions booking_transactions_pkey 
   CONSTRAINT     l   ALTER TABLE ONLY public.booking_transactions
    ADD CONSTRAINT booking_transactions_pkey PRIMARY KEY (id);
 X   ALTER TABLE ONLY public.booking_transactions DROP CONSTRAINT booking_transactions_pkey;
       public            postgres    false    207         3           2606    73844    flights flights_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.flights DROP CONSTRAINT flights_pkey;
       public            postgres    false    203         5           2606    82182    roles roles_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_pkey;
       public            postgres    false    204         7           2606    82191    users users_ffp_number_key 
   CONSTRAINT     [   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_ffp_number_key UNIQUE (ffp_number);
 D   ALTER TABLE ONLY public.users DROP CONSTRAINT users_ffp_number_key;
       public            postgres    false    205         9           2606    82189    users users_pkey 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (email);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    205         A           2606    82242 ;   booking_flights booking_flights_booking_transaction_id_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.booking_flights
    ADD CONSTRAINT booking_flights_booking_transaction_id_fkey FOREIGN KEY (booking_transaction_id) REFERENCES public.booking_transactions(id);
 e   ALTER TABLE ONLY public.booking_flights DROP CONSTRAINT booking_flights_booking_transaction_id_fkey;
       public          postgres    false    2875    209    207         B           2606    82247 .   booking_flights booking_flights_flight_id_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.booking_flights
    ADD CONSTRAINT booking_flights_flight_id_fkey FOREIGN KEY (flight_id) REFERENCES public.flights(id);
 X   ALTER TABLE ONLY public.booking_flights DROP CONSTRAINT booking_flights_flight_id_fkey;
       public          postgres    false    203    2867    209         C           2606    82252 A   booking_passengers booking_passengers_booking_transaction_id_fkey    FK CONSTRAINT     ?   ALTER TABLE ONLY public.booking_passengers
    ADD CONSTRAINT booking_passengers_booking_transaction_id_fkey FOREIGN KEY (booking_transaction_id) REFERENCES public.booking_transactions(id);
 k   ALTER TABLE ONLY public.booking_passengers DROP CONSTRAINT booking_passengers_booking_transaction_id_fkey;
       public          postgres    false    207    211    2875         @           2606    82207    users users_role_id_fkey    FK CONSTRAINT     w   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_role_id_fkey;
       public          postgres    false    205    2869    204                                                                                                                                                                                                                                                                                                                                                                                     3017.dat                                                                                            0000600 0004000 0002000 00000000073 14317501275 0014251 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	1	8	VND	2850000
2	2	8	VND	2850000
3	3	8	VND	2850000
\.


                                                                                                                                                                                                                                                                                                                                                                                                                                                                     3019.dat                                                                                            0000600 0004000 0002000 00000000405 14317501275 0014252 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	1	MR	LY	VIET ANH	1986-10-30	P	B12345678	9876543210
2	1	MISS	LY	HA MY	0001-01-01			
3	2	MR	LY	VIET ANH	1986-10-30	P	B12345678	9876543210
4	2	MISS	LY	HA MY	0001-01-01			
5	3	MR 	LY	VIET ANH	1986-10-30	P	B12345678	9876543210
6	3	MISS	LY	HA MY	0001-01-01			
\.


                                                                                                                                                                                                                                                           3015.dat                                                                                            0000600 0004000 0002000 00000000555 14317501275 0014254 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	2W2KD6	BOOKED	\N	\N	anhlv11@fsoft.com.vn	0915636882	VND	5700000	2022-10-03 23:35:03.129811	2022-10-03 23:35:05.153769
2	AMF88K	BOOKED	\N	\N	anhlv11@fsoft.com.vn	0915636882	VND	5700000	2022-10-06 11:14:41.419059	2022-10-06 11:14:43.462159
3	ASDYJL	BOOKED	\N	\N	anhlv11@fsoft.com.vn	0915636882	VND	5700000	2022-10-06 11:14:46.497372	2022-10-06 11:14:48.539062
\.


                                                                                                                                                   3011.dat                                                                                            0000600 0004000 0002000 00000002544 14317501275 0014250 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        1	HAN	SGN	2022-10-10 10:00:00	ECONOMY	VN209	120	t	10	VND	1000000	test	2022-09-18 15:40:38.650112	\N	\N
3	HAN	SGN	2022-10-10 11:00:00	ECONOMY	VN210	120	t	10	VND	1000000	test	2022-09-18 16:44:58.984884	\N	\N
4	HAN	SGN	2022-10-10 11:00:00	ECONOMY	VN211	120	t	10	VND	1000000	test	2022-09-18 16:45:19.91495	\N	\N
5	HAN	SGN	2022-10-10 11:00:00	ECONOMY	VN212	120	t	10	VND	1500000	test	2022-09-18 16:45:54.562725	\N	\N
6	HAN	SGN	2022-10-10 11:00:00	ECONOMY	VN213	120	t	10	VND	1750000	test	2022-09-18 16:46:00.879596	\N	\N
9	HAN	SGN	2022-10-10 11:00:00	BUSINESS	VN214	120	t	10	VND	3750000	test	2022-09-18 17:59:38.133089	\N	\N
10	HAN	SGN	2022-10-10 11:00:00	BUSINESS	VN215	120	t	10	VND	3750000	test	2022-09-18 20:02:03.619756	\N	\N
11	HAN	SGN	2022-10-10 11:00:00	BUSINESS	VN216	120	t	10	VND	3750000	test	2022-09-18 21:27:11.124935	\N	\N
12	HAN	SGN	2022-10-10 11:00:00	BUSINESS	VN217	120	t	10	VND	3750000	test	2022-09-20 00:03:21.987789	\N	\N
13	HAN	SGN	2022-10-10 12:00:00	BUSINESS	VN218	120	t	10	VND	3750000	test	2022-09-20 00:03:44.55438	\N	\N
8	HAN	SGN	2022-10-10 11:00:00	BUSINESS	VN213	130	t	6	VND	2850000	test	2022-09-18 16:46:25.622138	test	2022-10-06 11:14:47.515099
14	HAN	SGN	2022-10-10 12:00:00	BUSINESS	VN219	120	t	10	VND	3750000	test	2022-09-30 00:18:36.51234	\N	\N
15	HAN	SGN	2022-10-10 12:00:00	BUSINESS	VN220	120	t	10	VND	3750000	test	2022-09-30 00:18:43.438614	\N	\N
\.


                                                                                                                                                            3012.dat                                                                                            0000600 0004000 0002000 00000000166 14317501275 0014247 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        ADMINISTRATOR	Administrator	t	2022-10-03 16:24:21.434132	\N
SUBSCRIBER	Subcriber	t	2022-10-03 16:24:43.193741	\N
\.


                                                                                                                                                                                                                                                                                                                                                                                                          3013.dat                                                                                            0000600 0004000 0002000 00000001215 14317501275 0014244 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        anhlv11@fsoft.com.vn	ADMINISTRATOR	MR	LY	VIET ANH	1986-02-08	0915636882	\N	\N	\N	rkBMPhH6MNEmRNY6	77efa95e83f2843e5dc9608e84944288dc0420811d42dd90333f2793cf5f6b67	t	2022-10-03 16:44:13.607858	\N	\N
test@fsoft.com.vn	ADMINISTRATOR	MRS	TEST	VIET ANH	1999-10-30	123456789	P	B12345678	9876543210	LHTgYRnonlAOBc3z	71e08ad8aefe6a6ee59c53e7a8c60dec0df180735030da8c884ff776544bc668	t	2022-10-06 11:19:15.729068	\N	\N
subscriber@fsoft.com.vn	SUBSCRIBER	MRS	TEST	SUBSCRIBER	1999-10-30	123456789	P	B12345678	9876543211	WB3MfGFfUKsB9roe	f7edea6d35c85de6b523bec491ea05375558ef1e28688daf1cf0c6fec389f917	t	2022-10-06 11:20:52.43056	2022-10-06 11:53:07.798016	\N
\.


                                                                                                                                                                                                                                                                                                                                                                                   restore.sql                                                                                         0000600 0004000 0002000 00000030164 14317501275 0015375 0                                                                                                    ustar 00postgres                        postgres                        0000000 0000000                                                                                                                                                                        --
-- NOTE:
--
-- File paths need to be edited. Search for $$PATH$$ and
-- replace it with the path to the directory containing
-- the extracted data files.
--
--
-- PostgreSQL database dump
--

-- Dumped from database version 12.10 (Debian 12.10-1.pgdg110+1)
-- Dumped by pg_dump version 14.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE db_go_mock_project;
--
-- Name: db_go_mock_project; Type: DATABASE; Schema: -; Owner: -
--

CREATE DATABASE db_go_mock_project WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


\connect db_go_mock_project

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: booking_flights; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.booking_flights (
    id bigint NOT NULL,
    booking_transaction_id bigint,
    flight_id bigint,
    currency_code character varying(3) NOT NULL,
    price double precision NOT NULL
);


--
-- Name: booking_flights_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.booking_flights_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: booking_flights_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.booking_flights_id_seq OWNED BY public.booking_flights.id;


--
-- Name: booking_passengers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.booking_passengers (
    id bigint NOT NULL,
    booking_transaction_id bigint,
    title character varying(8),
    last_name character varying(25),
    first_name character varying(25),
    date_of_birth date,
    travel_doc_type character varying(1),
    travel_doc_number character varying(16),
    ffp_number character varying(16)
);


--
-- Name: booking_passengers_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.booking_passengers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: booking_passengers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.booking_passengers_id_seq OWNED BY public.booking_passengers.id;


--
-- Name: booking_transactions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.booking_transactions (
    id bigint NOT NULL,
    reservation_code character varying(8),
    status character varying(32),
    error_code integer,
    error_message character varying(256),
    contact_email character varying(64) NOT NULL,
    contact_phone character varying(16) NOT NULL,
    currency_code character varying(3),
    total_price double precision,
    create_date timestamp without time zone DEFAULT now(),
    modify_date timestamp without time zone
);


--
-- Name: booking_transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.booking_transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: booking_transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.booking_transactions_id_seq OWNED BY public.booking_transactions.id;


--
-- Name: flights; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.flights (
    id bigint NOT NULL,
    origin_airport_code character varying(3) NOT NULL,
    destination_airport_code character varying(3) NOT NULL,
    departure_date_time timestamp without time zone NOT NULL,
    booking_class character varying(16) NOT NULL,
    flight_number character varying(8) NOT NULL,
    flight_duration integer NOT NULL,
    enabled boolean DEFAULT true NOT NULL,
    seat_available integer DEFAULT 10 NOT NULL,
    currency_code character varying(3) NOT NULL,
    price double precision DEFAULT 1000000 NOT NULL,
    create_by character varying(64),
    create_date timestamp without time zone DEFAULT now(),
    modify_by character varying(64),
    modify_date timestamp without time zone
);


--
-- Name: flights_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.flights_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: flights_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.flights_id_seq OWNED BY public.flights.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.roles (
    id character varying(32) NOT NULL,
    name character varying(32),
    enabled boolean DEFAULT true NOT NULL,
    create_date timestamp without time zone DEFAULT now(),
    modify_date timestamp without time zone
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    email character varying(64) NOT NULL,
    role_id character varying(32),
    title character varying(8),
    last_name character varying(25),
    first_name character varying(25),
    date_of_birth date,
    phone character varying(16),
    travel_doc_type character varying(1),
    travel_doc_number character varying(16),
    ffp_number character varying(16),
    password_salt character varying(16),
    secure_password character varying(256),
    enabled boolean DEFAULT true NOT NULL,
    create_date timestamp without time zone DEFAULT now(),
    modify_date timestamp without time zone,
    last_login_date timestamp without time zone
);


--
-- Name: booking_flights id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_flights ALTER COLUMN id SET DEFAULT nextval('public.booking_flights_id_seq'::regclass);


--
-- Name: booking_passengers id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_passengers ALTER COLUMN id SET DEFAULT nextval('public.booking_passengers_id_seq'::regclass);


--
-- Name: booking_transactions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_transactions ALTER COLUMN id SET DEFAULT nextval('public.booking_transactions_id_seq'::regclass);


--
-- Name: flights id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flights ALTER COLUMN id SET DEFAULT nextval('public.flights_id_seq'::regclass);


--
-- Data for Name: booking_flights; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.booking_flights (id, booking_transaction_id, flight_id, currency_code, price) FROM stdin;
\.
COPY public.booking_flights (id, booking_transaction_id, flight_id, currency_code, price) FROM '$$PATH$$/3017.dat';

--
-- Data for Name: booking_passengers; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.booking_passengers (id, booking_transaction_id, title, last_name, first_name, date_of_birth, travel_doc_type, travel_doc_number, ffp_number) FROM stdin;
\.
COPY public.booking_passengers (id, booking_transaction_id, title, last_name, first_name, date_of_birth, travel_doc_type, travel_doc_number, ffp_number) FROM '$$PATH$$/3019.dat';

--
-- Data for Name: booking_transactions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.booking_transactions (id, reservation_code, status, error_code, error_message, contact_email, contact_phone, currency_code, total_price, create_date, modify_date) FROM stdin;
\.
COPY public.booking_transactions (id, reservation_code, status, error_code, error_message, contact_email, contact_phone, currency_code, total_price, create_date, modify_date) FROM '$$PATH$$/3015.dat';

--
-- Data for Name: flights; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.flights (id, origin_airport_code, destination_airport_code, departure_date_time, booking_class, flight_number, flight_duration, enabled, seat_available, currency_code, price, create_by, create_date, modify_by, modify_date) FROM stdin;
\.
COPY public.flights (id, origin_airport_code, destination_airport_code, departure_date_time, booking_class, flight_number, flight_duration, enabled, seat_available, currency_code, price, create_by, create_date, modify_by, modify_date) FROM '$$PATH$$/3011.dat';

--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.roles (id, name, enabled, create_date, modify_date) FROM stdin;
\.
COPY public.roles (id, name, enabled, create_date, modify_date) FROM '$$PATH$$/3012.dat';

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (email, role_id, title, last_name, first_name, date_of_birth, phone, travel_doc_type, travel_doc_number, ffp_number, password_salt, secure_password, enabled, create_date, modify_date, last_login_date) FROM stdin;
\.
COPY public.users (email, role_id, title, last_name, first_name, date_of_birth, phone, travel_doc_type, travel_doc_number, ffp_number, password_salt, secure_password, enabled, create_date, modify_date, last_login_date) FROM '$$PATH$$/3013.dat';

--
-- Name: booking_flights_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.booking_flights_id_seq', 3, true);


--
-- Name: booking_passengers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.booking_passengers_id_seq', 6, true);


--
-- Name: booking_transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.booking_transactions_id_seq', 3, true);


--
-- Name: flights_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.flights_id_seq', 15, true);


--
-- Name: booking_flights booking_flights_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_flights
    ADD CONSTRAINT booking_flights_pkey PRIMARY KEY (id);


--
-- Name: booking_passengers booking_passengers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_passengers
    ADD CONSTRAINT booking_passengers_pkey PRIMARY KEY (id);


--
-- Name: booking_transactions booking_transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_transactions
    ADD CONSTRAINT booking_transactions_pkey PRIMARY KEY (id);


--
-- Name: flights flights_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.flights
    ADD CONSTRAINT flights_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: users users_ffp_number_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_ffp_number_key UNIQUE (ffp_number);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (email);


--
-- Name: booking_flights booking_flights_booking_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_flights
    ADD CONSTRAINT booking_flights_booking_transaction_id_fkey FOREIGN KEY (booking_transaction_id) REFERENCES public.booking_transactions(id);


--
-- Name: booking_flights booking_flights_flight_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_flights
    ADD CONSTRAINT booking_flights_flight_id_fkey FOREIGN KEY (flight_id) REFERENCES public.flights(id);


--
-- Name: booking_passengers booking_passengers_booking_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.booking_passengers
    ADD CONSTRAINT booking_passengers_booking_transaction_id_fkey FOREIGN KEY (booking_transaction_id) REFERENCES public.booking_transactions(id);


--
-- Name: users users_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- PostgreSQL database dump complete
--

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            
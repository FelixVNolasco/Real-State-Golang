--
-- PostgreSQL database dump
--

-- Dumped from database version 12.14
-- Dumped by pg_dump version 12.14

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
-- Name: agents; Type: TABLE; Schema: public; Owner: felixvnolasco
--

CREATE TABLE public.agents (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text NOT NULL,
    last_name text NOT NULL,
    phone_number text NOT NULL,
    email text NOT NULL
);


ALTER TABLE public.agents OWNER TO felixvnolasco;

--
-- Name: agents_id_seq; Type: SEQUENCE; Schema: public; Owner: felixvnolasco
--

CREATE SEQUENCE public.agents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.agents_id_seq OWNER TO felixvnolasco;

--
-- Name: agents_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: felixvnolasco
--

ALTER SEQUENCE public.agents_id_seq OWNED BY public.agents.id;


--
-- Name: house_details; Type: TABLE; Schema: public; Owner: felixvnolasco
--

CREATE TABLE public.house_details (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    house_id bigint,
    rooms bigint,
    bathrooms text,
    parking_lot text,
    sq_mt text
);


ALTER TABLE public.house_details OWNER TO felixvnolasco;

--
-- Name: house_details_id_seq; Type: SEQUENCE; Schema: public; Owner: felixvnolasco
--

CREATE SEQUENCE public.house_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.house_details_id_seq OWNER TO felixvnolasco;

--
-- Name: house_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: felixvnolasco
--

ALTER SEQUENCE public.house_details_id_seq OWNED BY public.house_details.id;


--
-- Name: house_galleries; Type: TABLE; Schema: public; Owner: felixvnolasco
--

CREATE TABLE public.house_galleries (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    house_id bigint
);


ALTER TABLE public.house_galleries OWNER TO felixvnolasco;

--
-- Name: house_galleries_id_seq; Type: SEQUENCE; Schema: public; Owner: felixvnolasco
--

CREATE SEQUENCE public.house_galleries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.house_galleries_id_seq OWNER TO felixvnolasco;

--
-- Name: house_galleries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: felixvnolasco
--

ALTER SEQUENCE public.house_galleries_id_seq OWNED BY public.house_galleries.id;


--
-- Name: houses; Type: TABLE; Schema: public; Owner: felixvnolasco
--

CREATE TABLE public.houses (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title text NOT NULL,
    price text NOT NULL,
    location text NOT NULL,
    description text NOT NULL,
    status boolean DEFAULT true,
    agent_id bigint,
    house_details_id bigint,
    photo_url text
);


ALTER TABLE public.houses OWNER TO felixvnolasco;

--
-- Name: houses_id_seq; Type: SEQUENCE; Schema: public; Owner: felixvnolasco
--

CREATE SEQUENCE public.houses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.houses_id_seq OWNER TO felixvnolasco;

--
-- Name: houses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: felixvnolasco
--

ALTER SEQUENCE public.houses_id_seq OWNED BY public.houses.id;


--
-- Name: photos; Type: TABLE; Schema: public; Owner: felixvnolasco
--

CREATE TABLE public.photos (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    house_gallery_id bigint,
    photo_url text
);


ALTER TABLE public.photos OWNER TO felixvnolasco;

--
-- Name: photos_id_seq; Type: SEQUENCE; Schema: public; Owner: felixvnolasco
--

CREATE SEQUENCE public.photos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.photos_id_seq OWNER TO felixvnolasco;

--
-- Name: photos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: felixvnolasco
--

ALTER SEQUENCE public.photos_id_seq OWNED BY public.photos.id;


--
-- Name: agents id; Type: DEFAULT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.agents ALTER COLUMN id SET DEFAULT nextval('public.agents_id_seq'::regclass);


--
-- Name: house_details id; Type: DEFAULT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.house_details ALTER COLUMN id SET DEFAULT nextval('public.house_details_id_seq'::regclass);


--
-- Name: house_galleries id; Type: DEFAULT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.house_galleries ALTER COLUMN id SET DEFAULT nextval('public.house_galleries_id_seq'::regclass);


--
-- Name: houses id; Type: DEFAULT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.houses ALTER COLUMN id SET DEFAULT nextval('public.houses_id_seq'::regclass);


--
-- Name: photos id; Type: DEFAULT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.photos ALTER COLUMN id SET DEFAULT nextval('public.photos_id_seq'::regclass);


--
-- Data for Name: agents; Type: TABLE DATA; Schema: public; Owner: felixvnolasco
--

COPY public.agents (id, created_at, updated_at, deleted_at, first_name, last_name, phone_number, email) FROM stdin;
1	2023-04-22 21:00:04.15068+00	2023-04-22 21:00:04.15068+00	2023-04-23 19:23:13.246631+00	Felix Enrique	Vega Nolasco	5593029304	felix@correo.com
2	2023-04-23 19:14:07.950839+00	2023-04-23 19:14:07.950839+00	2023-04-23 19:23:36.612359+00	Felix Enrique	Vega Nolasco	5593029304	felix@correo.com
3	2023-04-23 19:15:34.681045+00	2023-04-23 22:52:47.008259+00	\N	Ana Erika	Vega Nolasco	+52 55 1333 1232	aevega3@gmail.com
\.


--
-- Data for Name: house_details; Type: TABLE DATA; Schema: public; Owner: felixvnolasco
--

COPY public.house_details (id, created_at, updated_at, deleted_at, house_id, rooms, bathrooms, parking_lot, sq_mt) FROM stdin;
3	2023-04-24 18:35:55.602176+00	2023-04-24 18:35:55.602176+00	\N	3	2	1.5	0	69.36 m2
4	2023-04-24 18:36:21.12285+00	2023-04-24 18:36:21.12285+00	\N	4	3	1.5	0	104.04 m2
5	2023-04-24 18:37:05.680132+00	2023-04-24 18:37:05.680132+00	\N	5	2	1.5	1	82.36 m2
6	2023-04-24 18:37:30.077194+00	2023-04-24 18:37:30.077194+00	\N	6	3	1.5	1	122.66 m2
7	2023-04-24 18:40:39.353731+00	2023-04-24 18:40:39.353731+00	\N	7	3	2.5	2	125.84 m2
8	2023-04-24 18:41:48.009617+00	2023-04-24 18:41:48.009617+00	\N	8	3	2.5	2	159.49 m2
9	2023-04-24 18:43:22.947079+00	2023-04-24 18:43:22.947079+00	\N	9	2	1.5	2	82.6 m2
10	2023-04-24 18:44:15.616393+00	2023-04-24 18:44:15.616393+00	\N	10	2	1.5	3	82.6 m2
11	2023-04-24 18:44:56.791351+00	2023-04-24 18:44:56.791351+00	\N	11	3	3.5	3	178.70 m2
12	2023-04-24 18:45:25.033129+00	2023-04-24 18:45:25.033129+00	\N	12	3	3.5	3	235.47 m2
\.


--
-- Data for Name: house_galleries; Type: TABLE DATA; Schema: public; Owner: felixvnolasco
--

COPY public.house_galleries (id, created_at, updated_at, deleted_at, house_id) FROM stdin;
\.


--
-- Data for Name: houses; Type: TABLE DATA; Schema: public; Owner: felixvnolasco
--

COPY public.houses (id, created_at, updated_at, deleted_at, title, price, location, description, status, agent_id, house_details_id, photo_url) FROM stdin;
3	2023-04-23 22:36:53.31758+00	2023-04-25 19:55:22.51732+00	\N	Super 1000	$934,000.00 MXN	Chalco, Estado de México	Te presento la casa más vendida y solicitada en Chalco Estado de México, una hermosa casa que puedes ampliar y decorar a tu gusto, es la casa que más personas han adquirido, y la cual se puede ajustar a tu crédito!. Ven y conoce el mejor fraccionamiento, por ubicación y precio, servicios y cercanía de todo lo que estás buscando, ✨‼️AGENDA TU, CITA HOY MISMO✨‼️	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682450361/RealState/super-1000.webp
4	2023-04-23 22:37:12.397654+00	2023-04-25 19:55:31.767757+00	\N	Super 1000 Plus	$1,212,000.00 MXN	Chalco, Estado de México	Te vas a enamorar de esta hermosa casa, tiene todo lo que necesitas. Es una casa que tiene espacios amplios y cómodos con la mejor distribución de espacios, ven y conoce la maravillosa opción que tengo para ti❤️‼️✨Te invito a que conozcas la casa más vendida y solicitada en una zona segura y cercana a todos los servicios✨‼️	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682451244/RealState/super-1000-plus.webp
12	2023-04-23 22:38:53.207954+00	2023-04-25 19:59:44.802636+00	\N	Royal 8000 Plus	$3,573,000.00 MXN	Chalco, Estado de México	La casa top🔝🔟 en nuestro FRACCIONAMIENTO EXCLUSIVO Y RESIDENCIAL, ven y conoce la mejor casa, atención a gente que quiere vivir como “MILLONARIO”	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/c_scale,w_304/v1682451502/RealState/royal-8000-plus.webp
5	2023-04-23 22:37:32.901679+00	2023-04-25 19:55:45.327795+00	\N	Grand 2000	$1,114,000.00 MXN	Chalco, Estado de México	Casa de 2 recámaras con posibilidad de crecimiento, ven u conoce el modelo más solicitado en Chacó Estado de México. Con una gran ubicación a solo 8 min del centro de chalco. ¡Agenda tu cita! Tengo casas disponibles al mejor 🏆 precio. Tenemos casa con terrenos excedentes hasta de 31 mts2. No sabes lo increíbles que están!!	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682451738/RealState/grand-2000.webp
6	2023-04-23 22:37:53.931171+00	2023-04-25 19:55:54.012691+00	\N	Grand 2000 Plus	$1,506,000.00 MXN	Chalco, Estado de México	Ven a conocer la casa con espacios increíbles en Chalco, es un excelente momento para que adquieras tu casa. “Agenda una cita”	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682451409/RealState/grand-2000-plus.webp
7	2023-04-23 22:38:05.685505+00	2023-04-25 19:56:03.087758+00	\N	Garden 6000	$1,312,000.00 MXN	Chalco, Estado de México	Casa con 2 cajones de estacionamiento, ideal para aquellas personas que tienen mascotas o gustan de un jardín amplio para reuniones familiares.	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682451451/RealState/garden-6000.webp
8	2023-04-23 22:38:13.13546+00	2023-04-25 19:56:12.570129+00	\N	Garden 8000	$1,522,000.00 MXN	Chalco, Estado de México	Linda casa con 3 estacionamientos, es una maravillosa casa para quien busca espacios amplios para que jueguen sus pequeños y mascotas, tiene un amplio jardín y estacionamientos frente de la casa. Ven u descubre todas las opciones increíbles que tengo para ti… no esperes más!!	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682451463/RealState/garden-8000.webp
9	2023-04-23 22:38:27.455932+00	2023-04-25 19:56:20.793685+00	\N	Premier 6000	$1,984,000.00 MXN	Chalco, Estado de México	Te presento una casa de 3 recámaras, que puedes ampliar, y hacer crecer tus espacios, es una excelente opción para aquellas personas que les gusta vivir en comodidad y seguridad con su familia. No esperes más y descubre la manera mas fácil en la que puedes cambiar tu ¡vida y de tu familia!	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682451784/RealState/premier-6000.webp
10	2023-04-23 22:38:33.451521+00	2023-04-25 19:56:33.598967+00	\N	Premier 6000 Plus	$2,456,000.00 MXN	Chalco, Estado de México	Tengo a la venta esta hermosa casa, te enamoraras de ella, es única en el Estado de México que te brindará lo que tu familia y tú buscan. Casa moderna con los mejores acabados y de lujo, ven y conoce las hermosas ubicaciones en Chalco. Carretera Mixquic-Chalco. ‼️✨AGENDA TU CITA HOY MISMO✨‼️	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682451434/RealState/premier-6000-plus.webp
11	2023-04-23 22:38:41.509674+00	2023-04-25 19:56:51.208711+00	\N	Royal 8000	$2,629,000.00 MXN	Chalco, Estado de México	En Residencial Iztac hay un lugar hermoso en el que podrás vivir experiencias en compañía de tu familia y amigos. Tenemos diferentes casas para diferentes presupuestos que se pueden ajustar a tu economía, no hay mejor inversión que un bien inmueble, visita *RESIDENCIAL IZTAC* y conoce las mejores casas del Estado de México, no esperes más y ‼️“AGENDA UNA CITA”‼️	t	3	\N	https://res.cloudinary.com/dhyxqmnua/image/upload/v1682451470/RealState/royal-8000.webp
\.


--
-- Data for Name: photos; Type: TABLE DATA; Schema: public; Owner: felixvnolasco
--

COPY public.photos (id, created_at, updated_at, deleted_at, house_gallery_id, photo_url) FROM stdin;
\.


--
-- Name: agents_id_seq; Type: SEQUENCE SET; Schema: public; Owner: felixvnolasco
--

SELECT pg_catalog.setval('public.agents_id_seq', 3, true);


--
-- Name: house_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: felixvnolasco
--

SELECT pg_catalog.setval('public.house_details_id_seq', 12, true);


--
-- Name: house_galleries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: felixvnolasco
--

SELECT pg_catalog.setval('public.house_galleries_id_seq', 1, true);


--
-- Name: houses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: felixvnolasco
--

SELECT pg_catalog.setval('public.houses_id_seq', 15, true);


--
-- Name: photos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: felixvnolasco
--

SELECT pg_catalog.setval('public.photos_id_seq', 2, true);


--
-- Name: agents agents_pkey; Type: CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.agents
    ADD CONSTRAINT agents_pkey PRIMARY KEY (id);


--
-- Name: house_details house_details_pkey; Type: CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.house_details
    ADD CONSTRAINT house_details_pkey PRIMARY KEY (id);


--
-- Name: house_galleries house_galleries_pkey; Type: CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.house_galleries
    ADD CONSTRAINT house_galleries_pkey PRIMARY KEY (id);


--
-- Name: houses houses_pkey; Type: CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT houses_pkey PRIMARY KEY (id);


--
-- Name: photos photos_pkey; Type: CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (id);


--
-- Name: idx_agents_deleted_at; Type: INDEX; Schema: public; Owner: felixvnolasco
--

CREATE INDEX idx_agents_deleted_at ON public.agents USING btree (deleted_at);


--
-- Name: idx_house_details_deleted_at; Type: INDEX; Schema: public; Owner: felixvnolasco
--

CREATE INDEX idx_house_details_deleted_at ON public.house_details USING btree (deleted_at);


--
-- Name: idx_house_galleries_deleted_at; Type: INDEX; Schema: public; Owner: felixvnolasco
--

CREATE INDEX idx_house_galleries_deleted_at ON public.house_galleries USING btree (deleted_at);


--
-- Name: idx_houses_deleted_at; Type: INDEX; Schema: public; Owner: felixvnolasco
--

CREATE INDEX idx_houses_deleted_at ON public.houses USING btree (deleted_at);


--
-- Name: idx_photos_deleted_at; Type: INDEX; Schema: public; Owner: felixvnolasco
--

CREATE INDEX idx_photos_deleted_at ON public.photos USING btree (deleted_at);


--
-- Name: houses fk_agents_houses; Type: FK CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.houses
    ADD CONSTRAINT fk_agents_houses FOREIGN KEY (agent_id) REFERENCES public.agents(id);


--
-- Name: photos fk_house_galleries_photos; Type: FK CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT fk_house_galleries_photos FOREIGN KEY (house_gallery_id) REFERENCES public.house_galleries(id);


--
-- Name: house_details fk_houses_house_details; Type: FK CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.house_details
    ADD CONSTRAINT fk_houses_house_details FOREIGN KEY (house_id) REFERENCES public.houses(id);


--
-- Name: house_galleries fk_houses_house_gallery; Type: FK CONSTRAINT; Schema: public; Owner: felixvnolasco
--

ALTER TABLE ONLY public.house_galleries
    ADD CONSTRAINT fk_houses_house_gallery FOREIGN KEY (house_id) REFERENCES public.houses(id);


--
-- PostgreSQL database dump complete
--


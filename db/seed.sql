CREATE DATABASE govideo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    date VARCHAR(100) NOT NULL,
    banner VARCHAR(255),
    thumbnail VARCHAR(255),
    home_page VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_events_created ON events(created_at);

INSERT INTO events (name, location, date, banner, thumbnail, home_page, description) VALUES
(
    'GopherCon 2025',
    'San Francisco, USA',
    'July 15-17, 2025',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gophercon.com',
    'The largest Go programming conference in North America, featuring talks from industry experts and core Go team members.'
),
(
    'GoWest Summit',
    'Vancouver, Canada',
    'August 10-12, 2025',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gowest.dev',
    'Western North America''s premier Go conference, focusing on cloud-native development and microservices.'
),
(
    'GoEurope Conference',
    'Berlin, Germany',
    'September 5-7, 2025',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://goeurope.dev',
    'Europe''s largest Go conference, bringing together developers from across the continent to share knowledge and experiences.'
),
(
    'GoCon Japan',
    'Tokyo, Japan', 
    'October 12-14, 2025',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gocon.jp',
    'Japan''s largest Go conference, featuring talks in both Japanese and English about Go development in Asia.'
),
(
    'GoLatin',
    'São Paulo, Brazil',
    'November 8-10, 2025', 
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://golatin.dev',
    'Latin America''s premier Go conference, showcasing the growing Go ecosystem in South America.'
),
(
    'GopherCon UK',
    'London, UK',
    'August 21-23, 2025',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gophercon.co.uk', 
    'The UK''s main Go conference, bringing together developers from across the British Isles and Europe.'
),
(
    'GopherCon Singapore',
    'Singapore',
    'December 5-7, 2025',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gophercon.sg',
    'Southeast Asia''s largest Go conference, focusing on cloud computing and distributed systems.'
),
(
    'GoAustralia',
    'Sydney, Australia',
    'March 15-17, 2026',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://goaustralia.dev',
    'Australia''s primary Go conference, featuring talks about Go in enterprise and startup environments.'
),
(
    'GopherCon India',
    'Bangalore, India',
    'February 22-24, 2026',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gopherconindia.com',
    'India''s biggest Go conference, highlighting the growing adoption of Go in the Indian tech industry.'
),
(
    'GoNordic',
    'Stockholm, Sweden',
    'May 18-20, 2026',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gonordic.dev',
    'The Nordic region''s Go conference, bringing together developers from Scandinavian countries.'
),
(
    'GoCon Russia',
    'Moscow, Russia',
    'April 8-10, 2026',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gocon.ru',
    'Russia''s main Go conference, featuring talks about systems programming and performance optimization.'
),
(
    'GoMENA',
    'Dubai, UAE',
    'January 15-17, 2026',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gomena.dev',
    'The Middle East and North Africa''s Go conference, showcasing Go adoption in the region''s tech sector.'
),
(
    'GoAfrica',
    'Cape Town, South Africa',
    'June 20-22, 2026',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://goafrica.dev',
    'Africa''s premier Go conference, highlighting the growing Go community across the continent.'
);

CREATE TABLE speakers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    home_page VARCHAR(255),
    github VARCHAR(255),
    twitter VARCHAR(255),
    linkedin VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_speakers_created ON speakers(created_at);

INSERT INTO speakers (name, avatar, home_page, github, twitter, linkedin) VALUES
(
    'Sarah Chen',
    'https://placehold.co/200x200',
    'https://sarahchen.dev',
    'https://github.com/sarahchen',
    'https://twitter.com/sarahchen',
    'https://linkedin.com/in/sarahchen'
),
(
    'James Wilson',
    'https://placehold.co/200x200',
    'https://jameswilson.tech',
    'https://github.com/jwilson',
    'https://twitter.com/jwilson',
    'https://linkedin.com/in/jameswilson'
),
(
    'Maria Rodriguez',
    'https://placehold.co/200x200',
    'https://maria-rodriguez.com',
    'https://github.com/mrodriguez',
    'https://twitter.com/mrodriguez',
    'https://linkedin.com/in/mariarodriguez'
),
(
    'Alex Kumar',
    'https://placehold.co/200x200',
    'https://alexkumar.dev',
    'https://github.com/alexkumar',
    'https://twitter.com/alexkumar',
    'https://linkedin.com/in/alexkumar'
),
(
    'Emily Zhang',
    'https://placehold.co/200x200',
    'https://emilyzhang.io',
    'https://github.com/ezhang',
    'https://twitter.com/ezhang',
    'https://linkedin.com/in/emilyzhang'
),
(
    'Michael Patel',
    'https://placehold.co/200x200',
    'https://mpatel.dev',
    'https://github.com/mpatel',
    'https://twitter.com/mpatel',
    'https://linkedin.com/in/michaelpatel'
),
(
    'Lisa Anderson',
    'https://placehold.co/200x200',
    'https://lisaanderson.tech',
    'https://github.com/landerson',
    'https://twitter.com/landerson',
    'https://linkedin.com/in/lisaanderson'
),
(
    'David Kim',
    'https://placehold.co/200x200',
    'https://davidkim.dev',
    'https://github.com/dkim',
    'https://twitter.com/dkim',
    'https://linkedin.com/in/davidkim'
),
(
    'Sophie Martin',
    'https://placehold.co/200x200',
    'https://sophiemartin.dev',
    'https://github.com/smartin',
    'https://twitter.com/smartin',
    'https://linkedin.com/in/sophiemartin'
),
(
    'Carlos Santos',
    'https://placehold.co/200x200',
    'https://carlossantos.dev',
    'https://github.com/csantos',
    'https://twitter.com/csantos',
    'https://linkedin.com/in/carlossantos'
),
(
    'Anna Kowalski',
    'https://placehold.co/200x200',
    'https://annakowalski.dev',
    'https://github.com/akowalski',
    'https://twitter.com/akowalski',
    'https://linkedin.com/in/annakowalski'
),
(
    'Thomas Schmidt',
    'https://placehold.co/200x200',
    'https://tschmidt.io',
    'https://github.com/tschmidt',
    'https://twitter.com/tschmidt',
    'https://linkedin.com/in/thomasschmidt'
),
(
    'Nina Patel',
    'https://placehold.co/200x200',
    'https://ninapatel.dev',
    'https://github.com/npatel',
    'https://twitter.com/npatel',
    'https://linkedin.com/in/ninapatel'
),
(
    'Robert Chang',
    'https://placehold.co/200x200',
    'https://robertchang.tech',
    'https://github.com/rchang',
    'https://twitter.com/rchang',
    'https://linkedin.com/in/robertchang'
),
(
    'Emma Brown',
    'https://placehold.co/200x200',
    'https://emmabrown.dev',
    'https://github.com/ebrown',
    'https://twitter.com/ebrown',
    'https://linkedin.com/in/emmabrown'
),
(
    'Lucas Silva',
    'https://placehold.co/200x200',
    'https://lucassilva.dev',
    'https://github.com/lsilva',
    'https://twitter.com/lsilva',
    'https://linkedin.com/in/lucassilva'
),
(
    'Julia Lee',
    'https://placehold.co/200x200',
    'https://julialee.tech',
    'https://github.com/jlee',
    'https://twitter.com/jlee',
    'https://linkedin.com/in/julialee'
),
(
    'Marco Rossi',
    'https://placehold.co/200x200',
    'https://marcorossi.dev',
    'https://github.com/mrossi',
    'https://twitter.com/mrossi',
    'https://linkedin.com/in/marcorossi'
),
(
    'Hannah Kim',
    'https://placehold.co/200x200',
    'https://hannahkim.dev',
    'https://github.com/hkim',
    'https://twitter.com/hkim',
    'https://linkedin.com/in/hannahkim'
),
(
    'Daniel Garcia',
    'https://placehold.co/200x200',
    'https://danielgarcia.tech',
    'https://github.com/dgarcia',
    'https://twitter.com/dgarcia',
    'https://linkedin.com/in/danielgarcia'
);

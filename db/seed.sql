CREATE DATABASE govideo CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE govideo;
CREATE TABLE events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    date_start DATETIME NOT NULL,
    date_end DATETIME NOT NULL,
    banner VARCHAR(255),
    thumbnail VARCHAR(255),
    home_page VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_events_created ON events(created_at);

INSERT INTO events (name, location, date_start, date_end, banner, thumbnail, home_page, description) VALUES
(
    'GopherCon 2025',
    'San Francisco, USA',
    '2025-07-15',
    '2025-07-17',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gophercon.com',
    'The largest Go programming conference in North America, featuring talks from industry experts and core Go team members.'
),
(
    'GoWest Summit',
    'Vancouver, Canada',
    '2025-08-10',
    '2025-08-12',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gowest.dev',
    'Western North America''s premier Go conference, focusing on cloud-native development and microservices.'
),
(
    'GoEurope Conference',
    'Berlin, Germany',
    '2025-09-05',
    '2025-09-07',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://goeurope.dev',
    'Europe''s largest Go conference, bringing together developers from across the continent to share knowledge and experiences.'
),
(
    'GoCon Japan',
    'Tokyo, Japan', 
    '2025-10-12',
    '2025-10-14',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gocon.jp',
    'Japan''s largest Go conference, featuring talks in both Japanese and English about Go development in Asia.'
),
(
    'GoLatin',
    'São Paulo, Brazil',
    '2025-11-08',
    '2025-11-10',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://golatin.dev',
    'Latin America''s premier Go conference, showcasing the growing Go ecosystem in South America.'
),
(
    'GopherCon UK',
    'London, UK',
    '2025-08-21',
    '2025-08-23',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gophercon.co.uk', 
    'The UK''s main Go conference, bringing together developers from across the British Isles and Europe.'
),
(
    'GopherCon Singapore',
    'Singapore',
    '2025-12-05',
    '2025-12-07',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gophercon.sg',
    'Southeast Asia''s largest Go conference, focusing on cloud computing and distributed systems.'
),
(
    'GoAustralia',
    'Sydney, Australia',
    '2026-03-15',
    '2026-03-17',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://goaustralia.dev',
    'Australia''s primary Go conference, featuring talks about Go in enterprise and startup environments.'
),
(
    'GopherCon India',
    'Bangalore, India',
    '2026-02-22',
    '2026-02-24',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gopherconindia.com',
    'India''s biggest Go conference, highlighting the growing adoption of Go in the Indian tech industry.'
),
(
    'GoNordic',
    'Stockholm, Sweden',
    '2026-05-18',
    '2026-05-20',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gonordic.dev',
    'The Nordic region''s Go conference, bringing together developers from Scandinavian countries.'
),
(
    'GoCon Russia',
    'Moscow, Russia',
    '2026-04-08',
    '2026-04-10',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gocon.ru',
    'Russia''s main Go conference, featuring talks about systems programming and performance optimization.'
),
(
    'GoMENA',
    'Dubai, UAE',
    '2026-01-15',
    '2026-01-17',
    'https://placehold.co/800x400',
    'https://placehold.co/400x300',
    'https://gomena.dev',
    'The Middle East and North Africa''s Go conference, showcasing Go adoption in the region''s tech sector.'
),
(
    'GoAfrica',
    'Cape Town, South Africa',
    '2026-06-20',
    '2026-06-22',
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

CREATE TABLE talks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    duration VARCHAR(50) NOT NULL,
    speaker_id INT NOT NULL,
    event_id INT NOT NULL,
    thumbnail VARCHAR(255),
    video_provider VARCHAR(50),
    video_id VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (speaker_id) REFERENCES speakers(id),
    FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE INDEX idx_talks_created ON talks(created_at);
CREATE INDEX idx_talks_speaker ON talks(speaker_id);
CREATE INDEX idx_talks_event ON talks(event_id);

INSERT INTO talks (title, duration, speaker_id, event_id, thumbnail, video_provider, video_id) VALUES
('Building Scalable Microservices with Go', '45 minutes', 1, 1, 'https://placehold.co/400x300', 'youtube', 'JKuha6vsYAE'),
('Advanced Concurrency Patterns', '30 minutes', 2, 1, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go for System Programming', '45 minutes', 3, 2, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Testing Strategies in Go', '30 minutes', 4, 2, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Building High-Performance APIs', '45 minutes', 5, 3, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go in the Cloud', '30 minutes', 6, 3, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Memory Management in Go', '45 minutes', 7, 4, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Effective Error Handling', '30 minutes', 8, 4, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go for AI and Machine Learning', '45 minutes', 9, 5, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('WebAssembly and Go', '30 minutes', 10, 5, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Optimizing Go Applications', '45 minutes', 11, 6, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go Modules Deep Dive', '30 minutes', 12, 6, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Building CLI Tools in Go', '45 minutes', 13, 7, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go for Game Development', '30 minutes', 14, 7, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Secure Coding in Go', '45 minutes', 15, 8, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Database Design with Go', '30 minutes', 16, 8, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go for IoT Applications', '45 minutes', 17, 9, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('RESTful API Best Practices', '30 minutes', 18, 9, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('gRPC in Production', '45 minutes', 19, 10, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go Generics in Practice', '30 minutes', 20, 10, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Debugging Go Applications', '45 minutes', 1, 11, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go for Blockchain', '30 minutes', 2, 11, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Performance Profiling in Go', '45 minutes', 3, 12, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Microservices Architecture', '30 minutes', 4, 12, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go and Kubernetes', '45 minutes', 5, 13, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Testing Microservices', '30 minutes', 6, 13, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go for Data Processing', '45 minutes', 7, 1, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Building Real-time Systems', '30 minutes', 8, 2, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Go and GraphQL', '45 minutes', 9, 3, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA'),
('Dependency Injection in Go', '30 minutes', 10, 4, 'https://placehold.co/400x300', 'youtube', 'YEKjSzIwAdA');

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

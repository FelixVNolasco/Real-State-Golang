
# ReadMe


## Docker Command
docker run --name nuestroHogar -p 5432:5432 -e POSTGRES_USER=felixvnolasco -e POSTGRES_PASSWORD=holamundo -d postgres:12-alpine


# Generate PostgreSQL database dump
pg_dump -U felixvnolasco state > dump.sql
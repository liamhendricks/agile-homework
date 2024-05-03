You will need `docker`, `docker compose` and `npm` to run everything. If you are running on OSX, you may need to remove the
`network_mode: "host"` line in the docker-compose.yml (if you want to use localhost). You can configure some host / port values in
`.env` file.

Run `make setup` to download dependencies and create docker images.
Run `make local` to start the backend.
If you want to use the frontend:
```
cd sw-frontend
npm install
npm run start
```
Otherwise refer to the postman collection and use the backend endpoints to look at the data.

Cheers ^^

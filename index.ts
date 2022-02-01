import { config } from "https://deno.land/x/dotenv/mod.ts";
import {Application} from "https://deno.land/x/oak/mod.ts";
import {oakCors} from "https://deno.land/x/cors/mod.ts";

//use config for access to backend secrets
config();

const app = new Application();

app.use(oakCors({
    origin: "http://localhost:3000"
}));

//connect to mongoDB

await app.listen({port: 8000});
console.log("server running on port 8000")
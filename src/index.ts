import { Elysia } from "elysia";
import { cors } from "@elysiajs/cors";
import { msgpack } from "elysia-msgpack";

import { auth } from "@/lib/auth";

const app = new Elysia()
    .use(
        cors({
            origin: process.env.BETTER_AUTH_URL as string,
            methods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
            credentials: true,
            allowedHeaders: ["Content-Type", "Authorization"],
        }),
    )
    .use(msgpack())
    .mount(auth.handler)
    .get("/", () => "Hello Elysia")
    .listen(process.env.LISTENING_PORT as string);

console.log(
    `Aurchat backend is running at ${app.server?.hostname}:${app.server?.port}`,
);

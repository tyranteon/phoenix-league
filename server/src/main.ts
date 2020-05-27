import {Client} from "pg";
import {WebAPI} from "./webapi";
import {Constants} from "./constants";

class Program {
    static async run() {
        const client = new Client({
            user: Constants.PGUSER,
            host: "localhost",
            database: "postgres",
            password: Constants.PGPASSWORD,
            port: 5432,
        });

        await client.connect();

        let webAPI = new WebAPI(client);

        console.log("Initialized");
    }
}

Program.run();
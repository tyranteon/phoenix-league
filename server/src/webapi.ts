import express = require("express");
import passport = require("passport");

import {Strategy as SteamStrategy} from "passport-steam";
import session = require('express-session');
import cors = require("cors");

import {Client} from "pg";
import {NextFunction, Request, Response} from "express";
import {Constants} from "./constants";
import * as http from "http";
import WebSocket = require("ws");

interface BaseUser {
    profileURL: string,
    avatar: {
        small: string,
        medium: string,
        full: string,
    },
    displayName: string,
}

interface User extends BaseUser {
    steamID: string,
    mmr: number,
    isVetted: boolean,
}

function serializeUser(u: BaseUser): string {
    return JSON.stringify({
        profileURL: u.profileURL,
        avatar: u.avatar,
        displayName: u.displayName,
    })
}

function profileToUser(profile: any): BaseUser {
    return {
        profileURL: profile._json.profileurl,
        displayName: profile.displayName,
        avatar: {
            small: profile._json.avatar,
            medium: profile._json.avatarmedium,
            full: profile._json.avatarfull,
        },
    }
}

export class WebAPI {
    websockets = new Map<string, WebSocket>();

    constructor(private pgClient: Client) {
        this.init();
    }

    async getUserFromDB(steamID: string, profile: any): Promise<User> {
        let d = await this.pgClient.query("SELECT * FROM users WHERE user_id = $1", [steamID]);
        if(d.rows.length == 0) {
            // TODO: Have the profile only be parsed in here!
            const newUser = profileToUser(profile);
            await this.pgClient.query("INSERT INTO users (user_id, profile_json, mmr, is_vetted, override_vetting) VALUES ($1, $2, -1, FALSE, FALSE)", [
                steamID, serializeUser(newUser)
            ]);
            return {
                ... newUser,
                steamID,
                mmr: -1,
                isVetted: false,
            };
        }

        const result = d.rows[0];

        let baseUser: BaseUser = JSON.parse(result.profile_json);
        return {
            ... baseUser,
            steamID: steamID,
            mmr: result.mmr,
            isVetted: result.override_vetting || result.is_vetted
        }
    }

    async verifyUser(identifier: any, profile: any, done: any) {
        console.log("Verifying user...");
        const identifierRegex = /^https?:\/\/steamcommunity\.com\/openid\/id\/(\d+)$/;
        let res = identifierRegex.exec(identifier);
        if(res === null) {
            console.error("SteamID not found or invalid");
            return done(null, false);
        }

        let steamID = res[1];

        let user = await this.getUserFromDB(steamID, profile);

        return done(null, user);
    }

    init() {
        passport.serializeUser((user, done) => {
            done(null, user);
        });

        passport.deserializeUser((obj, done) => {
            done(null, obj);
        });

        passport.use(new SteamStrategy({
            returnURL: 'http://localhost:3000/auth/steam/return',
            realm: 'http://localhost:3000/',
            apiKey: Constants.APIKEY,
        }, this.verifyUser.bind(this)));

        let app = express();

        app.use(cors({
            origin: ["https://www.phoenix-league.net", "http://localhost:8080"],
            credentials: true,
        }));

        const sessionParser = session({
            secret: 'secret smaug',
            resave: true,
            saveUninitialized: true
        });

        app.use(sessionParser);

        app.use(passport.initialize());
        app.use(passport.session());

        app.get('/auth/steam/return',
            passport.authenticate('steam'),
            (req, res) => {
                res.write("<h1>Success</h1>");
                res.end();
            });

        app.get("/user", ensureAuthenticated, (req, res) => {
            res.set("Content-Type", "application/json");
            res.write(JSON.stringify(req.user));
            res.end();
        });

        let server = http.createServer(app);
        const wss = new WebSocket.Server({
            server,
            verifyClient(info, done) {
                console.log("Parsing session from request...");
                sessionParser(info.req as Request, {} as Response, () => {
                    console.log("Session is parsed");
                    done(true);
                });

            }
        });

        wss.on("connection", (ws, req) => {
            let session = (<Request>req).session;

            if(session === undefined || session.passport === undefined) {
                return;
            }

            ws.onclose = () => {
                this.websockets.delete(session!.passport.user.steamID)
            };

            this.websockets.set(session.passport.user.steamID, ws);
            // console.log("Connected!", (<Request>req).session);
            // ws.send("Well, hello there!");
        });

        server.listen(3000);
    }
}

function ensureAuthenticated(req: Request, res: Response, next: NextFunction) {
    if (req.isAuthenticated()) {
        return next();
    }
    res.sendStatus(403);
    res.end();
}
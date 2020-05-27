interface Bucket {
    min: number;
    max: number;
    players: Player[]
}

interface Player {
    mmr: number;
}

class Matchmaker {
    buckets: Bucket[] = [];
    foundMatches: Player[][] = [];

    constructor() {
        this.initBuckets();
    }

    initBuckets() {
        for (let i = 0; i < 10000; i += 1000) {
            this.buckets.push({
                min: i,
                max: i + 1000,
                players: []
            })
        }
    }

    findBucket(mmr: number): Bucket {
        for(let bucket of this.buckets) {
            if(mmr >= bucket.min && mmr < bucket.max) {
                return bucket;
            }
        }
        // TODO: Fix this for players above max MMR
        return this.buckets[this.buckets.length-1];
    }

    addPlayer(player: Player) {
        let bucket = this.findBucket(player.mmr);
        bucket.players.push(player);
        if(bucket.players.length == 10) {
            this.foundMatches.push(bucket.players);
            bucket.players = [];
        }
    }
}

let mm = new Matchmaker();

let lastFoundLen = 0;
let lastFoundIndex = 0;

function logMatch(players: Player[]) {
    console.log("Match: \n  " + players.map(p => p.mmr).join("\n  ") + "\n");
}

for(let i = 0; i < 100; i++) {
    mm.addPlayer({
        mmr: Math.floor(Math.random() * 10000)
    });
    if(mm.foundMatches.length > lastFoundLen) {
        lastFoundLen = mm.foundMatches.length;

        logMatch(mm.foundMatches[lastFoundLen-1]);

        lastFoundIndex = i;
    }
}
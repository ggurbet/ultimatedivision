// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Exposes Division Club type. */
export class DivisionClub {
    /** Receives icon, position, name, games, wins, draws,
     * defeats, goalDifference, points as string parameters. */
    constructor(
        public position: string,
        public club: {
            name: string;
            icon: string;
        },
        public games: string,
        public wins: string,
        public draws: string,
        public defeats: string,
        public goalDifference: string,
        public points: string,
    ) { };
};

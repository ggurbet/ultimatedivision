// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Goal is entity that describes scored goal. */
export class Goal {
    /** Goal contains of player card and minute when was scored goal. */
    constructor(
        public card: string,
        public minute: number,
    ) { };
};

/** Team describes football team entity. */
export class Team {
    /** Team contains of summary goals number, goals array and userId. */
    constructor(
        public quantityGoals: number,
        public goals: Goal[] | null,
        public userId: string,
    ) { };
};

/** Match exposes match domain entity. */
export class Match {
    /** Contains of firstTeamGoalsCrored and secondTeamGoalsScored. */
    constructor(
        public firstTeam: Team,
        public secondTeam: Team,
    ) { };
};

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

const DEFAULT_VALUE = 0;

/** backend club */
export class Club {
    public id: string = '';
    public name: string = '';
    public createdAt: string = '';
}

/** backend squad */
export class Squad {
    public id: string = '';
    public clubId: string = '';
    public formation: number = DEFAULT_VALUE;
    public tactic: number = DEFAULT_VALUE;
    public captainId: string = '';
};


/** backend card  */
export class SquadCard {
    public squadId: string = '';
    public cardId: string = '';
    public position: number = DEFAULT_VALUE;
}

/** foe drag and drop implementation */
export class Options {
    /** options implementation */
    constructor(
        public chosedCard: number = DEFAULT_VALUE,
        public showCardSeletion: boolean = false,
        public dragStart: number | null = null,
        public dragTarget: number | null = null,
    ) { }
}
/** club reducer state  */
export class ClubState {
    public clubs: Club = new Club();
    public squad: Squad = new Squad();
    public squadCards: SquadCard[] = [];
    public options: Options = new Options();
}


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
}

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
        public dragTarget: number | null = null
    ) {}
}

/** club reducer state  */
export class ClubState {
    public clubs: Club = new Club();
    public squad: Squad = new Squad();
    public squadCards: SquadCard[] = [];
    public options: Options = new Options();
}

export type FormationsType =
  | '4-4-2'
  | '4-2-4'
  | '4-2-2-2'
  | '4-3-1-2'
  | '4-3-3'
  | '4-2-3-1'
  | '4-3-2-1'
  | '4-1-3-2'
  | '5-3-2'
  | '4-5-2';
export type TacticsType = 'attack' | 'defence' | 'balanced';

/* eslint-disable no-magic-numbers */
export enum Formations {
    '4-4-2' = 1,
    '4-2-4' = 2,
    '4-2-2-2' = 3,
    '4-3-1-2' = 4,
    '4-3-3' = 5,
    '4-2-3-1' = 6,
    '4-3-2-1' = 7,
    '4-1-3-2' = 8,
    '5-3-2' = 9,
    '4-5-2' = 10,
}

export enum Tactic {
    'attack' = 1,
    'defence' = 2,
    'balanced' = 3,
}

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.


/** class for our getters to get label and value while mapping */
export class CardField {
    /** label and value for mapping */
    constructor(
        public label: string,
        public value: string | number,
    ) { }
}

/* eslint-disable */
/** player stats implementation */
export class CardStats {
    /** main stat with substats */
    constructor(
        public title: string = '',
        public abbreviated: string = '',
        public fields: CardField[] = []
    ) { }
    public average: number = this.fields
        .map(item => +item.value)
        .reduce((prev, current) => prev + current) / this.fields.length;
    /** abbreviated title of card stat name */
    get abbr(): string {
        return this.title.slice(0, 3);
    }
    /** stat giagram color depend on avarage stat value */
    get color(): string {
        const STATISTIC_UPPER_BOUND = 90;
        const STATISTIC_LOWER_BOUND = 50;

        const STATISTIC_UPPER_BOUND_COLOR = '#3CCF5D';
        const STATISTIC_MEDIUM_BOUND_COLOR = '#E8EC16';
        const STATISTIC_LOWER_BOUND_COLOR = '#FF4200';

        switch (true) {
            case this.average >= STATISTIC_UPPER_BOUND:
                return STATISTIC_UPPER_BOUND_COLOR;
            case this.average >= STATISTIC_LOWER_BOUND:
                return STATISTIC_MEDIUM_BOUND_COLOR;
            default:
                return STATISTIC_LOWER_BOUND_COLOR;
        }
    }
}

/** Card base implementation */
export class Card {
    /** Card fields */
    constructor(
        public id: string,
        public playerName: number,
        public quality: number,
        public pictureType: number,
        public height: number,
        public weight: number,
        public skinColor: number,
        public hairStyle: number,
        public hairColor: number,
        public accessories: string,
        public dominantFoot: number,
        public isTatoos: boolean,
        public userId: string,
        public tactics: number,
        public positioning: number,
        public composure: number,
        public aggression: number,
        public vision: number,
        public awareness: number,
        public crosses: number,
        public physique: number,
        public acceleration: number,
        public runningSpeed: number,
        public reactionSpeed: number,
        public agility: number,
        public stamina: number,
        public strength: number,
        public jumping: number,
        public balance: number,
        public technique: number,
        public dribbling: number,
        public ballControl: number,
        public weakFoot: number,
        public skillMoves: number,
        public finesse: number,
        public curve: number,
        public volleys: number,
        public shortPassing: number,
        public longPassing: number,
        public forwardPass: number,
        public offense: number,
        public finishingAbility: number,
        public shotPower: number,
        public accuracy: number,
        public distance: number,
        public penalty: number,
        public freeKicks: number,
        public corners: number,
        public headingAccuracy: number,
        public defence: number,
        public offsideTrap: number,
        public sliding: number,
        public tackles: number,
        public ballFocus: number,
        public interceptions: number,
        public vigilance: number,
        public goalkeeping: number,
        public reflexes: number,
        public diving: number,
        public handling: number,
        public sweeping: number,
        public throwing: number,
    ) { }
    /** Using in footballerCard in info block */
    get mainInfo() {
        return [
            new CardField('name', this.playerName),
            // To do: at this momenty nation does not exist
            new CardField('nation', 'this.nation'),
            new CardField('skills', '5'),
            new CardField('weak foot', this.weakFoot),
            new CardField('intl. rep', '5'),
            new CardField('foot', this.dominantFoot),
            new CardField('height', this.height),
            new CardField('nation', this.weight),
            // To do: at this momenty revision does not exist or it is designer mistake or it is quality
            new CardField('revision', 'rare'),
            // To do: create method to convert attack and defence values into this scale
            new CardField('def. wr', 'low'),
            new CardField('arr. wr', 'high'),
            // next fields does not exist in card at this moment
            new CardField('added on', '2020-09-10'),
            new CardField('origin', 'na'),
            new CardField('r. Face', 'low'),
            new CardField('b. type', ''),
            new CardField('age', '36 years old'),
        ];
    }

    /** Using in diagramm area in footballerCard */
    get diagramArea() {
        // To do: need to get real min and max values to convert into diagram value
        // To do: this fields does not exist
        return [
            new CardField('physical', 688),
            new CardField('mental', 688),
            new CardField('skill', 688),
            new CardField('cham. style', 688),
            new CardField('base stats', 688),
            new CardField('in game stats', 688),
        ];
    }

    /** returns fields for card stats area in footballerCard */
    get statsArea() {
        return [
            new CardStats('tactics', 'tac', [
                new CardField('positioning', this.positioning),
                new CardField('composure', this.composure,),
                new CardField('aggression', this.aggression),
                new CardField('vision', this.vision),
                new CardField('awareness', this.awareness),
                new CardField('crosses', this.crosses),
            ]),
            new CardStats('physique', 'phy', [
                new CardField('acceleration', this.acceleration),
                new CardField('running speed', this.runningSpeed),
                new CardField('reaction speed', this.reactionSpeed),
                new CardField('agility', this.agility),
                new CardField('stamina', this.stamina),
                new CardField('strength', this.strength),
                new CardField('jumping', this.jumping),
                new CardField('balance', this.jumping),
            ]),
            new CardStats('technique', 'tec', [
                new CardField('dribbing', this.dribbling),
                new CardField('ball fontrol', this.ballControl),
                new CardField('weak foot', this.weakFoot),
                new CardField('skill moves', this.skillMoves),
                new CardField('finesse', this.finesse),
                new CardField('curve', this.curve),
                new CardField('volleys', this.volleys),
                new CardField('short passing', this.shortPassing),
                new CardField('long passing', this.longPassing),
                new CardField('forward pass', this.forwardPass),
            ]),
            new CardStats('offence', 'off', [
                new CardField('finishing ability', this.finishingAbility),
                new CardField('shot power', this.shotPower),
                new CardField('accuracy', this.accuracy),
                new CardField('distance', this.distance),
                new CardField('penalty', this.penalty),
                new CardField('free Kicks', this.freeKicks),
                new CardField('corners', this.corners),
                new CardField('heading accuracy', this.headingAccuracy),
            ]),
            new CardStats('defence', 'def', [
                new CardField('offside trap', this.offsideTrap),
                new CardField('tackles', this.tackles),
                new CardField('ball focus', this.ballFocus),
                new CardField('interceptions', this.interceptions),
                new CardField('vigilance', this.vigilance),
            ]),
            new CardStats('goalkeeping', 'gk', [
                new CardField('diving', this.diving),
                new CardField('handling', this.handling),
                new CardField('sweeping', this.sweeping),
                new CardField('throwing', this.throwing),
            ]),
        ];
    }
}

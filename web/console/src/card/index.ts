// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import diamond from '@static/img/MarketPlacePage/marketPlaceCardsGroup/diamond2.svg';
import gold from '@static/img/MarketPlacePage/marketPlaceCardsGroup/gold2.svg';
import silver from '@static/img/MarketPlacePage/marketPlaceCardsGroup/silver2.svg';
import wood from '@static/img/MarketPlacePage/marketPlaceCardsGroup/wood2.svg';

import diamondShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/diamondShadow.svg';
import goldShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/goldShadow.svg';
import silverShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/silverShadow.svg';
import woodShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/woodShadow.svg';

import currentBid
    from '@static/img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/bid.svg';
import minimumPrice
    from '@static/img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/minimum.svg';
import purchased
    from '@static/img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/purchased.svg';

import confirmIcon from '@static/img/MarketPlacePage/MyCard/ok.svg';
import priceGoldIcon from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';
import playerFace from '@static/img/MarketPlacePage/marketPlaceCardsGroup/player.svg';
import priceIcon
    from '@static/img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/price.svg';

import checked from '@static/img/FootballerCardPage/checked.svg';
import star from '@static/img/FootballerCardPage/star.svg';


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
    get average() {
        return Math.round(this.fields
            .map(item => +item.value)
            .reduce((prev, current) => prev + current) / this.fields.length);
    }
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

export interface ICard {
    id: string,
    playerName: string,
    quality: string,
    pictureType: number,
    height: number,
    weight: number,
    skinColor: number,
    hairStyle: number,
    hairColor: number,
    accessories: number[],
    dominantFoot: string,
    isTattoos: boolean,
    status: number,
    type: string,
    userId: string,
    tactics: number,
    positioning: number,
    composure: number,
    aggression: number,
    vision: number,
    awareness: number,
    crosses: number,
    physique: number,
    acceleration: number,
    runningSpeed: number,
    reactionSpeed: number,
    agility: number,
    stamina: number,
    strength: number,
    jumping: number,
    balance: number,
    technique: number,
    dribbling: number,
    ballControl: number,
    weakFoot: number,
    skillMoves: number,
    finesse: number,
    curve: number,
    volleys: number,
    shortPassing: number,
    longPassing: number,
    forwardPass: number,
    offense: number,
    finishingAbility: number,
    shotPower: number,
    accuracy: number,
    distance: number,
    penalty: number,
    freeKicks: number,
    corners: number,
    headingAccuracy: number,
    defence: number,
    offsideTrap: number,
    sliding: number,
    tackles: number,
    ballFocus: number,
    interceptions: number,
    vigilance: number,
    goalkeeping: number,
    reflexes: number,
    diving: number,
    handling: number,
    sweeping: number,
    throwing: number
}

/** Card base implementation */
export class Card {
    id: string = '';
    playerName: string = '';
    quality: string = '';
    pictureType: number = 0;
    height: number = 0;
    weight: number = 0;
    skinColor: number = 0;
    hairStyle: number = 0;
    hairColor: number = 0;
    accessories: number[] = [];
    dominantFoot: string = '';
    isTattoos: boolean = false;
    status: number = 0;
    type: string = '';
    userId: string = '';
    tactics: number = 0;
    positioning: number = 0;
    composure: number = 0;
    aggression: number = 0;
    vision: number = 0;
    awareness: number = 0;
    crosses: number = 0;
    physique: number = 0;
    acceleration: number = 0;
    runningSpeed: number = 0;
    reactionSpeed: number = 0;
    agility: number = 0;
    stamina: number = 0;
    strength: number = 0;
    jumping: number = 0;
    balance: number = 0;
    technique: number = 0;
    dribbling: number = 0;
    ballControl: number = 0;
    weakFoot: number = 0;
    skillMoves: number = 0;
    finesse: number = 0;
    curve: number = 0;
    volleys: number = 0;
    shortPassing: number = 0;
    longPassing: number = 0;
    forwardPass: number = 0;
    offense: number = 0;
    finishingAbility: number = 0;
    shotPower: number = 0;
    accuracy: number = 0;
    distance: number = 0;
    penalty: number = 0;
    freeKicks: number = 0;
    corners: number = 0;
    headingAccuracy: number = 0;
    defence: number = 0;
    offsideTrap: number = 0;
    sliding: number = 0;
    tackles: number = 0;
    ballFocus: number = 0;
    interceptions: number = 0;
    vigilance: number = 0;
    goalkeeping: number = 0;
    reflexes: number = 0;
    diving: number = 0;
    handling: number = 0;
    sweeping: number = 0;
    throwing: number = 0;
    /** Card fields */
    constructor(
        card: Partial<ICard> = {}
    ) {
        Object.assign(this, card);
    }

    /** returns background type and shadow type according to quality */
    get style() {

        switch (this.quality) {
            case 'wood':
                return {
                    background: wood,
                    shadow: woodShadow,
                };
            case 'silver':
                return {
                    background: silver,
                    shadow: silverShadow,
                };
            case 'gold':
                return {
                    background: gold,
                    shadow: goldShadow,
                };
            case 'diamond':
                return {
                    background: diamond,
                    shadow: diamondShadow,
                };
        };
    }
    /** will be replaced by backend face implementation */
    get face() {
        return playerFace
    }

    /**TODO: for testing, will be replaced */
    get cardPrice() {
        const prp = 75;
        const pr = 'PR: 1,142,000 - 15,000,000';
        const updated = 16;
        const price = '10,868,000';
        /** get stat giagram color depend on price value  */
        const PRICE_UPPER_BOUND = 80;
        const PRICE_MEDIUM_BOUND = 70;
        const PRICE_LOWER_BOUND = 50;

        const PRICE_UPPER_BOUND_COLOR = '#1898D7';
        const PRICE_MEDIUM_BOUND_COLOR = '#3CCF5D';
        const PRICE_LOWER_BOUND_COLOR = '#E86C27';
        const PRICE_DEFAULT_BOUND_COLOR = '#FF4200';
        let color: string;

        switch (true) {
            case prp >= PRICE_UPPER_BOUND:
                color = PRICE_UPPER_BOUND_COLOR;
                break;
            case prp >= PRICE_MEDIUM_BOUND:
                color = PRICE_MEDIUM_BOUND_COLOR;
                break;
            case prp >= PRICE_LOWER_BOUND:
                color = PRICE_LOWER_BOUND_COLOR;
                break;
            default:
                color = PRICE_DEFAULT_BOUND_COLOR;
        }

        return {
            prp,
            color,
            pr,
            updated,
            price
        }
    }

    /** Using in footballerCard in info block */
    get infoBlock() {
        return [
            new CardField('name', this.playerName),
            // TODO: at this momenty nation does not exist
            new CardField('nation', 'this.nation'),
            new CardField('skills', '5'),
            new CardField('weak foot', this.weakFoot),
            new CardField('intl. rep', '5'),
            new CardField('foot', this.dominantFoot),
            new CardField('height', this.height),
            new CardField('nation', this.weight),
            // TODO: at this momenty revision does not exist or it is designer mistake or it is quality
            new CardField('revision', 'rare'),
            // TODO: create method to convert attack and defence values into this scale
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
        // TODO: need to get real min and max values to convert into diagram value
        // TODO: this fields does not exist
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
};
/** Cards domain entity */
export class CardsPage {
    /** default Cards initial values */
    constructor(
        public cards: Card[],
        public page: {
            offset: number,
            limit: number,
            currentPage: number,
            pageCount: number,
            totalCount: number
        },
    ) { };
};

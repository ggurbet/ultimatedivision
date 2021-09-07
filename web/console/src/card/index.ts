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

/** Card base implementation */
export class Card {
    /** Card fields */
    constructor(
        public id: string,
        public playerName: string,
        public quality: string,
        public pictureType: number,
        public height: number,
        public weight: number,
        public skinColor: number,
        public hairStyle: number,
        public hairColor: number,
        public accessories: number[],
        public dominantFoot: string,
        public isTatoos: boolean,
        public status: number,
        public type: string,
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

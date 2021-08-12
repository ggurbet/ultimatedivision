// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/* eslint-disable */
import diamond from '@static/img/MarketPlacePage/marketPlaceCardsGroup/diamond2.svg';
import gold from '@static/img/MarketPlacePage/marketPlaceCardsGroup/gold2.svg';
import silver from '@static/img/MarketPlacePage/marketPlaceCardsGroup/silver2.svg';
import wood from '@static/img/MarketPlacePage/marketPlaceCardsGroup/wood2.svg';

import diamondShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/diamondShadow.svg'
import goldShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/goldShadow.svg'
import silverShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/silverShadow.svg'
import woodShadow from '@static/img/MarketPlacePage/marketPlaceCardsGroup/woodShadow.svg'

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

/** player stats implementation */
export class CardStats {
    /** main stat with substats */
    constructor(
        public title: string = '',
        public abbreviated: string = '',
        public fields: CardStatsField[] = []
    ) { }
    public average: number = this.fields
        .map(item => item.value)
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

/** Main Player information implementation */
export class CardMainInfo {
    /** main card datas */
    constructor(
        public lastName: string,
        public price: number,
        public playerFace: string,
        public priceIcon: string,
        public priceGoldIcon: string,
        public confirmIcon: string,
        public bgType: number,
    ) { }
    /** backgroundtype picture that depend on quality */
    get backgroundType() {
        const qualities = [
            diamond, gold, silver, wood
        ];
        let background = qualities[this.bgType];
        return background;
    };
    get shadowType() {
        const qualities = [
            diamondShadow, goldShadow, silverShadow, woodShadow
        ];
        let shadow = qualities[this.bgType];
        return shadow;
    };
    /** get image with price status depend on price status */
    get priceStatus() {
        const statuses = [
            currentBid, minimumPrice, purchased, currentBid
        ];
        let status = statuses[this.bgType];
        return status;
    };
}

/** implementation field that uses in CardStats class */
export class CardStatsField {
    /** subStat: label(name) and depend value */
    constructor(
        public label: string = '',
        public value: number = 0
    ) { }
}

/** implementation that uses in overall info*/
export class CardInfoField {
    /** overall label name, depend value and depend icon */
    constructor(
        public label: string = '',
        public value: string = '',
        public icon?: string
    ) { }
}

/** price information implementation */
export class CardPriceField {
    /** price and depend value */
    constructor(
        public label: string = '',
        public value: number | string
    ) { }
}

/** price ID implementation */
export class CardPriceId {
    /** label(id) and depend value */
    constructor(
        public label: string = '',
        public value: number | string
    ) { }
}

/** another price information implementation */
export class CardPricePRP {
    /** label(prp) and depend value */
    constructor(
        public label: string = '',
        public value: number = 0
    ) { }
}

/** base price class */
export class CardPrice {
    /** price datas */
    constructor(
        public id: CardPriceField,
        public price: CardPriceField,
        public prp: CardPricePRP,
        public updated: CardPriceField,
        public pr: CardPriceField,
    ) { }
    /** get stat giagram color depend on price value  */
    get color() {
        const PRICE_UPPER_BOUND = 80;
        const PRICE_MEDIUM_BOUND = 70;
        const PRICE_LOWER_BOUND = 50;

        const PRICE_UPPER_BOUND_COLOR = '#1898D7';
        const PRICE_MEDIUM_BOUND_COLOR = '#3CCF5D';
        const PRICE_LOWER_BOUND_COLOR = '#E86C27';
        const PRICE_DEFAULT_BOUND_COLOR = '#FF4200';
        switch (true) {
            case this.prp.value >= PRICE_UPPER_BOUND:
                return PRICE_UPPER_BOUND_COLOR;
            case this.prp.value >= PRICE_MEDIUM_BOUND:
                return PRICE_MEDIUM_BOUND_COLOR;
            case this.prp.value >= PRICE_LOWER_BOUND:
                return PRICE_LOWER_BOUND_COLOR;
            default:
                return PRICE_DEFAULT_BOUND_COLOR;
        }
    }
}

/** base diagram of player stats implementation */
export class Diagram {
    /** player stats datas */
    constructor(
        public id: string,
        public name: string,
        public min: number,
        public max: number,
        public value: number,
    ) { }
}

/** football field implementation  */
export class FotballFieldInformationLine {
    /** football field datas */
    constructor(
        public id: string = '',
        public title: string = '',
        public options: string[] = []
    ) { }
}

/** Card base implementation */
export class Card {
    /** constructor has private bgType for test */
    constructor(private bgType: number) { }
    mainInfo = new CardMainInfo(
        'Ronalculus',
        1000000,
        playerFace,
        priceIcon,
        priceGoldIcon,
        confirmIcon,
        this.bgType,
    )
    overalInfo = [
        new CardInfoField('name', 'Albert Ronalculus'),
        new CardInfoField('nation', 'Portugal 🇵🇹'),
        new CardInfoField('skills', '5', star),
        new CardInfoField('weak foot', '5', star),
        new CardInfoField('intl. rep', '5', star),
        new CardInfoField('foot', 'right'),
        new CardInfoField('height', '187'),
        new CardInfoField('nation', '87'),
        new CardInfoField('revision', 'rare'),
        new CardInfoField('def. wr', 'low'),
        new CardInfoField('arr. wr', 'high'),
        new CardInfoField('added on', '2020-09-10'),
        new CardInfoField('origin', 'na'),
        new CardInfoField('r. Face', 'low'),
        new CardInfoField('b. type', '', checked),
        new CardInfoField('age', '36 years old')
    ]
    stats = [
        new CardStats('tactics', 'tac', [
            new CardStatsField('positioning', 100),
            new CardStatsField('composure', 95,),
            new CardStatsField('aggression', 98),
            new CardStatsField('vision', 98),
            new CardStatsField('awareness', 99),
            new CardStatsField('crosses', 98),
        ]),
        new CardStats('physique', 'phy', [
            new CardStatsField('acceleration', 26),
            new CardStatsField('running speed', 25),
            new CardStatsField('reaction speed', 45),
            new CardStatsField('agility', 31),
            new CardStatsField('stamina', 40),
            new CardStatsField('strength', 35),
            new CardStatsField('jumping', 28),
            new CardStatsField('balance', 42),
        ]),
        new CardStats('technique', 'tec', [
            new CardStatsField('dribbing', 26),
            new CardStatsField('ball fontrol', 26),
            new CardStatsField('weak foot', 26),
            new CardStatsField('skill moves', 26),
            new CardStatsField('finesse', 26),
            new CardStatsField('curve', 26),
            new CardStatsField('volleys', 26),
            new CardStatsField('short passing', 26),
            new CardStatsField('long passing', 26),
            new CardStatsField('forward pass', 26),
        ]),
        new CardStats('offence', 'off', [
            new CardStatsField('finishing ability', 42),
            new CardStatsField('shot power', 42),
            new CardStatsField('accuracy', 42),
            new CardStatsField('distance', 42),
            new CardStatsField('penalty', 42),
            new CardStatsField('free Kicks', 42),
            new CardStatsField('corners', 42),
            new CardStatsField('heading accuracy', 42),
        ]),
        new CardStats('defence', 'def', [
            new CardStatsField('offside trap', 74),
            new CardStatsField('tackles', 74),
            new CardStatsField('ball focus', 74),
            new CardStatsField('interceptions', 74),
            new CardStatsField('vigilance', 74),
        ]),
        new CardStats('goalkeeping', 'gk', [
            new CardStatsField('diving', 84),
            new CardStatsField('handling', 84),
            new CardStatsField('sweeping', 84),
            new CardStatsField('throwing', 84),
        ])
    ]
    price = new CardPrice(
        new CardPriceId('id', 1),
        new CardPriceField('price', '11,400,00'),
        new CardPricePRP('prp', 75),
        new CardPriceField('updated', 16),
        new CardPriceField('pr', '1,142,000 - 15,000,000'),
    )
    diagram = [
        new Diagram('1', 'physical', 100, 800, 688),
        new Diagram('2', 'mental', 100, 800, 688),
        new Diagram('3', 'skill', 100, 800, 688),
        new Diagram('4', 'chem. style', 100, 800, 688),
        new Diagram('5', 'base stats', 100, 800, 688),
        new Diagram('6', 'in game stats', 100, 800, 688),
    ]
}
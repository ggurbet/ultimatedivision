/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

/* eslint-disable */
import diamond from '../img/MarketPlacePage/marketPlaceCardsGroup/diamond2.png';
import gold from '../img/MarketPlacePage/marketPlaceCardsGroup/gold2.png';
import silver from '../img/MarketPlacePage/marketPlaceCardsGroup/silver2.png';
import wood from '../img/MarketPlacePage/marketPlaceCardsGroup/wood2.png';

import currentBid
    from '../img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/bid.png';
import minimumPrice
    from '../img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/minimum.png';
import purchased
    from '../img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/purchased.png';

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

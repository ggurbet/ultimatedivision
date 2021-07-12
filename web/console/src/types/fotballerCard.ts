/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */


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
export class CardStats {
    public average: number = this.fields
        .map(item => item.value)
        .reduce((prev, current) => prev + current) / this.fields.length;
    constructor(
        public title: string = '',
        public abbreviated: string = '',
        public fields: CardStatsField[] = []
    ) {
    }
    get abbr(): string {
        return this.title.slice(0, 3);
    }
    get color(): string {
        const STATISTIC_UPPER_BOUND = 90;
        const STATISTIC_LOWER_BOUND = 50;

        const STATISTIC_UPPER_BOUND_COLOR = '#3CCF5D';
        const STATISTIC_MEDIUM_BOUND_COLOR = '#E8EC16';
        const STATISTIC_LOWER_BOUND_COLOR = '#FF4200';

        switch (true) {
            case (this.average >= STATISTIC_UPPER_BOUND):
                return STATISTIC_UPPER_BOUND_COLOR;
            case (this.average >= STATISTIC_LOWER_BOUND):
                return STATISTIC_MEDIUM_BOUND_COLOR;
            default:
                return STATISTIC_LOWER_BOUND_COLOR;
        }
    }
}

export class CardMainInfo {
    constructor(
        public lastName: string,
        public price: number,
        public playerFace: string,
        public priceIcon: string,
        public priceGoldIcon: string,
        public confirmIcon: string,
    ) { }
    get backgroundType() {
        /*
        * bakgroundtype picture that depend on quality
        */
        const qualities = [
            diamond, gold, silver, wood
        ];
        let background = qualities[Math.floor(Math.random()
            * qualities.length)];
        return background;
    };
    get priceStatus() {
        /*
        * get image with price status depend on price status
        */
        const statuses = [
            currentBid, minimumPrice, purchased
        ];
        let status = statuses[Math.floor(Math.random()
            * statuses.length)];
        return status;
    };
}

export class CardStatsField {
    constructor(
        public label: string = '',
        public value: number = 0
    ) { }
}

export class CardInfoField {
    constructor(
        public label: string = '',
        public value: string = '',
        public icon?: string
    ) { }
}
export class CardPriceField {
    constructor(
        public label: string = '',
        public value: number | string
    ) { }
}
export class CardPriceId {
    constructor(
        public label: string = '',
        public value: number | string
    ) { }
}
export class CardPricePRP {
    constructor(
        public label: string = '',
        public value: number = 0
    ) { }
}

export class CardPrice {
    constructor(
        public id: CardPriceField,
        public price: CardPriceField,
        public prp: CardPricePRP,
        public updated: CardPriceField,
        public pr: CardPriceField,
    ) { }

    get color() {
        const PRICE_UPPER_BOUND = 80;
        const PRICE_MEDIUM_BOUND = 70;
        const PRICE_LOWER_BOUND = 50;

        const PRICE_UPPER_BOUND_COLOR = '#1898D7';
        const PRICE_MEDIUM_BOUND_COLOR = '#3CCF5D';
        const PRICE_LOWER_BOUND_COLOR = '#E86C27';
        const PRICE_DEFAULT_BOUND_COLOR = '#FF4200';
        switch (true) {
            case (this.prp.value >= PRICE_UPPER_BOUND):
                return PRICE_UPPER_BOUND_COLOR;
            case (this.prp.value >= PRICE_MEDIUM_BOUND):
                return PRICE_MEDIUM_BOUND_COLOR;
            case (this.prp.value >= PRICE_LOWER_BOUND):
                return PRICE_LOWER_BOUND_COLOR;
            default:
                return PRICE_DEFAULT_BOUND_COLOR;
        }
    }
}

export class Diagram {
    constructor(
        public id: string,
        public name: string,
        public min: number,
        public max: number,
        public value: number,
    ) { }
}

export class FotballFieldInformationLine {
    constructor(
        public id: string = '',
        public title: string = '',
        public options: string[] = []
    ) { }
}

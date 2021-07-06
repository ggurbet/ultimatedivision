/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

export class CardStats {
    public average: number = this.fields
        .map(item => item.value)
        .reduce((prev, current) => prev + current) / this.fields.length;
    constructor(
        public title: string = '',
        public fields: CardStatsField[] = []
    ) {
    }
    get abbr (): string {
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
        public id: string ='',
        public title: string = '',
        public options: string[] = []
    ) { }
}

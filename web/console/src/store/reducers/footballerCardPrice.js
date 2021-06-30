/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

export class PriceCard {
    constructor(fields) {
        this.fields = fields;
    }
    get color() {
        switch (true) {
        case (this.fields.prp.value >= 80):
            return '#1898D7';
        case (this.fields.prp.value >= 70):
            return '#3CCF5D';
        case (this.fields.prp.value >= 50):
            return '#E86C27';
        default:
            return '#FF4200';
        }
    }
}
class CardPriceField {
    constructor(key, value) {
        this.label = key;
        this.value = value;
    }
}

const priceAreaData = new PriceCard({
    id: new CardPriceField('id', 1),
    price: new CardPriceField('Price', '11,400,00'),
    prp: new CardPriceField('PRP', 75),
    updated: new CardPriceField('updated', 16),
    pr:new CardPriceField('PR', '1,142,000 - 15,000,000'),
});

/* eslint-disable */
export const cardPriceReducer = (cardState = priceAreaData, action) => {
    return cardState;
}
/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import confirmIcon from '../../static/img/MarketPlacePage/MyCard/ok.svg';
import priceGoldIcon from '../../static/img/MarketPlacePage/MyCard/goldPrice.svg';
import playerFace from '../../static/img/MarketPlacePage/marketPlaceCardsGroup/player.svg';
import priceIcon
    from '../../static/img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/price.svg';

import checked from '../../static/img/FootballerCardPage/checked.svg';
import star from '../../static/img/FootballerCardPage/star.svg';

import {
    CardInfoField, CardMainInfo, CardPrice, CardPriceField,
    CardPriceId, CardPricePRP, CardStats, CardStatsField, Diagram,
} from '../../types/fotballerCard';

/* eslint-disable */

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
/** create list of player cards (implementation for test)*/
function cardList(count: number) {
    let list: Card[] = [];
    while (count > 0) {
        list.push(new Card(0), new Card(1), new Card(2), new Card(3));
        count--;
    }
    return list;

}

export const cardReducer = (cardState = cardList(100)) => {
    return cardState;
};

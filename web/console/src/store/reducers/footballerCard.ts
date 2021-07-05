/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import player from '../../img/MarketPlacePage/marketPlaceCardsGroup/player.png';
import price
    from '../../img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/price.png';

import diamond from '../../img/MarketPlacePage/marketPlaceCardsGroup/diamond2.png';
import gold from '../../img/MarketPlacePage/marketPlaceCardsGroup/gold2.png';
import silver from '../../img/MarketPlacePage/marketPlaceCardsGroup/silver2.png';
import wood from '../../img/MarketPlacePage/marketPlaceCardsGroup/wood2.png';

import currentBid
    from '../../img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/bid.png';
import minimumPrice
    from '../../img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/minimum.png';
import purchased
    from '../../img/MarketPlacePage/marketPlaceCardsGroup/marketPlaceFootballerCard/purchased.png';

import star from '../../img/FootballerCardPage/star.png';
import checked from '../../img/FootballerCardPage/checked.png';

import { CardStats } from '../../types/fotballerCard';
import { CardStatsField } from '../../types/fotballerCard';
import { CardInfoField } from '../../types/fotballerCard';
import { CardPrice } from "../../types/fotballerCard";
import { CardPriceField } from '../../types/fotballerCard';
import { CardPriceId } from '../../types/fotballerCard';
import { CardPricePRP } from '../../types/fotballerCard';
import { Diagram } from '../../types/fotballerCard';

class Card {
    mainInfo = {
        price: 1000000,
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
        },
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
        },
        facePicture: player,
        pricePicture: price,
    };
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
        new CardStats('tactics', [
            new CardStatsField('positioning', 100),
            new CardStatsField('composure', 95,),
            new CardStatsField('aggression', 98),
            new CardStatsField('vision', 98),
            new CardStatsField('awareness', 99),
            new CardStatsField('crosses', 98),
        ]),
        new CardStats('physique', [
            new CardStatsField('acceleration', 26),
            new CardStatsField('running speed', 25),
            new CardStatsField('reaction speed', 45),
            new CardStatsField('agility', 31),
            new CardStatsField('stamina', 40),
            new CardStatsField('strength', 35),
            new CardStatsField('jumping', 28),
            new CardStatsField('balance', 42),
        ]),
        new CardStats('technique', [
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
        new CardStats('offence', [
            new CardStatsField('finishing ability', 42),
            new CardStatsField('shot power', 42),
            new CardStatsField('accuracy', 42),
            new CardStatsField('distance', 42),
            new CardStatsField('penalty', 42),
            new CardStatsField('free Kicks', 42),
            new CardStatsField('corners', 42),
            new CardStatsField('heading accuracy', 42),
        ]),
        new CardStats('defence', [
            new CardStatsField('offside trap', 74),
            new CardStatsField('tackles', 74),
            new CardStatsField('ball focus', 74),
            new CardStatsField('interceptions', 74),
            new CardStatsField('vigilance', 74),
        ]),
        new CardStats('goalkeeping', [
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

function cardlist(count: number) {
    let quantity = count;
    const list = [];
    while (quantity > 0) {
        list.push(new Card());
        quantity--;
    }
    return list;
}

export const cardReducer = (cardState = cardlist(100)) => {
    return cardState;
};

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction } from 'react';
import { useDispatch } from 'react-redux';

import { LootboxStats } from '@/app/types/lootbox';

import { openLootbox } from '@/app/store/actions/lootboxes';


import wood from '@static/img/StorePage/BoxCard/wood.svg';
import silver from '@static/img/StorePage/BoxCard/silver.svg';
import gold from '@static/img/StorePage/BoxCard/gold.svg';
import diamond from '@static/img/StorePage/BoxCard/diamond.svg';
import coin from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';
import { LootboxCardQuality } from './LootboxCardQuality';

import './index.scss';

export const LootboxCard: React.FC<{ data: LootboxStats; handleOpening: Dispatch<SetStateAction<boolean>> }> = ({ data, handleOpening }) => {
    const dispatch = useDispatch();
    const qualities = [
        {
            name: 'Wood',
            icon: wood,
        },
        {
            name: 'Silver',
            icon: silver,
        },
        {
            name: 'Gold',
            icon: gold,
        },
        {
            name: 'Diamond',
            icon: diamond,
        },
    ];

    const handleAnimation = async() => {
        // TODO: need add id lootbox from BD after be create endpoint fetch lootboxex.
        await dispatch(openLootbox({ id: data.id, type: data.type }));
        handleOpening(true);
    };

    return (
        <div className="box-card">
            <div className="box-card__wrapper">
                <div className="box-card__description">
                    <div
                        className="box-card__icon"
                        style={{ backgroundImage: `url(${data.icon})` }}
                    />
                    <h2 className="box-card__title">{data.type === 'Regular Box' ? 'Regular Box': 'Cool box' }</h2>
                    <div className="box-card__quantity">
                        <span className="box-card__quantity-label">Cards</span>
                        <span className="box-card__quantity-value">{data.quantity}</span>
                    </div>
                </div>
                <div className="box-card__qualities">
                    {data.dropChance.map((item, index) =>
                        <LootboxCardQuality
                            label={qualities[index]}
                            chance={item}
                            key={index}
                        />
                    )}
                    <button
                        className="box-card__button"
                        onClick={handleAnimation}
                    >
                        <span className="box-card__button-text">OPEN</span>
                        <span className="box-card__button-value">
                            <img src={coin} alt="coin" />
                            {data.price}
                        </span>
                    </button>
                </div>
            </div>
        </div>
    );
};

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Dispatch, SetStateAction, useState } from 'react';
import { useDispatch } from 'react-redux';
import { toast } from 'react-toastify';

import { LootboxCardQuality } from './LootboxCardQuality';

import { RegistrationPopup } from '@/app/components/common/Registration/Registration';

import coin from '@static/img/MarketPlacePage/MyCard/goldPrice.svg';
import diamond from '@static/img/StorePage/BoxCard/diamond.svg';
import gold from '@static/img/StorePage/BoxCard/gold.svg';
import silver from '@static/img/StorePage/BoxCard/silver.svg';
import wood from '@static/img/StorePage/BoxCard/wood.svg';

import { UnauthorizedError } from '@/api';
import { openLootbox } from '@/app/store/actions/lootboxes';
import { LootboxStats } from '@/app/types/lootbox';

import './index.scss';

export const LootboxCard: React.FC<{
    data: LootboxStats;
    handleOpening: Dispatch<SetStateAction<boolean>>;
}> = ({ data, handleOpening }) => {
    /** Indicates if registration required. */
    const [isRegistrationRequired, setIsRegistrationRequired] = useState(false);

    /** Closes Registration popup componnet. */
    const closeRegistrationPopup = () => {
        setIsRegistrationRequired(false);
    };

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
        try {
            await dispatch(openLootbox({ id: data.id, type: data.type }));

            handleOpening(true);
        } catch (error: any) {
            if (error instanceof UnauthorizedError) {
                setIsRegistrationRequired(true);

                return;
            };

            toast.error('Failed to open lootbox', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    if (isRegistrationRequired) {
        return <RegistrationPopup closeRegistrationPopup={closeRegistrationPopup} />;
    };

    return (
        <div className="box-card">
            <div className="box-card__wrapper">
                <div className="box-card__description">
                    <img className="box-card__icon" src={data.icon} alt="box" />
                    <h2 className="box-card__title">
                        {data.type === 'Regular Box'
                            ? 'Regular Box'
                            : 'Cool box'}
                    </h2>
                    <div className="box-card__quantity">
                        <span className="box-card__quantity-label">Cards</span>
                        <span className="box-card__quantity-value">
                            {data.quantity}
                        </span>
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
